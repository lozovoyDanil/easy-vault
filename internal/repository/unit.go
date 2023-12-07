package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"main.go/internal/model"
)

type UnitSQLite struct {
	db *sqlx.DB
}

func NewUnitSQLite(db *sqlx.DB) *UnitSQLite {
	return &UnitSQLite{db: db}
}

func (r *UnitSQLite) UnitBelongsToUser(userId, unitId int) error {
	var count int

	query := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s u
		INNER JOIN %s gu ON u.id=gu.unit_id
		INNER JOIN %s sg ON gu.group_id=sg.group_id
		INNER JOIN %s us ON sg.space_id=us.space_id
		WHERE us.user_id=$1 AND u.id=$2`,
		unitTable, groupTable, groupInSpaceTable, userSpacesTable)
	err := r.db.Get(&count, query, userId, unitId)
	if err != nil {
		return err
	}
	if count != 0 {
		return nil
	}

	query = fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s usu
		WHERE usu.user_id=$1`,
		userUnitsTable)
	err = r.db.Get(&count, query, userId)
	if count != 0 {
		return nil
	}

	return err
}

func (r *UnitSQLite) GroupUnits(groupId int) ([]model.StorageUnit, error) {
	var units []model.StorageUnit

	query := fmt.Sprintf(`
		SELECT u.name, u.isOccupied, u.lastUsed, u.busyUntil 
		FROM %s u INNER JOIN %s gu 
		ON u.id=gu.unit_id 
		WHERE group_id = $1`,
		unitTable, unitInGroupTable)
	err := r.db.Select(&units, query, groupId)
	if err != nil {
		return nil, err
	}

	return units, nil
}

func (r *UnitSQLite) UnitById(unitId int) (model.StorageUnit, error) {
	var unit model.StorageUnit

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", unitTable)
	err := r.db.Get(&unit, query, unitId)
	if err != nil {
		return model.StorageUnit{}, err
	}

	return unit, nil
}

func (r *UnitSQLite) CreateUnit(groupId int, unit model.StorageUnit) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (name, isOccupied, lastUsed, busyUntil) VALUES ($1, $2, $3, $4)", unitTable)
	row, err := tx.Exec(query, unit.Name, unit.IsOccupied, unit.LastUsed, unit.BusyUntil)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	unitId, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (group_id, unit_id) VALUES ($1, $2)", unitInGroupTable)
	_, err = tx.Exec(query, groupId, unitId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(unitId), tx.Commit()
}

func (r *UnitSQLite) UpdateUnit(unitId int, input model.UpdateUnitInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.IsOccupied != nil {
		setValues = append(setValues, fmt.Sprintf("isOccupied=$%d", argId))
		args = append(args, *input.IsOccupied)
		argId++
	}
	if input.LastUsed != nil {
		setValues = append(setValues, fmt.Sprintf("lastUsed=$%d", argId))
		args = append(args, *input.LastUsed)
		argId++
	}
	if input.BusyUntil != nil {
		setValues = append(setValues, fmt.Sprintf("busyUntil=$%d", argId))
		args = append(args, *input.BusyUntil)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s u SET %s WHERE u.id=$%d", unitTable, setQuery, argId)
	args = append(args, unitId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *UnitSQLite) DeleteUnit(unitId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", unitTable)
	_, err := r.db.Exec(query, unitId)

	return err
}

func (r *UnitSQLite) ReservedUnits(userId int) ([]model.StorageUnit, error) {
	var units []model.StorageUnit

	query := fmt.Sprintf(`
		SELECT u.name, u.isOccupied, u.lastUsed, u.busyUntil
		FROM %s u
		INNER JOIN %s usu ON u.id=usu.unit_id
		WHERE usu.user_id=$1`,
		unitTable, userUnitsTable)
	err := r.db.Get(&units, query, userId)

	return units, err
}
