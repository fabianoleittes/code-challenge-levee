package main

import (
	"os"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/infrastructure"
	"github.com/fabianoleittes/code-challenge-levee/infrastructure/database"
	"github.com/fabianoleittes/code-challenge-levee/infrastructure/log"
	"github.com/fabianoleittes/code-challenge-levee/infrastructure/router"
	"github.com/fabianoleittes/code-challenge-levee/infrastructure/validation"
)

func main() {
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		DbSQL(database.InstancePostgres).
		DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGorillaMux).
		Start()
}
