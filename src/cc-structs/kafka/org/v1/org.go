package v1

// User represents a user in the system
type User struct {
	Id    int64
	Email string
}

// Account represents an account/environment
type Account struct {
	Id   string
	Name string
}

// Organization represents an organization
type Organization struct {
	Id   int64
	Name string
	Plan *Plan
}

// Plan represents a service plan
type Plan struct {
	ProductLevel ProductLevel
}

// ProductLevel represents different product levels
type ProductLevel int

const (
	ProductLevel_TEAM ProductLevel = iota
	ProductLevel_ENTERPRISE
)

// PrincipalContext represents the principal context for requests
type PrincipalContext struct {
	Principal    *User
	Environments []*Account
	Organization *Organization
}