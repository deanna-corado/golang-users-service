go mod init

go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

go get -u github.com/gin-gonic/gin

go get -u golang.org/x/crypto/bcrypt

go get -u github.com/golang-jwt/jwt/v5

go get github.com/joho/godotenv

go install github.com/githubnemo/CompileDaemon@latest

# use daemon

compiledaemon --command="./user-service"

# gormigrate for migrations

go get github.com/go-gormigrate/gormigrate/v2@latest

# gormigrate for migrations

go get github.com/go-gormigrate/gormigrate/v2@latest

# resty

go get github.com/go-resty/resty/v2
