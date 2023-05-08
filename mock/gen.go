package mock

//go:generate go run github.com/derision-test/go-mockgen/cmd/go-mockgen --disable-formatting -f github.com/ionos-cloud/streams -i Source -o source.go
//go:generate go run github.com/derision-test/go-mockgen/cmd/go-mockgen --disable-formatting -f github.com/ionos-cloud/streams -i Sink -o sink.go
//go:generate go run github.com/derision-test/go-mockgen/cmd/go-mockgen --disable-formatting -f github.com/ionos-cloud/streams -i Table -o table.go
