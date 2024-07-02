package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("oerbnabuift1b4n")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" &&\n@request.auth.id ?= user.id")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" &&\n@request.auth.id ?= user.id")

		collection.CreateRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" &&\n@request.auth.id ?= user.id")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\"")

		// add
		new_user := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ljqizq90",
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
		}`), new_user); err != nil {
			return err
		}
		collection.Schema.AddField(new_user)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("oerbnabuift1b4n")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		collection.ViewRule = nil

		collection.CreateRule = nil

		collection.UpdateRule = nil

		collection.DeleteRule = nil

		// remove
		collection.Schema.RemoveField("ljqizq90")

		return dao.SaveCollection(collection)
	})
}
