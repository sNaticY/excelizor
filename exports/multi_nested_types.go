package exports

// MultiNestedTypes exported from multi_nested_types.xlsx
type MultiNestedTypes struct {
	Id        int32                        `json:"Id"`
	DictTest5 map[string]map[string]string `json:"DictTest5"`
	ListTest3 [][]string                   `json:"ListTest3"`
	ListTest4 [][]string                   `json:"ListTest4"`
	ListTest5 [][]string                   `json:"ListTest5"`
	DictTest6 map[string][]string          `json:"DictTest6"`
	DictTest7 map[string][]string          `json:"DictTest7"`
	ListTest6 []map[string]string          `json:"ListTest6"`
	ListTest7 []map[string]string          `json:"ListTest7"`
}
