// Package tools @Description 类型转换相关
// @Author 小游
// @Date 2021/04/10
package tools

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
	"time"
)

// Interface2String 接口转string
func Interface2String(data interface{}) string {
	if data == nil {
		return ""
	}
	res, ok := data.(string)
	if !ok {
		return ""
	} else {
		return res
	}
}

// Interface2Int 接口转int
func Interface2Int(data interface{}) int32 {
	if data == nil {
		return 0
	}
	res, ok := data.(int32)
	if ok {
		return res
	} else {
		return 0
	}
}

// Interface2Bool 接口转bool
func Interface2Bool(data interface{}) bool {
	if data == nil {
		return false
	}
	res, ok := data.(bool)
	if ok {
		return res
	} else {
		return false
	}
}

// Time2String 时间转string
func Time2String(times time.Time, showHour bool) string {
	var cstZone = time.FixedZone("CST", 8*3600)
	if showHour {
		return times.In(cstZone).Format("2006-01-02 15:04:05")
	}
	return times.In(cstZone).Format("2006-01-02")
}

// Str2Time string转时间
func Str2Time(times string, hour bool) time.Time {
	if loc, err := time.LoadLocation("Local"); err == nil {
		var theTime time.Time
		// 是否为带时间的转换
		if hour {
			theTime, err = time.ParseInLocation("2006-01-02 15:04:05", times, loc)
		} else {
			theTime, err = time.ParseInLocation("2006-01-02", times, loc)
		}
		// 判断是否转换成功
		if err == nil {
			return theTime
		}
	}
	return time.Now()
}

// Primitive2Struct primitive.D数据转结构体
func Primitive2Struct(data interface{}, result interface{}) error {
	bsonBytes, err := bson.Marshal(data)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(bsonBytes, result)
	if err != nil {
		return err
	}
	return nil
}

// Struct2Bson 结构体数据转bson.m
func Struct2Bson(data interface{}, result interface{}) error {
	return Primitive2Struct(data, result)
}

// Str2Int 字符串转int
func Str2Int(dataS string) int {
	data, err := strconv.Atoi(dataS)
	if err != nil {
		return 0
	}
	return data
}

// Str2ObjectID string转objectID(如果转换不了)
func Str2ObjectID(data string) primitive.ObjectID {
	if id, err := primitive.ObjectIDFromHex(data); err == nil {
		return id
	} else {
		return primitive.ObjectID{}
	}
}

// Int2Str int转字符串
func Int2Str(data int) string {
	return strconv.Itoa(data)
}

// IsInIntArray 判断int 数组中是否存在某个值
func IsInIntArray(data []int, key int) bool {
	for _, v := range data {
		if key == v {
			return true
		}
	}
	return false
}

// IsInStringArray 判断string 数组中是否存在某个值
func IsInStringArray(data []string, key string) bool {
	for _, v := range data {
		if key == v {
			return true
		}
	}
	return false
}

// Interface2Float 接口转float64
func Interface2Float(data interface{}) float64 {
	if data == nil {
		return 0
	}
	res, ok := data.(float64)
	if !ok {
		return 0
	} else {
		return res
	}
}

// InterfaceFloat2Int float64接口转int
func InterfaceFloat2Int(data interface{}) int {
	if data == nil {
		return 0
	}
	res, ok := data.(float64)
	if !ok {
		return 0
	} else {
		return int(res)
	}
}

// InterfaceFloat2String float64接口转字符串
func InterfaceFloat2String(data interface{}) string {
	if data == nil {
		return ""
	}
	res, ok := data.(float64)
	if !ok {
		return ""
	} else {
		return strconv.Itoa(int(res))
	}
}

// String2IntArray 把字符串批量转换为int数组
func String2IntArray(data string, sep string) []int {
	strArray := strings.Split(data, sep)
	var returnData []int
	for _, v := range strArray {
		returnData = append(returnData, Str2Int(v))
	}
	return returnData
}

// String2ObjectIDArray 把字符串批量转换为ObjectID
func String2ObjectIDArray(data string, sep string) []primitive.ObjectID {
	strArray := strings.Split(data, sep)
	var returnData []primitive.ObjectID
	for _, v := range strArray {
		if re := Str2ObjectID(v); !re.IsZero() {
			returnData = append(returnData, re)
		}
	}
	return returnData
}

// Str2Bool 字符串转bool值
func Str2Bool(data string) bool {
	if data == "true" {
		return true
	}
	return false
}
