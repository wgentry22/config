package config

import (
  "fmt"
  "github.com/mitchellh/mapstructure"
  "io"
)

const (
  defaultConfigName      = `config`
  defaultConfigExtension = `yaml`
  defaultConfigPath      = `.`
)

// Interface is an abstraction around the underlying configuration
type Interface interface {
  // Get fills the structPtr with configuration values found at key
  Get(key string, structPtr interface{}) error
}

// Options represents the information required to create an Interface instance
type Options struct {
  configName      string
  configExtension string
  configPaths     []string
  reader          io.Reader
  hooks           []mapstructure.DecodeHookFunc
}

// Error returns a string representation of the error state during Init and MustInit
func (o *Options) Error() string {
  if o.reader != nil {
    return fmt.Sprintf("failed to read %s config from io.Reader", o.configExtension)
  }

  return fmt.Sprintf("failed to read %s.%s from %s", o.configName, o.configExtension, o.configPaths)
}

// Option functions allow for customization of the associated Options
type Option func(*Options)

func defaultOptions() *Options {
  return &Options{
    configName:      defaultConfigName,
    configExtension: defaultConfigExtension,
    configPaths:     []string{defaultConfigPath},
    reader:          nil,
    hooks:           []mapstructure.DecodeHookFunc{},
  }
}

// Name is an Option that sets the name of the configuration file to search for
func Name(configName string) Option {
  return func(options *Options) {
    if configName != "" {
      options.configName = configName
    }
  }
}

// Extension is an Option that sets the file extension of the configuration file to search for
func Extension(configExtension string) Option {
  return func(options *Options) {
    if configExtension != "" {
      options.configExtension = configExtension
    }
  }
}

// Paths is an Option that sets where on the filesystem Interface should look for your configuration file
func Paths(configPaths ...string) Option {
  return func(options *Options) {
    if len(configPaths) > 0 {
      options.configPaths = configPaths
    }
  }
}

// Reader is an Option that populates the Interface from an io.Reader
// Helpful in testing.
func Reader(reader io.Reader) Option {
  return func(options *Options) {
    if reader != nil {
      options.reader = reader
    }
  }
}

// Hooks is an Option which allows additional customization of how to unmarshal the underlying
// configuration into a struct
func Hooks(hooks ...mapstructure.DecodeHookFunc) Option {
  return func(options *Options) {
    if len(hooks) > 0 {
      options.hooks = append(options.hooks, hooks...)
    }
  }
}
