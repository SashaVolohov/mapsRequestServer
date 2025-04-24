# Maps Service
Allows you to work with maps via API.

## Config
| Type | Description |
| ---- | ----------- |
| port | The port the server will listen on |
| invalidRequestErrorText | Message about invalid client parameters |

## REST API
| Type | Method | Description |
| ---- | ----------- | ---------- |
| /api/:key/:value/:lifeTime | PUT | Create a new key with the specified value and lifetime. |
| /api/:key | GET | Gets the value for the given key |
| /api/:key | DELETE | Deletes the value for the given key |

## Build
```
go build cmd/main.go
```

## Build and run
```
go run cmd/main.go
```
