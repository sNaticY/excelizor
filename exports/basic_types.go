package exports

// BasicTypes exported from basic_types.xlsx
type BasicTypes struct {
	Id         int32  `json:"Id"`
	NumberTest int32  `json:"NumberTest"`
	StringTest string `json:"StringTest"`
	BoolTest   bool   `json:"BoolTest"`
}
