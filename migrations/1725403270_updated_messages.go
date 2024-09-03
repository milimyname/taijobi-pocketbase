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

		collection, err := dao.FindCollectionByNameOrId("abkwr5075z4bem0")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id =\nuser)")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id =\nuser)")

		collection.CreateRule = types.Pointer("@request.auth.id != \"\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id =\nuser)")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" && (@request.auth.role ~ \"admin\" || @request.auth.id =\nuser)")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("abkwr5075z4bem0")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" &&\n@request.auth.id ?= user.id")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" &&\n@request.auth.id ?= user.id")

		collection.CreateRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" &&\n@request.auth.id ?= user.id")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" &&\n@request.auth.id ?= user.id")

		return dao.SaveCollection(collection)
	})
}
