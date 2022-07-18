// Copyright 2022 wgentry22. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config_test

import (
  "fmt"
  "github.com/wgentry22/config"
  "strings"
)

func ExampleInit() {
  type configData struct {
    Value string `mapstructure:"value"`
  }

  yamlReader := strings.NewReader(`
data:
  value: Hello, World!
`)

  readerOption := config.Reader(yamlReader)
  // yaml is the default value
  configExtensionOption := config.Extension("yaml")

  options := []config.Option{readerOption, configExtensionOption}

  // Could also use the following
  // configFileOption := config.Name("myConfigFileName")
  // configPathsOption := config.Paths("/my/absolute/path/to/config", "./my/relative/path/to/config")
  // options := []config.Option{configFileOption, configPathsOption, configExtensionOption}

  conf, err := config.Init(options...)
  if err != nil {
    panic(err)
  }

  var confData configData
  if err := conf.Get(`data`, &confData); err != nil {
    panic(err)
  }

  fmt.Println(confData.Value)
}
