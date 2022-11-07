module github.com/wasmcloud/actor-tinygo/example

go 1.17

require (
	github.com/wasmcloud/actor-tinygo v0.1.3
	github.com/wasmcloud/tinygo-msgpack v0.1.4 // indirect
)

require github.com/oneitfarm/uppercase v0.0.0

require github.com/wasmcloud/tinygo-cbor v0.1.0 // indirect

replace github.com/oneitfarm/uppercase => ../uppercase/tinygo
