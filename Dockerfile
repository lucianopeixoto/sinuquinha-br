# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

COPY data/* ./data/

RUN go build -o ./sinuquinhabr

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app /

EXPOSE 443
EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT [ "./sinuquinhabr" ]