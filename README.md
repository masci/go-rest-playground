# go-rest-playground

`go-rest-playground` is a web service exposing a REST API for booking a class in a gym or studio.

## Install

There are no binary distributions available, so you need a working Go environment to build and
run the service locally:
```sh
$ go get github.com/masci/go-rest-playground && go-rest-playground
```

## Usage

The service can use an in-memory storage (used by default) or a SQLite database file on disk.

To start the in-memory version:
```sh
$ go-rest-playground
Using in-memory storage, all data will be lost on exit
```

To persist data, launch the program passing a valid path to a file with `-use-db`:
```sh
$ go-rest-playground -use-db=./.db
```

## CRUD operations

With the service running you can perform the following operations.

Add a class:
```sh
$ curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Crossfit","start_date":"2022-01-29T00:00:00Z", "end_date": "2022-02-28T00:00:00Z", "capacity": 100}' \
  http://localhost:3333/classes
```

Get the list of classes:
```sh
$ curl -s localhost:3333/classes/ | jq
[
  {
    "ID": "CR6769",
    "name": "Crossfit",
    "start_date": "2022-01-29T00:00:00Z",
    "end_date": "2022-02-28T00:00:00Z",
    "capacity": 100
  }
]
```

Book a class (date must be in the availability range, capacity isn't limited):
```sh
$ curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"customer":"Jane Doe","date":"2022-01-30T00:00:00Z", "class": "FB0001"}' \
  http://localhost:3333/bookings
{"ID":1,"date":"2022-01-30T00:00:00Z","customer":"Jane Doe","class":"FB0001"}
```

## Development

Testing can be very opinionated so I decided to go with the standard library, without
using any framework or external dependency. To execute the test suite, from the root
of the source tree run:
```sh
$ go test ./...
```

Tests run with the volatile storage implementation by default, if you want to run
the same test with the SQLite implementation run them with:
```sh
$ go test ./... -args -storage-type sqlite
```
In a CI environment we would want to run them both.

`chi` was used to implement the HTTP router, along with the helpers to render the
request and response payloads.

`sqlx` was used to talk to the SQLite database.

## Architecture

The code is organized in three packages:
- `models` provides the data types of the data model
- `storage` provides the functionalities to organize and persist data
- `main` implements the REST API service

The code layout reflects the overall architecture:

```
          ┌──────────────────────────────────────────────────────┐
          │                                                      │
          │                                                      ▼
┌───────────────────┐       ┌───────────────────┐      ┌───────────────────┐
│                   │       │                   │      │                   │
│    API service    │──────▶│      Storage      │─────▶│    Data models    │
│                   │       │                   │      │                   │
└───────────────────┘       └───────────────────┘      └───────────────────┘
                                      ▲
                                      │
                        ┌─────────────┴─────────────┐
                        │                           │
                        │                           │
              ┌───────────────────┐       ┌───────────────────┐
              │      In-mem       │       │      SQLite       │
              │  implementation   │       │  implementation   │
              └───────────────────┘       └───────────────────┘
```

## Limitations

- The web service doesn't provide any authentication/authorization feature
- The happy code path was usually assumed, leaving out some error handling