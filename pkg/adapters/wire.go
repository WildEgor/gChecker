package adapters

import "github.com/google/wire"

var AdaptersSet = wire.NewSet(
	NewPingAdapter,
	wire.Bind(new(IPingAdapter), new(*PingAdapter)),
)
