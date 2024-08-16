# Let's Build a Go version of Laravel

## Introduction
### Introduction
- Let's build a Go version of Laravel
  - Laravel is the most popular framework for PHP
  - Easy database access
  - A lot of functionality right out of the box
  - Easy to install
- Disadvantages with Laravel
  - Relatively slow, because it's written in PHP
  - Massive external dependencies (which sometimes disappear)
  - Vast number of bad resources online
  - Horizontally scalable, but requires a lot of resources
  - Individual installations require a lot of system resources
- Biggest advantage with Laravel
  - Install an application with one command
- What we're going to do
  - We'll build a Go module, named Celeritas (which is Latin for Speed)
  - We'll also build a command line application
  - With one command, we'll have a Go project that includes the following Features
    - Easy access to multiple Databases (we'll include Postgres and MySQL, but have support for Mongo, CockroachDB and SQLite)
    - Complete user authentication (for web and api) with one command
    - Database migrations
    - Multiple HTML templating engines (Go templates and Jet)
    - Session support
    - Generate handlers (like Laravel controllers and models with one command)
    - Support for easy to write middleware
    - Form validation
    - CSRF protection
    - Encryption
    - Response utilities (XML, JSON, download a file)
    - Support for multiple caching back ends (Redis and BadgerDB)
    - Sending email via SMTP, SendGrid, MailGun and Sparkpost
    - Remember me and password resets for user authentication
### A bit about me
### How to ask for help
### Install Go
- Download Go [here](https://go.dev/dl/)
- WinGet
  - Install: ```winget install --id GoLang.Go```
  - Upgrade: ```winget upgrade --id GoLang.Go```
- Verify  
  ```go version```
### Installing an IDE
- Visual Studio Code
  - Install: ```winget install --id Microsoft.VisualStudioCode```
  - Upgrade: ```winget upgrade --id Microsoft.VisualStudioCode```
  - Add extensions:
    - [Go]
      - Also, press `shift+ctl` and search for `Go: Install/Update Tools`
        - Click on it, select all associated checkboxes and click OK to install them
    - [goTemplate-syntax]
- GoLand
  - Install using toolbox: ```winget install --id JetBrains.Toolbox```
  - Install directly: ```winget install --id JetBrains.GoLand```
### Install Make
- Install: ```winget install -e --id GnuWin32.Make```

## Getting Started with the Celeritas Module
### Setting up our project structure
- Create root folders for entire app
  ```shell
  md go-laravel
  cd go-laravel
  ```
- Create files and folders
  ```shell
  md celeritas
  md myapp
  ```
- Initialize
  ```shell
  cd celeritas
  go mod init github.com/johnwr-response/celeritas
  cd ..
  md myapp
  cd myapp
  go mod init myapp
  cd ..
  ni celeritas/celeritas.go -type file -Value "package celeritas`n`n"
  ni myapp/main.go -type file -Value "package main`n`n"
  cd myapp
  go get github.com/johnwr-response/celeritas
  go run .
  cd ..
  ```
### Keeping our application and package in sync with Make
  ```shell
  ni myapp/Makefile -type file
  ni Makefile -type file
  make build
  make clean
  make start
  make stop
  make restart
  make test
  ```
### Starting work on Celeritas
  ```shell
  ni celeritas/types.go -type file -Value "package celeritas`n`n"
  ni celeritas/helpers.go -type file -Value "package celeritas`n`n"
  ```
### Creating application folders
  ```shell
  ni myapp/init-celeritas.go -type file -Value "package main`n`n"
  ```
### Creating and reading the .env file
- GoDotEnv - A Go port of Ruby's dotenv library (Loads environment variables from .env files)
  [GitHub](https://github.com/joho/godotenv)
  ```shell
  cd celeritas
  go get github.com/joho/godotenv
  cd ..
  ```
  ```shell
  ni myapp/.env -type file
  ```
### Creating logs
### Setting up Celeritas configuration
### Getting a simple web server up and running
- Chi - A lightweight, idiomatic and composable router for building Go HTTP services.
  [link](https://github.com/go-chi/chi)
  ```shell
  ni celeritas/routes.go -type file -Value "package celeritas`n`n"
  cd celeritas
  go get -u github.com/go-chi/chi/v5
  go get -u github.com/go-chi/chi/v5/middleware
  cd ..
  ```

## Rendering Pages
### Implementing a page renderer in the Celeritas package (for Go templates)
  ```shell
  md celeritas/render
  ni celeritas/render/render.go -type file -Value "package render`n`n"
  ```
### A note about configuration
`Important!`
- In the next lecture, I neglected to mention that in order to get the Go rendering functional, you will need to make a change to your .env file, as follows:
  ```
  # template engine: go or jet
  # RENDERER=jet
  RENDERER=go
  ```
- Just make sure that the value for RENDERER is set to go (in lowercase).
### Trying out the Go render functionality
  ```shell
  ni myapp/routes.go -type file -Value "package main`n`n"
  ni myapp/handlers/handlers.go -type file -Value "package handlers`n`n"
  ```
### Cleaning up Celeritas
### Adding Jet support to our page rendering package
- Jet Template Engine for Go
  [link](https://github.com/CloudyKit/jet)
  ```shell
  cd celeritas
  go get github.com/CloudyKit/jet/v6
  cd ..
  ```
### Rendering a Jet Template
### Working with Jet Templates
- Jet Template Engine extension
  - Visual Studio Code
    - Search for Jet Template Engine
  - GoLand
    - There is no support for the jet template in GoLand
  ```shell
  md views/layouts
  ni views/layouts/base.jet -type file
  ni views/home.jet -type file
  ```

## Testing
### Testing the render package
- Create files and folders
  ```shell
  md celeritas/render/testdata/views
  ni celeritas/render/testdata/views/home.page.gohtml -type file -Value "Hello world`n`n"
  ni celeritas/render/testdata/views/home.jet -type file -Value "Hello jet`n`n"
  ni celeritas/render/setup_test.go -type file -Value "package render`n`n"
  ni celeritas/render/render_test.go -type file -Value "package render`n`n"
  ```
- Run tests for render package
  `go test .\celeritas\render\.`
- Run tests for viewing full coverage report render package
  `(go test '-coverprofile=coverage.out' .\celeritas\render\.) -and (go tool cover '-html=coverage.out')`
### Writing more tests for the render package
- Create files and folders
  ```shell
  ni celeritas/Makefile -type file
  ```
### Simplifying our tests using Table Tests

## Sessions
### Implementing Sessions in Celeritas
  ```shell
  ni views/jet-template.jet -type file
  ```
### Choosing and Installing a session package
- SCS - HTTP Session Management for Go. 
  [link](https://github.com/alexedwards/scs)
  ```shell
  cd celeritas
  go get github.com/alexedwards/scs/v2
  cd ..
  md celeritas/session
  ni celeritas/session/session.go -type file -Value "package session`n`n"
  ```
### Adding session middleware
  ```shell
  ni celeritas/middleware.go -type file -Value "package celeritas`n`n"
  ```
### Verifying that sessions work with myapp
  ```shell
  ni views/sessions.jet -type file
  ```
### Reading data from the session and passing it to the Jet template
### Writing tests for the session package
  ```shell
  ni celeritas/session/setup_test.go -type file -Value "package session`n`n"
  ni celeritas/session/session_test.go -type file -Value "package session`n`n"
  ```
### Checking our Coverage
  ```shell
  make test_celeritas_cover
  ```

## Installing MariaDB, Postgres and Redis using Docker
### Installing Docker
[Install Docker Desktop on Windows](https://docs.docker.com/desktop/install/windows-install/)
  ```shell
  winget install -e --id Docker.DockerDesktop
  ```
### Bringing up and tearing down a development environment using docker-compose
- Create files and folders
  ```shell
  md docker/db-data
  ni docker/docker-compose.yml -type file
  ```
- Bring up
  ```shell
  cd docker
  docker-compose up -d
  cd ..
  ```
- Stop
  ```shell
  cd docker
  docker-compose down
  cd ..
  ```
- Control with make
  - Bring up `make start_compose`
  - Stop `make stop_compose`

## Adding support for Postgres to Celeritas
### Getting started with Postgres
- pgx - PostgresSQL driver and toolkit for Go
  [GitHub](https://github.com/jackc/pgx)
  ```shell
  cd celeritas
  go get github.com/jackc/pgx/v5
  go get github.com/jackc/pgconn
  go get github.com/jackc/pgx/stdlib
  cd ..
  ni celeritas/driver.go -type file -Value "package celeritas`n`n"
  ```
### Building a Postgres connection string and connecting to the database
### Trying out our database connection
  ```postgresql
  CREATE TABLE users(id SERIAL, first_name varchar);
  INSERT INTO users values(1,'John');
  ```
### Adding ORM like functionality to our application with upper/db
- Create files and folders
  ```shell
  ni myapp/data/models.go -type file -Value "package data`n`n"
  ```
- upper/db - A productive data access layer for Go
  [GitHub](https://upper.io/v4/)
  ```shell
  cd myapp
  go get -u github.com/upper/db/v4/adapter/postgresql
  go get -u github.com/upper/db/v4/adapter/mysql
  cd ..
  ```
### Creating a real users table and a user model
- Create files and folders
  ```shell
  ni users.sql -type file
  ni myapp/data/user.go -type file -Value "package data`n`n"
  ni myapp/data/token.go -type file -Value "package data`n`n"
  ```
### Additional database functions for the User type
### Finishing up the database functions for the User model
### Inserting a user
### Testing other database functions on the User model
### Creating a login page and handler
- Create files and folders
  ```shell
  ni myapp/handlers/auth-handlers.go -type file -Value "package handlers`n`n"
  ni views/login.jet -type file -Value "{{extends `u{0022}./layouts/base.jet`u{0022}}}`n`n{{block browserTitle()}}`n`n{{end}}`n`n{{block css()}}`n`n{{end}}`n`n{{block pageContent()}}`n`n{{end}}`n`n{{block js()}}`n`n{{end}}"
  ```
### Creating the post handler for logging in
### Adding functions to the Tokens model
- Create files and folders
  ```shell
  ri user.sql
  ni auth_tables.sql -type file
  ```

## Testing Models
### Writing tests for models.go
  ```shell
  ni myapp/data/setup_test.go -type file -Value "package data`n`n"
  ni myapp/data/models_test.go -type file -Value "package data`n`n"
  ```
- Sql driver mock for Golang - Sql mock driver for golang to test database interactions
  [GitHub](https://github.com/DATA-DOG/go-sqlmock) 
  ```shell
  cd myapp
  go get github.com/DATA-DOG/go-sqlmock
  cd ..
  ```
### Getting started with our integration tests
- DockerTest - Write better integration tests! DockerTest helps you boot up ephemeral docker images for your Go tests with minimal work.
  [GitHub](https://github.com/ory/dockertest)
  ```shell
  cd myapp
  go get github.com/ory/dockertest/v3
  go get github.com/ory/dockertest/v3/docker
  cd ..
  ```
- Create files and folders
  ```shell
  ni myapp/data/integration_test.go -type file -Value "package data`n`n"
  ```
### Creating tables in our test docker image, and running some tests
  ```shell
  cd myapp/data
  go test . --tags integration --count=1
  cd ../..
  ```
### Continuing to write integration tests
### Finishing up our integration tests
- Run tests and check coverage: `make test_myapp_integration_cover`
- Run tests and display coverage: `make test_myapp_integration_coverage`
- NOTE! Had to change the way time is handled to ensure UTC is always used
  - Replaced every usage of `time.Now()` with `time.Now().UTC()`
  - Also changed timestamp fields in database to default to UTC by replacing `now()` with `(now() at time zone 'utc')`
### Cleaning up our tests

## Starting work on the Celeritas Command Line application
### Setting up a simple CLI package in Celeritas
- Color - Color package for Go (golang)
  [GitHub](https://github.com/fatih/color)
  ```shell
  cd celeritas
  go get github.com/fatih/color
  cd ..
  ```
- Create files and folders
  ```shell
  md celeritas/cmd/cli
  ni celeritas/cmd/cli/main.go -type file -Value "package main`n`n"
  ```
- To build cli: `make build_cli`
- To run cli: `myapp/celeritas.exe`
### Adding support for migrations to the Celeritas package
- Migrate - Database migrations. CLI and Golang library.
  [GitHub](https://github.com/golang-migrate/migrate)
  ```shell
  cd celeritas
  go get github.com/golang-migrate/migrate/v4
  go get github.com/golang-migrate/migrate/v4/database/mysql
  go get github.com/golang-migrate/migrate/v4/database/postgres
  go get github.com/golang-migrate/migrate/v4/source/file
  cd ..
  ```
- Create files and folders
  ```shell
  ni celeritas/migrations.go -type file -Value "package celeritas`n`n"
  ```
### Starting work on "make migration" in our CLI application
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/make.go -type file -Value "package main`n`n"
  ```
- Test (Note the celeritas struct is not yet populated)
  ```shell
  make bild_cli
  myapp/celeritas.exe make migration test
  ```
### Using templates in our CLI
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/copy-files.go -type file -Value "package main`n`n"
  md celeritas/cmd/cli/templates/migrations
  ni celeritas/cmd/cli/templates/migrations/migration.postgres.up.sql -type file
  ni celeritas/cmd/cli/templates/migrations/migration.postgres.down.sql -type file
  ```
### Trying out our make migration functionality
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/helpers.go -type file -Value "package main`n`n"
  ```
- This creates up/down migrations in the migration folder under where the command is run. (`go-laravel/migrations/*`)
  ```shell
  make build_cli
  myapp/celeritas.exe make migration some_test_name
  ```
### Running migrations
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/migrate.go -type file -Value "package main`n`n"
  ```
### Trying out our "make migrate" commands with the Celeritas CLI
- Fixing error when using windows: `Error: parse "file://...myapp/migrations": invalid port ":\\...\\myapp" after host`
  ```shell
  make build_cli
  myapp/celeritas.exe make migration
  myapp/celeritas.exe make migration some_test_name
  myapp/celeritas.exe migrate
  myapp/celeritas.exe migrate down
  myapp/celeritas.exe help
  myapp/celeritas.exe version
  ```
### Getting started with Implementing "make auth" functionality
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/auth.go -type file -Value "package main`n`n"
  ni celeritas/cmd/cli/templates/migrations/auth_tables.postgres.sql -type file
  ```
### Trying out the make auth functionality
- Drop all tables in database
- Delete all files in migration folder
- Run `myapp/celeritas.exe make auth`
- Run `myapp/celeritas.exe make migrate down`
### Continuing with the "make auth" functionality in our command line program
- Create files and folders
  ```shell
  md celeritas/cmd/cli/templates/data
  ni celeritas/cmd/cli/templates/data/user.go.txt -type file -Value "package data`n`n"
  ni celeritas/cmd/cli/templates/data/token.go.txt -type file -Value "package data`n`n"
  ```
- Delete `myapp/data/user.go`, it will be recreated
- Delete `myapp/data/token.go`, it will be recreated
- Delete all migrations in migration folder
- Drop all tables in database
- Run `myapp/celeritas.exe make auth`
- Run `myapp/celeritas.exe migrate down`
- Run `myapp/celeritas.exe migrate up`
### Creating simple auth middleware, and adding it to the "make auth" command
- Create files and folders
  ```shell
  ni myapp/middleware/auth.go -type file -Value "package middleware`n`n"
  ni myapp/middleware/middleware.go -type file -Value "package middleware`n`n"
  ni myapp/middleware/auth-token.go -type file -Value "package middleware`n`n"
  ni celeritas/response-utils.go -type file -Value "package celeritas`n`n"
  ```
### Installing our auth middleware with the celeritas command line utility
- Create files and folders
  ```shell
  md celeritas/cmd/cli/templates/middleware
  ni celeritas/cmd/cli/templates/middleware/auth.go.txt -type file -Value "package middleware`n`n"
  ni celeritas/cmd/cli/templates/middleware/auth-token.go.txt -type file -Value "package middleware`n`n"
  ```
### Trying out our improved make auth functionality
- Rebuild cli `make build_cli`
- Delete `user.go` and `token.go` in the `myapp/data` folder
- Delete `auth.go` and `auth-token.go` files in the `myapp/middleware` folder
- Delete all files in migration folder
- Drop all tables in database
  `drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens; drop table if exists schema_migrations;`
- Run `myapp/celeritas.exe make auth`
### Implementing "make handler" functionality
- StrCase - A golang package for converting to snake_case or CamelCase
  [GitHub](https://github.com/iancoleman/strcase)
  ```shell
  cd celeritas
  go get -u github.com/iancoleman/strcase
  cd ..
  ```
- Create files and folders
  ```shell
  md celeritas/cmd/cli/templates/handlers
  ni celeritas/cmd/cli/templates/handlers/handler.go.txt -type file -Value "package handlers`n`n"
  ```
- Try it out
  `make build_cli`
  `myapp/celeritas.exe make handler testHandler`
### Implementing "make model" functionality
- go-pluralize - Pluralize and singularize any word (golang adaptation of https://www.npmjs.com/package/pluralize)
  [GitHub](https://github.com/gertd/go-pluralize)
  ```shell
  cd celeritas
  go get -u github.com/gertd/go-pluralize
  cd ..
  ```
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/templates/data/model.go.txt -type file -Value "package data`n`n"
  ```
- Try it out
  `make build_cli`
  `myapp/celeritas.exe make model test`
### Adding database stores to our sessions package
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/session.go -type file -Value "package main`n`n"
  ni celeritas/cmd/cli/templates/migrations/postgres_session.sql -type file
  ni celeritas/cmd/cli/templates/migrations/mysql_session.sql -type file
  ```
### Adding support for database session store to the celeritas project
  ```shell
  cd celeritas
  go get github.com/alexedwards/scs/pgxstore
  go get github.com/alexedwards/scs/mysqlstore
  cd ..
  ```
- To add a user in the database, use: [http://localhost:4000/create-user](http://localhost:4000/create-user)
### Supporting MySQL/MariaDB with"make auth"
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/templates/migrations/auth_tables.mysql.sql -type file
  ```

## Validation
### Creating a validation package
- goValidator - Go Package of validators and sanitizers for strings, numerics, slices and structs
  [GitHub](https://github.com/asaskevich/govalidator)
  ```shell
  cd celeritas
  go get github.com/asaskevich/govalidator
  cd ..
  ```
- Create files and folders
  ```shell
  ni celeritas/validator.go -type file -Value "package celeritas`n`n"
  ```
### Trying out our validation
- Try it out
  [Link](http://localhost:4000/update-user/1)
### Adding validation to models
### Trying out our model validation
- Try it out
  [Link](http://localhost:4000/update-user/1)
### Building a simple form and performing validation on it
- Create files and folders
  ```shell
  ni views/form.jet -type file -Value "{{extends `u{0022}./layouts/base.jet`u{0022}}}`n`n{{block browserTitle()}}`n`n{{end}}`n`n{{block css()}}`n`n{{end}}`n`n{{block pageContent()}}`n`n{{end}}`n`n{{block js()}}`n`n{{end}}"
  ni myapp/handlers/form-val-handlers.go -type file -Value "package handlers`n`n"
  ```
### Building our PostForm handler with validation

## Helper utilities
### Helper functions for the routes file
- Create files and folders
  ```shell
  ni myapp/convenience.go -type file -Value "package main`n`n"
  ```
### Helper functions for handlers
- Create files and folders
  ```shell
  ni myapp/handlers/convenience.go -type file -Value "package handlers`n`n"
  ```

## Response utilities, Encryption and more
### JSON, XML, and other response types
### Creating handlers for our response types
### Creating the routes and links for our response types
### Encryption/Decryption
### Generating and getting our encryption key
- Try it out
  ```shell
  make build_cli
  myapp/celeritas.exe make key
  ```
### Trying out our encryption functionality
- Try it out
  [Link](http://localhost:4000/crypto)

## Implementing a Redis cache
### Installing the necessary package and getting started
- Redigo - Go client for Redis
  [GitHub](https://github.com/gomodule/redigo)
  ```shell
  cd celeritas
  go get github.com/gomodule/redigo/redis
  cd ..
  ```
- Create files and folders
  ```shell
  md celeritas/cache
  ni celeritas/cache/cache.go -type file -Value "package cache`n`n"
  ```
### Connecting to Redis
### Completing the rest of the cache functions
### Testing the cache package
- MiniRedis - Pure Go Redis server for Go unittests
  [GitHub](https://github.com/alicebob/miniredis)
  ```shell
  cd celeritas
  go get github.com/alicebob/miniredis/v2
  cd ..
  ```
- Create files and folders
  ```shell
  ni celeritas/cache/setup_test.go -type file -Value "package cache`n`n"
  ni celeritas/cache/cache_test.go -type file -Value "package cache`n`n"
  ```
- Try it out
  ```shell
  go test -cover -v ./celeritas/cache/...
  ```
### Trying out the cache in myapp
- Create files and folders
  ```shell
  ni views/cache.jet -type file -Value "{{extends `u{0022}./layouts/base.jet`u{0022}}}`n`n{{block browserTitle()}}`n`n{{end}}`n`n{{block css()}}`n`n{{end}}`n`n{{block pageContent()}}`n`n{{end}}`n`n{{block js()}}`n`n{{end}}"
  ni myapp/handlers/cache-handlers.go -type file -Value "package handlers`n`n"
  ```
### Finishing up our cache page in myapp

## An aside: Redis sessions, CSRF and debugging info
### Adding a Redis store to our sessions package
  ```shell
  cd celeritas
  go get github.com/alexedwards/scs/redisstore
  cd ..
  ```
### CSRF Protection
- NoSurf - CSRF protection middleware for Go
  [GitHub](https://github.com/justinas/nosurf)
  ```shell
  cd celeritas
  go get github.com/justinas/nosurf
  cd ..
  cd myapp
  go get github.com/justinas/nosurf
  cd ..
  ```
### Speeding up templates
- Create files and folders
  ```shell
  ni celeritas/utils.go -type file -Value "package celeritas`n`n"
  ```

## Implementing a BadgerDB cache
### Installing the necessary package and implementing necessary functions
- BadgerDB - Fast key-value DB in Go.
  [GitHub](https://github.com/dgraph-io/badger)
  ```shell
  cd celeritas
  go get github.com/dgraph-io/badger/v4
  cd ..
  ```
- Create files and folders
  ```shell
  ni celeritas/cache/badger_cache.go -type file -Value "package cache`n`n"
  ```
### Updating setup_test.go to create a Badger database for our tests
- Create files and folders
  ```shell
  md celeritas/cache/testdata/tmp
  ```
### Writing and running tests for our Badger cache
  ```shell
  ni celeritas/cache/badger_cache_test.go -type file -Value "package cache`n`n"
  ```
### Connecting to Badger
- cron - a cron library for go.
  [GitHub](https://github.com/robfig/cron)
  ```shell
  cd celeritas
  go get github.com/robfig/cron/v3
  cd ..
  ```
### Trying out the Badger cache

## Sending Email
### Getting started sending email using SMTP
- Create files and folders
  ```shell
  md celeritas/mailer
  ni celeritas/mailer/mail.go -type file -Value "package mailer`n`n"
  ```
### Important Note
- In the next lecture there is a call to go get a package called `go mail`
- The author of that package has changed the way it works for the better.
- This will install a version that matches the one used in this lecture.
  - Replace `go get github.com/ainsleyclark/go-mail`
  - With `go get github.com/ainsleyclark/go-mail@v1.0.3`
- Either use that and update the code for newer version later or update code right away
### Adding the necessary packages, and completing sending email via SMTP
- Go Simple Mail - Golang package for send email. Support keep alive connection, TLS and SSL. Easy for bulk SMTP.
  [GitHub](https://github.com/xhit/go-simple-mail)
- go-premailer - Inline styling for html mail in golang
  [GitHub](https://github.com/vanng822/go-premailer)
  ```shell
  cd celeritas
  go get github.com/xhit/go-simple-mail/v2
  go get github.com/vanng822/go-premailer/premailer
  cd ..
  ```
### Sending email using Mailgun, SparkPost and more
- Go Mail - A cross-platform mail driver for GoLang. Featuring Mailgun, Postal, Postmark, SendGrid, SparkPost & SMTP.
  [GitHub](https://github.com/ainsleyclark/go-mail)
  ```shell
  cd celeritas
  go get -u github.com/ainsleyclark/go-mail
  cd ..
  ```
### Connecting Celeritas to our mailer package
### Trying out or mailer package
- MailTrap Email Delivery Platform is the toolset to test, send, and control your emails in one place.
  [Link](https://mailtrap.io/)
- Sample Credentials
  - Host: `sandbox.smtp.mailtrap.io`
  - Port: `25, 465, 587 or 2525`
  - Username: `25853d08526311`
  - Password: `399982fbb4cbe9`
- MailHog is an email-testing tool with a fake SMTP server underneath, installed here as a docker container
  [Link](https://github.com/mailhog/MailHog) 
- Starting and stopping mailHog
  ````
  cd docker
  docker-compose up mailHog -d
  docker-compose down mailHog 
  ````
- Usage:
  - [Sendmail](`localhost:1025`)
  - [Web interface](http://localhost:8025/ "MailHog web interface")
### Sending mail using an API
### Adding "make mail" to the CLI
- Create files and folders
  ```shell
  md celeritas/cmd/cli/templates/mailer
  ni celeritas/cmd/cli/templates/mailer/mail.html.gohtml -type file
  ni celeritas/cmd/cli/templates/mailer/mail.plain.gohtml -type file
  ```
- Try it out
  ```shell
  make build_cli
  myapp/celeritas.exe make mail test2
  ```
### Testing mail
- Adding DockerTest in Celeritas as well
  ```shell
  cd celeritas
  go get github.com/ory/dockertest/v3
  go get github.com/ory/dockertest/v3/docker
  cd ..
  ```
- Create files and folders
  ```shell
  md celeritas/mailer/testdata/mail
  ni celeritas/mailer/testdata/mail/test.html.gohtml -type file
  ni celeritas/mailer/testdata/mail/test.plain.gohtml -type file
  ni celeritas/mailer/setup_test.go -type file -Value "package mailer`n`n"
  ni celeritas/mailer/mail_test.go -type file -Value "package mailer`n`n"
  ```
- Run simple test coverage on mailer package
  ```shell
  (go test '-coverprofile=coverage.out' ./celeritas/mailer/.) -and (go tool cover '-html=coverage.out')
  ```

## Remember Me & Password Resets
### Setting up models and middleware for "remember me" functionality
- Create files and folders
  ```shell
  ni myapp/data/remember_token.go -type file -Value "package data`n`n"
  ni myapp/middleware/remember.go -type file -Value "package middleware`n`n"
  ```
### Updating the auth handlers for remember me functionality
### Trying out the remember me functionality
### Password resets
### Handling a password reset request
- go-alone - A simple to use, high-performance, Go (golang) MAC signer.
  [GitHub](https://github.com/bwmarrin/go-alone)
  ```shell
  cd celeritas
  go get github.com/bwmarrin/go-alone
  cd ..
  ```
- Create files and folders
  ```shell
  md celeritas/urlSigner
  ni celeritas/urlSigner/signer.go -type file -Value "package urlSigner`n`n"
  ```
### Sending a password reset link via email
- Create email template
  ```shell
  myapp/celeritas.exe make mail password-reset
  ```
### Validating our signed link, and displaying the password reset form
### Resetting the user's password
### Updating the "make auth" functionality in the Celeritas CLI
- Create files and folders
  ```shell
  md celeritas/cmd/cli/templates/views
  ```

## Implementing Celeritas new <myapp>
### Starting work on "celeritas new" in the CLI
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/new.go -type file -Value "package main`n`n"
  ```
### Sanitizing the project name
- Try it out
  ```shell
  make build_cli
  myapp/celeritas.exe new myapp
  myapp/celeritas.exe new github.com/something/myotherapp
  ```
### Cloning a (currently non-existent) repository right in Go
- go-git - A highly extensible Git implementation in pure Go.
  [GitHub](https://github.com/go-git/go-git)
  ```shell
  cd celeritas
  go get github.com/go-git/go-git/v5
  cd ..
  ```
### Creating a skeleton application
- Copied necessary files from myapp and views, and removed unnecessary test code into a new repository to be pushed
### Pushing our skeleton application to GitHub
- Create the new git repository
  ```shell
  git init
  git add .
  git commit -m "Initial entry"
  ```
- Go to GitHub and create a new public repository without initializing it
- Make sure you have the GitHub cli installed, and logged in
  ```powershell
  winget install --id GitHub.cli
  gh auth login
  ```
- Set the remote and push the new repository
  ```shell
  git remote add origin https://github.com/john-wraa/celeritas-app.git
  git branch -M main
  git push -u origin main
  ```
### Trying out the code that clones a remote GitHub repository
- Try it out
  ```shell
  make dist_cli
  cd dist
  ./celeritas.exe new test
  ```
### Removing the .git directory and creating a .env file
- Create files and folders
  ```shell
  ni celeritas/cmd/cli/templates/env.txt -type file
  ```
- Remove files from last test and try it out again
  ```shell
  rm ./dist/test -r -force
  make dist_cli
  cd dist
  ./celeritas.exe new test
  ```
### Creating the correct Makefile




## Where to go from here
