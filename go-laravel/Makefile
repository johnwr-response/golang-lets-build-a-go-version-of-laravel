SHELL=cmd
BINARY_NAME=celeritasApp.exe

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

test:
	@echo Testing...
	@go test ./myapp/...
	@echo Done!

run: build
	@echo Staring Celeritas...
	set dummy=dummy&& start /B .\dist\${BINARY_NAME}
## @.\dist\${BINARY_NAME} & ## this one wilL not run in the background and thus prevail the logs
	@echo Celeritas started!

start: run

stop:
	@taskkill /IM ${BINARY_NAME} /F  /FI "MemUsage gt 2"
	@echo Stopped Celeritas

restart: stop start
