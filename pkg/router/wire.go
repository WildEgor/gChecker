package router

import (
	"github.com/WildEgor/checker/pkg/handlers"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(NewRouter, handlers.HandlersSet)
