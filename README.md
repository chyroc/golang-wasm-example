# golang-wasm-example

本项目的线上地址：https://blog.chyroc.cn/golang-wasm-example/

## build wasm
```
GOARCH=wasm GOOS=js /path/go/bin/go build -o example.wasm wasm.go
```

## server html
```
npm install http-server -g // 安装依赖
http-server
```

然后访问：http://127.0.0.1:8080 ，就可以看见加1减1的界面了~