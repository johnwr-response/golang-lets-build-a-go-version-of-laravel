SHELL=cmd
BINARY_NAME=celeritasApp.exe
CLI_NAME=celeritas.exe

## build: builds all binaries
build: clean build_myapp
	@echo All binaries built!

## build_myapp: builds the application
build_myapp:
	@go build -o dist/${BINARY_NAME} ./myapp
	@echo Celeritas built!

## clean: cleans all binaries and runs go clean
clean:
	@echo Cleaning...
	@echo y | DEL /S dist
	@go clean
	@echo Cleaned and deleted binaries

run: build
	@echo Staring Celeritas...
	set dummy=dummy&& start /B .\dist\${BINARY_NAME}
	@echo Celeritas started!

run_foreground: build
	@echo Staring Celeritas...
	@.\dist\${BINARY_NAME} &
	@echo Celeritas started in foreground!


start: run

start_fg: run_foreground

stop:
	@taskkill /IM ${BINARY_NAME} /F  /FI "MemUsage gt 2"
	@echo Stopped Celeritas

restart: stop start

restart_fg: stop start_fg

start_compose:
	docker-compose --project-directory docker up -d

stop_compose:
	docker-compose --project-directory docker down

## build_cli: builds the command line tool celeritas and copies it to myapp
build_cli:
	@go build -o ./myapp/${CLI_NAME} ./celeritas/cmd/cli

## build_cli: builds the command line tool dist directory
dist_cli:
	@go build -o ./dist/${CLI_NAME} ./celeritas/cmd/cli

test: test_myapp test_celeritas
	@echo All tests run!

test_cover: test_myapp_cover test_celeritas_cover
	@echo All tests run!

test_coverage: test_myapp_coverage test_celeritas_coverage
	@echo All tests run!

test_myapp:
	@echo Testing myapp...
	@go test -v ./myapp/...
	@echo Done!

test_myapp_integration:
	@echo Testing myapp...
	@go test -v ./myapp/... --tags integration --count=1
	@echo Done!

test_myapp_cover:
	@echo Testing myapp...
	@go test '-coverprofile=coverage.out' ./myapp/...
	@go tool cover '-html=coverage.out'
	@echo Done!

test_myapp_integration_cover:
	@echo Testing myapp...
	@go test '-coverprofile=coverage.out' ./myapp/... --tags integration --count=1
	@go tool cover '-html=coverage.out'
	@echo Done!

test_myapp_integration_coverage:
	@echo Testing myapp...
	@go test -cover ./myapp/... --tags integration --count=1
	@echo Done!

test_myapp_coverage:
	@echo Testing myapp...
	@go test -cover ./myapp/...
	@echo Done!

test_celeritas:
	@echo Testing celeritas...
	@go test -v ./celeritas/...
	@echo Done!

test_celeritas_cover:
	@echo Testing celeritas...
	@go test '-coverprofile=coverage.out' ./celeritas/...
	@go tool cover '-html=coverage.out'
	@echo Done!

test_celeritas_coverage:
	@echo Testing celeritas...
	@go test -cover ./celeritas/...
	@echo Done!
