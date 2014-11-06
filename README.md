ruslog During the development :)
======

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
