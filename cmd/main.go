package main

import (
	"flag"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/tuhalang/authen/bootstrap"
)

func main() {
	configFile := flag.String("conf", "config.yml", "The application configuration file")
	flag.Parse()
	bootstrap.Run(*configFile)
}
