package exports

type SimpleTypes struct {
	Id int32 `json:"Id"`
	NumberTest int32 `json:"NumberTest"`
	StringTest string `json:"StringTest"`
	FloatTest float32 `json:"FloatTest"`
	BoolTest bool `json:"BoolTest"`
	DictTest1 map[string]float32 `json:"DictTest1"`
	DictTest2 map[string]float32 `json:"DictTest2"`
	ListTest1 []string `json:"ListTest1"`
	ListTest2 []string `json:"ListTest2"`
	DictTest3 map[string]map[string]string `json:"DictTest3"`
	DictTest4 map[string]map[string]string `json:"DictTest4"`
	DictTest5 map[string]map[string]string `json:"DictTest5"`
	ListTest3 [][]string `json:"ListTest3"`
	ListTest4 [][]string `json:"ListTest4"`
	ListTest5 [][]string `json:"ListTest5"`
	DictTest6 map[string][]string `json:"DictTest6"`
	DictTest7 map[string][]string `json:"DictTest7"`
	ListTest6 []map[string]string `json:"ListTest6"`
	ListTest7 []map[string]string `json:"ListTest7"`
}