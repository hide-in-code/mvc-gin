# mvc-gin
gin封装的mvc框架(gin+gorm+redigo)，由于不太喜欢前后分离的方式，所以这里会继续采用bootstrap来解决前端渲染。

### 目录说明
    .
    ├── cmd         执行脚本，计划任务的目录
    ├── component   组件，对外抛出连接实例，自身不处理逻辑
    │   ├── mysql   
    │   └── redis
    ├── config      配置目录，.ini的配置
    ├── controllers 控制器
    │   └── site    
    ├── middleware  中间件
    ├── modelgen    gorm代码生成模块
    ├── models      model目录
    ├── route       router
    └── views       视图模板目录
        ├── site
        └── test

