package model

type Common struct {
	// A very short time to suspend between two commands
	Shim  int     `json:"SHIM"`
	Scale float64 `json:"SCALE"`
}

const (
	CommonItemTypeInt     = "int"
	CommonItemTypeFloat64 = "float64"
	CommonItemTypeString  = "string"
)

// CommonItem: Metadata of config
type CommonItem struct {
	Name string
	Type string
}
