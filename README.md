# Practice for leetcode

## Running databases:

### mysql

```
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<password> \
    -e MYSQL_DATABASE=<database> -e MYSQL_USER=<user>  \
    -e MYSQL_PASSWORD=<password> mysql:latest

```

### sql server

```
docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=<password>" -p 1433:1433 \
    mcr.microsoft.com/mssql/server:2019-latest

```

## problems

* [go](go)
* [python](python)
* [shell](shell)
* [sql](sql)
