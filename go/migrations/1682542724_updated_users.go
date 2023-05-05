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

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("smr3ajhd")

		// add
		new_admin := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "innacpv6",
			"name": "admin",
			"type": "bool",
			"required": false,
			"unique": false,
			"options": {}
		}`), new_admin)
		collection.Schema.AddField(new_admin)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// add
		del_role := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "smr3ajhd",
			"name": "role",
			"type": "select",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"admin",
					"profesor",
					"alumno"
				]
			}
		}`), del_role)
		collection.Schema.AddField(del_role)

		// remove
		collection.Schema.RemoveField("innacpv6")

		return dao.SaveCollection(collection)
	})
}
