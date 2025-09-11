package v1

// Sku represents product SKU types
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

// String returns the string representation of the SKU
func (s Sku) String() string {
	names := []string{"BASIC", "STANDARD", "PREMIUM"}
	if int(s) < len(names) {
		return names[s]
	}
	return "UNKNOWN"
}