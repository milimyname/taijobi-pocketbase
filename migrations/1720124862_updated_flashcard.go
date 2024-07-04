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

		// update
		edit_partOfSpeech := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vxg4kcop",
			"name": "partOfSpeech",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"verb",
					"adjective",
					"interjection",
					"pronoun",
					"particle",
					"unknown"
				]
			}
		}`), edit_partOfSpeech); err != nil {
			return err
		}
		collection.Schema.AddField(edit_partOfSpeech)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("b48svadkeybl4f2")
		if err != nil {
			return err
		}

		// update
		edit_partOfSpeech := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vxg4kcop",
			"name": "partOfSpeech",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"verb",
					"adjective",
					"unknown"
				]
			}
		}`), edit_partOfSpeech); err != nil {
			return err
		}
		collection.Schema.AddField(edit_partOfSpeech)

		return dao.SaveCollection(collection)
	})
}
