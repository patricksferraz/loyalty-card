/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Netflix/go-env"
	"github.com/c-4u/loyalty-card/app/rest"
	"github.com/c-4u/loyalty-card/infra/db"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// restCmd represents the rest command
func restCmd() *cobra.Command {
	var conf Config

	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		log.Fatal(err)
	}

	restCmd := &cobra.Command{
		Use:   "rest",
		Short: "Run rest Service",

		Run: func(cmd *cobra.Command, args []string) {
			orm, err := db.NewDbOrm(*conf.Db.DsnType, *conf.Db.Dsn)
			if err != nil {
				log.Fatal(err)
			}

			if *conf.Db.Debug {
				orm.Debug(true)
			}

			if *conf.Db.Migrate {
				orm.Migrate()
			}
			defer orm.Db.Close()

			rest.StartRestServer(orm, *conf.RestPort)
		},
	}

	return restCmd
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(basepath + "/../.env")
		if err != nil {
			log.Printf("Error loading .env files")
		}
	}

	rootCmd.AddCommand(restCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
