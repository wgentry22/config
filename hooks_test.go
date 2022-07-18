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
  "time"
)

const (
  expectedDuration = 5 * time.Second
)

type configWithDuration struct {
  Timeout time.Duration `mapstructure:"timeout"`
}

type HooksTestSuite struct {
  suite.Suite
}

func (h *HooksTestSuite) TestParseDuration() {
  configData := strings.NewReader(`
data:
  timeout: 5s
`)

  var conf config.Interface
  require.NotPanics(h.T(), func() {
    conf = config.MustInit(config.Reader(configData))
  })

  var c configWithDuration
  err := conf.Get(`data`, &c)
  require.Nil(h.T(), err)
  require.Equal(h.T(), expectedDuration, c.Timeout)
}

func (h *HooksTestSuite) TestParseDuration_Error() {
  configData := strings.NewReader(`
data:
  timeout: notADuration
`)
  expectedErrorText := "unable to convert notADuration into a time.Duration"

  var conf config.Interface
  require.NotPanics(h.T(), func() {
    conf = config.MustInit(config.Reader(configData))
  })

  var c configWithDuration
  err := conf.Get(`data`, &c)
  require.NotNil(h.T(), err)
  require.True(h.T(), strings.Contains(err.Error(), expectedErrorText))
}

func TestHooksTestSuite(t *testing.T) {
  suite.Run(t, new(HooksTestSuite))
}
