# golang-wasm-example

本项目的线上地址：https://blog.chyroc.cn/golang-wasm-example/

## build go
```
git clone http://github.com/golang/go /path/go
cd /path/go/src
./make.bash
```

## build wasm and server
```
git clone https://github.com/Chyroc/golang-wasm-example
GOARCH=wasm GOOS=js /path/go/bin/go build -o example.wasm wasm.go
npm install http-server -g // 安装依赖
http-server
```

然后访问：http://127.0.0.1:8080 ，就可以看见加1减1的界面了~