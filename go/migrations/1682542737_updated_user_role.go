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

		collection, err := dao.FindCollectionByNameOrId("7l42f9e4ng1frj7")
		if err != nil {
			return err
		}

		// update
		edit_user := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wuddjpzf",
			"name": "user",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": [
					"email"
				]
			}
		}`), edit_user)
		collection.Schema.AddField(edit_user)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("7l42f9e4ng1frj7")
		if err != nil {
			return err
		}

		// update
		edit_user := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wuddjpzf",
			"name": "user",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": [
					"email"
				]
			}
		}`), edit_user)
		collection.Schema.AddField(edit_user)

		return dao.SaveCollection(collection)
	})
}
