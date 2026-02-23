package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/hashicorp/go-plugin"
	"github.com/hexawx/hexawx/core"
)

// Notre implémentation concrète du Driver
type DummyDriver struct{}

// Init implémente la méthode requise par l'interface core.Driver
func (d *DummyDriver) Init(config map[string]string) error {
	// Ici, tu pourrais lire des valeurs comme config["unit"]
	fmt.Println("🔌 [Dummy Plugin] Initialisé avec succès")
	return nil
}

func (p *DummyDriver) Name() (string, error) {
	return "DummyDriver", nil
}

func (d *DummyDriver) Fetch() (core.WeatherRecord, error) {
	return core.WeatherRecord{
		Timestamp:   time.Now(),
		Temperature: 20.0 + rand.Float64()*10,
		Humidity:    40.0 + rand.Float64()*20,
	}, nil
}

func main() {
	driver := &DummyDriver{}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: core.Handshake,
		Plugins: map[string]plugin.Plugin{
			"driver": &core.DriverPlugin{Impl: driver},
		},
	})
}
