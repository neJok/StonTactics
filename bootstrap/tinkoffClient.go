package bootstrap

import (
	"github.com/nikita-vanyasin/tinkoff"
)

func NewTinkoffClient(terminalKey string, terminalPassword string) *tinkoff.Client {
	client := tinkoff.NewClient(terminalKey, terminalPassword)
	return client
}
