//+build wireinject

package main

import (
	"github.com/goml/gomlet/server"
	"github.com/google/wire"
)

func setup() (*server.Server, func(), error) {
	wire.Build(
		applicationSet,
		server.Set,
	)
	return nil, nil, nil
}
