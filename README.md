# golang-wasm-example

本项目的线上地址：https://chyroc.cn/golang-wasm-example/

国内用户可以访问：https://chyroc.coding.me/golang-wasm-example/

## dep

* go 1.11([start from here](https://golang.org/doc/install))

## build wasm and server
```
git clone https://github.com/Chyroc/golang-wasm-example
export GO111MODULE=on && go mod vendor
npm install http-server -g // 安装依赖
./run.sh
```

然后访问：http://127.0.0.1:8080 ，就可以看见加1和坦克游戏的界面了~

## 子项目

### 加1（[点击访问](https://chyroc.cn/golang-wasm-example/plus-one/), [国内地址](https://chyroc.coding.me/golang-wasm-example/plus-one/)）

![plus-one](http://recordit.co/08BYIUCJ5X.gif)

### 坦克游戏（[点击访问](https://chyroc.cn/golang-wasm-example/tank/), [国内地址](https://chyroc.coding.me/golang-wasm-example/tank/)）

![tank](http://g.recordit.co/Uq3qxUgVu1.gif)

### 随机生成头像（[点击访问](http://chyroc.cn/golang-wasm-example/generate_avatar/), [国内地址](https://chyroc.coding.me/golang-wasm-example/generate_avatar/)）

![image](https://user-images.githubusercontent.com/15604894/41845539-67eb9798-78a6-11e8-95e3-d8300d855eda.png)
