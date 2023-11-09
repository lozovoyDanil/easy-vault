package repository

/*
import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)
*/

/*
func check(db *sqlx.DB, userId, target int, table, prop string) error {
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE %s = ?", table, prop)

	var res int
	err := db.Get(&res, query, target)
	if err != nil {
		return err
	}
	if res != userId {
		return errors.New("access forbiden: user ownership violation")
	}

	return nil
}

func checkFun(db *sqlx.DB, target int, fun func() string) error {
	query := fun()

	var res int
	err := db.Get(&res, query)
	if err != nil {
		return err
	}
	if res != target {
		return errors.New("access forbiden: user ownership violation")
	}

	return nil
}
*/

/*
	err := check(r.db, userId, spaceId, userSpacesTable, "space_id")
	if err != nil {
		return err
	}

	err = checkFun(r.db, userId, func() string {
		return fmt.Sprintf("SELECT user_id FROM %s WHERE space_id = ?", userSpacesTable)
	})
	if err != nil {
		return err
	}

*/
