package adapters

import "github.com/google/wire"

//
var Set = wire.NewSet(
	NewPingAdapter,
	wire.Bind(new(IPingAdapter), new(*PingAdapter)),
)
