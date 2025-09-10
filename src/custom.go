package main

import (
  "context"
  "encoding/json"
  "fmt"
  "os"
  "time"

  log "./go-kit"
  fflags_main "./cc-fflags"
  corev1 "./cc-structs/core/v1"
  orgv1 "./cc-structs/org/v1"
  
  fflags "./static"
  fflagsmock "./mock"
)

func custom_bool_value() {
  var ffs *fflags_main.FFlags // Mock FFlags instance
  ff := fflags.NewFeatureFlags(ffs, log.NewLogfmtLogger(os.Stderr), nil)

  enabled, err := ff.EnableMyCustomFeatureFlag(context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
    Principal:    &orgv1.User{Id: 123},
    Environments: []*orgv1.Account{{Id: "t456"}},
    Organization: &orgv1.Organization{Plan: &orgv1.Plan{ProductLevel: orgv1.ProductLevel_TEAM}},
  }))
  if err != nil {
    panic(err)
  }
  if enabled {
    fmt.Println("custom boolean flag with cluster enabled")
  } else {
    fmt.Println("custom boolean flag with cluster disabled")
  }
}

func custom_int_value() {
  var ffs *fflags_main.FFlags // Mock FFlags instance
  ff := fflags.NewFeatureFlags(ffs, log.NewLogfmtLogger(os.Stderr), nil)

  intValue, err := ff.EnableMyCustomIntFeatureFlag(context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
    Principal:    &orgv1.User{Id: 123},
    Environments: []*orgv1.Account{{Id: "t456"}},
    Organization: &orgv1.Organization{Plan: &orgv1.Plan{ProductLevel: orgv1.ProductLevel_TEAM}},
  }))
  if err != nil {
    panic(err)
  }
  if intValue == 123 {
    fmt.Println("custom int value 123")
  } else {
    fmt.Println("custom int value not 123")
  }
}

func custom_str_value() {
  var ffs *fflags_main.FFlags // Mock FFlags instance
  ff := fflags.NewFeatureFlags(ffs, log.NewLogfmtLogger(os.Stderr), nil)

  key, err := ff.EnableMyCustomStringFeatureFlag(context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
    Principal:    &orgv1.User{Id: 123},
    Environments: []*orgv1.Account{{Id: "t456"}},
    Organization: &orgv1.Organization{Plan: &orgv1.Plan{ProductLevel: orgv1.ProductLevel_TEAM}},
  }))
  if err != nil {
    panic(err)
  }
  if key == "enabled" {
    fmt.Println("custom string value enabled")
  } else {
    fmt.Println("custom string value disabled")
  }
}

func custom_double_value() {
  var ffs *fflags_main.FFlags // Mock FFlags instance
  ff := fflags.NewFeatureFlags(ffs, log.NewLogfmtLogger(os.Stderr), nil)

  dValue, err := ff.EnableMyCustomDoubleFeatureFlag(context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
    Principal:    &orgv1.User{Id: 123},
    Environments: []*orgv1.Account{{Id: "t456"}},
    Organization: &orgv1.Organization{Plan: &orgv1.Plan{ProductLevel: orgv1.ProductLevel_TEAM}},
  }))
  if err != nil {
    panic(err)
  }
  if dValue == 123.123 {
    fmt.Println("custom double value enabled")
  } else {
    fmt.Println("custom double value disabled")
  }
}
