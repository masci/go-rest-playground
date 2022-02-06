# go-rest-playground


## CRUD operations

```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Crossfit","start_date":"2020-01-29T00:00:00Z", "end_date": "2020-01-29T00:00:00Z", "capacity": 100}' \
  http://localhost:3333/classes
```