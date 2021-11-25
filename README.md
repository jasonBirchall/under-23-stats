# Football stats for players under 23 years old

Note: This repository is in its "Getting started" phase. I haven't actually got an outline of the requirements, but I know I want to collect and store a number of metrics about football players under 23. This will then present an API for quick lookup of data.

The main intention of this repository is to get me familiar with Gin and RESTful API's using Go.

## How to run locally

```bash
go run main.go
```

_in another terminal_

```bash
curl 'http://localhost:8080/players/Emile%20Smith-Rowe'
```