# ASIALUXE BILLING ````GOFAX````

### USE
```
go get -u github.com/gin-gonic/gin
```
```
go get -u gorm.io/gorm
```
```
go get -u gorm.io/driver/postgres
```
```
go get github.com/joho/godotenv
```
```
go get github.com/go-playground/validator/v10
```
```
go get github.com/gin-contrib/cors
```
```
go get github.com/google/uuid
```

### DATABASE `POSTGRESSQL`
Supports the latest version of Postgres!
Uses `gorm` package for ORM.

### CACHE OPERATIONS
This document outlines the basic operations for working with cache in Golang. The cache can be used to store, retrieve, and delete data in a backend application.
Uses `.bin` files!
### Cache Set
Use the `Set` method to store data in the cache. You can set a specific expiration time for the cached data.
```
cache.Set("cache_key", data, 1*time.Hour)
```
### Cache GET
Use the `Get` method to retrieve data from the cache.
```
cache.Get("cache_key")
```
### Cache Delete
To delete the data in the cache, use the `Delete` method.
```
cache.Delete("cache_key")
```
## Update packages
```
go mod tidy
```
## BUILD FOR LINUX 64bit
```
GOOS=linux GOARCH=amd64 go build -o ../myapp
```
```
chmod +x myapp
```
```
./myapp 
```
## BUILD FOR LOCALHOST
```
go build -o ../myapp
```
```
chmod +x myapp
```
```
./myapp 
```

## SERVER DEPLOY
Create service file
```
sudo nano /etc/systemd/system/myapp.service
```

```
[Unit]
Description=Go MyApp Application
After=network.target

[Service]
ExecStart=/home/myapp/myapp
WorkingDirectory=/home/myapp
Restart=always
User=ubuntu

[Install]
WantedBy=multi-user.target
```
```
sudo systemctl daemon-reload
```
```
sudo systemctl enable myapp
```
```
sudo systemctl start myapp
```
```
sudo systemctl status myapp
```
```
sudo systemctl stop myapp
```
```
sudo systemctl restart myapp
```