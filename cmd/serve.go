package cmd

/*
Copyright © 2022 Jan-Philipp Schmeißer <dev@marla.one>
This code is owned by marla.one
*/

import (
	"fmt"
	"github.com/marlaone/website/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the marla.one website",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var logger *zap.Logger
		if viper.GetString("app.env") == "debug" {
			logger, _ = zap.NewDevelopment()
		} else {
			logger, _ = zap.NewProduction()
		}
		defer func() {
			if err := logger.Sync(); err != nil {
				log.Fatal(fmt.Errorf("failed to sync logger: %v", err))
			}
		}()

		httpServer := server.NewHttpServer(logger)

		logger.Info(
			"server listening",
			zap.String(
				"address",
				fmt.Sprintf(
					"%s://%s:%d",
					viper.GetString("http.protocol"),
					viper.GetString("http.host"),
					viper.GetInt("http.port"),
				),
			),
		)
		if err := httpServer.Serve(); err != nil {
			logger.Fatal("starting http server failed", zap.Error(err))
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
