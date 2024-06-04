# Invoice Management System

Invoice backend written in golang.

![screenshot for generated Invoice](https://firebasestorage.googleapis.com/v0/b/creadable-22c39.appspot.com/o/Screenshot%20from%202024-06-02%2000-42-22.png?alt=media&token=5ae1dce7-ec93-4f85-9315-6e0f6fd1c52a)

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

### Sample PDF Generated

Find the sample PDF generated at ./1717027122.pdf
