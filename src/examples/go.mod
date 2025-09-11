module examples

go 1.21

require (
  github.com/confluentinc/cc-fflags v0.0.0-00010101000000-000000000000
  github.com/confluentinc/cc-structs v0.0.0-00010101000000-000000000000  
  github.com/go-kit/kit v0.0.0-00010101000000-000000000000
)

// Reference parent directory mocks for examples
replace (
  github.com/confluentinc/cc-fflags => ../cc-fflags
  github.com/confluentinc/cc-structs => ../cc-structs  
  github.com/go-kit/kit => ../go-kit
)

// Reference all parent directory local packages for examples  
replace confluent-test/launchdarkly => ../launchdarkly
replace confluent-test/flags => ../flags
replace confluent-test/static => ../static
replace confluent-test/mock => ../mock
replace confluent-test/fakeflags => ../fakeflags
replace confluent-test/idgenmock => ../idgenmock
replace confluent-test/metricnoop => ../metricnoop
replace confluent-test/ctxlog => ../ctxlog
replace confluent-test/types => ../types