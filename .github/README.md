# CI/CD Configuration

This directory contains the GitHub Actions workflows for automated testing and validation of the Go feature flags testing framework.

## Workflows

### CI Workflow (ci.yml)

The main CI workflow runs on every push and pull request to main/master branches and includes:

#### Build Jobs
- **Multi-version testing**: Tests against Go 1.21 and 1.22
- **Dependency management**: Downloads and verifies Go modules
- **Code formatting**: Ensures code is properly formatted with `gofmt`
- **Static analysis**: Runs `go vet` for code analysis
- **Compilation**: Builds all packages and main executables
- **Race condition detection**: Builds with race detection enabled

#### Testing Jobs  
- **Unit tests**: Runs all tests with race detection
- **Coverage reporting**: Generates test coverage reports
- **Mock validation**: Tests that all mock dependencies work correctly
- **Import validation**: Ensures only local mocks are used (no external dependencies)

#### Code Quality Jobs
- **Linting**: Uses golangci-lint for comprehensive code analysis
- **Security scanning**: Runs Gosec security scanner
- **SARIF reporting**: Uploads security scan results

## Features

### ✅ Automated Validation
- Builds succeed with zero external dependencies
- All mocks function correctly with random value generation
- Code follows Go best practices and formatting standards
- Security vulnerabilities are detected and reported

### ✅ Multi-Environment Testing
- Tests against multiple Go versions for compatibility
- Validates both development and production build scenarios
- Ensures consistent behavior across different environments

### ✅ Comprehensive Coverage
- Unit tests for all mock packages
- Integration tests for feature flag functionality
- Performance benchmarks for mock services
- Import dependency validation

## Local Development

To run the same validations locally:

```bash
# From the src/ directory
go mod tidy
go vet ./...
gofmt -s -l .
go build ./...
go test -race ./...
golangci-lint run
```

## Configuration Files

- `.golangci.yml`: Linting configuration with custom rules
- `go.mod`: Module definition with local mock dependencies
- `go.sum`: Dependencies checksum (intentionally minimal)