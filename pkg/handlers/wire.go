package handlers

import (
	"github.com/WildEgor/checker/pkg/adapters"
	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(NewHealthCheckHandler, adapters.AdaptersSet)
