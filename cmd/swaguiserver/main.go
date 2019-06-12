package main

import (
	"github.com/labstack/echo/v4"
	"github.com/maxstreese/swaguiutils/pkg/swaguihandler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

const (
	SettingNameAddr       = "addr"
	SettingNameDocUrl     = "doc.url"
	SettingNameHideTopbar = "hide.topbar"
)

func main() {
	configureViper()
	rootCmd := newRootCmd()
	rootCmd.Execute()
}

func configureViper() {
	viper.SetConfigName("settings")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("goswagui")
	viper.AutomaticEnv()
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swaguiserver",
		Short: "SwagUiServer is a server that serves the Swagger UI",
		Run: func(cmd *cobra.Command, args []string) {
			runRootCmd(cmd, args)
		},
	}

	cmd.Flags().String("addr", ":8080",
		"the address of the service")
	cmd.Flags().String("doc.url", "",
		"url to the open api document that should be shown be default")
	cmd.Flags().Bool("hide.topbar", false,
		"disable the topbar of the Swagger UI")

	viper.BindPFlags(cmd.Flags())

	return cmd
}

func runRootCmd(cmd *cobra.Command, args []string) {
	addr := viper.GetString(SettingNameAddr)
	docUrl := viper.GetString(SettingNameDocUrl)
	hideTopbar := viper.GetBool(SettingNameHideTopbar)

	swagUiHandler := swaguihandler.New(docUrl, hideTopbar)

	e := echo.New()

	for _, path := range swaguihandler.Paths {
		e.GET(path, swagUiHandler.ServeEcho)
	}

	e.Start(addr)
}
