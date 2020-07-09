# markov-bot

[![Build Status](https://travis-ci.com/cjonesy/markov-bot.svg?branch=master)](https://travis-ci.com/cjonesy/markov-bot) [![Go Report Card](https://goreportcard.com/badge/github.com/cjonesy/markov-bot)](https://goreportcard.com/report/github.com/cjonesy/markov-bot) [![Release](https://img.shields.io/github/release/cjonesy/markov-bot.svg)](https://github.com/cjonesy/markov-bot/releases/latest)

A Slack bot that uses markov chains to respond.

# Usage

### Basic Usage
```
Usage:
  markov-bot [command]

Available Commands:
  help        Help about any command
  start       Start the bot
  version     Print the version number

Flags:
  -h, --help   help for markov-bot

Use "markov-bot [command] --help" for more information about a command.
```

### To start the bot:
```shell script
markov-bot start -t <SLACK_TOKEN> -c /path/to/corpus.txt
```

# Development

### Compiling
```shell script
make build
```

### Running Tests
To run all the standard tests:
```shell script
make test
```

### Releasing
This project is using [goreleaser](https://goreleaser.com). GitHub release creation is automated using Travis CI. New releases are automatically created when new tags are pushed to the repo.
```shell script
$ TAG=0.1.0 make tag
```

## How to contribute
This project has some clear Contribution Guidelines and expectations that you can read here ([CONTRIBUTING](CONTRIBUTING.md)).

The contribution guidelines outline the process that you'll need to follow to get a patch merged.

And you don't just have to write code. You can help out by writing documentation, tests, or even by giving feedback about this work.

Thank you for contributing!