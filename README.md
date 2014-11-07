ruslog During the development :)
======

[![Build Status](http://img.shields.io/travis/dogenzaka/ruslog.svg?style=flat)](https://travis-ci.org/dogenzaka/ruslog)
[![Coverage](http://img.shields.io/codecov/c/github/dogenzaka/ruslog.svg?style=flat)](https://codecov.io/github/dogenzaka/ruslog)

wrap logrush

# What is ruslog?

logging library that extend the logrus

# Features

- [x] Manage multiple loggers. 
- [x] Manage of an Appender and Formatter. 
- [x] File appender
- [ ] Asynchronous log output


# Requirements

- [logrus](https://github.com/sirupsen/logrus)
- [rotator](https://github.com/dogenzaka/rotator)

# Getting started

```sh
$ go get github.com/dogenzaka/ruslog
```

# Example

TODO

# Default appender and formatter

- Formatter
  - `Default`
  - `Text`
  - `Json`

- Appender
  - `Default`

## Add

TODO

# Development

## Test

```go
$ make install-test
$ make test
```

# License

MIT License
