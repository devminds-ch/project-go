# Go Training Project by [devminds GmbH](https://devminds.ch)

This Go project is used for DevOps CI/CD and Git trainings.

The project contains an application providing a CLI to calculate the sum of two numbers:

```bash
Go training project by devminds GmbH
This Go project is used for DevOps CI/CD and Git trainings.
The project contains an application providing a CLI to calculate the sum of two numbers.

Usage:
  go_training_project [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  sum         Show the sum of two numbers on the console
  version     Print the version number of the Go Training Project

Flags:
  -h, --help   help for go_training_project

Use "go_training_project [command] --help" for more information about a command.
```


## Directory structure

```
├── calculate   Go package containing the calculations including corresponding tests
├── cmd         Go package containing the application CLI
└── docs        Project documentation
```


## Build and test instructions

### Build Go application

Build the Go application:

```bash
go build -v ./...
```

### Run Go static analysis

Check if source code is formatted properly:

```bash
gofmt -d .
```

Run [golangci-lint](https://golangci-lint.run/):

```bash
golangci-lint run
```

Note: the `golangci-lint` executable is available within the `devcontainer`.


### Run Go tests

Run the Go tests:

```bash
go test ./...
```
