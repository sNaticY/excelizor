package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Xlsx struct {
	Name     string
	FileName string
	Template *Field
	Data     []*Field
	keymap   map[int]*Field
}

func (x *Xlsx) Init(fileName string, name string) {
	x.Name = name
	x.FileName = fileName
	x.Data = make([]*Field, 0)
	x.keymap = make(map[int]*Field)
}

func (x *Xlsx) Parse(rows [][]string) {
	x.Template = new(Field)
	if ok, _ := x.Template.Init(x.Name, "struct"); ok {
		x.Template.ParseSubFieldsDefs(rows[1], rows[2])
		for i := 3; i < len(rows); i++ {
			field := x.Template.Copy()
			id, _ := strconv.Atoi(rows[i][0])
			if _, ok2 := x.keymap[id]; !ok2 {
				field.ParseDatas(id, rows[i])
				field.SetLevel(4)
				x.Data = append(x.Data, field)
				x.keymap[id] = field
			} else {
				log.Fatalln("Parse", x.Name, "failed, Id", id, "is duplicated")
			}

		}
		i := 0
		for i < len(x.Template.Fields) {
			v := x.Template.Fields[i]
			if strings.HasPrefix(v.Type, "//") {
				x.Template.Fields = append(x.Template.Fields[:i], x.Template.Fields[i+1:]...)
			} else {
				i++
			}
		}
	} else {
		log.Fatalln("Parse", x.Name, "head field")
	}
}

func (x *Xlsx) Print() {
	for k, v := range x.Data {
		fmt.Print(k, " ")
		v.Print()
	}
}
