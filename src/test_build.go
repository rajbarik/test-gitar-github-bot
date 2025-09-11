package main

import (
	"context"
	"fmt"
	"os"

	"confluent-test/flags"
	"confluent-test/launchdarkly"
	"confluent-test/mock"
	"confluent-test/static"
	fflags "github.com/confluentinc/cc-fflags"
	corev1 "github.com/confluentinc/cc-structs/kafka/core/v1"
	orgv1 "github.com/confluentinc/cc-structs/kafka/org/v1"
	productv1 "github.com/confluentinc/cc-structs/kafka/product/core/v1"
	"github.com/go-kit/kit/log"
)

func testBuild() {
	fmt.Println("=== Testing Mock Dependencies Build ===")

	// Test logger
	logger := log.NewJSONLogger(os.Stdout)
	logger.Log("msg", "Logger test successful")

	// Test LaunchDarkly client
	ld, err := launchdarkly.New("test-key", 0, logger)
	if err != nil {
		fmt.Printf("LaunchDarkly mock failed: %v\n", err)
		return
	}
	fmt.Println("✓ LaunchDarkly mock created successfully")

	// Test feature flags config service
	configService := fflags.NewConfigService(ld, logger, nil)

	// Test with context
	ctx := context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, &orgv1.PrincipalContext{
		Principal: &orgv1.User{
			Id:    123,
			Email: "test@test.com",
		},
		Environments: []*orgv1.Account{{Id: "test-env", Name: "test"}},
		Organization: &orgv1.Organization{Id: 1, Name: "test-org"},
	})

	// Test cluster with product SKU
	cluster := &fflags.Cluster{
		Type: "kafka",
		ID:   "test-cluster",
		Sku:  productv1.Sku_STANDARD,
	}

	// Test boolean flag
	enabled, err := configService.BoolValue(ctx, flags.MyFeatureFlag, cluster)
	if err != nil {
		fmt.Printf("Bool flag test failed: %v\n", err)
		return
	}
	fmt.Printf("✓ Boolean flag test: %v\n", enabled)

	// Test string flag
	strValue, err := configService.StringValue(ctx, flags.MyStringFeatureFlag, &fflags.EventsParams{AppName: "test"})
	if err != nil {
		fmt.Printf("String flag test failed: %v\n", err)
		return
	}
	fmt.Printf("✓ String flag test: %s\n", strValue)

	// Test static feature flags
	var ffs *fflags.FFlags
	staticFlags := static.NewFeatureFlags(ffs, logger, nil)

	customEnabled, err := staticFlags.EnableMyCustomFeatureFlag(ctx)
	if err != nil {
		fmt.Printf("Custom flag test failed: %v\n", err)
		return
	}
	fmt.Printf("✓ Custom boolean flag test: %v\n", customEnabled)

	// Test mock service
	mockService := mock.NewMockConfigService()
	intValue, err := mockService.IntValue(ctx, "test-key", nil)
	if err != nil {
		fmt.Printf("Mock service test failed: %v\n", err)
		return
	}
	fmt.Printf("✓ Mock service int test: %d\n", intValue)

	fmt.Println("=== All Mock Dependencies Working Successfully! ===")
}

func main() {
	testBuild()
}
