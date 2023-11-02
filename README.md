```shell
golang-basic go mod init golang-basic
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

```shell
go get github.com/golang-jwt/jwt
```

```shell
go get github.com/stretchr/testify
``

### Instalação e uso do mockgen
- https://www.youtube.com/watch?v=pTGdAhnFP3c
- criar pastas test/mocks
- lembrar de criar o arquivo destination antes de rodar o comando
-  mockgen -source=model/repository/user_repository.go -destination=test/mocks/user_repository_mock.go -package=mocks
-  mockgen -source=model/user_domain_interface.go -destination=test/mocks/user_domain_interface_mock.go -package=mocks
```shell
go install go.uber.org/mock/mockgen@latest
```

### https://www.mongodb.com/compatibility/docker
```shell
export MONGODB_VERSION=6.0-ubi8
docker run --name mongodb -d mongodb/mongodb-community-server:$MONGODB_VERSION
```

- GORM
```shell
go get -u gorm.io/gorm
```

- Mysql 
```shell
go get -u gorm.io/driver/mysqldb
```


### remove all container
```shell
docker container rm -f $(docker container ls -aq)
```

### Stopping and Removing All Containers
```shell
sudo docker ps -aq | sudo xargs docker stop | sudo xargs docker rm

sudo docker stop $(docker ps -a -q)
sudo docker rm $(docker ps -a -q)

```

### docker sudo
```shell
sudo usermod -aG docker $USER
sudo reboot
```
