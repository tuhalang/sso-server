package main

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/tuhalang/authen/bootstrap"
)

func main() {
	bootstrap.App("config.yml")
}
