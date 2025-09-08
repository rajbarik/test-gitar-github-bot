package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"

	fflags "github.com/confluentinc/cc-fflags"
	"github.com/confluentinc/cc-fflags/launchdarkly"
	corev1 "github.com/confluentinc/cc-structs/kafka/core/v1"
	orgv1 "github.com/confluentinc/cc-structs/kafka/org/v1"
	
	"./flags"
)

func cluster_bool_value(ld *launchdarkly.LDClient, logger log.Logger) {
	featureFlags := fflags.NewConfigService(ld, logger, nil)
	ctx := context.WithValue(context.Background(), corev1.ContextKeyRequestPrincipalContext, 
	&orgv1.PrincipalContext{
		Principal: &orgv1.User{
			Id:    123, 
			Email: "raj",
		},
		Environments: []*orgv1.Account{{Id: "t456", Name: "test-account"}},
		Organization: &orgv1.Organization{Id: 23, Name: "test-org"},
	})

	cluster := &fflags.Cluster{
		Type: "kafka",
		ID:   "pkc-foobar",
	}
	enabled, err := featureFlags.BoolValue(ctx, flags.MyFeatureFlag, cluster)
	if err != nil {
		panic(err)
	}
	if enabled {
		fmt.Println("flag with cluster enabled")
	} else {
		fmt.Println("flag with cluster disabled")
	}
	enabled1, err := featureFlags.BoolValue(ctx, "my-feature-flag", cluster)
	if err != nil {
		panic(err)
	}
	if enabled1 {
		fmt.Println("flag with cluster enabled")
	} else {
		fmt.Println("flag with cluster disabled")
	}
}

func sequence_id_bool_value(ld *launchdarkly.LDClient, logger log.Logger) {
	eventsFeatureFlags := fflags.NewEventsConfigService(ld, logger, nil)
	enabled, err := eventsFeatureFlags.BoolValue(context.Background(), flags.MyFeatureFlag, 
		&fflags.EventsParams{AppName: "test", SequenceId: 1})
	if err != nil {
		panic(err)
	}
	if enabled {
		fmt.Println("flag with sequence id enabled")
	} else {
		fmt.Println("flag with sequence id disabled")
	}
	enabled1, err := eventsFeatureFlags.BoolValue(context.Background(), "my-feature-flag", 
		&fflags.EventsParams{AppName: "test", SequenceId: 1})
	if err != nil {
		panic(err)
	}
	if enabled1 {
		fmt.Println("flag with sequence id enabled")
	} else {
		fmt.Println("flag with sequence id disabled")
	}
}

func org_resource_id_bool_value(ld *launchdarkly.LDClient, logger log.Logger) {
	eventsFeatureFlags := fflags.NewEventsConfigService(ld, logger, nil)
	enabled, err := eventsFeatureFlags.BoolValue(context.Background(), flags.MyFeatureFlag, 
			&fflags.EventsParams{AppName: "test", OrgResourceId: "test"})
	if err != nil {
		panic(err)
	}
	if enabled {
		fmt.Println("flag with org resource id enabled")
	} else {
		fmt.Println("flag with org resource id disabled")
	}
	enabled1, err := eventsFeatureFlags.BoolValue(context.Background(), "my-feature-flag", 
			&fflags.EventsParams{AppName: "test", OrgResourceId: "test"})
	if err != nil {
		panic(err)
	}
	if enabled1 {
		fmt.Println("flag with org resource id enabled")
	} else {
		fmt.Println("flag with org resource id disabled")
	}
}

func org_resource_id_bool_value_with_reassoc(ld *launchdarkly.LDClient, logger log.Logger) {
	eventsFeatureFlags := fflags.NewEventsConfigService(ld, logger, nil)
	flag_key := "my-feature-flag"
	enabled, err := eventsFeatureFlags.BoolValue(context.Background(), flag_key, 
			&fflags.EventsParams{AppName: "test", OrgResourceId: "test"})
	if err != nil {
		panic(err)
	}
	if enabled {
		fmt.Println("flag with org resource id (with flag_key var) enabled")
	} else {
		fmt.Println("flag with org resource id (with flag_key var) disabled")
	}
}

func main() {
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

	cluster_bool_value(ld, logger)
	sequence_id_bool_value(ld, logger)
	org_resource_id_bool_value(ld, logger)
	org_resource_id_bool_value_with_reassoc(ld, logger)
} 