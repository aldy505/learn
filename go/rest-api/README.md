# Learn Go - REST API

My attempt of creating a full featured REST API using Go.

Just to be clear, you don't have to follow my way of doing project structure, you can do anything you want.

And about importing stuffs, the imports basically came from the module name from `go.mod` file which is `github.com/aldy505/learn-go`, then followed by the folder/package name. For example, the on folder routes, it'd be `github.com/aldy505/learn-go/routes`, not caring about what's the file name is.

This is one thing I learned after throwing myself in Go for a bit of a long time: Everything (struct, function, interfaces) is exported as long as you create a name with the first letter being an uppercase.

`func Hello() {}` -- This is exported, you can access it from anywhere.
`func world() {}` -- This is not exported, you can only access it from the same file.

Packages/modules:
- https://github.com/cristalhq/jwt
- https://github.com/Masterminds/squirrel
- https://github.com/jackc/pgx
- https://github.com/gorilla/csrf
- https://github.com/jordan-wright/email
- https://github.com/go-redis/redis/v8
- https://github.com/levigross/grequests
- https://github.com/dwin/goArgonPass
- https://github.com/unrolled/secure
- https://github.com/hoisie/mustache