module confluent-test

go 1.21

require (
	github.com/confluentinc/cc-fflags v0.0.0-00010101000000-000000000000
	github.com/confluentinc/cc-structs v0.0.0-00010101000000-000000000000  
	github.com/go-kit/kit v0.0.0-00010101000000-000000000000
)

// All dependencies are now mocked locally - no external dependencies required
replace (
	github.com/confluentinc/cc-fflags => ./cc-fflags
	github.com/confluentinc/cc-structs => ./cc-structs  
	github.com/go-kit/kit => ./go-kit
)