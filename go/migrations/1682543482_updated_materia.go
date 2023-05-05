package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("7c90jeww18zefjw")
		if err != nil {
			return err
		}

		collection.Name = "matter"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("7c90jeww18zefjw")
		if err != nil {
			return err
		}

		collection.Name = "materia"

		return dao.SaveCollection(collection)
	})
}
