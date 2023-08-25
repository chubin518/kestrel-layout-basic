
# kestrel-layout-basic
基于Golang的应用脚手架

```
kestrel-layout-basic
├── buildinfo
│   └── buildinfo.go
├── cmd
│   ├── main.go
│   ├── wire.go
│   └── wire_gen.go
├── conf
│   ├── dev.yaml
│   ├── prod.yaml
│   ├── stage.yaml
│   └── test.yaml
├── internal
│   ├── app
│   │   ├── app.go
│   │   ├── hook.go
│   │   └── routes.go
│   ├── handler
│   │   └── user.go
│   ├── model
│   │   └── user.go
│   └── service
│       └── user.go
├── pkg
│   ├── config
│   │   └── config.go
│   ├── env
│   │   └── env.go
│   ├── graceful
│   │   ├── builder.go
│   │   ├── func.go
│   │   ├── graceful.go
│   │   ├── lifecycle.go
│   │   ├── options.go
│   │   └── routes.go
│   ├── logging
│   │   ├── context.go
│   │   ├── logging.go
│   │   ├── options.go
│   │   └── utils.go
│   ├── middleware
│   │   ├── recovery.go
│   │   ├── request-id.go
│   │   └── request-logging.go
│   ├── result
│   │   ├── page.go
│   │   └── result.go
│   └── utils
│       ├── number-utils
│       │   └── number-utils.go
│       ├── path-utils
│       │   └── path-utils.go
│       └── time-utils
│           └── time-utils.go
├── webui
│   ├── dist
│       ├── assets
│   │   └── index.html
│   ├── webui.go
├── deploy.sh
├── go.mod
└── readme.md
```

### 使用nunu创建项目

```
# nunu官网
https://github.com/go-nunu/nunu/

# 安装nunu

go install github.com/go-nunu/nunu@latest

# 创建项目
nunu new projectName -r https://github.com/chubin518/kestrel-layout-basic.git

```