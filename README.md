# Invoice Management System

Invoice backend written in golang.

## Frameworks & Tools

- GoFiber (Golang Web Framework)
- Golang Migrate (Database Migration Tool)
- SQLC (Type safe SQL go-code generatoe)
- Auth0 (Authentication and Authorization)
- [jung-kurt/gofpdf](github.com/jung-kurt/gofpdf). (PDF Generating Library)

## Features

- Create Invoice
- Download Invoice
- Get Invoice/Invoices (RBAC)
- Create company
- etc

## Getting Started

- clone repo `git clone https://github.com/kevinkimutai/nvoice-management-system.git`
- Run all `Makefile` commands to get DB up and running
- Setup [Auth0](https://auth0.com/docs/libraries#backend).
- run `go run cmd/main.go` inside dir
