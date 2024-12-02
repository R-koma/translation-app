# Translation app backend

## Gin

```
go get -u github.com/gin-gonic/gin
```

## Hot Reload

```
go install github.com/air-verse/air@latest
air init
air
```

## env

```
go get github.com/joho/godotenv
```

## GORM

sqlite: test postgres: production

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/postgres
```

## Migration

after creating migrations/migration.go, implement the command below.

```
go run migrations/migration.go
```

## JWT

```
go get -u github.com/golang-jwt/jwt/v5
```

## Secret Key

Create secret key

```
openssl rand -hex 32
```

## CORS

```
go get github.com/gin-contrib/cors
```
