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

		// update
		edit_description := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "hbxzwvzq",
			"name": "description",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": 100,
				"pattern": ""
			}
		}`), edit_description); err != nil {
			return err
		}
		collection.Schema.AddField(edit_description)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("l0kibsvt0ol53vc")
		if err != nil {
			return err
		}

		// update
		edit_description := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "hbxzwvzq",
			"name": "description",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": 50,
				"pattern": ""
			}
		}`), edit_description); err != nil {
			return err
		}
		collection.Schema.AddField(edit_description)

		return dao.SaveCollection(collection)
	})
}
