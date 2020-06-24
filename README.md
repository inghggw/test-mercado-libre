**Nivel 1**
- go run cmd/main.go

**Ejecutar el makefile para correr el programa**

- make dev

**Ejecutar con coverage los test**

- go test -coverprofile fmtcoverage.html fmt

**Analizar la cobertura de los test**

- go test -coverprofile cover.out ./...
- go tool cover -html=cover.out
