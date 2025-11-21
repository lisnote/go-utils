# HTTP Echo

`http-echo` is a simple HTTP server that echoes back the details of the incoming request in a JSON format. It's a useful tool for debugging and inspecting HTTP requests.

## Overview

The server listens on a specified address and port. When it receives a request, it captures various details such as the request method, path, query parameters, headers, and body, and then sends them back as a JSON response.

## Usage

Run the server and specify the address and port to listen on.

```bash
http-echo.exe 127.0.0.1:8080
```

Once the server is running, you can send any HTTP request to it, and it will respond with the request details.

### Example Request

```bash
curl -X POST "http://127.0.0.1:8080/test?param1=value1" -H "Content-Type: application/json" -d '{"key": "value"}'
```

### Example Response

```json
{
  "data": { "key": "value" },
  "headers": {
    "Accept": ["*/*"],
    "Content-Length": ["16"],
    "Content-Type": ["application/json"],
    "User-Agent": ["curl/8.10.1"]
  },
  "host": "127.0.0.1:8080",
  "method": "POST",
  "path": "/test",
  "query": { "param1": ["value1"] }
}
```
