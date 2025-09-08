package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"

	fflags "github.com/confluentinc/cc-fflags"
	fsflags "github.com/confluentinc/cc-fflags/static"
	"github.com/confluentinc/cc-fflags/launchdarkly"
	corev1 "github.com/confluentinc/cc-structs/kafka/core/v1"
	orgv1 "github.com/confluentinc/cc-structs/kafka/org/v1"
	productv1 "github.com/confluentinc/cc-structs/kafka/product/core/v1"
    org_flags "./flags"
)

func string_flag_key_test(ffs *fflags.FFlags) {
	logger := log.NewJSONLogger(os.Stdout)
	launchdarklyKey, ok := os.LookupEnv("LAUNCHDARKLY_KEY")
	if !ok {
		fmt.Println("missing required env var LAUNCHDARKLY_KEY")
		os.Exit(1)
	}

	ld, err := launchdarkly.New(launchdarklyKey, 5*time.Second, logger)
	if err != nil {
		panic(err)
	}
	featureFlags := fflags.NewConfigService(ld, logger, nil)

	// Test user/account/org targeting
	ctx := context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
		Principal: &orgv1.User{
			Id:    123, // when matching "key" you have to click so it says "string" not "num"
			Email: "cody@confluent.io",
		},
		Environments: []*orgv1.Account{{Id: "t456", Name: "test-account"}},
		Organization: &orgv1.Organization{Id: 23, Name: "test-org"},
	})

	cluster := &fflags.Cluster{
		Type: "kafka",
		ID:   "pkc-foobar",
	}
	enabled1, err := featureFlags.BoolValue(ctx, "my-feature-flag", cluster)
	if err != nil {
		panic(err)
	}
	if enabled1 {
		fmt.Println("boolean flag with cluster enabled")
	} else {
		fmt.Println("boolean flag with cluster disabled")
	}

	// Test flag with JsonVariation.
	raw, err := featureFlags.JsonValue(ctx, "my-json-feature-flag", &fflags.Cluster{Sku: productv1.Sku_BASIC})
	if err != nil {
		panic(err)
	}
	var skus []string
	err = json.Unmarshal(raw, &skus)
	if err != nil {
		panic(err)
	}
	for _, s := range skus {
		sku := productv1.Sku(productv1.Sku_value[s])
		fmt.Println(sku.String())
		fmt.Println(sku == productv1.Sku_STANDARD)
	}

	// Test events config
	eventsFeatureFlags := fflags.NewEventsConfigService(ld, logger, nil)
	enabled2, err := eventsFeatureFlags.BoolValue(context.Background(), "my-feature-flag", &fflags.EventsParams{AppName: "test", SequenceId: 1})
	if err != nil {
		panic(err)
	}
	if enabled2 {
		fmt.Println("boolean flag with sequence id enabled")
	} else {
		fmt.Println("boolean flag with sequence id disabled")
	}

	enabled3, err := eventsFeatureFlags.BoolValue(context.Background(), "my-feature-flag", &fflags.EventsParams{AppName: "test", OrgResourceId: "test"})
	if err != nil {
		panic(err)
	}
	if enabled3 {
		fmt.Println("boolean flag with org resource id enabled")
	} else {
		fmt.Println("boolean flag with org resource id disabled")
	}

	key, err := eventsFeatureFlags.StringValue(context.Background(), "my-str-feature-flag", &fflags.EventsParams{AppName: "test"})
	if err != nil {
		panic(err)
	}
	if key == "enabled" {
		fmt.Println("string value enabled")
	} else {
		fmt.Println("string value disabled")
	}
	intValue, err := eventsFeatureFlags.IntValue(context.Background(), "my-int-feature-flag", &fflags.EventsParams{AppName: "test"})
	if err != nil {
		panic(err)
	}
	if intValue == 123 {
		fmt.Println("int value 123")
	} else {
		fmt.Println("int value not 123")
	}
	dValue, err := eventsFeatureFlags.IntValue(context.Background(), "my-double-feature-flag", &fflags.EventsParams{AppName: "test"})
	if err != nil {
		panic(err)
	}
	if dValue == 123 {
		fmt.Println("double value 123")
	} else {
		fmt.Println("double value not 123")
	}

	// Test connect config
	connectFeatureFlags := fflags.NewConnectConfigService(ld, logger, nil)
	enabled4, err := connectFeatureFlags.BoolValue(context.Background(), "my-feature-flag", &fflags.Cluster{Type: fflags.ConnectType}, &fflags.Connect{ErrorMappingId: 123})
	if err != nil {
		panic(err)
	}
	if enabled4 {
		fmt.Println("boolean flag with org resource id (with flag_key var) enabled")
	} else {
		fmt.Println("boolean flag with org resource id (with flag_key var) disabled")
	}
	
	ff := fsflags.NewFeatureFlags(ffs, log.NewLogfmtLogger(os.Stderr), nil)

	enabled5, err := ff.EnableMyCustomFeatureFlag(context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
		Principal:    &orgv1.User{Id: 123},
		Environments: []*orgv1.Account{{Id: "t456"}},
		Organization: &orgv1.Organization{Plan: &orgv1.Plan{ProductLevel: orgv1.ProductLevel_TEAM}},
	}))
	if err != nil {
		panic(err)
	}
	if enabled5 {
		fmt.Println("custom boolean flag with cluster enabled")
	} else {
		fmt.Println("custom boolean flag with cluster disabled")
	}
	cintValue, err := ff.EnableMyCustomIntFeatureFlag(context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
		Principal:    &orgv1.User{Id: 123},
		Environments: []*orgv1.Account{{Id: "t456"}},
		Organization: &orgv1.Organization{Plan: &orgv1.Plan{ProductLevel: orgv1.ProductLevel_TEAM}},
	}))
	if err != nil {
		panic(err)
	}
	if cintValue == 123 {
		fmt.Println("custom int value 123")
	} else {
		fmt.Println("custom int value not 123")
	}
	key1, err := ff.EnableMyCustomStringFeatureFlag(context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
		Principal:    &orgv1.User{Id: 123},
		Environments: []*orgv1.Account{{Id: "t456"}},
		Organization: &orgv1.Organization{Plan: &orgv1.Plan{ProductLevel: orgv1.ProductLevel_TEAM}},
	}))
	if err != nil {
		panic(err)
	}
	if key1 == "enabled" {
		fmt.Println("custom string value enabled")
	} else {
		fmt.Println("custom string value disabled")
	}
	dbValue, err := ff.EnableMyCustomDoubleFeatureFlag(context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
		Principal:    &orgv1.User{Id: 123},
		Environments: []*orgv1.Account{{Id: "t456"}},
		Organization: &orgv1.Organization{Plan: &orgv1.Plan{ProductLevel: orgv1.ProductLevel_TEAM}},
	}))
	if err != nil {
		panic(err)
	}
	if dbValue == 123.123 {
		fmt.Println("custom double value enabled")
	} else {
		fmt.Println("custom double value disabled")
	}
}

func const_flag_key_test(ffs *fflags.FFlags) {
	logger := log.NewJSONLogger(os.Stdout)
	launchdarklyKey, ok := os.LookupEnv("LAUNCHDARKLY_KEY")
	if !ok {
		fmt.Println("missing required env var LAUNCHDARKLY_KEY")
		os.Exit(1)
	}

	ld, err := launchdarkly.New(launchdarklyKey, 5*time.Second, logger)
	if err != nil {
		panic(err)
	}
	featureFlags := fflags.NewConfigService(ld, logger, nil)

	// Test user/account/org targeting
	ctx := context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
		Principal: &orgv1.User{
			Id:    123, // when matching "key" you have to click so it says "string" not "num"
			Email: "cody@confluent.io",
		},
		Environments: []*orgv1.Account{{Id: "t456", Name: "test-account"}},
		Organization: &orgv1.Organization{Id: 23, Name: "test-org"},
	})

	cluster := &fflags.Cluster{
		Type: "kafka",
		ID:   "pkc-foobar",
	}
	enabled1, err := featureFlags.BoolValue(ctx, org_flags.MyFeatureFlag, cluster)
	if err != nil {
		panic(err)
	}
	if enabled1 {
		fmt.Println("boolean flag with cluster enabled")
	} else {
		fmt.Println("boolean flag with cluster disabled")
	}

	// Test flag with JsonVariation.
	raw, err := featureFlags.JsonValue(ctx, org_flags.MyJsonFeatureFlag, &fflags.Cluster{Sku: productv1.Sku_BASIC})
	if err != nil {
		panic(err)
	}
	var skus []string
	err = json.Unmarshal(raw, &skus)
	if err != nil {
		panic(err)
	}
	for _, s := range skus {
		sku := productv1.Sku(productv1.Sku_value[s])
		fmt.Println(sku.String())
		fmt.Println(sku == productv1.Sku_STANDARD)
	}

	// Test events config
	eventsFeatureFlags := fflags.NewEventsConfigService(ld, logger, nil)
	enabled2, err := eventsFeatureFlags.BoolValue(context.Background(), org_flags.MyFeatureFlag, &fflags.EventsParams{AppName: "test", SequenceId: 1})
	if err != nil {
		panic(err)
	}
	if enabled2 {
		fmt.Println("boolean flag with sequence id enabled")
	} else {
		fmt.Println("boolean flag with sequence id disabled")
	}

	enabled3, err := eventsFeatureFlags.BoolValue(context.Background(), org_flags.MyFeatureFlag, &fflags.EventsParams{AppName: "test", OrgResourceId: "test"})
	if err != nil {
		panic(err)
	}
	if enabled3 {
		fmt.Println("boolean flag with org resource id enabled")
	} else {
		fmt.Println("boolean flag with org resource id disabled")
	}

	key, err := eventsFeatureFlags.StringValue(context.Background(), org_flags.MyStringFeatureFlag, &fflags.EventsParams{AppName: "test"})
	if err != nil {
		panic(err)
	}
	if key == "enabled" {
		fmt.Println("string value enabled")
	} else {
		fmt.Println("string value disabled")
	}
	intValue, err := eventsFeatureFlags.IntValue(context.Background(), org_flags.MyIntFeatureFlag, &fflags.EventsParams{AppName: "test"})
	if err != nil {
		panic(err)
	}
	if intValue == 123 {
		fmt.Println("int value 123")
	} else {
		fmt.Println("int value not 123")
	}
	dValue, err := eventsFeatureFlags.IntValue(context.Background(), org_flags.MyDoubleFeatureFlag, &fflags.EventsParams{AppName: "test"})
	if err != nil {
		panic(err)
	}
	if dValue == 123 {
		fmt.Println("double value 123")
	} else {
		fmt.Println("double value not 123")
	}

	// Test connect config
	connectFeatureFlags := fflags.NewConnectConfigService(ld, logger, nil)
	enabled4, err := connectFeatureFlags.BoolValue(context.Background(), org_flags.MyFeatureFlag, &fflags.Cluster{Type: fflags.ConnectType}, &fflags.Connect{ErrorMappingId: 123})
	if err != nil {
		panic(err)
	}
	if enabled4 {
		fmt.Println("boolean flag with org resource id (with flag_key var) enabled")
	} else {
		fmt.Println("boolean flag with org resource id (with flag_key var) disabled")
	}
}
