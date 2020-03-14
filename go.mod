module github.com/Alter/blog

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.0
	github.com/boombuler/barcode v1.0.0 // indirect
	github.com/cosiner/argv v0.0.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-delve/delve v1.4.0 // indirect
	github.com/go-ini/ini v1.51.1
	github.com/go-openapi/spec v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.7 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.3.4
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-runewidth v0.0.8 // indirect
	github.com/peterh/liner v1.2.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/cobra v0.0.6 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/unknwon/com v1.0.1
	go.starlark.net v0.0.0-20200306205701-8dd3e2ee1dd5 // indirect
	golang.org/x/arch v0.0.0-20191126211547-368ea8f32fff // indirect
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	golang.org/x/tools v0.0.0-20200129045341-207d3de1faaf // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.51.1 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	github.com/Alter/blog/conf => /Users/lingfengchen/go/Project/blog/pkg/conf
	github.com/Alter/blog/middleware => /Users/lingfengchen/go/Project/blog/middleware
	github.com/Alter/blog/models => /Users/lingfengchen/go/Project/blog/models
	github.com/Alter/blog/pkg/setting => /Users/lingfengchen/go/Project/blog/pkg/setting
	github.com/Alter/blog/routers => /Users/lingfengchen/go/Project/blog/routers
)
