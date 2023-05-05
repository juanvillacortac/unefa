package server

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func disableAdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Path(), "/_") {
				return apis.NewForbiddenError("You are not allowed to access this resource", nil)
			}
			return next(c)
		}
	}
}

func StartServer(production bool) {
	homeDir, _ := os.UserHomeDir()
	path := fmt.Sprintf("%s/.unefa", homeDir)
	app := pocketbase.NewWithConfig(&pocketbase.Config{
		DefaultDataDir: path,
	})

	if production {
		app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
			e.Router.Use(disableAdminMiddleware())
			return nil
		})
	} else {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
			Automigrate: true, // auto creates migration files when making collection changes
			Dir:         fmt.Sprintf("%s/go/migrations", path),
		})
	}

	if production {
		app.RootCmd.SetArgs([]string{"serve", "--http", "127.0.0.1:9080"})
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
