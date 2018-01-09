local SimpleTypes = {
    [1001] = {
        Id = 1001,
        NumberTest = 1,
        StringTest = "这是测试字符串1",
        FloatTest = 8.9, 
        DictTest1 = {
            Item1 = 3.33,
            Item2 = 4.4400000000000004,
            Item3 = 5.55, 
        }, 
        DictTest2 = {
            Item1 = 2.22,
            Item2 = 3.33, 
        },
        ListTest1 = {
            [1] = "第一个字符串",
            [2] = "第二个字符串",
        },
        ListTest2 = {
            [1] = "第一个字符串",
            [2] = "second string",
            [3] = "third string",
        }, 
        DictTest3 = { 
            SubDict1 = {
                Item1 = "DictTest3[\"SubDIct1\"][\"Item1\"}",
                Item2 = "DictTest3[\"SubDIct1\"][\"Item2\"}", 
            }, 
            SubDict2 = {
                Item1 = "DictTest3[\"SubDIct2\"][\"Item1\"}",
                Item2 = "DictTest3[\"SubDIct2\"][\"Item2\"}", 
            }, 
        }, 
        DictTest4 = { 
            SubDict1 = {
                item1 = "seges",
                item2 = "fesgs",
                item3 = "fesge", 
            }, 
            SubDict2 = {
                item1 = "123",
                item2 = "121", 
            }, 
        }, 
        DictTest5 = { 
            Subdict1 = {
                item1 = "asd",
                item2 = "sdf", 
            }, 
            Subdict2 = {
                item1 = "qwe",
                item2 = "wer", 
            }, 
        },
        ListTest3 = {
            [1] = {
                [1] = "ListTest3[0][0]",
                [2] = "LIstTest3[0][1]",
            },
            [2] = {
                [1] = "ListTest3[1][0]",
                [2] = "ListTest3{1][1]",
            },
        },
        ListTest4 = {
            [1] = {
                [1] = "list[0][0]",
                [2] = "List[0][1]",
            },
            [2] = {
                [1] = "list[1][0]",
                [2] = "list[1][1]",
            },
        },
        ListTest5 = {
            [1] = {
                [1] = "123",
                [2] = "234",
                [3] = "345",
            },
            [2] = {
                [1] = "234",
                [2] = "345",
                [3] = "456",
            },
        }, 
        DictTest6 = {
            sublist1 = {
                [1] = "dictTest6[\"sublist1\"][0]",
                [2] = "dictTest6[\"sublist1\"][1]",
            },
            sublist2 = {
                [1] = "dictTest6[\"sublist2\"][0]",
                [2] = "dictTest6[\"sublist2\"][1]",
            }, 
        }, 
        DictTest7 = {
            sublist1 = {
                [1] = "zxc",
                [2] = "xcv",
                [3] = "cvb",
            },
            sublist2 = {
                [1] = "vbn",
                [2] = "bnm",
                [3] = "mnb",
            }, 
        },
        ListTest6 = { 
            [1] = {
                item1 = "fse",
                item2 = "fsd", 
            }, 
            [2] = {
                item1 = "tyu",
                item2 = "poi", 
            },
        },
        ListTest7 = { 
            [1] = {
                item0_1 = "fse",
                item0_2 = "fsd", 
            }, 
            [2] = {
                item1_1 = "qwe",
                item1_2 = "qse", 
            },
        },
    },
    [1002] = {
        Id = 1002,
        NumberTest = 2,
        StringTest = "这是测试字符串2",
        FloatTest = 9.1, 
        DictTest1 = {
            Item1 = 4.44,
            Item2 = 5.55,
            Item3 = 6.66, 
        }, 
        DictTest2 = {
            Item1 = 2.22,
            Item2 = 3.33, 
        },
        ListTest1 = {
            [1] = "第一个字符串",
            [2] = "第二个字符串",
        },
        ListTest2 = {
            [1] = "第一个字符串",
            [2] = "second string",
            [3] = "third string",
        }, 
        DictTest3 = { 
            SubDict1 = {
                Item1 = "DictTest3[\"SubDIct1\"][\"Item1\"}",
                Item2 = "DictTest3[\"SubDIct1\"][\"Item2\"}", 
            }, 
            SubDict2 = {
                Item1 = "DictTest3[\"SubDIct2\"][\"Item1\"}",
                Item2 = "DictTest3[\"SubDIct2\"][\"Item2\"}", 
            }, 
        }, 
        DictTest4 = { 
            SubDict1 = {
                item1 = "seges",
                item2 = "fesgs",
                item3 = "fesge", 
            }, 
            SubDict2 = {
                item1 = "123",
                item2 = "121", 
            }, 
        }, 
        DictTest5 = { 
            Subdict1 = {
                item1 = "asd",
                item2 = "sdf", 
            }, 
            Subdict2 = {
                item1 = "qwe",
                item2 = "wer", 
            }, 
        },
        ListTest3 = {
            [1] = {
                [1] = "ListTest3[0][0]",
                [2] = "LIstTest3[0][1]",
            },
            [2] = {
                [1] = "ListTest3[1][0]",
                [2] = "ListTest3{1][1]",
            },
        },
        ListTest4 = {
            [1] = {
                [1] = "list[0][0]",
                [2] = "List[0][1]",
            },
            [2] = {
                [1] = "list[1][0]",
                [2] = "list[1][1]",
            },
        },
        ListTest5 = {
            [1] = {
                [1] = "123",
                [2] = "234",
                [3] = "345",
            },
            [2] = {
                [1] = "234",
                [2] = "345",
                [3] = "456",
            },
        }, 
        DictTest6 = {
            sublist1 = {
                [1] = "dictTest6[\"sublist1\"][0]",
                [2] = "dictTest6[\"sublist1\"][1]",
            },
            sublist2 = {
                [1] = "dictTest6[\"sublist2\"][0]",
                [2] = "dictTest6[\"sublist2\"][1]",
            }, 
        }, 
        DictTest7 = {
            sublist1 = {
                [1] = "zxc",
                [2] = "xcv",
                [3] = "cvb",
            },
            sublist2 = {
                [1] = "vbn",
                [2] = "bnm",
                [3] = "mnb",
            }, 
        },
        ListTest6 = { 
            [1] = {
                item1 = "fse",
                item2 = "fsd", 
            }, 
            [2] = {
                item1 = "tyu",
                item2 = "poi", 
            },
        },
        ListTest7 = { 
            [1] = {
                item0_1 = "fse",
                item0_2 = "fsd", 
            }, 
            [2] = {
                item1_1 = "qwe",
                item1_2 = "qse", 
            },
        },
    },
}

return SimpleTypes