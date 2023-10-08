# HTTP Mock NetLogger

<p align="left">
    <a href="https://github.com/emmanuelgautier/httpmock-netlogger/actions/workflows/ci.yml"><img src="https://github.com/emmanuelgautier/httpmock-netlogger/actions/workflows/ci.yml/badge.svg?branch=main&event=push" alt="CI Tasks for Go Cli template"></a>
    <a href="https://goreportcard.com/report/github.com/emmanuelgautier/httpmock-netlogger"><img src="https://goreportcard.com/badge/github.com/emmanuelgautier/httpmock-netlogger" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/emmanuelgautier/httpmock-netlogger"><img src="https://pkg.go.dev/badge/www.github.com/emmanuelgautier/httpmock-netlogger" alt="PkgGoDev"></a>
</p>

A simple HTTP Server mock which log every network connections events.

## Usage

1. Clone or download this repository:

```bash
git clone https://github.com/emmanuelgautier/httpmock-netlogger.git
cd httpmock-netlogger
```

2. Build the CLI tool:

```bash
go build -o httpmock-netlogger
```

3. Run the CLI tool with the `--help` flag to see the available commands:

```bash
./httpmock-netlogger --help
```

You should see output similar to the following:

```
A simple Go CLI template.

Usage:
  httpmock-netlogger [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  serve       Create a HTTP Server
  help        Help about any command

Flags:
  -h, --help   help for httpmock-netlogger

Use "httpmock-netlogger [command] --help" for more information about a command.
```

4. Run the `serve` subcommand:

```bash
./httpmock-netlogger serve -p 8080
```

## License

This HTTP server mock net logger is open-source and available under the MIT License. Contributions and improvements are welcome!

## Author

[Emmanuel Gautier](https://www.emmanuelgautier.com/)
