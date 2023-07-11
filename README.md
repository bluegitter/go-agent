# 编译静态资源

```
cd web-app
npm run build
```

# 编译执行文件

```
linux
GOOS=linux GOARCH=amd64 go build -o sysinfo main.go

windows
GOOS=windows GOARCH=amd64 go build -o sysinfo main.go

macos
GOOS=darwin GOARCH=amd64 go build -o sysinfo main.go
```

# 运行

```
./sysinfo
```
