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

		collection, err := dao.FindCollectionByNameOrId("hjfbuvlt23mnbbh")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id = userId)")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id = userId)")

		collection.CreateRule = types.Pointer("@request.auth.id != \"\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id = userId)")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id = userId)")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("hjfbuvlt23mnbbh")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\"")

		collection.CreateRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\"")

		return dao.SaveCollection(collection)
	})
}
