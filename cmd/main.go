package main

import (
	"github.com/huuthuan-nguyen/klik-dokter/app"
	"github.com/huuthuan-nguyen/klik-dokter/app/utils"
)

func main() {
	config := utils.ReadConfig() // read config from env
	application := app.NewApp(config)
	application.Run()
}
