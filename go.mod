module github.com/lepingbeta/go-common-v2-dh-oauth2-server

replace (
	github.com/lepingbeta/go-common-v2-dh-log => ../go-common-v2-dh-log
	github.com/lepingbeta/go-common-v2-dh-redis => ../go-common-v2-dh-redis
)

go 1.22.1

require (
	github.com/lepingbeta/go-common-v2-dh-log v0.0.0-00010101000000-000000000000
	github.com/lepingbeta/go-common-v2-dh-redis v0.0.0-00010101000000-000000000000
)

require github.com/gomodule/redigo v1.9.2 // indirect
