package api

import ()

var REGION_DEFAULT_ZONE = map[int]int{
	666888:  666888,
	1000001: 5001,
	1000002: 1,
	1000003: 7001,
	1000004: 3001,
	1000005: 6001,
	1000006: 110100,
	1000007: 8001,
	1000008: 1001,
	1000009: 8100,
}

var REGION2ZONES = map[int][]uint32{
	666888:  []uint32{666888},
	1000001: []uint32{4001, 5001, 9001},
	1000002: []uint32{1},
	1000003: []uint32{7001},
	1000004: []uint32{3001},
	1000005: []uint32{6001},
	1000006: []uint32{110100},
	1000007: []uint32{8001},
	1000008: []uint32{1001},
	1000009: []uint32{8100},
}

var ZONE2VALIDSUBNET = map[int]string{
	99999:  "1",  //10.1
	2001:   "0",  //10.0
	3001:   "1",  //10.1
	4001:   "2",  //10.2
	5001:   "3",  //10.3
	6001:   "12", //10.12
	7001:   "14", //10.14
	8001:   "16", //10.16
	8100:   "24", //10.24
	9001:   "20", //10.20
	110100: "18", //10.18
	666888: "22", // 10.22
}

var names_code = map[string]map[string]int{
	"Zone": {
		"pre":          666888,
		"cn-bj1-01":    1001,
		"cn-bj2-02":    4001,
		"cn-bj2-03":    5001,
		"cn-bj2-04":    9001,
		"cn-zj-01":     1,
		"cn-sh-01":     8001,
		"cn-gd-01":     2001,
		"cn-gd-02":     7001,
		"hk-01":        3001,
		"us-ca-01":     6001,
		"cn-sh2-01":    8100,
		"cn-inspur-01": 110100,
	},
	"ChargeType": {
		"Dynamic": 1,
		"Month":   2,
		"Year":    3,
		"Trial":   5,
		"Day":     8,
	},
	"Region2DefaultZone": {
		"666888":  666888,
		"1000001": 5001,
		"1000002": 1,
		"1000003": 7001,
		"1000004": 3001,
		"1000005": 6001,
		"1000006": 110100,
		"1000007": 8001,
		"1000008": 1001,
		"1000009": 8100,
	},
}

func getCodeByName(group, name string) int {
	_map, ok := names_code[group]
	if !ok {
		return 0
	}
	code, ok := _map[name]
	if !ok {
		return 0
	}
	return code
}

func getNameByCode(group string, code int) string {
	_map, ok := names_code[group]
	if !ok {
		return ""
	}
	for k, v := range _map {
		if v == code {
			return k
		}
	}
	return ""
}
