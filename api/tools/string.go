// Package tools @Description 和字符串处理有关的
// @Author 小游
// @Date 2021/04/10
package tools

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/axgle/mahonia"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// FindMatch 寻找匹配的正则表达式
func FindMatch(reg string, content string) string {
	r, err := regexp.Compile(reg)
	if err != nil {
		return ""
	}
	if data := r.FindStringSubmatch(content); data != nil && len(data) > 1 {
		return data[1]
	}
	return ""
}

// 把Unicode转string编码
func u2s(form string) (to string, err error) {
	if strings.Index(form, "\\u") == -1 {
		return form, nil
	}
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		_ = binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}

// FindAllMatch 寻找所有的匹配
func FindAllMatch(content string, reg string, isChange bool) []string {
	r, err := regexp.Compile(reg)
	if err != nil {
		return nil
	}
	//寻找所有匹配的
	list := r.FindAllStringSubmatch(content, -1)
	//创建数组
	lists := new([]string)
	for _, v := range list {
		if isChange {
			str, err := u2s(v[1])
			if err != nil {
				*lists = append(*lists, "")
			} else {
				*lists = append(*lists, str)
			}
		} else {
			*lists = append(*lists, v[1])
		}
	}
	return *lists
}

// ReplaceUnicode 替换一些常见的Unicode字符
func ReplaceUnicode(str string) string {
	str = strings.Replace(str, "\\u0026", "&", -1)
	str = strings.Replace(str, "\\n", "<br>", -1)
	return str
}

// Float2String float64类型的数据转换为string
func Float2String(num float64) string {
	return strconv.Itoa(int(num))
}

// Unix2Time unix时间戳转换为字符串
func Unix2Time(times int) string {
	tm := time.Unix(int64(times), 0)
	return tm.Format("2006-01-02 15:04:05")
}

// ReplaceXss 替换<>避免xss攻击
func ReplaceXss(text string) string {
	text = strings.Replace(text, "<", "&lt;", -1)
	text = strings.Replace(text, ">", "&gt;", -1)
	return text
}

// CoverGBKToUTF8 把GBK转换为utf-8编码
func CoverGBKToUTF8(src string) string {
	return mahonia.NewDecoder("gbk").ConvertString(src)
}
