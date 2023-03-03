# gin-api-frame
开箱即用，轻松上手的基于gin的开源项目骨架

--------------------------------------------

## 1.目录说明
```text
├─ logistics_go
│  ├─ app          
│     ├── dao  // 数据处理层
│     ├── entity //写请求的和返回的数据的结构体
│     ├── global //全局的变量或者常量
│       ├── const //常量
│       ├── variable //变量
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

```text
1.目录和文件命名均采用小写字母加下划线组合的形式
2.变量采用驼峰法命名
3.项目分controller,service，dao,model四层
4.参数校验请写在validator层，务必对请求参数做校验
5.上线代码请务必删除fmt.println之类代码，日志组件基于logrus封装好了，项目中可以直接使用
```


## 3.使用的常用扩展包
>参考：https://juejin.cn/user/3210229684391048

- 1.日志使用

github.com/sirupsen/logrus

github.com/lestrrat-go/file-rotatelogs

- 2.配置文件读取

github.com/spf13/viper

- 3.数据库

gorm.io/gorm

- 4.redis

github.com/go-redis/redis/v8
