# CodeArena Golang 项目模板

```text
├── api					      ## 接口层代码
├── conf				      ## 配置层【使用echo swagger zap日志组件 viper】
│   ├── conf.go
│   ├── config.toml
│   ├── logger.go
│   └── swagger.go
├── consts				    ## 常量池
│   └── setting.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── embeds           	## embeds文件嵌入
│   ├── code.go
│   └── static			  ## 替换echo banner
│       └── banner.txt

├── main.go				  ## 入口文件
├── middware			  ## 中间件
│   ├── auth.go			## 检验中间件
│   └── jwks.go			## 生成读取cert
├── model				    ## 实体类
├── server
│   ├── adapter.go		## 数据库orm
│   └── router.go		  ## echo 路由
└── utils				      ## 工具类
    ├── file.go
    └── viper.go

```

- main分支：未集成orm框架

- zorm分支：集成zorm框架，示例连接mysql

  zorm文档：https://www.yuque.com/u27016943/nrgi00/zorm#TVZkr
