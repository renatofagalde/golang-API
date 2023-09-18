```shell
golang-basicgo mod init golang-basic
```

```shell
go get -u github.com/gin-gonic/gin
```
```shell
go get -u github.com/go-playground/validator/v10
```
```shell
go get -u go.uber.org/zap
```

### https://www.mongodb.com/compatibility/docker
```shell
export MONGODB_VERSION=6.0-ubi8
docker run --name mongodb -d mongodb/mongodb-community-server:$MONGODB_VERSION
```