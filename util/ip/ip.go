package ip

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getIpSegRange(userSegIp, offset uint8) (int, int) {
	var ipSegMax uint8 = 255
	netSegIp := ipSegMax << offset
	segMinIp := netSegIp & userSegIp
	segMaxIp := userSegIp&(255<<offset) | ^(255 << offset)
	return int(segMinIp), int(segMaxIp)
}

func getIpSeg1Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 8 {
		segIp, _ := strconv.Atoi(ipSegs[0])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[0])
	return getIpSegRange(uint8(ipSeg), uint8(8-maskLen))
}

func getIpSeg2Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 16 {
		segIp, _ := strconv.Atoi(ipSegs[1])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[1])
	return getIpSegRange(uint8(ipSeg), uint8(16-maskLen))
}

func getIpSeg3Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 24 {
		segIp, _ := strconv.Atoi(ipSegs[2])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[2])
	return getIpSegRange(uint8(ipSeg), uint8(24-maskLen))
}

func getIpSeg4Range(ipSegs []string, maskLen int) (int, int) {
	ipSeg, _ := strconv.Atoi(ipSegs[3])
	segMinIp, segMaxIp := getIpSegRange(uint8(ipSeg), uint8(32-maskLen))
	return segMinIp, segMaxIp
}

func ParseCIDR(cidr string) (uint32, uint32) {
	ip := strings.Split(cidr, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen, _ := strconv.Atoi(strings.Split(cidr, "/")[1])
	seg1MinIp, seg1MaxIp := getIpSeg1Range(ipSegs, maskLen)
	seg2MinIp, seg2MaxIp := getIpSeg2Range(ipSegs, maskLen)
	seg3MinIp, seg3MaxIp := getIpSeg3Range(ipSegs, maskLen)
	seg4MinIp, seg4MaxIp := getIpSeg4Range(ipSegs, maskLen)

	minIP := strconv.Itoa(seg1MinIp) + "." + strconv.Itoa(seg2MinIp) + "." + strconv.Itoa(seg3MinIp) + "." + strconv.Itoa(seg4MinIp)
	maxIP := strconv.Itoa(seg1MaxIp) + "." + strconv.Itoa(seg2MaxIp) + "." + strconv.Itoa(seg3MaxIp) + "." + strconv.Itoa(seg4MaxIp)

	return IP2Long(minIP), IP2Long(maxIP)
}

func IfOverlap(cidr1 string, cidr2 string) bool {
	min1, max1 := ParseCIDR(cidr1)
	min2, max2 := ParseCIDR(cidr2)
	if min1 > max2 || min2 > max1 {
		return false
	}
	return true
}

func IfInclude(cidr string, cidr_included string) bool {
	min, max := ParseCIDR(cidr)
	min_included, max_included := ParseCIDR(cidr_included)
	if min_included >= min && max_included <= max {
		return true
	}
	return false
}

func IP2Long(ipstr string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)

	return
}

func Long2IP(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

func GetMaskBit(mask string) int {
	masks := strings.Split(mask, ".")
	bit_count := 0
	for _, _m := range masks {
		i_m, _ := strconv.Atoi(_m)
		data := byte(i_m)
		var a byte
		for i := 0; i < 8; i++ {
			a = data
			data <<= 1
			data >>= 1

			switch a {
			case data:
				continue
			default:
				bit_count += 1
			}

			data <<= 1
		}
	}
	return bit_count
}

func GetCIDR(gateway, mask string) string {
	mask_bit := GetMaskBit(mask)
	min_ip := Long2IP(IP2Long(gateway) - 1)
	return fmt.Sprintf("%s/%d", min_ip, mask_bit)
}
