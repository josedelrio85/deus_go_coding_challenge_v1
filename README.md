# DEUS Go coding Challenge v1

Resolution of the proposed technical test.

Created using go version `1.17`.

## How to run it?

Clone the project, get into the correct folder and execute the command:

```bash
go run main.go
```

### Ingesting endpoint

To ingest an event, send a POST request using the URL `http://localhost:8080/setevent` and a valid payload:

```json
{
  "url": "http://web1.com/foo",
  "uuid":"4e9ca7cc-2e11-4b8d-adc4-c77e6606dbff"
}
```

```bash
# Example of curl request to ingest an event

curl --request POST 'http://localhost:8080/setevent' -d '{"url": "http://web1.com/foo", "uuid":"4e9ca7cc-2e11-4b8d-adc4-c77e6606dbff"}'
```

You will get a response with this format if the process succeed:

```json
{
  "status":200,
  "success":true
}
```

### Serving endpoint

To serve the number of distinct visitors of any given page, send a GET request using the URL `http://localhost:8080/getevent?url=http://web1.com/foo"` whith a valid querystring:

```bash
curl --request GET 'http://localhost:8080/getevent?url=http://web1.com/foo'
```
You will get a response with this format:

```json
{
  "status":200,
  "success":true,
  "data": {
    "unique_visitors":2
  }
}
```