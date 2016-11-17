package scaler

import (
	"github.com/chronojam/prometheus-scaler/config"
)

var Drivers = make(map[string]Driver)

type Driver func(options config.RegistrableScalarConfig) (ScalableResource, error)

func Register(name string, driver Driver) {
	if driver == nil {
		panic("scaler: could not register nil driver.")
	}

	if _, dup := Drivers[name]; dup {
		panic("scaler: could not register duplicate driver: " + name)
	}

	Drivers[name] = driver
}

// A scalableresource is something that can be scaled.
type ScalableResource interface {
	Names() []string

	// To decrease, just pass a negative value here.
	Scale(count int) error
}
