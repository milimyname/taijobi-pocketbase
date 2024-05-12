package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("feedbacks")
		if err != nil {
			return err
		}

		// this returns a pointer so we can directly modify the underlaying field data
		nameField := collection.Schema.GetFieldByName("name")
		nameField.Required = false

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		// Optional: Add rollback or undo logic here if the migration can be reversed
		return nil
	})
}
