package exports

// NestedTypes exported from nested_types.xlsx
type NestedTypes struct {
	Id        int32              `json:"Id"`
	DictTest1 map[string]float32 `json:"DictTest1"`
	ListTest2 []float32          `json:"ListTest2"`
}
