package cmd

import (
	"casbin-gin/cmd/api"
	"errors"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	//命令行指令简单使用方法（一行）
	Use: "casbin-gin",
	//命令别名
	Aliases: []string{"cg"},
	//APP简短描述
	Short: "casbin-gin",
	//APP详细描述
	Long: `Gin集成Casbin进行权限管理的示例Web程序!`,
	//隐藏默认的使用说明
	//SilenceUsage: true,
	Version: "0.0.1",
	//用于传参校验（如下：只是传参不为空校验）
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	//
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	//命令业务逻辑
	//Run不带错误返回值，如果产生错误需要返回可用RunE
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	//自定义子命令
	rootCmd.AddCommand(api.StartCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
