FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

ENV GOARCH=amd64

RUN go build -o /go/bin/app main.go

## Deploy
FROM gcr.io/distroless/base-debian11

COPY --from=build /go/bin/app /app

EXPOSE 8089

CMD ["/app"]