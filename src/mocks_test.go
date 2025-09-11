package main

import (
	"context"
	"testing"

	fflags "github.com/confluentinc/cc-fflags"
	"confluent-test/launchdarkly"
	"confluent-test/static"
	"confluent-test/mock"
	"confluent-test/flags"
	"github.com/go-kit/kit/log"
)

func TestLaunchDarklyMock(t *testing.T) {
	logger := log.NewJSONLogger(nil)
	client, err := launchdarkly.New("test-key", 0, logger)
	if err != nil {
		t.Fatalf("Failed to create LaunchDarkly mock: %v", err)
	}
	
	// Test boolean variation
	result := client.BoolVariation("test-flag", nil, false)
	t.Logf("Boolean variation result: %v", result)
	
	// Test string variation
	strResult := client.StringVariation("test-flag", nil, "default")
	if strResult == "" {
		t.Error("String variation returned empty string")
	}
	t.Logf("String variation result: %s", strResult)
}

func TestConfigServiceMock(t *testing.T) {
	logger := log.NewJSONLogger(nil)
	ld, _ := launchdarkly.New("test-key", 0, logger)
	
	configService := fflags.NewConfigService(ld, logger, nil)
	ctx := context.Background()
	
	// Test boolean value
	boolVal, err := configService.BoolValue(ctx, flags.MyFeatureFlag, nil)
	if err != nil {
		t.Fatalf("BoolValue failed: %v", err)
	}
	t.Logf("Boolean flag result: %v", boolVal)
	
	// Test string value
	strVal, err := configService.StringValue(ctx, flags.MyStringFeatureFlag, nil)
	if err != nil {
		t.Fatalf("StringValue failed: %v", err)
	}
	if strVal == "" {
		t.Error("StringValue returned empty string")
	}
	t.Logf("String flag result: %s", strVal)
	
	// Test int value
	intVal, err := configService.IntValue(ctx, flags.MyIntFeatureFlag, nil)
	if err != nil {
		t.Fatalf("IntValue failed: %v", err)
	}
	t.Logf("Int flag result: %d", intVal)
}

func TestStaticFeatureFlags(t *testing.T) {
	logger := log.NewJSONLogger(nil)
	var ffs *fflags.FFlags
	
	staticFlags := static.NewFeatureFlags(ffs, logger, nil)
	ctx := context.Background()
	
	// Test custom boolean flag
	boolResult, err := staticFlags.EnableMyCustomFeatureFlag(ctx)
	if err != nil {
		t.Fatalf("EnableMyCustomFeatureFlag failed: %v", err)
	}
	t.Logf("Custom boolean flag result: %v", boolResult)
	
	// Test custom int flag
	intResult, err := staticFlags.EnableMyCustomIntFeatureFlag(ctx)
	if err != nil {
		t.Fatalf("EnableMyCustomIntFeatureFlag failed: %v", err)
	}
	t.Logf("Custom int flag result: %d", intResult)
	
	// Test custom string flag
	strResult, err := staticFlags.EnableMyCustomStringFeatureFlag(ctx)
	if err != nil {
		t.Fatalf("EnableMyCustomStringFeatureFlag failed: %v", err)
	}
	if strResult == "" {
		t.Error("EnableMyCustomStringFeatureFlag returned empty string")
	}
	t.Logf("Custom string flag result: %s", strResult)
}

func TestMockConfigService(t *testing.T) {
	mockService := mock.NewMockConfigService()
	ctx := context.Background()
	
	// Test all value types
	boolVal, err := mockService.BoolValue(ctx, "test-key", nil)
	if err != nil {
		t.Fatalf("BoolValue failed: %v", err)
	}
	t.Logf("Mock boolean result: %v", boolVal)
	
	strVal, err := mockService.StringValue(ctx, "test-key", nil)
	if err != nil {
		t.Fatalf("StringValue failed: %v", err)
	}
	t.Logf("Mock string result: %s", strVal)
	
	intVal, err := mockService.IntValue(ctx, "test-key", nil)
	if err != nil {
		t.Fatalf("IntValue failed: %v", err)
	}
	t.Logf("Mock int result: %d", intVal)
	
	jsonVal, err := mockService.JsonValue(ctx, "test-key", nil)
	if err != nil {
		t.Fatalf("JsonValue failed: %v", err)
	}
	if len(jsonVal) == 0 {
		t.Error("JsonValue returned empty result")
	}
	t.Logf("Mock JSON result: %s", string(jsonVal))
}

func BenchmarkMockServices(b *testing.B) {
	logger := log.NewJSONLogger(nil)
	ld, _ := launchdarkly.New("test-key", 0, logger)
	configService := fflags.NewConfigService(ld, logger, nil)
	ctx := context.Background()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = configService.BoolValue(ctx, "benchmark-flag", nil)
	}
}