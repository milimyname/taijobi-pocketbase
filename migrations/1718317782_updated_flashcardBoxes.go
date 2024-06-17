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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("l0kibsvt0ol53vc")
		if err != nil {
			return err
		}

		// add
		new_flashcards := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "il2kuk1v",
			"name": "flashcards",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "b48svadkeybl4f2",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), new_flashcards); err != nil {
			return err
		}
		collection.Schema.AddField(new_flashcards)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("l0kibsvt0ol53vc")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("il2kuk1v")

		return dao.SaveCollection(collection)
	})

}
