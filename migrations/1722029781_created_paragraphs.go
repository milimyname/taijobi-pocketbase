package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "exbuwsccnr3ennr",
			"created": "2024-07-26 21:36:21.286Z",
			"updated": "2024-07-26 21:36:21.286Z",
			"name": "paragraphs",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "haivcfyw",
					"name": "files",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"mimeTypes": [],
						"thumbs": [],
						"maxSelect": 99,
						"maxSize": 5242880,
						"protected": false
					}
				},
				{
					"system": false,
					"id": "xxgckyzg",
					"name": "user",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "1idrgvvn",
					"name": "formatted_ai_data",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 200000000000
					}
				},
				{
					"system": false,
					"id": "1ft6932t",
					"name": "ocr_data",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 200000000000
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("exbuwsccnr3ennr")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
