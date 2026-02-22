package main

import (
	"fmt"
	"time"

	"github.com/HexaWX/HexaWX/core"
	"github.com/hashicorp/go-plugin"
)

// Notre implémentation concrète du Driver
type DummyDriver struct{}

// Init implémente la méthode requise par l'interface core.Driver
func (d *DummyDriver) Init(config map[string]string) error {
	// Ici, tu pourrais lire des valeurs comme config["unit"]
	fmt.Println("🔌 [Dummy Plugin] Initialisé avec succès")
	return nil
}

func (d *DummyDriver) Fetch() (core.WeatherRecord, error) {
	return core.WeatherRecord{
		Timestamp:   time.Now(),
		Temperature: 22.5,
		Humidity:    50.0,
	}, nil
}

func main() {
	driver := &DummyDriver{}

	// La configuration doit correspondre à celle définie dans le package core
	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "GOWX_PLUGIN",
		MagicCookieValue: "hello",
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"driver": &core.DriverPlugin{Impl: driver},
		},
	})
}
