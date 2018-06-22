# golang-wasm-example

本项目的线上地址：https://blog.chyroc.cn/golang-wasm-example/

国内用户可以访问：https://chyroc.coding.me/golang-wasm-example/

## build go
```
git clone http://github.com/golang/go /path/go
cd /path/go/src
./make.bash
```

## build wasm and server
```
git clone https://github.com/Chyroc/golang-wasm-example
npm install http-server -g // 安装依赖
./run.sh
```

然后访问：http://127.0.0.1:8080 ，就可以看见加1和坦克游戏的界面了~

## 子项目

### 加1（[点击访问](https://blog.chyroc.cn/golang-wasm-example/plus-one/)）

![plus-one](http://recordit.co/08BYIUCJ5X.gif)

### 坦克游戏（[点击访问](https://blog.chyroc.cn/golang-wasm-example/tank/)）

![tank](http://g.recordit.co/Uq3qxUgVu1.gif)
