package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"main.go/internal/model"
)

type SpaceSQLite struct {
	db *sqlx.DB
}

func NewSpaceSQLite(db *sqlx.DB) *SpaceSQLite {
	return &SpaceSQLite{db: db}
}

func (r *SpaceSQLite) AllSpaces() ([]model.Space, error) {
	var spaces []model.Space

	query := fmt.Sprintf("SELECT s.id, s.name, s.addr, s.numOfGroups, s.size, s.numOfFree FROM %s s INNER JOIN %s us ON s.id=us.space_id", spaceTable, userSpacesTable)
	err := r.db.Select(&spaces, query)

	return spaces, err
}

func (r *SpaceSQLite) UserSpaces(id int) ([]model.Space, error) {
	var spaces []model.Space

	query := fmt.Sprintf("SELECT s.id, s.name, s.addr, s.numOfGroups, s.size, s.numOfFree FROM %s s INNER JOIN %s us ON s.id=us.space_id WHERE us.user_id=$1", spaceTable, userSpacesTable)
	err := r.db.Select(&spaces, query, id)

	return spaces, err
}

func (r *SpaceSQLite) SpaceById(spaceId int) (model.Space, error) {
	var space model.Space

	query := fmt.Sprintf("SELECT s.id, s.name, s.addr, s.numOfGroups, s.size, s.numOfFree FROM %s s INNER JOIN %s us ON s.id=us.space_id WHERE us.space_id=$1", spaceTable, userSpacesTable)
	err := r.db.Get(&space, query, spaceId)

	return space, err
}

func (r *SpaceSQLite) CreateSpace(userId int, space model.Space) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s(name, addr,numOfGroups, size, numOfFree) VALUES($1,$2,$3,$4,$5) RETURNING id", spaceTable)
	row := tx.QueryRow(query, space.Name, space.Addr, space.NumOfGroups, space.Size, space.NumOfFree)
	var id int
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s(user_id, space_id) VALUES($1, $2)", userSpacesTable)
	_, err = tx.Exec(query, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *SpaceSQLite) UpdateSpace(userId, spaceId int, input model.UpdateSpaceInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Addr != nil {
		setValues = append(setValues, fmt.Sprintf("addr=$%d", argId))
		args = append(args, *input.Addr)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s s SET %s FROM %s us WHERE s.id = us.space_id AND us.space_id = $%d AND us.user_id = $%d",
		spaceTable, setQuery, userSpacesTable, argId, argId+1)
	args = append(args, spaceId, userId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *SpaceSQLite) DeleteSpace(userId, spaceId int) error {
	//* Starting transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE space_id=$1 AND user_id=$2", userSpacesTable)
	res, err := tx.Exec(query, spaceId, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	// If res.RowsAffected() returns 0, this means that eather
	// space does not exist or user does not own it.
	if r, _ := res.RowsAffected(); r == 0 {
		tx.Rollback()
		return errors.New("access forbiden or object does not exist")
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id = ?", spaceTable)
	_, err = tx.Exec(query, spaceId)
	if err != nil {
		tx.Rollback()
		return err
	}

	//! All the groups and units must be deleted too.
	//TODO: Add calls to delete them, or implement them here.

	return tx.Commit()
}
