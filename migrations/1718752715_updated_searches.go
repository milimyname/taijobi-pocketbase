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

		collection, err := dao.FindCollectionByNameOrId("wv3yy8dyn1m8gqg")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" && user = @request.auth.id")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" && user = @request.auth.id")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && user = @request.auth.id")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" && user = @request.auth.id")

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_VW2mwzE` + "`" + ` ON ` + "`" + `searches` + "`" + ` (\n  ` + "`" + `flashcard` + "`" + `,\n  ` + "`" + `user` + "`" + `\n)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// update
		edit_flashcard := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "haqb1hpb",
			"name": "flashcard",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "b48svadkeybl4f2",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_flashcard); err != nil {
			return err
		}
		collection.Schema.AddField(edit_flashcard)

		// update
		edit_user := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "cqbqzsz9",
			"name": "user",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_user); err != nil {
			return err
		}
		collection.Schema.AddField(edit_user)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("wv3yy8dyn1m8gqg")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" && userId = @request.auth.id")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" && userId = @request.auth.id")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && userId = @request.auth.id")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" && userId = @request.auth.id")

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_VW2mwzE` + "`" + ` ON ` + "`" + `searches` + "`" + ` (\n  ` + "`" + `flashcardId` + "`" + `,\n  ` + "`" + `userId` + "`" + `\n)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// update
		edit_flashcard := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "haqb1hpb",
			"name": "flashcardId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "b48svadkeybl4f2",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_flashcard); err != nil {
			return err
		}
		collection.Schema.AddField(edit_flashcard)

		// update
		edit_user := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "cqbqzsz9",
			"name": "userId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_user); err != nil {
			return err
		}
		collection.Schema.AddField(edit_user)

		return dao.SaveCollection(collection)
	})
}
