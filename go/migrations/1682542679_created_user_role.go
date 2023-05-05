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
			"id": "7l42f9e4ng1frj7",
			"created": "2023-04-26 20:57:59.084Z",
			"updated": "2023-04-26 20:57:59.084Z",
			"name": "user_role",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "cdxtvbpn",
					"name": "matter",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "7c90jeww18zefjw",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": [
							"name"
						]
					}
				},
				{
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
				},
				{
					"system": false,
					"id": "xljxyjsy",
					"name": "role",
					"type": "select",
					"required": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"profesor",
							"alumno"
						]
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_P4bYZ3q` + "`" + ` ON ` + "`" + `user_role` + "`" + ` (\n  ` + "`" + `matter` + "`" + `,\n  ` + "`" + `user` + "`" + `\n)"
			],
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

		collection, err := dao.FindCollectionByNameOrId("7l42f9e4ng1frj7")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
