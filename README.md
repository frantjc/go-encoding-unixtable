# go-encoding-unixtable [![CI](https://github.com/frantjc/go-encoding-unixtable/actions/workflows/push.yml/badge.svg?branch=main&event=push)](https://github.com/frantjc/go-encoding-unixtable/actions)

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
