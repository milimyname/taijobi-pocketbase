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
		new_searches := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "hvwg2dzo",
			"name": "searches",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "wv3yy8dyn1m8gqg",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), new_searches); err != nil {
			return err
		}
		collection.Schema.AddField(new_searches)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("b48svadkeybl4f2")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("hvwg2dzo")

		return dao.SaveCollection(collection)
	})
}
