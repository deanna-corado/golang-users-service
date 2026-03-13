go mod init

go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

go get -u github.com/gin-gonic/gin

go get -u golang.org/x/crypto/bcrypt

go get -u github.com/golang-jwt/jwt/v5

go get github.com/joho/godotenv

go install github.com/githubnemo/CompileDaemon@latest

compiledaemon --command="./golang-users-service"

go get github.com/go-gormigrate/gormigrate/v2@latest

go get github.com/go-gormigrate/gormigrate/v2@latest

go get github.com/go-resty/resty/v2
