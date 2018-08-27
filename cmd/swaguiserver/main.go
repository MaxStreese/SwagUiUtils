package main

import (
	"github.com/labstack/echo"
	"github.com/maxstreese/swaguiutils/pkg/swaguihandler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "github.com/ogier/pflag"
)

type Settings struct {
	addr        *string
	docUrl      *string
	hideTopebar *bool
}

func main() {
	settings := newSettings()
	rootCmd := newRootCmd(settings)
	rootCmd.Execute()
}

func newSettings() Settings {
	return Settings{new(string), new(string), new(bool)}
}

func newRootCmd(s Settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swaguiserver",
		Short: "SwagUiServer is a server that serves the Swagger UI",
		Run: func(cmd *cobra.Command, args []string) {
			runRootCmd(cmd, args, s)
		},
	}

	cmd.Flags().StringVar(s.addr, "addr", ":8080",
		"the address of the service")
	cmd.Flags().StringVar(s.docUrl, "doc-url", "",
		"url to the open api document that should be shown be default")
	cmd.Flags().BoolVar(s.hideTopebar, "hide-topbar", false,
		"disable the topbar of the Swagger UI")

	viper.BindPFlags(cmd.Flags())

	return cmd
}

func runRootCmd(cmd *cobra.Command, args []string, s Settings) {
	swagUiHandler := swaguihandler.New(*s.docUrl, *s.hideTopebar)

	e := echo.New()
	e.GET("/*", swagUiHandler.ServeEcho)
	e.Start(*s.addr)
}
