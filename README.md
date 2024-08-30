# A Very Simple BFF

## Getting Started

### Create a Resource File

The resources should be registered in the file `resources/resources.json` in the following format:

```json
{
    "swapi-people": {
        "url":"https://swapi.dev/api/people/:personId/",
        "method":"GET"
    }
}
```

### Create a Query File

The queries should be registered in the file `queries/queries.json` in the following format:

```json
{
    "swapi-people": {
        "resource": "swapi-people", // configured resource
        "queryParams": [ // query params to add to request
            {
                "name": "format",
                "type": "constant",
                "value": "json"
            },
            {
                "name": "resourceQueryParamKey",
                "type": "querystring", // query param forwarded from request
                "value": "receivedQueryParam"
            }
        ],
        "pathParams": [ // path params to add to request
            {
                "name": "personId",
                "type": "querystring", // path param from query string
                "value": "otherReceivedQueryParam"
            }
        ],
        "headers": [ // headers to add to request
            {
                "name": "Authorization",
                "type": "constant",
                "value": "Bearer 1234"
            }
        ],
        "returns": ["name"] // fields to return from response
    }
}
```

Params can be of type `constant`, `querystring`, or `header`.

### Start the Server

```bash
go run cmd/main.go
```

### Make a Request

```bash
curl -X GET "http://localhost:8080/swapi-people?receivedQueryParam=aValue&otherReceivedQueryParam=1"
```

