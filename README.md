# go-encoding-unixtable [![CI](https://github.com/frantjc/go-encoding-unixtable/actions/workflows/push.yml/badge.svg?branch=main&event=push)](https://github.com/frantjc/go-encoding-unixtable/actions) [![godoc](https://pkg.go.dev/badge/github.com/frantjc/go-encoding-unixtable.svg)](https://pkg.go.dev/github.com/frantjc/go-encoding-unixtable) [![goreportcard](https://goreportcard.com/badge/github.com/frantjc/go-encoding-unixtable)](https://goreportcard.com/report/github.com/frantjc/go-encoding-unixtable) ![license](https://shields.io/github/license/frantjc/go-encoding-unixtable)

Golang module akin to `encoding/json` but for encoding structs to a unixtable format for readability:

```sh
Id   Name               Age
1    Obi-Wan Kenobi     35
2    General Grievous   50
```

Useful for displaying CLI output (think `kubectl get pods`).

## install

```sh
go get github.com/frantjc/go-encoding-unixtable
```

## use

See [`example`](example).
