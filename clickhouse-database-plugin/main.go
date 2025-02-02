package main

import (
	dbplugin "github.com/hashicorp/vault/sdk/database/dbplugin/v5"
	"log"
	"os"
	clickhouse "vault-plugin-database-clickhouse"
)

func main() {
	err := Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// Run instantiates a Clickhouse object, and runs the RPC server for the plugin
func Run() error {
	dbType, err := clickhouse.New()
	if err != nil {
		return err
	}

	dbplugin.Serve(dbType.(dbplugin.Database))

	return nil
}
