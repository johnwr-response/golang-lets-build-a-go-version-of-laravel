# golang-lets-build-a-go-version-of-laravel
Build a reusable Go module and command line application that makes building a web application simple, fast and secure.

## What you'll learn
- How to build a reusable package in Go
- How to integrate multiple database types into a Go application
- How to build a complete User authentication system (web and API) in Go
- How to build a caching system using Go and Redis
- How to implement a caching system using Go and BadgerDB
- How to build a command line tool that writes code for you
- How to automate database migrations in Go
- How to integrate multiple template rendering engines into a single application
- How to write unit tests in Go
- How to write integration tests in Go

## Requirements
- A basic understanding of Go
- A basic understanding of SQL databases
- A Windows, Mac, or Linux computer
- An internet connection
- Docker
- Visual Studio Code (or the IDE of your choice)

## Course content
- Introduction
- Getting Started with the `Celeritas` Module
- Rendering Pages
- Testing
- Sessions
- Installing MariaDB, Postgres and Redis using Docker
- Adding support for Postgres to `Celeritas`
- Testing Models
- Starting work on the `Celeritas` Command Line application
- Validation
- Helper utilities
- Response utilities, Encryption and more
- Implementing a Redis cache
- An aside: Redis sessions, CSRF and debugging info
- Implementing a BadgerDB cache
- Sending Email
- Remember Me & Password Resets
- Implementing `Celeritas` new <myapp>
- Where to go from here

## Description
Laravel is one of the most popular web application frameworks in the PHP world, and with good reason. It's easy to use, well-designed, and lets developers work on their applications without worrying about re-inventing the wheel every time they start a project. Go, often referred to as Golang, is one of the most popular programming languages in the world, and has been used to create systems at Netflix, American Express, and many other well known companies. It's extremely fast, type safe, and designed from the ground up to be used on the web.

This course is all about taking some of the most useful features found in Laravel, and implement similar functionality in Go. Since Go is compiled and type safe, web applications written in this language are typically much, much faster, and far less error-prone that a similar application written in Laravel/PHP.

The key features we'll work on in this course include:
- Implementing an Object Relation Mapper (ORM) that is database agnostic, and offers much of the functionality found in Laravel Eloquent ORM.
- A fully functional Database Migration system
- Building a fully featured user authentication system that can be installed with a single command, which includes:
    - A password reset system
    - Session based authentication (for web based applications)
    - Token based authentication (for APIs and systems built with front ends like React and Vue)
- A fully featured templating system (using both Go templates and Jet templates)
- A complete caching system that supports Redis and Badger
- Easy session management, with cookie, database (MySQL and Postgres), Redis stores
- Simple response types for HTML, XML, JSON, and file downloads
- Form validation
- JSON validation
- A complete mailing system which supports SMTP servers, and third party APIs including MailGun, SparkPost, and SendGrid
- A command line application which allows for easy generation of emails, handlers, database models
- Finally, the command line application will allow us to create a ready-to-go web application by tying a single command: `celeritas new <myproject>`

## Who this course is for:
- Developers with some knowledge of Go, and Laravel developers who want faster, more scalable applications

## Project app: go-laravel
- Folder is [go-laravel](go-laravel)
- Built in Go version 1.22.5
- Uses GoDotEnv [link](https://github.com/joho/godotenv)
- Uses the chi router [link](https://github.com/go-chi/chi)
- Uses the jet template engine [link](https://github.com/CloudyKit/jet)
  - Indirectly uses FastPrinter from same vendor to support write values in io.Writer without allocation
    [FastPrinter](https://github.com/CloudyKit/fastprinter)
