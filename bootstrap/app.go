package bootstrap

import (
	"stontactics/mongo"

	"github.com/nikita-vanyasin/tinkoff"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	TinkoffClient *tinkoff.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	app.TinkoffClient = NewTinkoffClient(app.Env.TinkoffTerminalKey, app.Env.TinkoffTerminalPassword)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
