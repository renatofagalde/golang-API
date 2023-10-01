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

### https://github.com/mongodb/mongo-go-driver
```shell
go get -u go.mongodb.org/mongo-driver/mongo
```


```shell
go get github.com/google/uuid
```


### https://www.mongodb.com/compatibility/docker
```shell
export MONGODB_VERSION=6.0-ubi8
docker run --name mongodb -d mongodb/mongodb-community-server:$MONGODB_VERSION
```

### remove all container
```shell
docker container rm -f $(docker container ls -aq)
```

### Stopping and Removing All Containers
```shell
docker ps -aq | xargs docker stop | xargs docker rm
```