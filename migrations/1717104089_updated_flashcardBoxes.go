package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("l0kibsvt0ol53vc")
		if err != nil {
			return err
		}

		// add
		new_quizCount := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "0qqndlxe",
			"name": "quizCount",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": true
			}
		}`), new_quizCount); err != nil {
			return err
		}
		collection.Schema.AddField(new_quizCount)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("l0kibsvt0ol53vc")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("0qqndlxe")

		return dao.SaveCollection(collection)
	})
}
