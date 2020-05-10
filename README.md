# GO Hello

This a really basic Hello World server in GOLang. You can start it up using the `run.sh` script; this will start a server on localhost:8080 and you can receive a hello world json response over localhost:8080/helloworld

Build artefacts are placed in the bin/ directory, this directory will be created after the first build operation.

## Structure

- The main command is under `cmd/hello`
- The internal code is under the `pkg` directory