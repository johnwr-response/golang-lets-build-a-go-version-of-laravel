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
  ```
  cd celeritas
  go get github.com/joho/godotenv
  cd ..
  ```
  ```shell
  ni myapp/.env -type file
  ```
### Creating logs




## Rendering Pages
## Testing
## Sessions
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
