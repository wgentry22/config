// Copyright 2022 wgentry22. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config_test

import (
  "github.com/stretchr/testify/require"
  "github.com/stretchr/testify/suite"
  "github.com/wgentry22/config"
  "strings"
  "testing"
)

const (
  configPath = `./internal/testdata`
)

type testConfig struct {
  Key string `mapstructure:"key"`
}

type ConfigTestSuite struct {
  suite.Suite
  defaultOptions []config.Option
}

func (c *ConfigTestSuite) SetupTest() {
  c.defaultOptions = []config.Option{config.Paths(configPath), config.Name("config")}
}

func (c *ConfigTestSuite) TestReadYamlConfigFile() {
  options := append(c.defaultOptions, config.Extension("yaml"))

  conf, err := config.Init(options...)
  require.Nil(c.T(), err)
  require.NotNil(c.T(), conf)
}

func (c *ConfigTestSuite) TestReadJsonConfigFile() {
  options := append(c.defaultOptions, config.Extension("json"))

  conf, err := config.Init(options...)
  require.Nil(c.T(), err)
  require.NotNil(c.T(), conf)
}

func (c *ConfigTestSuite) TestYamlJsonReader() {
  yamlReader := strings.NewReader(``)
  options := []config.Option{config.Extension("yaml"), config.Reader(yamlReader)}

  conf, err := config.Init(options...)
  require.Nil(c.T(), err)
  require.NotNil(c.T(), conf)
}

func (c *ConfigTestSuite) TestReadJsonReader() {
  jsonReader := strings.NewReader(`{}`)
  options := []config.Option{config.Extension("json"), config.Reader(jsonReader)}

  conf, err := config.Init(options...)
  require.Nil(c.T(), err)
  require.NotNil(c.T(), conf)
}

func (c *ConfigTestSuite) TestErrorReturned_WhenReadingFromReader() {
  expected := "failed to read json config from io.Reader"

  jsonReader := strings.NewReader(`{`)
  options := []config.Option{config.Extension("json"), config.Reader(jsonReader)}

  conf, err := config.Init(options...)
  require.Nil(c.T(), conf)
  require.NotNil(c.T(), err)
  require.Equal(c.T(), err.Error(), expected)
}

func (c *ConfigTestSuite) TestErrorReturned_WhenReadingFromFile() {
  expected := "failed to read bad.json from [./internal/testdata]"
  options := []config.Option{config.Extension("json"), config.Name("bad"), config.Paths(configPath)}

  conf, err := config.Init(options...)
  require.Nil(c.T(), conf)
  require.NotNil(c.T(), err)
  require.Equal(c.T(), err.Error(), expected)
}

func (c *ConfigTestSuite) TestMustInitDoesNotPanic_WhenNoErrorReadingConfig() {
  jsonReader := strings.NewReader(`{}`)
  options := []config.Option{config.Extension("json"), config.Reader(jsonReader)}

  require.NotPanics(c.T(), func() {
    config.MustInit(options...)
  })
}

func (c *ConfigTestSuite) TestMustInitPanics_WhenErrorReadingConfig() {
  jsonReader := strings.NewReader(`{`)
  options := []config.Option{config.Extension("json"), config.Reader(jsonReader)}

  require.Panics(c.T(), func() {
    config.MustInit(options...)
  })
}

func (c *ConfigTestSuite) TestGet() {
  expected := "value"
  options := append(c.defaultOptions, config.Extension("yaml"))

  conf, err := config.Init(options...)
  require.Nil(c.T(), err)
  require.NotNil(c.T(), conf)

  var tc testConfig
  err = conf.Get(`options`, &tc)
  require.Nil(c.T(), err)
  require.Equal(c.T(), tc.Key, expected)
}

func (c *ConfigTestSuite) TestGetReturnsError_WhenNoStructPointerProvided() {
  expected := "expected a pointer but got config_test.testConfig"

  options := append(c.defaultOptions, config.Extension("yaml"))

  conf, err := config.Init(options...)
  require.Nil(c.T(), err)
  require.NotNil(c.T(), conf)

  var tc testConfig
  err = conf.Get(`options`, tc)
  require.Equal(c.T(), err.Error(), expected)
}

func TestConfigTestSuite(t *testing.T) {
  suite.Run(t, new(ConfigTestSuite))
}
