# HTTP from TCP

A [boot.dev](https://boot.dev) guided project that builds an HTTP server from scratch using raw TCP in Go — no `net/http`, just sockets.

## Goal

Understand how HTTP works under the hood by implementing it layer by layer:

- Reading raw bytes from a TCP connection
- Parsing HTTP request lines, headers, and bodies
- Writing valid HTTP responses

## Run

```bash
go run main.go
```

## Requirements

- Go 1.21+
