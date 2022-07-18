// Copyright 2022 wgentry22. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
  "github.com/mitchellh/mapstructure"
  "github.com/stretchr/testify/require"
  "github.com/stretchr/testify/suite"
  "strings"
  "testing"
)

type OptionsTestSuite struct {
  suite.Suite
}

func (o *OptionsTestSuite) TestNameOption() {
  expected := "anotherName"
  option := Name(expected)

  opt := defaultOptions(option)
  require.NotEqual(o.T(), opt.configName, defaultConfigName)
  require.Equal(o.T(), opt.configName, expected)
}

func (o *OptionsTestSuite) TestNameOption_WhenProvidedConfigIsEmpty() {
  option := Name("")

  opt := defaultOptions(option)
  require.Equal(o.T(), opt.configName, defaultConfigName)
}

func (o *OptionsTestSuite) TestExtensionOption() {
  expected := "json"
  option := Extension(expected)

  opt := defaultOptions(option)
  require.NotEqual(o.T(), opt.configExtension, defaultConfigExtension)
  require.Equal(o.T(), opt.configExtension, expected)
}

func (o *OptionsTestSuite) TestExtensionOption_WhenProvidedExtensionIsEmpty() {
  option := Extension("")

  opt := defaultOptions(option)
  require.Equal(o.T(), opt.configExtension, defaultConfigExtension)
}

func (o *OptionsTestSuite) TestPathsOption() {
  expected := []string{"/tmp", "/app", "/etc"}
  option := Paths(expected...)

  opt := defaultOptions(option)
  require.NotEqual(o.T(), opt.configPaths, []string{defaultConfigPath})
  require.Equal(o.T(), opt.configPaths, expected)
}

func (o *OptionsTestSuite) TestPathsOption_WhenProvidedPathsIsEmpty() {
  var expected []string
  option := Paths(expected...)

  opt := defaultOptions(option)
  require.Equal(o.T(), opt.configPaths, []string{defaultConfigPath})
}

func (o *OptionsTestSuite) TestReaderOption() {
  reader := strings.NewReader(``)
  option := Reader(reader)

  opt := defaultOptions(option)
  require.NotNil(o.T(), opt.reader)
  require.Same(o.T(), opt.reader, reader)
}

func (o *OptionsTestSuite) TestReaderOption_WhenProvidedReaderIsNil() {
  option := Reader(nil)

  opt := defaultOptions(option)
  require.Nil(o.T(), opt.reader)
}

func (o *OptionsTestSuite) TestHooksOption() {
  additionalHooks := []mapstructure.DecodeHookFunc{func() {}, func() {}, func() {}}

  option := Hooks(additionalHooks...)

  opt := defaultOptions(option)
  require.Equal(o.T(), len(opt.hooks), len(additionalHooks) + len(allHooks))
}

func (o *OptionsTestSuite) TestHooksOption_WhenProvidedHooksIsEmpty() {
  var additionalHooks []mapstructure.DecodeHookFunc

  option := Hooks(additionalHooks...)

  opt := defaultOptions(option)
  require.Equal(o.T(), len(opt.hooks), len(allHooks))
}

func TestOptionsTestSuite(t *testing.T) {
  suite.Run(t, new(OptionsTestSuite))
}
