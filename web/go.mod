module github.com/freightcms/carriers/web

go 1.23.4

require (
	github.com/freightcms/carriers/db v1.1.0
	github.com/freightcms/carriers/db/mongodb v0.0.0-20250515175618-20ea3c516410
	github.com/freightcms/carriers/models v0.0.0-20250515201019-999a2e521f36
	github.com/freightcms/common/models v1.1.0
	github.com/freightcms/organizations/web v0.0.0-20250319134210-79a6e808531e
	github.com/graphql-go/graphql v0.8.1
)

require (
	github.com/freightcms/locations/models v1.1.0 // indirect
	github.com/freightcms/locations/web v1.0.1 // indirect
	github.com/freightcms/organizations/db v0.0.0-20250319134210-79a6e808531e // indirect
	github.com/freightcms/organizations/db/mongodb v0.0.0-20250319134210-79a6e808531e // indirect
	github.com/freightcms/organizations/models v0.0.0-20250319134210-79a6e808531e // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/labstack/echo/v4 v4.13.3 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	go.mongodb.org/mongo-driver v1.17.3 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)

replace github.com/freightcms/carriers/models => ../models

replace github.com/freightcms/carriers/db => ../db
