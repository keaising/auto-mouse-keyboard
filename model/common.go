package model

type Common struct {
	// 两条指令之间的暂停时间
	Shim  int     `json:"SHIM"`
	Scale float64 `json:"SCALE"`
}

const (
	CommonItemTypeInt     = "int"
	CommonItemTypeFloat64 = "float64"
	CommonItemTypeString  = "string"
)

// 标注每条配置的名字和类型，需要反序列化到 Common 上
type CommonItem struct {
	Name string
	Type string
}
