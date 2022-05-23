# Redirect

This demonstrates how to http redirect using an API service

## Usage

Run the micro API

```bash
cd redirect/
# micro --registry=mdns api
# micro --registry=mdns --server_address=localhost:8080 run
```

Run the redirect API

```bash
go run main.go --registry=mdns
```

Make request
```
curl -v http://localhost:8080/redirect/url
```

Should return

```
HTTP/1.1 301 Moved Permanently
Location: https://google.com
```

// not work