module todo-service

go 1.24.4

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/jmoiron/sqlx v1.4.0
	github.com/labstack/echo/v4 v4.13.4
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.73.0
	todo/auth-service v0.0.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/labstack/echo-jwt/v4 v4.3.1 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace todo/auth-service => ../auth-service
