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

		collection, err := dao.FindCollectionByNameOrId("b48svadkeybl4f2")
		if err != nil {
			return err
		}

		// add
		new_copiedFlashcard := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "tiiyij1i",
			"name": "copiedFlashcard",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "b48svadkeybl4f2",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_copiedFlashcard); err != nil {
			return err
		}
		collection.Schema.AddField(new_copiedFlashcard)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("b48svadkeybl4f2")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("tiiyij1i")

		return dao.SaveCollection(collection)
	})
}
