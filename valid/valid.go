package valid

import (
	"net"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// IsSlice 是否为 slice
func IsSlice(i any) bool {
	return reflect.ValueOf(i).Kind() == reflect.Slice
}

// IsMap 是否为 map
func IsMap(i any) bool {
	return reflect.ValueOf(i).Kind() == reflect.Map
}

// IsArray 是否为 array
func IsArray(i any) bool {
	return reflect.ValueOf(i).Kind() == reflect.Array
}

// IsEmail 是否为 email
func IsEmail(input string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, input)
	return match && err == nil
}

// IsBankCardNo 验证是否为大陆银行卡号
func IsBankCardNo(cardNumber string) bool {
	if len(cardNumber) != 16 && len(cardNumber) != 19 {
		return false
	}
	var cardArr []int
	for _, c := range cardNumber {
		if c < '0' || c > '9' {
			return false
		}
		cardArr = append(cardArr, int(c-'0'))
	}
	if len(cardArr) == 16 {
		sum := 0
		for i := len(cardArr) - 1; i >= 0; i-- {
			if i%2 == 0 {
				cardArr[i] *= 2
				if cardArr[i] > 9 {
					cardArr[i] -= 9
				}
			}
			sum += cardArr[i]
		}
		return sum%10 == 0
	} else {
		sum := 0
		for i := len(cardArr) - 1; i >= 0; i-- {
			if (len(cardArr)-i)%2 == 0 {
				cardArr[i] *= 2
				if cardArr[i] > 9 {
					cardArr[i] -= 9
				}
			}
			sum += cardArr[i]
		}
		return sum%10 == 0
	}
}

// IsIDCard 验证身份证号(18 或 15 位)
func IsIDCard(idCard string) bool {
	if len(idCard) != 15 && len(idCard) != 18 {
		return false
	}
	if len(idCard) == 18 {
		return isIDCard18(idCard)
	} else {
		return isIDCard15(idCard)
	}
}

// isIDCard18 验证 18 位身份证号
func isIDCard18(idCard string) bool {
	// 18位身份证号码正则表达式，根据规则来编写
	regExp := "^[1-9]\\d{5}(19|20)\\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$"
	// 利用正则表达式匹配身份证号码
	match, err := regexp.MatchString(regExp, idCard)
	if err != nil {
		// 匹配过程出错，返回false
		return false
	}
	if !match {
		// 身份证号码不符合规则，返回false
		return false
	}
	// 解析身份证号码中的年、月、日
	year, _ := strconv.Atoi(idCard[6:10])
	month, _ := strconv.Atoi(idCard[10:12])
	day, _ := strconv.Atoi(idCard[12:14])
	// 判断年份是否合法
	if year < 1900 || year > time.Now().Year() {
		return false
	}
	// 判断月份是否合法
	if month < 1 || month > 12 {
		return false
	}
	// 判断日期是否合法
	if day < 1 || day > 31 {
		return false
	}
	// 对身份证号码的最后一位进行校验
	// 根据身份证号码的规则，最后一位可能是数字0-9，也可能是字符X（表示10）
	// 将字符X转换成数字10进行校验
	lastChar := idCard[len(idCard)-1]
	var lastNum int
	if lastChar == 'X' || lastChar == 'x' {
		lastNum = 10
	} else {
		lastNum, _ = strconv.Atoi(string(lastChar))
	}
	// 对身份证号码的前17位进行加权和校验
	// 加权系数，根据规则固定
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 计算加权和
	sum := 0
	for i := 0; i < len(weights); i++ {
		num, _ := strconv.Atoi(string(idCard[i]))
		sum += num * weights[i]
	}

	// 计算校验码
	checkCode := sum % 11
	checkCodeMap := map[int]string{
		0:  "1",
		1:  "0",
		2:  "10", // 身份证最后一位是X，加权求和是10
		3:  "9",
		4:  "8",
		5:  "7",
		6:  "6",
		7:  "5",
		8:  "4",
		9:  "3",
		10: "2",
	}
	// 校验身份证号码的最后一位
	return checkCodeMap[checkCode] == strconv.Itoa(lastNum)
}

// isIDCard15 验证 15 位身份证号
func isIDCard15(idCard string) bool {
	// 验证是否为15位数字
	if match, _ := regexp.MatchString(`^\d{15}$`, idCard); !match {
		return false
	}

	// 将身份证号前两位转换成省份代码
	provinceCode, err := strconv.Atoi(idCard[:2])
	if err != nil || provinceCode < 11 || provinceCode > 91 {
		return false
	}

	// 验证生日是否正确
	year := strconv.Itoa(1900 + int(idCard[6]-'0')*10 + int(idCard[7]-'0'))
	month := idCard[8:10]
	day := idCard[10:12]
	if match, _ := regexp.MatchString(`^(19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])$`, year+month+day); !match {
		return false
	}

	return true
}

// IsIPv4 是否为 ipv4 地址
func IsIPv4(i string) bool {
	ip := net.ParseIP(i)
	return ip != nil && ip.To4() != nil
}

// IsIPv6 是否为 ipv6 地址
func IsIPv6(i string) bool {
	ip := net.ParseIP(i)
	return ip != nil && ip.To4() == nil
}

// IsMobile 验证是否为手机号码
func IsMobile(phone string) bool {
	match, _ := regexp.MatchString(`^1[3456789]\d{9}$`, phone)
	return match
}
