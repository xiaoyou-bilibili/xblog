// Package database @Description mongo常用的常量
// @Author 小游
// @Date 2021/04/10
package database

// MongoNoResult mongo没有找到数据时的报错内容
const MongoNoResult = "mongo: no documents in result"

// FieldSum mongo统计字段的值
type FieldSum struct {
	ID    string `bson:"_id"`
	Total int64  `bson:"total"`
}
