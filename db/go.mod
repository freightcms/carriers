module github.com/freightcms/carriers/db

go 1.23.4

require github.com/freightcms/carriers/models v0.0.0-20250515201019-999a2e521f36

require (
	github.com/freightcms/common/models v1.1.0 // indirect
	github.com/freightcms/locations/models v1.1.0 // indirect
	github.com/freightcms/organizations/models v0.0.0-20250319134210-79a6e808531e // indirect
)

replace github.com/freightcms/carriers/models => ../models
