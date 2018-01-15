package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type xlsx struct {
	Name     string
	FileName string
	Template *xField
	Data     []*xField
	keymap   map[int]*xField
}

func (x *xlsx) Init(fileName string, name string) {
	x.Name = name
	x.FileName = fileName
	x.Data = make([]*xField, 0)
	x.keymap = make(map[int]*xField)
}

func (x *xlsx) Parse(rows [][]string) {
	x.Template = new(xField)
	if ok, _ := x.Template.Init(x.Name, "struct", ""); ok {
		x.Template.ParseSubFieldsDefs(rows[1], rows[2], rows[3])
		for i := 4; i < len(rows); i++ {
			field := x.Template.Copy()

			// comment row
			if strings.HasPrefix(rows[i][0], "//") || rows[i][0] == "" {
				continue
			}
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
			if strings.HasPrefix(v.Type, "//") || (v.Tag != "" && v.Tag != params.tag) {
				x.Template.Fields = append(x.Template.Fields[:i], x.Template.Fields[i+1:]...)
			} else {
				i++
			}
		}
	} else {
		log.Fatalln("Parse", x.Name, "head field")
	}
}

func (x *xlsx) Print() {
	for k, v := range x.Data {
		fmt.Print(k, " ")
		v.Print()
	}
}
