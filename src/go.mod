module confluent-test

go 1.21

// All dependencies are now mocked locally - no external dependencies required

replace (
	github.com/confluentinc/cc-fflags => ./cc-fflags
	github.com/confluentinc/cc-structs/kafka/core/v1 => ./cc-structs
	github.com/confluentinc/cc-structs/kafka/org/v1 => ./cc-structs
	github.com/confluentinc/cc-structs/kafka/product/core/v1 => ./cc-structs
	github.com/go-kit/kit/log => ./go-kit
)