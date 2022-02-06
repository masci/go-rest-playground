# go-rest-playground

`go-rest-playground` is a web service exposing a REST API for booking a class in a gym or studio.

## Install

There are no binary distributions available, so you need a working Go environment to build and
run the service locally:
```sh
$ go get github.com/masci/go-rest-playground
```

## Usage

The service can use an in-memory storage (used by default) or a SQLite database file on disk.

To start the in-memory version:
```sh
$ go run .
Using in-memory storage, all data will be lost on exit
```

To persist data, launch the program passing a valid path to a file with `-use-db`:
```sh
$ go run . -use-db=./.db
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
