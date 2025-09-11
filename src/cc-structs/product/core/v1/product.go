package v1

// Re-export from parent package for import compatibility

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
