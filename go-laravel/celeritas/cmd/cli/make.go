package main

import (
	"errors"
	"fmt"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
	case "migrate":
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
	}
	return nil
}
