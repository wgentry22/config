// Copyright 2022 wgentry22. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
  "fmt"
  "github.com/mitchellh/mapstructure"
  "github.com/spf13/viper"
  "reflect"
  "strings"
)

type viperConfigImpl struct {
  runtimeViper  *viper.Viper
  decoderConfig viper.DecoderConfigOption
  topLevelKeys  []string
}

// Get implements config.Interface#Get
func (v *viperConfigImpl) Get(key string, structPtr interface{}) error {
  if reflect.TypeOf(structPtr).Kind() != reflect.Pointer {
    return fmt.Errorf("expected a pointer but got %T", structPtr)
  }

  return v.runtimeViper.UnmarshalKey(key, structPtr, v.decoderConfig)
}

// Keys implements config.Interface#Keys
func (v *viperConfigImpl) Keys() []string {
  return v.topLevelKeys
}

// Init initializes a config.Interface using the provided config.Option's
// An error is returned when there is an issue reading in the configuration
func Init(options ...Option) (Interface, error) {
  o := defaultOptions(options...)

  runtimeViper := viper.New()
  runtimeViper.SetConfigType(o.configExtension)

  if o.reader != nil {
    if err := runtimeViper.ReadConfig(o.reader); err != nil {
      return nil, o
    }
  } else {
    runtimeViper.SetConfigName(o.configName)
    for _, configPath := range o.configPaths {
      runtimeViper.AddConfigPath(configPath)
    }

    if err := runtimeViper.ReadInConfig(); err != nil {
      return nil, o
    }
  }

  setOfKeys := make(map[string]bool)

  for _, key := range runtimeViper.AllKeys() {
    if strings.Contains(key, ".") {
      split := strings.Split(key, ".")
      k := split[0]
      if _, ok := setOfKeys[k]; !ok {
        setOfKeys[k] = true
      }
    }
  }

  topLevelKeys := make([]string, 0, len(setOfKeys))

  for key := range setOfKeys {
    topLevelKeys = append(topLevelKeys, key)
  }

  decoderConfig := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(o.hooks...))

  return &viperConfigImpl{
    runtimeViper:  runtimeViper,
    decoderConfig: decoderConfig,
    topLevelKeys:  topLevelKeys,
  }, nil
}

// MustInit initializes a config.Interface using the provided config.Option's
// MustInit panics when there is an issue reading in the configuration
func MustInit(options ...Option) Interface {
  if c, err := Init(options...); err != nil {
    panic(err)
  } else {
    return c
  }
}
