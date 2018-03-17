package paramcheck

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var MapRegion2InvaildNet = map[int][]string{

	0:       []string{"10.254.0.0/16"},                                     //全网禁用
	1000001: []string{"10.19.192.0/24", "10.9.192.0/24", "10.10.192.0/24"}, //北京二
	1000003: []string{"10.13.192.0/24"},                                    //广东
	1000004: []string{"10.8.192.0/24"},                                     //香港
	1000005: []string{"10.11.192.0/24"},                                    //加州
	1000007: []string{"10.15.192.0/24"},                                    //上海一
	1000009: []string{"10.23.192.0/24"},                                    //上海二
	//预发布，测试使用
	666888: []string{"10.254.0.0/16", "10.19.192.0/24", "10.9.192.0/24", "10.10.192.0/24", "10.13.192.0/24", "10.8.192.0/24", "10.11.192.0/24", "10.15.192.0/24", "10.23.192.0/24"},
}

const (
	NAME_LENGTH = 63
)

func CheckExtendInfo(value string, defaultValue string, illegalMode bool) (string, error) {
	lenLimit := NAME_LENGTH
	value = strings.TrimSpace(value)
	defaultValue = strings.TrimSpace(defaultValue)

	if value == "" {
		if defaultValue == "" {
			return "", fmt.Errorf("")
		}
		value = defaultValue
	}
	if len(value) > lenLimit || !matchNamePattern(value, illegalMode) {
		return value, fmt.Errorf("")
	} else {
		return value, nil
	}
}

/*
 * illegalMode = true, 匹配到非法字符，返回 false, 否则 true
 * illegalMode = false, 匹配合法字符失败，返回 false, 否则 true
 */
func matchNamePattern(name string, illegalMode bool) bool {
	illegalChars := `^.*[\'\"\\].*$`
	legalChars := `^[A-Za-z0-9-_\.\p{Han}]+$`

	var re *regexp.Regexp = nil
	var err error = nil
	if illegalMode {
		re, err = regexp.Compile(illegalChars)
	} else {
		re, err = regexp.Compile(legalChars)
	}
	if err != nil {
		return false
	}
	// re.MatchString(name) ^ illegalMode
	return re.MatchString(name) != illegalMode
}
func CheckPrivteIP(network string) bool {
	// 匹配 10.0.0.0 --10.255.255.255
	r1 := `^10(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d|10)){3}$`
	// 匹配 172.16.0.0—172.31.255.255
	r2 := `^172\.([1][6-9]|[2]\d|3[01])(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){2}$`
	// 匹配 192.168.0.0-192.168.255.255
	r3 := `^192\.168(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){2}$`
	info := strings.Split(network, "/")
	if len(info) != 2 {
		return false
	}
	ip := info[0]
	len, err := strconv.Atoi(info[1])
	if err != nil || len <= 0 || len > 32 {
		return false
	}
	match1, _ := regexp.MatchString(r1, ip)
	match2, _ := regexp.MatchString(r2, ip)
	match3, _ := regexp.MatchString(r3, ip)
	return match1 || match2 || match3

}

func IsInArray(item interface{}, array interface{}) bool {
	targetValue := reflect.ValueOf(array)
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == item {
				return true
			}
		}
	}
	return false
}

//将点分10进制ip转化为2进制
func converToBin(ips []string) int {
	value, _ := strconv.Atoi(ips[0])
	value = value << 24
	var i uint32
	for i = 1; i < uint32(len(ips)); i++ {
		tmp, _ := strconv.Atoi(ips[1])
		tmp = tmp << (24 - i*8)
		value = value | tmp

	}
	return value

}

//判断某个ip是否在某网段中
func IsInRange(ip string, net string) bool {
	ips := strings.Split(ip, ".")
	ipAddr := converToBin(ips)
	tmp_type := strings.Split(net, "/")[1]
	Type, _ := strconv.Atoi(tmp_type)
	//子网掩码
	mask := 0xFFFFFFFF << (32 - uint32(Type))
	//得到网段
	netIPs := strings.Split(net, "/")[0]
	netIp := strings.Split(netIPs, ".")
	netAddr := converToBin(netIp)
	return (mask & ipAddr) == (mask & netAddr)

}

//在某个region下输入的网段是否是被禁止的 -
func NetIsValid(mynet string, regionId int) error {
	ivaildnets, err := MapRegion2InvaildNet[regionId]
	allregion := MapRegion2InvaildNet[0] //把全网禁止的加入进来
	if err != true {
		return errors.New("region is not vaild !")
	}
	for _, r := range allregion {
		ivaildnets = append(ivaildnets, r)
	}
	for _, net := range ivaildnets {
		ip := strings.Split(mynet, "/")[0]
		flag := IsInRange(ip, net)
		if flag == true {
			msg := fmt.Sprintf("%s", ivaildnets)
			return errors.New(msg)
		}

	}
	return nil

}
