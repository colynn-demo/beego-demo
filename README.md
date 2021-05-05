# beego-demo

## Requirements
* go 1.15+
* beego 2.0+
* mysql 5.7+
* node 10.0+
* yarn 

## Run

```sh
$ git clone git@github.com:colynn-demo/beego-demo.git
$ cd beego-demo
# backend - beego
$ go run main.go

# frontend - vue
$ cd web 
$ yarn install
$ yarn serve
```

## Changelog
* [0.3.1](https://github.com/colynn-demo/beego-demo/tree/0.3.1) 基于vue创建示例前端项目，完成user用户接口的前端实现及联调验证
* [0.3.0](https://github.com/colynn-demo/beego-demo/tree/0.3.0) 基于user表完成用户的增删改查的示例接口
* [0.2.0](https://github.com/colynn-demo/beego-demo/tree/0.2.0)  通过`bee`创建beego初始化项目，并完成mysql数据库的初始化
