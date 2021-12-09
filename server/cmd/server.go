package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/n-arsenic/mongostudy/server/internal/router"
	"github.com/n-arsenic/mongostudy/server/internal/storage"

	"github.com/caarlos0/env/v6"
	"github.com/spf13/cobra"
)

type SrvConfig struct {
	SrvHost string `env:"SERVER_HOST" envDefault:":80"`
}

type Config struct {
	storage.DBConfig
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
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, err := storage.NewMongoDB(ctx, config.DBConfig)
			if err != nil {
				return err
			}
			// wait quit , handle quit - close mongo
			log.Printf("start server")
			return http.ListenAndServe(config.SrvHost, router.NewSampleRouter(db))
		},
		SilenceErrors: true,
	}

	if err := cmd.Execute(); err != nil {
		log.Fatalf("failed to execute cmd: %v", err)
	}
}
