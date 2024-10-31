# t2e

This program converts the specified time to epoch time.

## Getting Started

### Go

```bash-session
$ go install
```

### Homebrew

```bash-session
$ brew install 135yshr/tap/t2e
```

## Usage

When executed without arguments, the current time is converted to epic time.

```bash-session
$ t2e
1729987200
```

If an argument is specified, the specified time is converted to epoch time.

```bash-session
$ t2e '2024-10-31 12:00:00'
1729987200
```

If the `--add-min` option is specified, the current time is converted to epoch time by adding the specified number of minutes.

```bash-session
$ t2e --add-min 5
```

## Prerequisites

- [Go](https://go.dev/) 1.23 or later
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/)
- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)

## Development

```bash-session
$ git clone https://github.com/135yshr/t2e
$ cd ./t2e
$ pre-commit install
$ make test
```

## Contributing

This project is an open-source endeavor that thrives on your active participation. We're always on the lookout for individuals interested in contributing to the project's growth. If you have any ideas or improvements, no matter how small, they are welcome. Feel free to submit a pull request at any time. We're eagerly awaiting your collaboration!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
