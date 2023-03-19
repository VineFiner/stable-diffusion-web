# stable-diffusion-web

- gin 安装依赖

```
go mod init healthCheck

go mod tidy

GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -o target/main main.go
```

- 云函数部署

```
export imageurl="registry.us-east-1.aliyuncs.com/vine/stable-diffusion-web:$(date +%F-%H-%M-%S)

s deploy -t ./s.yaml -a fc-access --use-local -y
```