# MiniDB

MiniDB is a minimal, production-inspired relational database written in Go.

## Features
- SQL interface
- Tables and schemas
- Hash indexes
- WAL-based persistence
- Read Committed transactions
- Interactive REPL
- HTTP demo app
- Unit tests

## Architecture
Parser → Executor → Tables → Indexes → WAL → Disk

## Run
go run cmd/repl/main.go  
go run cmd/web/main.go

## Tests
go test ./...
