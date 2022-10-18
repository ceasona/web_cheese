# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```



## 步骤：

1. 安装kratos 、protoc

2. 创建项目模板

   ```
   kratos new helloworld
   ```

3. 自定义proto模板

   ```
   # 生成proto模板
   kratos proto add api/digital_collection/v1/user.proto
   # 生成proto源码
   kratos proto client api/digital_collection/v1/user.proto
   # 生成server模板
   kratos proto server api/digital_collection/v1/user.proto -t internal/service
   ```

4. biz目录

   业务逻辑的组装层，类似 DDD 的 domain 层

5. data目录

   业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口

   NewGormDB建立数据库连接 》NewData实例化Data结构体 》NewUserRepo具体应用

6. internal\service\service.go

   ```go
   var ProviderSet = wire.NewSet(NewStudentService)
   ```

7. 流程

   internal\service\user.go > internal\biz\user.go > internal\data\user.go

8. internal\service\user.go

    [包含部分逻辑代码，相当于view函数]

9.  internal\biz\user.go 

    [逻辑代码，相当于view函数]

10. internal\data\user.go

    [业务数据，相当于model]

11. internal\server

    服务Register，相当于最外层的url









## 参考资料：

- [Go微服务框架go-kratos实战学习03：使用 gorm 实现增删改查操作](https://www.cnblogs.com/jiujuan/p/16338305.html) 
- [Gorm官方文档](https://gorm.io/zh_CN/docs/query.html)
- [微服务电商项目](https://github.com/go-kratos/beer-shop)
- [前后端分离的管理系统](https://github.com/feihua/kratos-mall)
- [基于Kratos的集中式仓库项目模版](https://github.com/wjs1152283574/kratos-mono-repo)