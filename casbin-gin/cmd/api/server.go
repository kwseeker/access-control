package api

import (
	"casbin-gin/cmd/config"
	"github.com/spf13/cobra"
)

var (
	configYamlFile string

	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "casbin-gin server -c conf/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYamlFile, "config", "c", "conf/settings.yml",
		"Start server with provided configuration file")
	//AppRouters = append(AppRouters, router.InitRouter)
}

// 服务器启动前的准备工作
func setup() {
	//１组件配置解析及安装
	config.ParseAndSetup(configYamlFile)
	//Others...
}

func run() {
	//初始化路由

	//r := gin.New()
	////加载路由
	//server.Init(r)
	////注册中间件
	//r.Use(Db)                                     //数据库
	//r.Use((*jwt.GinJWTMiddleware).MiddlewareFunc) //JWT
	//r.Use(auth.Authorize())                       //Role
	//r.Use()                                       //Permission
	////服务器启动
	//if err := r.Run(":8000"); err != nil {
	//	fmt.Printf("Server start failed, err:%v\n", err)
	//}
}
