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
			"id": "noo1xxcyvic2b8v",
			"created": "2024-07-28 00:07:29.215Z",
			"updated": "2024-07-28 00:07:29.215Z",
			"name": "chats",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "oqdluzdh",
					"name": "messages",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "abkwr5075z4bem0",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "ftelepaa",
					"name": "user",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.id != \"\" &&\n@request.auth.id ?= user.id",
			"viewRule": "@request.auth.id != \"\" &&\n@request.auth.id ?= user.id",
			"createRule": null,
			"updateRule": "@request.auth.id != \"\" &&\n@request.auth.id ?= user.id",
			"deleteRule": "@request.auth.id != \"\" &&\n@request.auth.id ?= user.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("noo1xxcyvic2b8v")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
