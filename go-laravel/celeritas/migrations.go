package celeritas

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"path/filepath"
)

func (c *Celeritas) MigrateUp(dsn string) error {
	rootPath := filepath.ToSlash(c.RootPath)
	log.Println("MigrateUp DSN:", dsn)
	m, err := migrate.New("file://"+rootPath+"/migrations", dsn)
	if err != nil {
		log.Println("Error migrate.New: ", err.Error())
		return err
	}
	defer func(m *migrate.Migrate) {
		_, _ = m.Close()
	}(m)

	if err := m.Up(); err != nil {
		log.Println("Migration Up err:", err)
		return err
	}
	return nil
}

func (c *Celeritas) MigrateDownAll(dsn string) error {
	rootPath := filepath.ToSlash(c.RootPath)
	m, err := migrate.New("file://"+rootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer func(m *migrate.Migrate) {
		_, _ = m.Close()
	}(m)

	if err := m.Down(); err != nil {
		log.Println("Migration Down err:", err)
		return err
	}
	return nil
}

func (c *Celeritas) Steps(n int, dsn string) error {
	rootPath := filepath.ToSlash(c.RootPath)
	m, err := migrate.New("file://"+rootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer func(m *migrate.Migrate) {
		_, _ = m.Close()
	}(m)

	if err := m.Steps(n); err != nil {
		log.Println("Migration Steps err:", err)
		return err
	}
	return nil
}

func (c *Celeritas) MigrateForce(dsn string) error {
	rootPath := filepath.ToSlash(c.RootPath)
	m, err := migrate.New("file://"+rootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer func(m *migrate.Migrate) {
		_, _ = m.Close()
	}(m)

	if err := m.Force(-1); err != nil {
		log.Println("Migration Force err:", err)
		return err
	}
	return nil
}
