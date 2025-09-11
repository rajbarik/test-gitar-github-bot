package v1

// Re-export from parent package for import compatibility

type User struct {
	Id    int64
	Email string
}

type Account struct {
	Id   string
	Name string
}

type Organization struct {
	Id   int64
	Name string
	Plan *Plan
}

type Plan struct {
	ProductLevel ProductLevel
}

type ProductLevel int

const (
	ProductLevel_TEAM ProductLevel = iota
	ProductLevel_ENTERPRISE
)

type PrincipalContext struct {
	Principal    *User
	Environments []*Account
	Organization *Organization
}
