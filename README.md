# Confluent Test

This repository contains Go code for testing Confluent feature flags functionality.

## Project Structure

The Go files have been reorganized into a proper package structure under the `src/` directory:

```
.
├── src/                  # Main Go package directory
│   ├── flags/           # Feature flag constants package
│   │   └── constants.go
│   ├── main.go          # Main application entry point
│   ├── custom.go        # Custom feature flag functions
│   ├── bool_value.go    # Boolean feature flag examples
│   ├── string_value.go  # String feature flag examples
│   ├── int_value.go     # Integer feature flag examples
│   ├── go.mod           # Go module definition
│   └── README.md        # Detailed documentation
├── flags.yml            # Feature flags configuration
└── README.md            # This file
```

## Quick Start

1. Navigate to the src directory:
   ```bash
   cd src
   ```

2. Set up your environment:
   ```bash
   export LAUNCHDARKLY_KEY="your-launchdarkly-api-key"
   ```

3. Run the application:
   ```bash
   go mod tidy
   go run .
   ```

## Features

- **Centralized Flag Management**: All feature flag keys are defined as constants in the `flags` package
- **Modular Structure**: Code is organized into logical packages and files
- **Multiple Examples**: Different types of feature flag usage (boolean, string, integer, JSON)
- **Proper Error Handling**: Comprehensive error handling throughout the codebase

## Documentation

For detailed documentation about the package structure and usage, see [src/README.md](src/README.md).

## Dependencies

- Go 1.21+
- LaunchDarkly API key
- Confluent feature flags libraries
