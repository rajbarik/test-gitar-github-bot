package structs

import "context"

// Mock structs for github.com/confluentinc/cc-structs

// Core v1 package
type ContextKey string

const (
	ContextKeyRequestPrincipalContext ContextKey = "principal-context"
)

// Org v1 package structures
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

// Product core v1 package structures
type Sku int

const (
	Sku_BASIC Sku = iota
	Sku_STANDARD
	Sku_PREMIUM
)

var Sku_value = map[string]int32{
	"BASIC":    0,
	"STANDARD": 1,
	"PREMIUM":  2,
}

func (s Sku) String() string {
	names := []string{"BASIC", "STANDARD", "PREMIUM"}
	if int(s) < len(names) {
		return names[s]
	}
	return "UNKNOWN"
}