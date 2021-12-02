package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/n-arsenic/mongostudy/server/internal/router"

	"github.com/caarlos0/env/v6"
	"github.com/spf13/cobra"
)

type DBConfig struct {
	Host     string `env:"MONGO_HOST,required"`
	User     string `env:"MONGO_USER,required"`
	Password string `env:"MONGO_PASSWORD,required"`
	DBName   string `env:"MONGO_DB_NAME,required"`
}

type SrvConfig struct {
	SrvHost string `env:"SERVER_HOST" envDefault:":80"`
}

type Config struct {
	DBConfig
	SrvConfig
}

func Execute() {
	var config Config

	cmd := &cobra.Command{
		Use:   "server",
		Short: "mongo study server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := env.Parse(&config.DBConfig); err != nil {
				return fmt.Errorf("failed to parse db config: %v", err)
			}
			if err := env.Parse(&config.SrvConfig); err != nil {
				return fmt.Errorf("failed to parse server config: %v", err)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return http.ListenAndServe(config.SrvHost, router.New())
		},
		SilenceErrors: true,
	}

	if err := cmd.Execute(); err != nil {
		log.Fatalf("failed to execute cmd: %v", err)
	}
}
