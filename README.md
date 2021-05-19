# GoNowQs
## This golang frame
## Under development


## Swagger 
``` shell
# github.com/swaggo/swag

# admin api document
# swag init -g http/admin/controller/base.go --exclude http/blog -o ./http/admin/swagger

# blog api document
# swag init -g http/blog/v1/controller/base.go --exclude http/admin -o ./http/blog/v1/swagger

swag init
```


## Frame Structure
```
|--cli              cmd appliction
|--config           config message function
|--errorcode        errorcode message function
|--http             about http appliction
|  |--admin         Backstage
|  |--middleware    middleware
|  |--routers       router
|--language         language message function
|--log              log file
|--models           model
|  |--mysql         about mysql model
|  |--redis         about redis model
|--utils            utils function
go.mod
LICENSE
main.go
README.md
```