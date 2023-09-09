# CodeArena Golang 项目模板

```text
## 接口层代码
├── api		
## 配置层【使用echo swagger zap日志组件 viper】
├── conf						
│   ├── conf.go
│   ├── config.toml
│   ├── logger.go
│   └── swagger.go
## 常量池
├── consts						
│   └── setting.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
## embeds文件嵌入
├── embeds           	
│   ├── code.go
│   └── static		
## 替换echo banner
│       └── banner.txt
## 入口文件
├── main.go		
## 中间件
├── middware					
│   ├── auth.go			## 检验中间件
│   └── jwks.go			## 生成读取cert
## 实体类
├── model						
├── server
│   ├── adapter.go		## 数据库orm
│   └── router.go		## echo 路由
## 工具类
└── utils						
    ├── file.go
    └── viper.go

```