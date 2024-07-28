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

		collection, err := dao.FindCollectionByNameOrId("abkwr5075z4bem0")
		if err != nil {
			return err
		}

		// add
		new_chat := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vhcdbtrq",
			"name": "chat",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "noo1xxcyvic2b8v",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_chat); err != nil {
			return err
		}
		collection.Schema.AddField(new_chat)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("abkwr5075z4bem0")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("vhcdbtrq")

		return dao.SaveCollection(collection)
	})
}
