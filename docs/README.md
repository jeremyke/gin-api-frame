# gin-api-frame
>gin重构的物流系统

## 1.目录说明
```text
├─ gin-api-frame
│  ├─ app          
│     ├── dao  //（再三犹豫，觉得还是有必要）数据库的curd
│     ├── entity //写请求的和返回的数据的结构体
│     ├── global //全局的结构体或者常量
│     ├── http //http服务
│       ├── controller //控制器层
│       ├── middleware //中间件
│       ├── validator //校验层
│     ├── model//数据库的ORM
│     ├── service//服务层，写逻辑
│  ├─ bootstrap        //写服务起来是初始化的逻辑
│  ├─ cmd      //项目入口
│     ├── api //api服务入口
│  ├─ docs   //文档
│     ├── ...
│  ├─ routers          //路由
│     ├── ...
│  ├─ storage           //写日志的
│     ├── logs
│  ├─ utils          //工具
│  ├─ vendor          //扩展包
```
## 2.项目约束

1.目录和文件命名均采用小写字母加下划线组合的形式

## 3.使用的常用扩展包
>参考：https://juejin.cn/user/3210229684391048

 - 1.日志使用 
 
 github.com/sirupsen/logrus
 
 github.com/lestrrat-go/file-rotatelogs
 
 - 2.配置文件读取 
 
 github.com/spf13/viper
 
 - 3.数据库
 
 gorm




