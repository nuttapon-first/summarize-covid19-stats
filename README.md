# Overview

Summarize COVID-19 stats Project

## Requirements

- [Go](https://go.dev/) `1.x`
- [Docker Command](https://www.docker.com/)

## Installation

```bash
$ go mod download
```

## Testing

```bash
# test specific directory
$ cd specific directory
$ go test -v
# test all directory
$ go test ./... -v
```

## Usage

```bash
# local machine default port 8899
## for another port define in local.env
$ go run main.go
```

## Build and run on container

- You can build and run docker images using Makefile or following docker command.
- ⚠️ Make sure you already setup env file correctly.

```bash
# build docker image
$ docker build -t summarize-covid19-stats:latest -f Dockerfile .
# run container default port 8089
## run commands in project root directory to provide environment for container
$ docker run --rm -p 8899:8899 --env-file ./local.env --name summarize-covid19-stats summarize-covid19-stats:latest
```

## API

### Get covid-19 summary data

| Method | Url            | Description            |
| ------ | -------------- | ---------------------- |
| GET    | /covid/summary | get covid summary data |

#### Reponse

##### Success

| Name     | type    | Description                    |
| -------- | ------- | ------------------------------ |
| success  | boolean | response status                |
| AgeGroup | Object  | summary data of each age group |
| Province | Object  | summary data of each province  |

##### Failure

| Name    | type    | Description     |
| ------- | ------- | --------------- |
| success | boolean | response status |
| message | string  | error message   |
