package util

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

//DoSign
/**
 * 加签
 *
 * @param map 参数map
 * @param apiKey key
 * @return 加签之后的值
 */
func DoSign(data map[string]any, apiKey string) (string, error) {

	m := make(map[string]any, len(data))
	for k, v := range data {
		s := fmt.Sprintf("%s", v)
		//not null
		if len(s) < 1 {
			continue
		}

		// not sign filed
		if k == "sign" {
			continue
		}
		m[k] = v
	}

	sortStr, err := paramsAsciiAsc(m)
	if err != nil {
		return "", err
	}
	sortStr = fmt.Sprintf("%s%s", apiKey, sortStr)
	md5Str := md5.Sum([]byte(sortStr))
	sortStr = fmt.Sprintf("%x", md5Str)
	sortStr = strings.ToLower(sortStr)
	return sortStr, nil
}

//VerifySign
/**
 * 验签
 *
 * @param map     参数map
 * @param apiKey key
 * @param signStr 接收方生成的sign
 * @return boolean
 */
func VerifySign(data map[string]any, apiKey, signStr string) (bool, error) {
	str, err := DoSign(data, apiKey)
	if err != nil {
		return false, err
	}

	if str == signStr {
		return true, nil
	} else {
		return false, fmt.Errorf("verifySign is failure")
	}
}

//ParamsAsciiAsc
/**
 * 对所传入的参数 进行ASCII码升序排列 最终生成字符串
 *
 * @param paramMap 参数mao
 * @return 字符串
 */
func paramsAsciiAsc(paramMap map[string]any) (string, error) {
	list := make([]string, 0, len(paramMap))
	for k, _ := range paramMap {
		list = append(list, k)
	}

	sort.Strings(list)

	r := strings.Builder{}
	for _, k := range list {
		r.WriteString(fmt.Sprintf("%v%v", k, paramMap[k]))
	}

	return r.String(), nil
}
