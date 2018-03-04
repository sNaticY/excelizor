package main

import (
	"errors"
	"strconv"
	"strings"
)

func convertToVertical(data [][]string) [][]string {
	ret := make([][]string, 0)
	for i := 0; i < len(data[0]); i++ {
		row := make([]string, 0)
		row = append(row, data[0][i])
		ret = append(ret, row)
	}
	for i := 1; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			ret[j] = append(ret[j], data[i][j])
		}
	}
	return ret
}

func trimData(data string) string {
	before := data
	for {
		data = strings.TrimPrefix(data, "{")
		data = strings.TrimSuffix(data, "}")
		if before == data {
			return data
		}
		before = data
	}

}

func splitName(data string) []string {
	subDatas := make([]string, 0)
	equal := strings.Index(data, "=")

	var n string
	var d string
	if equal == -1 {
		d = data
	} else {
		n = data[:equal]
		d = data[equal+1:]
	}
	subDatas = append(subDatas, n)
	subDatas = append(subDatas, d)
	return subDatas
}

func splitSubData(layer int, data string) []string {
	sept := ""
	for i := 1; i < layer; i++ {
		sept += "}"
	}
	sept += "|"

	subDatas := make([]string, 0)

	for {
		pos := strings.Index(data, sept)
		if pos == -1 {
			subDatas = append(subDatas, data)
			break
		} else {
			subData := data[0 : pos+layer-1]
			data = data[pos+layer:]
			subDatas = append(subDatas, subData)
		}
	}
	return subDatas
}

func handleData(dataType string, data string) (string, error) {
	var result string
	var retErr error
	switch dataType {
	case "int":
		ret, err := strconv.Atoi(data)
		result = strconv.Itoa(ret)
		retErr = err
	case "float":
		ret, err := strconv.ParseFloat(data, 32)
		result = strconv.FormatFloat(ret, 'f', 3, 32)
		retErr = err
	case "bool":
		ret, err := strconv.ParseBool(data)
		result = strconv.FormatBool(ret)
		retErr = err
	case "string":
		result = data
		retErr = nil
	default:
		retErr = errors.New("DataType " + dataType + " is invalid for data " + data)
	}
	return result, retErr
}

func name2lower2Camel(name string) (string, string) {
	dotIndex := strings.LastIndex(name, ".")
	lower := name[:dotIndex]
	return lower, name2Camel(lower)
}

func name2Camel(v string) string {
	if v == "" {
		return v
	}
	initial := strings.ToUpper(v[0:1])
	other := v[1:]
	for strings.Index(other, "_") != -1 {
		index := strings.Index(other, "_")
		replace := strings.ToUpper(other[index+1 : index+2])
		s := []string{other[:index], replace, other[index+2:]}
		other = strings.Join(s, "")
	}
	return initial + other
}

func upperInitialChar(str string) string {
	initial := strings.ToUpper(str[0:1])
	other := str[1:]
	return initial + other
}

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func csharpInherit(v string) string {
	if v == "" {
		return v
	}
	return ": " + name2Camel(v)
}

func golangInherit(v string) string {
	return name2Camel(v)
}
