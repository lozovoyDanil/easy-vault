package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/uptrace/bun"
	"main.go/internal/model"
)

type UnitSQLite struct {
	db *bun.DB
}

func NewUnitSQLite(db *bun.DB) *UnitSQLite {
	return &UnitSQLite{db: db}
}

func (r *UnitSQLite) UnitOwnerId(unitId int) (int, error) {
	var id int

	err := r.db.NewSelect().
		Table(unitTable).
		ColumnExpr("user_id").
		Where("id = ?", unitId).
		Scan(context.Background(), &id)

	return id, err
}

func (r *UnitSQLite) ManagerOwnsUnit(userId, unitId int) bool {
	var count int

	count, err := r.db.NewSelect().
		Table(unitTable).
		ColumnExpr("u.id").
		Join(fmt.Sprintf("INNER JOIN %s g ON g.id = u.group_id", groupTable)).
		Join(fmt.Sprintf("INNER JOIN %s s ON s.id = g.space_id", spaceTable)).
		Join(fmt.Sprintf("INNER JOIN %s us ON us.space_id = s.id", userSpacesTable)).
		Where("us.user_id = ?", userId).
		Where("u.id = ?", unitId).
		Count(context.Background())

	return count > 0 && err == nil
}

func (r *UnitSQLite) ManagerUnitsCount(userId int) (int, error) {
	var count int

	count, err := r.db.NewSelect().
		Table(unitTable).
		ColumnExpr("u.id").
		Join(fmt.Sprintf("INNER JOIN %s g ON g.id = u.group_id", groupTable)).
		Join(fmt.Sprintf("INNER JOIN %s s ON s.id = g.space_id", spaceTable)).
		Join(fmt.Sprintf("INNER JOIN %s us ON us.space_id = s.id", userSpacesTable)).
		Where("us.user_id = ?", userId).
		Count(context.Background())

	return count, err
}

func (r *UnitSQLite) GroupUnits(groupId int) ([]model.StorageUnit, error) {
	var units []model.StorageUnit

	err := r.db.NewSelect().
		Model(&units).
		Where("group_id = ?", groupId).
		Scan(context.Background())

	return units, err
}

func (r *UnitSQLite) UnitById(unitId int) (model.StorageUnit, error) {
	var unit model.StorageUnit

	err := r.db.NewSelect().
		Model(&unit).
		Where("id = ?", unitId).
		Scan(context.Background())

	return unit, err
}

func (r *UnitSQLite) CreateUnit(unit model.StorageUnit) (int, error) {
	_, err := r.db.NewInsert().
		Model(&unit).
		Exec(context.Background())

	return unit.Id, err
}

func (r *UnitSQLite) UpdateUnit(unitId int, input model.UnitInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.UserId != nil {
		setValues = append(setValues, "user_id=?")
		args = append(args, *input.UserId)
		argId++
	}
	if input.Name != nil {
		setValues = append(setValues, "name=?")
		args = append(args, *input.Name)
		argId++
	}
	if input.IsOccupied != nil {
		setValues = append(setValues, "isOccupied=?")
		args = append(args, *input.IsOccupied)
		argId++
	}
	if input.LastUsed != nil {
		setValues = append(setValues, "lastUsed=?")
		args = append(args, *input.LastUsed)
		argId++
	}
	if input.BusyUntil != nil {
		setValues = append(setValues, "busyUntil=?")
		args = append(args, *input.BusyUntil)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s AS u SET %s WHERE u.id=?", unitTable, setQuery)
	args = append(args, unitId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *UnitSQLite) DeleteUnit(unitId int) error {
	_, err := r.db.NewDelete().
		Table(unitTable).
		Where("id = ?", unitId).
		Exec(context.Background())

	return err
}

func (r *UnitSQLite) ReservedUnits(userId int) ([]model.StorageUnit, error) {
	var units []model.StorageUnit

	err := r.db.NewSelect().
		Model(&units).
		Where("u.user_id = ?", userId).
		Scan(context.Background())

	return units, err
}

func (r *UnitSQLite) LogHistory(log model.UnitHistory) error {
	_, err := r.db.NewInsert().
		Model(&log).
		Exec(context.Background())

	return err
}

func (r *UnitSQLite) UnitHistory(unitId int) ([]model.UnitHistory, error) {
	var logs []model.UnitHistory

	err := r.db.NewSelect().
		Model(&logs).
		Where("unit_id = ?", unitId).
		Scan(context.Background())

	return logs, err
}
