package main

import (
	"errors"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"os"
	"strings"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		dbType := cel.DB.DatabaseType
		if arg3 == "" {
			exitGracefully(errors.New("you must give the migration a name"))
		}

		// follow migration naming conventions: {current unix micro-time}_{whatever name}
		filename := fmt.Sprintf("%d_%s", time.Now().UTC().UnixMicro(), arg3)
		upFile := cel.RootPath + "/migrations/" + filename + "." + dbType + ".up.sql"
		downFile := cel.RootPath + "/migrations/" + filename + "." + dbType + ".down.sql"

		err := copyFileFromTemplate(fmt.Sprintf("templates/migrations/migration.%s.up.sql", dbType), upFile)
		if err != nil {
			exitGracefully(err)
		}
		err = copyFileFromTemplate(fmt.Sprintf("templates/migrations/migration.%s.down.sql", dbType), downFile)
		if err != nil {
			exitGracefully(err)
		}
	case "auth":
		err := doAuth()
		if err != nil {
			exitGracefully(err)
		}
	case "handler":
		if arg3 == "" {
			exitGracefully(errors.New("you must give the handler a name"))
		}
		// TODO: Somewhere down the road, get rid of the `/myapp` hardcoding
		fileName := fmt.Sprintf("%s/myapp/handlers/%s.go", cel.RootPath, strings.ToLower(arg3))
		if fileExists(fileName) {
			exitGracefully(errors.New(fileName + " already exists"))
		}
		data, err := templateFS.ReadFile("templates/handlers/handler.go.txt")
		if err != nil {
			exitGracefully(err)
		}
		handler := string(data)
		handler = strings.ReplaceAll(handler, "$HANDLER_NAME$", strcase.ToCamel(arg3))
		err = os.WriteFile(fileName, []byte(handler), 0644)
		if err != nil {
			exitGracefully(err)
		}
	case "model":
		if arg3 == "" {
			exitGracefully(errors.New("you must give the model a name"))
		}
		data, err := templateFS.ReadFile("templates/data/model.go.txt")
		if err != nil {
			exitGracefully(err)
		}
		model := string(data)
		pl := pluralize.NewClient()
		var modelName = arg3
		var tableName = arg3
		if pl.IsPlural(arg3) {
			modelName = pl.Singular(arg3)
			tableName = strings.ToLower(tableName)
		} else {
			tableName = strings.ToLower(pl.Plural(arg3))
		}
		// TODO: Somewhere down the road, get rid of the `/myapp` hardcoding
		fileName := fmt.Sprintf("%s/myapp/data/%s.go", cel.RootPath, strings.ToLower(modelName))
		model = strings.ReplaceAll(model, "$MODEL_NAME$", strcase.ToCamel(modelName))
		model = strings.ReplaceAll(model, "$TABLE_NAME$", tableName)
		err = copyDataToFile([]byte(model), fileName)
		if err != nil {
			exitGracefully(err)
		}

	}
	return nil
}
