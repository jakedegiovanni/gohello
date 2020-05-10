# GO Hello

This a really basic Hello World server written in GO.

Build artifacts are placed in the bin/ directory, this directory will be created if it doesn't already exist.

## Structure

- The entry point of the app is under `cmd/hello`. This directory will parse command line arguments and hand over to the business logic.
- The business logic is under the `pkg` directory; this is responsible for starting the server and listening and responding to network requests.

## Building

- You can execute the `build.sh` script to build the application.

## Running

- You can start the app up using the `run.sh` script, this will build the application and start the server on the default port. You can access the server on `localhost:8080` and you can receive a hello world JSON response by visiting `localhost:8080/helloworld`
- Alternatively, you can build the app manually and execute `bin/hello.exe` if you don't want to use the run script.
- `bin/hello.exe --help` also provides usage instructions

## Customizing the port

The default port of the app is 8080, if you want to change the port the app starts up on you can use one of the below commands:

- `./bin/hello.exe -port SOME_PORT`
- `./run.sh SOME_PORT`

## Cleanup

- Run `cleanup.sh` to clean your working directory.

## Testing

- Run `./test.sh` to run the tests inside the `cmd/` and `pkg/` directories.
