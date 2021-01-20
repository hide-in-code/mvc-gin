# mvc-gin
gin封装的mvc框架(gin+gorm+redigo)，本人不太喜欢前后分离的方式，所以这里会继续采用bootstrap来解决前端渲染。

### 写在前面
本人是PHP半路出家跑来写golang的，go也是半吊子，github上基于gin的优秀的封装非常多，而我这个写出来更多的像是一个玩具，没有经历过正式项目的考验，所以更多的是借鉴意义；我自己也想在这个过程中不断提高自己，同时打造一套自己顺手的"玩具"，当然如果我的代码能对您有零星半点的启发，那将是莫大的荣幸，如有建议和问题，欢迎联系 QQ：2969921454，项目开始于2021年1月20号。

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
### 开发记录
    20210120  首次提交，基本出现mvc雏形 
