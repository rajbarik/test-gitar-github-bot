module confluent-test

go 1.21

require (
	github.com/go-kit/kit v0.12.0
	github.com/confluentinc/cc-fflags v0.0.0-00010101000000-000000000000
	github.com/confluentinc/cc-structs v0.0.0-00010101000000-000000000000
)

// Replace with actual versions when available
replace (
	github.com/confluentinc/cc-fflags => ../vendor/github.com/confluentinc/cc-fflags
	github.com/confluentinc/cc-structs => ../vendor/github.com/confluentinc/cc-structs
) 