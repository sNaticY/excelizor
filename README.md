# excelizor

A simple tool that can be used to export .xlsx files to lua-table, json and their corresponding csharp classes and golang structs

[![Travis](https://travis-ci.org/sNaticY/excelizor.svg?branch=master)](https://travis-ci.org/sNaticY/excelizor) 
[![Go Report Card](https://goreportcard.com/badge/github.com/sNaticY/excelizor)](https://goreportcard.com/report/github.com/sNaticY/excelizor) 
[![GitHub release](https://img.shields.io/github/release/sNaticY/excelizor.svg)](https://github.com/sNaticY/excelizor/releases)
[![license](https://img.shields.io/github/license/sNaticY/excelizor.svg)](https://github.com/sNaticY/excelizor/blob/master/LICENSE)

## Installation

```bash
$ go get https://github.com/sNaticY/excelizor
```

or you could download release version directly at [Release Page](https://github.com/sNaticY/excelizor/releases)

## Usage

``` Text
Usage: excelizor -p <path> [-lua=<luaExportPath>] [-json=<luaExportPath>] [-csharp=<luaExportPath>] [-golang=<luaExportPath>]
  -csharp string
    	path to place exported .cs class files, export no .cs files if parameter is missing
  -golang string
    	path to place exported .go struct files, export no .go files if parameter is missing
  -json string
    	path to place exported .json files, export no .json files if parameter is missing
  -lua string
    	path to place exported .lua files, export no .lua files if parameter is missing
  -p path
    	[Required] Relative path of excel files folder
```

## Excel Content Format

![example](doc/example.png)

### Sheet

Only the first sheet in .xlsx file will be export 

> Default name in excel is "Sheet1"

### Head & Key

The first 3 rows in your excel is Head, 

First row is descriptions of each field, It will not be exported at all so you can fill it with everything you want. Second row is name of each field and the third row is type of each field.

The first field must be "Id-int" as key of every row.

### Basic Type

* int: `int32` in golang and `int` in csharp
* string: is `string`
* float: `float32` in golang and `float` in csharp
* bool: is `bool`. [true, T, 1, TRUE] is `true` and [false, F, 0, FALSE] is `false`
* comment: any string in type-row start with `//` is commet, that field in each row will be ignored

Excel

|      | 整型         | 字符串              | 浮点        | 布尔       | 注释(不导出)           |
| ---- | ---------- | ---------------- | --------- | -------- | ----------------- |
| Id   | NumberTest | StringTest       | FloatTest | BoolTest | can be empty      |
| Int  | int        | string           | float     | bool     | //comment         |
| 1001 | 345        | This is a string | 2,6       | true     | won't be exported |
| 1002 | nil        | nil              | nil       | nil      | nil               |

Export Json

``` json
[ 
    { 
        "Id": 1001,
        "NumberTest": 345,
        "StringTest": "This is a string",
        "FloatTest": 2.600,
        "BoolTest": true
    },
    { 
        "Id": 1002
    }
]
```

Export Lua

``` lua
local BasicTypes = {
    [1001] = {
        Id = 1001,
        NumberTest = 345,
        StringTest = "This is a string",
        FloatTest = 2.600,
        BoolTest = true,
    },
    [1002] = {
        Id = 1002,
    },
}

return BasicTypes
```

Export csharp class

``` csharp
using System.Collections.Generic;

namespace Configs
{
    public class BasicTypes 
    {
        public int Id;
        public int NumberTest;
        public string StringTest;
        public float FloatTest;
        public bool BoolTest;
    }
}
```

Export golang struct

``` go
package exports

type BasicTypes struct {
	Id int32 `json:"Id"`
	NumberTest int32 `json:"NumberTest"`
	StringTest string `json:"StringTest"`
	FloatTest float32 `json:"FloatTest"`
	BoolTest bool `json:"BoolTest"`
}
```

### Nested Type

* list<T>:δ  `List<T>` in csharp and `[]T` in golang, 
* dict<T>:δ `Dictionary<string, T>` in csharp and `map[string]T` in golang

> T is any Type such as `float` or `int`,  δ is count of column which it cost in table, can be [0, +∞), When δ == 0, the number of elements in the structure can be arbitrary and seperated with "|", otherwise, the maximum number of elements can not exceed the delta, only one element in every single cell.

Excel

|      | Spread Dictionary |       |       |       | Fold Dictionary      | Spread List |      |      |      | Fold List     |
| ---- | ----------------- | ----- | ----- | ----- | -------------------- | ----------- | ---- | ---- | ---- | ------------- |
| Id   | dictTest1         | Item1 | Item2 | Item3 | DictTest2            | ListTest1   |      |      |      | ListTest2     |
| int  | dict<float>:3     |       |       |       | dict<int>:0          | list<int>:3 |      |      |      | list<float>:0 |
| 2002 |                   | 4.44  | 5.55  | 6.66  | Item1=10 \| Item2=11 |             | 123  | 124  | 125  | 0.2\|0.4\|0.6 |
| 2003 |                   | 4.44  | nil   | 6.66  | nil                  | nil         |      |      |      | 1.3\|1.5\|1.7 |

Export Json

```json
[ 
    { 
        "Id": 2002, 
        "DictTest1": { 
            "Item1": 4.440,
            "Item2": 5.550,
            "Item3": 6.660 
        }, 
        "DictTest2": { 
            "Item1": 10,
            "Item2": 11 
        },
        "ListTest1": [ 
            123,
            124,
            125
        ],
        "ListTest2": [ 
            0.200,
            0.400,
            0.600
        ]
    },
    { 
        "Id": 2003, 
        "DictTest1": { 
            "Item1": 4.440,
            "Item3": 6.660 
        },
        "ListTest2": [ 
            1.300,
            1.500,
            1.700
        ]
    }
]
```

Export lua

```lua
local NestedTypes = {
    [2002] = {
        Id = 2002, 
        DictTest1 = {
            Item1 = 4.440,
            Item2 = 5.550,
            Item3 = 6.660, 
        }, 
        DictTest2 = {
            Item1 = 10,
            Item2 = 11, 
        },
        ListTest1 = {
            [0] = 123,
            [1] = 124,
            [2] = 125,
        },
        ListTest2 = {
            [0] = 0.200,
            [1] = 0.400,
            [2] = 0.600,
        },
    },
    [2003] = {
        Id = 2003, 
        DictTest1 = {
            Item1 = 4.440,
            Item3 = 6.660, 
        },
        ListTest2 = {
            [0] = 1.300,
            [1] = 1.500,
            [2] = 1.700,
        },
    },
}

return NestedTypes
```

Export csharp class

```csharp
using System.Collections.Generic;

namespace Configs
{
    public class NestedTypes 
    {
        public int Id;
        public Dictionary<string, float> DictTest1;
        public Dictionary<string, int> DictTest2;
        public List<int> ListTest1;
        public List<float> ListTest2;
    }
}
```

Export golang struct

```go
package exports

type NestedTypes struct {
	Id int32 `json:"Id"`
	DictTest1 map[string]float32 `json:"DictTest1"`
	DictTest2 map[string]int32 `json:"DictTest2"`
	ListTest1 []int32 `json:"ListTest1"`
	ListTest2 []float32 `json:"ListTest2"`
}
```

### Multi-nested Type

* list<ANY_TYPE_OR_NESTED_TYPE>:δ such as list<dict<string>:0>:5
* dict<ANY_TYPE_OR_NESTED_TYPE>:δ such as dict<list<dict<float>:0>:2>:3

> **MOST OF LOGIC IS EXACTLY THE SAME AS BEFORE. HARD TO EXPLAIN, LET'S SEE SOME EXAMPLES**

Excel

| KEY  | DICTINDICT1         |          |       |       |          |       |       |
| ---- | ------------------- | -------- | ----- | ----- | -------- | ----- | ----- |
| Id   | DictTest3           | SubDict1 | Item1 | Item2 | SubDict2 | Item1 | Item2 |
| int  | dict<dict<int>:2>:2 |          |       |       |          |       |       |
| 3001 |                     |          | 3111  | 3112  |          | 3121  | 3122  |
| 3002 |                     | nil      |       |       |          | 3221  | 3222  |

| DICTINDICT2           |                                 |                      | DICTINDICT3                              |
| --------------------- | ------------------------------- | -------------------- | ---------------------------------------- |
| DictTest4             | SubDict1                        | SubDict2             | DictTest5                                |
| dict<dict<float>:0>:2 |                                 |                      | dict<dict<string>:0>:0                   |
|                       | it1=31.11\|it2=31.12\|it3=31.13 | it1=31.21\|it2=31.22 | Subdict1={item1=asd\|item2=sdf}\|Subdict2={item1=qwe\|item2=wer} |
| nil                   |                                 |                      | nil                                      |

Export Json

``` json
[ 
    { 
        "Id": 3001, 
        "DictTest3": {  
            "SubDict1": { 
                "Item1": 3111,
                "Item2": 3112 
            }, 
            "SubDict2": { 
                "Item1": 3121,
                "Item2": 3122 
            } 
        }, 
        "DictTest4": {  
            "SubDict1": { 
                "It1": 31.110,
                "It2": 31.120,
                "It3": 31.130 
            }, 
            "SubDict2": { 
                "It1": 31.210,
                "It2": 31.220 
            } 
        }, 
        "DictTest5": {  
            "Subdict1": { 
                "Item1": "asd",
                "Item2": "sdf" 
            }, 
            "Subdict2": { 
                "Item1": "qwe",
                "Item2": "wer" 
            } 
        }
...

```

Export lua

```lua
local MultiNestedTypes = {
    [3001] = {
        Id = 3001, 
        DictTest3 = { 
            SubDict1 = {
                Item1 = 3111,
                Item2 = 3112, 
            }, 
            SubDict2 = {
                Item1 = 3121,
                Item2 = 3122, 
            }, 
        }, 
        DictTest4 = { 
            SubDict1 = {
                It1 = 31.110,
                It2 = 31.120,
                It3 = 31.130, 
            }, 
            SubDict2 = {
                It1 = 31.210,
                It2 = 31.220, 
            }, 
        }, 
        DictTest5 = { 
            Subdict1 = {
                Item1 = "asd",
                Item2 = "sdf", 
            }, 
            Subdict2 = {
                Item1 = "qwe",
                Item2 = "wer", 
            }, 
        },
...
```

Export csharp class

```csharp
using System.Collections.Generic;

namespace Configs
{
    public class MultiNestedTypes 
    {
        public int Id;
        public Dictionary<string, Dictionary<string, int>> DictTest3;
        public Dictionary<string, Dictionary<string, float>> DictTest4;
        public Dictionary<string, Dictionary<string, string>> DictTest5;
...
```

Export golang struct

```go
package exports

type MultiNestedTypes struct {
	Id int32 `json:"Id"`
	DictTest3 map[string]map[string]int32 `json:"DictTest3"`
	DictTest4 map[string]map[string]float32 `json:"DictTest4"`
	DictTest5 map[string]map[string]string `json:"DictTest5"`
...
```

> **FOR MORE EXAMPLES, PLEASE CHECK `excels/nested_types.xlsx` and `exports/nested_types.*`** 

## Other Features

### Customizable templates

All of exporting features are based on [go-template](https://golang.org/pkg/text/template/), so you can edit `templates/*.tmpl` to do anything you want

* csharp namespace
* csharp class inherit some base class
* golang package name
* ...

### Auto convert file name

We recommand your excel name is `full_lowercase_letters.xlsx`. When exporting csharp class file, it will auto convert your file name to `CamelFileName.cs` to adapt csharp code style. Exporting other files is not affected. 

All the class or struct name, even lua table name will also be `CamelFileName`

### Auto convert field name

In csharp class and golang struct, public field should start with a capital letter. If you start with lowercase letter, we will automatically convert to capital letters for you

### Vertical sheet

Sometimes you have only few rows in a sheet but many fields, it will be much easier to transpose your sheet. So we support vertical sheet.

|      | 整型         | 字符串              | 浮点        | 布尔        | 布尔        | 布尔        | ...            |
| ---- | ---------- | ---------------- | --------- | --------- | --------- | --------- | -------------- |
| Id   | NumberTest | StringTest       | FloatTest | BoolTest1 | BoolTest2 | BoolTest3 | 4,5,6,7,8,9,10 |
| int  | int        | string           | float     | bool      | bool      | bool      | ,,,            |
| 1001 | 345        | This is a string | 2.6       | true      | false     | true      | ,,,            |

Edit your sheet name (default is "Sheet1") to "Vertical" and then you can fill your cell like this

|      | Id         | int    | 1001             |
| ---- | ---------- | ------ | ---------------- |
| 整型   | NumberTest | int    | 345              |
| 字符串  | StringTest | string | This is a string |
| 浮点   | FloatTest  | float  | 2.6              |
| 布尔   | BoolTest1  | bool   | true             |
| 布尔   | BoolTest2  | bool   | false            |
| 布尔   | BoolTest3  | bool   | true             |
| 布尔   | BoolTest4  | bool   | false            |
| 布尔   | BoolTest5  | bool   | true             |
| 布尔   | BoolTest6  | bool   | false            |
| 布尔   | BoolTest7  | bool   | true             |
| 布尔   | BoolTest8  | bool   | false            |
| 布尔   | BoolTest9  | bool   | true             |
| 布尔   | BoolTest10 | bool   | false            |

Then everyhing still works well.

### Type check

That is a basic feature. We will check if the value is valid for the type, so if you fill `2..2`in a float cell, we will tell you.

### Comment row

If it is not enough for you to only use comment column ( field type start with "//" ) , You can insert a row which id start with "//". Everything in this row will be ignored.

|        | 整型             | 字符串              | 浮点        | 布尔       | 注释(不导出)           |
| ------ | -------------- | ---------------- | --------- | -------- | ----------------- |
| Id     | NumberTest     | StringTest       | FloatTest | BoolTest | can be empty      |
| Int    | int            | string           | float     | bool     | //comment         |
| 1001   | 345            | This is a string | 2,6       | true     | won't be exported |
| //1002 | ok, I can fill | everything       | because   | this row | will be ignored   |

## Contributing

If you are interested in contributing to the excelizor project, please make a PR.

## License

This project is licensed under the MIT License.

License can be found [here](https://github.com/sNaticY/excelizor/blob/master/LICENSE).

