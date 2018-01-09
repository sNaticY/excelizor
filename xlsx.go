package main

import (
	"fmt"
	"strconv"
)

type Xlsx struct {
	Name     string
	Template *Field
	Data     []*Field
}

func (x *Xlsx) Init(name string) {
	x.Name = name
	x.Data = make([]*Field, 0)
}

func (x *Xlsx) Parse(rows [][]string) {
	x.Template = new(Field)
	if ok, _ := x.Template.Init(x.Name, "struct"); ok {
		x.Template.ParseSubFieldsDefs(rows[1], rows[2])
		for i := 3; i < len(rows); i++ {
			field := x.Template.Copy()
			id, _ := strconv.Atoi(rows[i][0])
			field.ParseDatas(id, rows[i])
			field.SetLevel(4)
			x.Data = append(x.Data, field)
		}
	}
}

func (x *Xlsx) Print() {
	for k, v := range x.Data {
		fmt.Print(k, " ")
		v.Print()
	}
}
