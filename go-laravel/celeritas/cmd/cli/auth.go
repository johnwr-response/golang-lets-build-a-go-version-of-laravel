package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"time"
)

func doAuth() error {
	//create migrations
	dbType := cel.DB.DatabaseType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UTC().UnixMicro())
	upFile := cel.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := cel.RootPath + "/migrations/" + fileName + ".down.sql"
	log.Println(dbType, upFile, downFile)
	err := copyFileFromTemplate(fmt.Sprintf("templates/migrations/auth_tables.%s.sql", dbType), upFile)
	if err != nil {
		exitGracefully(err)
	}
	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// copy files over
	err = copyFileFromTemplate("templates/data/user.go.txt", cel.RootPath+"/myapp/data/user.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/data/token.go.txt", cel.RootPath+"/myapp/data/token.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/middleware/auth.go.txt", cel.RootPath+"/myapp/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/middleware/auth-token.go.txt", cel.RootPath+"/myapp/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow("  - users, tokens and remember_tokens migrations created and executed successfully")
	color.Yellow("  - users and token models created")
	color.Yellow("  - autgh middleware created")
	color.Yellow("")
	color.Yellow("Don't forget to add user and token models in data/models.com, and to add appropriate middleware to your routes!")

	return nil
}
