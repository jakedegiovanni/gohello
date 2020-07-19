# Go Hello

This a basic Hello World server written in GO that was created as a way to learn some Go as well as explore various ways to structure a vanilla server in the language. The only endpoint currently registered is found under `/world`, hitting it will return a "Hello World" message. You can customise the message it returns using `/world/SOME_MESSAGE`.

The architecture of the program is as follows:
```
                                                cmd <- dependencies injected
                                                 |
                                                \ /
 -----------------------------------------------------------------------------------------------
 |   Server                                                                                    |
 |                                                                                             |
 |     --------------------------------------------------------------------------              |
 |     |\   Handler...N                                                          \             |
 |     | \________________________________________________________________________\            |
 |     | |                                                                        |            |
 |     | |                                                                        |            |
 |     | |                                                                        |            |
 |     | |                                WorldHandler                            |            |
 |     | |                                                                        |            |
 |     \ |                                                                        |            |
 |      \|________________________________________________________________________|            |
 |                                                                                             |
 -----------------------------------------------------------------------------------------------
```

## Structure

- The entry point of the app is under `cmd/hello`. This directory will parse command line arguments and hand over to the business logic.
  - Dependencies need for runtime are created here and passed through to the business logic entry point
- The business logic is under the `internal/app` directory; this is responsible for starting the server and listening and responding to network requests.

## Building

- You can execute the `build.sh` script to build the application.
- Build artifacts are placed in the bin/ directory, this directory will be created if it doesn't already exist.

## Running

- You can start the app up using the `run.sh` script, this will build the application and start the server on the default port. You can access the server on `localhost:8080` and you can receive a hello world JSON response by visiting `localhost:8080/world`
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
