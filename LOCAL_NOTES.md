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







## Installing MariaDB, Postgres and Redis using Docker
## Adding support for Postgres to Celeritas
## Testing Models
## Starting work on the Celeritas Command Line application
## Validation
## Helper utilities
## Response utilities, Encryption and more
## Implementing a Redis cache
## An aside: Redis sessions, CSRF and debugging info
## Implementing a BadgerDB cache
## Sending Email
## Remember Me & Password Resets
## Implementing Celeritas new <myapp>
## Where to go from here
