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









## Getting Started with the Celeritas Module
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
