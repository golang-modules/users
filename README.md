# Manage users in mysql

## Run mysql using docker.
```

    docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=myrootpwd -e MYSQL_USER=mysqluser -e MYSQL_PASSWORD=mysqlpwd -d mysql

```

## DB Scrips to create user table.
```sql

    CREATE DATABASE main;

    USE main;

    CREATE TABLE `users` (
        `id`        int(6) unsigned NOT NULL AUTO_INCREMENT,
        `firstname` varchar(30) NOT NULL,
        `lastname`  varchar(30) NOT NULL,
        `email`     varchar(30) NOT NULL UNIQUE,
        `password`  varchar(120) NOT NULL,
        `active`    boolean,
        `created`   datetime NOT NULL,
        `updated`   datetime NOT NULL,
        PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

```

## Depandency Packages
```golang

go get github.com/go-sql-driver/mysql
go get golang.org/x/crypto/bcrypt

```

## Environmental variables required

```bash

export MYSQL_DB_URL="mysqluser:mysqlpwd@tcp(127.0.0.1:3306)/main"

```

## Manage User

### Create user
```go

func main() {

	user := models.User{}
	user.ID = 1
	user.FirstName = "FirstName"
	user.LastName = "LastName"
	user.Email = "first.last@email.com"
	user.Password = "password"

	uc := mysql.New()
	res, err := uc.Create(user)

	log.Println(res)
	log.Println(err)
}

```


### List users
```go

func main() {

    uc := mysql.New()
	res, err := uc.GetAll()

	log.Println(res)
	log.Println(err)
}

```

### Users details by id
```go

func main() {

    uc := mysql.New()
	res, err := uc.Get(1)

	log.Println(res)
	log.Println(err)
}

```

### Update User Details details by id
```go

func main() {

    user := models.User{}
    user.ID = 1
    user.FirstName = "FirstName Update"
    user.LastName = "LastName Update"
    user.Email = "first.updated@email.com"
    user.Password = "password"
    
    uc := mysql.New()
    ßres, err := uc.Update(user)

    log.Println(res)
    log.Println(err)
}

```

### Update User Stauts
```go

func main() {

    
    uc := mysql.New()
    ßres, err := uc.UpdateStatus(1, true)

    log.Println(res)
    log.Println(err)
}

```

### Update User password
```go

func main() {
    
    uc := mysql.New()
    res, err := uc.UpdatePassword(1, "newpasswored")

    log.Println(res)
    log.Println(err)
}

```

### Delete User
```go

func main() {
    
    uc := mysql.New()
    res, err := uc.Delete(1)

    log.Println(res)
    log.Println(err)
}

```
