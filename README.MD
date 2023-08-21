auto create migrations

`!bash
atlas migrate diff --env gorm 
`

run migrations

`!bash
atlas migrate apply --dir "file://migrations" --url "mysql://<user>:<pass>@<host>:<port>/<db>"
`
