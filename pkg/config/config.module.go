package config

import "github.com/google/wire"

// Set - Creates app configuration set
var Set = wire.NewSet(NewAppConfig)
