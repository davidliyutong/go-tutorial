package server

import (
	"github.com/spf13/pflag"
)

// ServerOptions contains the options while running a generic api server.
type ServerOptions struct {
	Healthz bool `json:"healthz"     mapstructure:"healthz"`
}

// NewServerOptions creates a new ServerOptions object with default parameters.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Healthz: true,
	}
}

// Validate checks validation of ServerOptions.
func (s *ServerOptions) Validate() []error {
	errors := []error{}

	return errors
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *ServerOptions) ApplyTo(c *Config) error {
	c.Healthz = s.Healthz

	return nil
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet.
func (s *ServerOptions) AddFlags(fs *pflag.FlagSet) {

	fs.BoolVar(&s.Healthz, "healthz", true, "Add self readiness check /healthz.")
}
