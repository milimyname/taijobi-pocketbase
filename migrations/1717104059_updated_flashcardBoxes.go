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
		new_kanjiCount := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ke9n2y8b",
			"name": "kanjiCount",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": true
			}
		}`), new_kanjiCount); err != nil {
			return err
		}
		collection.Schema.AddField(new_kanjiCount)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("l0kibsvt0ol53vc")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("ke9n2y8b")

		return dao.SaveCollection(collection)
	})
}
