// Copyright 2022 wgentry22. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
  "github.com/mitchellh/mapstructure"
  "github.com/rotisserie/eris"
  "reflect"
  "time"
)

var (
  allHooks = []mapstructure.DecodeHookFunc{
    durationDecodeHook(),
  }
)

func durationDecodeHook() mapstructure.DecodeHookFunc {
  return func(from, to reflect.Type, data interface{}) (interface{}, error) {
    if from.Kind() == reflect.String && to.String() == "time.Duration" {
      s := data.(string)
      if d, err := time.ParseDuration(s); err != nil {
        return nil, eris.Errorf("unable to convert %s into a time.Duration", data)
      } else {
        return d, nil
      }
    }
    return data, nil
  }
}
