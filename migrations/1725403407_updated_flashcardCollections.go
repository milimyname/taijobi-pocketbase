package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("23a40op2xcy6dh4")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.id != \"\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id =\nuserId)")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id =\nuserId)")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("23a40op2xcy6dh4")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\"")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\"")

		return dao.SaveCollection(collection)
	})
}
