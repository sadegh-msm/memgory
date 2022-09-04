# go-redis
> http/rest implementation of redis in golang

## introduction 
Here is an implementation of [redis](https://github.com/redis/redis) with a little twist.<br/>
redis is an in-memory database so It's so fast and its usage is for caching (most of the time) but it can be used for other purposes.

## Structure

### Middlewares
I also used some middlewares:
- logger (for logging user data, end points, time and time taken to serve the request)
- CORS
- bodyLimit (to stop user from adding big data for values in database)
- rateLimit (for stopping high loads on short time)

## End Points
there are 9 `end points` that you can use to work with database :

```go
func NewRouter(e *echo.Echo) *echo.Echo {
    e.POST("/set", cmd.Set) 
    e.GET("/get", cmd.Get) // get value from key by query parameters {"name" for database name} {"key" for data key}
    e.DELETE("/del", cmd.Del)
    e.POST("/use", cmd.UseDB)
    e.POST("/keyregex", cmd.KeyRegex)
    e.GET("/listdata", cmd.ListData) // get list of values from database name by query parameters {"name" for database name}
    e.GET("/listdb", cmd.ListDBs)    // get list of databases from storage name by query parameters {"name" for storage name}
    e.POST("/save", cmd.Save)
    e.POST("/load", cmd.Load)

    return e
}
````

### usage

- `/set` you can set a new (key, value) pair by database name, key and value.
- `/get` you can get data by passing database name and key.
- `/del` you can delete data by passing database name and key.
- `/use` you can set a pointer to passed database and if that doesn't exist will create a new one.
- `/keyregex` you can find a key by passing database name and expected piece of data.
- `/listdata` it will list all data from a database by passing database name.
- `/listdb` it will show all databases from a container by passing the name of the container.
- `/save` will save all data from a database to passed file path in body of request.
- `/load` will load all data to a database from passed file path in body of request.

if you pass a new database name the server will create a new database and set that as reference. </br>
also, if you pass an empty string as a database name it will use last used database as the reference and do the operation on that database.

## Up and Running

its so simple just run these commands, (I also used [echo](https://github.com/labstack/echo) in project you have to first install that by `go get github.com/labstack/echo/v4`)

```sh
go build
./go-redis
```

```sh
curl -X POST \
  http://localhost:8080/set \
  -H 'cache-control: no-cache' \
  -H 'postman-token: 7c9578fb-d646-a7a0-ebe5-831956b9a224' \
  -d '{
	"DbName": "default",
	"Key": "giga",
	"Value": "chad"
}'
```
