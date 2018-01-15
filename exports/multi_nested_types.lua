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
        ListTest3 = {
            [0] = {
                [0] = "ListTest3[0][0]",
                [1] = "LIstTest3[0][1]",
            },
            [1] = {
                [0] = "ListTest3[1][0]",
                [1] = "ListTest3{1][1]",
            },
        },
        ListTest4 = {
            [0] = {
                [0] = "list[0][0]",
                [1] = "List[0][1]",
            },
            [1] = {
                [0] = "list[1][0]",
                [1] = "list[1][1]",
            },
        },
        ListTest5 = {
            [0] = {
                [0] = "123",
                [1] = "234",
                [2] = "345",
            },
            [1] = {
                [0] = "234",
                [1] = "345",
                [2] = "456",
            },
        }, 
        DictTest6 = {
            Sublist1 = {
                [0] = "dictTest6[\"sublist1\"][0]",
                [1] = "dictTest6[\"sublist1\"][1]",
            },
            Sublist2 = {
                [0] = "dictTest6[\"sublist2\"][0]",
                [1] = "dictTest6[\"sublist2\"][1]",
            }, 
        }, 
        DictTest7 = {
            Sublist1 = {
                [0] = "zxc",
                [1] = "xcv",
                [2] = "cvb",
            },
            Sublist2 = {
                [0] = "vbn",
                [1] = "bnm",
                [2] = "mnb",
            }, 
        },
        ListTest6 = { 
            [0] = {
                Item1 = "fse",
                Item2 = "fsd", 
            }, 
            [1] = {
                Item1 = "tyu",
                Item2 = "poi", 
            },
        },
        ListTest7 = { 
            [0] = {
                Item0_1 = "fse",
                Item0_2 = "fsd", 
            }, 
            [1] = {
                Item1_1 = "qwe",
                Item1_2 = "qse", 
            },
        },
    },
    [3002] = {
        Id = 3002, 
        DictTest3 = { 
            SubDict2 = {
                Item1 = 3221,
                Item2 = 3222, 
            }, 
        },
        ListTest3 = {
            [0] = {
                [0] = "ListTest3[0][0]",
                [1] = "LIstTest3[0][1]",
            },
            [1] = {
                [0] = "ListTest3[1][0]",
                [1] = "ListTest3{1][1]",
            },
        },
        ListTest4 = {
            [0] = {
                [0] = "list[0][0]",
                [1] = "List[0][1]",
            },
            [1] = {
                [0] = "list[1][0]",
                [1] = "list[1][1]",
            },
        },
        ListTest5 = {
            [0] = {
                [0] = "123",
                [1] = "234",
                [2] = "345",
            },
            [1] = {
                [0] = "234",
                [1] = "345",
                [2] = "456",
            },
        }, 
        DictTest6 = {
            Sublist1 = {
                [0] = "dictTest6[\"sublist1\"][0]",
                [1] = "dictTest6[\"sublist1\"][1]",
            },
            Sublist2 = {
                [0] = "dictTest6[\"sublist2\"][0]",
                [1] = "dictTest6[\"sublist2\"][1]",
            }, 
        }, 
        DictTest7 = {
            Sublist1 = {
                [0] = "zxc",
                [1] = "xcv",
                [2] = "cvb",
            },
            Sublist2 = {
                [0] = "vbn",
                [1] = "bnm",
                [2] = "mnb",
            }, 
        },
        ListTest6 = { 
            [0] = {
                Item1 = "fse",
                Item2 = "fsd", 
            }, 
            [1] = {
                Item1 = "tyu",
                Item2 = "poi", 
            },
        },
        ListTest7 = { 
            [0] = {
                Item0_1 = "fse",
                Item0_2 = "fsd", 
            }, 
            [1] = {
                Item1_1 = "qwe",
                Item1_2 = "qse", 
            },
        },
    },
}

return MultiNestedTypes