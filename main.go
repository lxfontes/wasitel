package main

//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world server --out gen ./wit

import (
	"fmt"
	"math/rand"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.wasmcloud.dev/component/net/wasihttp"
)

var (
	wasiTransport = &wasihttp.Transport{}
	httpClient    = &http.Client{Transport: wasiTransport}
)

const name = "github.com/lxfontes/wasitel"

var tracer = otel.Tracer(name)

func init() {
	setupOTelSDK()
	wasihttp.HandleFunc(helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "hello")
	defer span.End()

	roll := 1 + rand.Intn(6)
	rollValueAttr := attribute.Int("dice.roll", roll)
	span.SetAttributes(rollValueAttr)

	fmt.Fprint(w, "Dice Rolled %d\n", roll)
}

func main() {}
