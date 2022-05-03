// Package module
// @Description
// @Author 小游
// @Date 2021/04/06
package module

import (
	"context"
	"github.com/dop251/goja"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"xBlog/internal/app/common"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// RegisterDatabase 注册数据库操作对象
func RegisterDatabase(vm *goja.Runtime) *goja.Object {
	blog := vm.NewObject()
	_ = blog.Set("newDb", NewDb)
	_ = blog.Set("regex", DatabaseRegex)
	_ = blog.Set("adminDeleteObject", common.AdminDeleteObjectData)
	_ = blog.Set("adminDelete", common.AdminDeleteData)
	return blog
}

// 当前上下文对象
var ctx = context.Background()

// Db 这个封装了所有的函数，js直接调用即可
type Db struct {
	Collection *mongo.Collection //集合名
	Name       string
}

// NewDb 初始化一个新的数据库连接
func NewDb(collection string) *Db {
	database.DbReconnect()
	return &Db{Collection: database.DB.Collection(collection)}
}

// FindOne 查询单条数据
func (db *Db) FindOne(option map[string]interface{}, callback func(err interface{}, result interface{})) {
	op := options.FindOne()
	// 最终查询结果
	var data = make(map[string]interface{})
	// 过滤条件
	var filter interface{}
	// 判断是否有过滤条件
	if _, ok := option["filter"]; ok {
		filter = option["filter"]
	} else {
		filter = bson.M{}
	}
	// 设置sort
	if _, ok := option["sort"]; ok {
		op.SetSort(option["sort"])
	}
	// 设置跳过
	if _, ok := option["skip"]; ok {
		op.SetSkip(option["skip"].(int64))
	}
	// 设置投影
	if _, ok := option["projection"]; ok {
		op.SetProjection(option["projection"])
	}
	// 执行查询
	res := db.Collection.FindOne(ctx, filter)
	// 解析数据
	if err := res.Decode(data); err != nil {
		callback(err.Error(), nil)
		return
	}
	// 如果无误，那么就触发回调函数
	callback(nil, data)
}

// FindMany 查询多条数据
func (db *Db) FindMany(option map[string]interface{}, callback func(err interface{}, result interface{})) {
	var data []map[string]interface{}
	op := options.Find()
	var filter interface{}
	if _, ok := option["filter"]; ok {
		filter = option["filter"]
	} else {
		filter = bson.M{}
	}
	// 判断是否有limit
	if _, ok := option["limit"]; ok {
		op.SetLimit(option["limit"].(int64))
	}
	// 设置sort
	if _, ok := option["sort"]; ok {
		op.SetSort(option["sort"])
	}
	// 设置跳过
	if _, ok := option["skip"]; ok {
		op.SetSkip(option["skip"].(int64))
	}
	// 设置投影
	if _, ok := option["projection"]; ok {
		op.SetProjection(option["projection"])
	}
	//开始查找
	cursor, err := db.Collection.Find(ctx, filter, op)
	if err != nil {
		callback(err.Error(), nil)
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	//取出cursor所有数据然后导出到result中
	err = cursor.All(ctx, &data)
	if err != nil {
		callback(err.Error(), nil)
	} else {
		callback(nil, data)
	}
}

// GetCount 获取记录数目(失败直接返回0)
func (db *Db) GetCount(filter interface{}) int64 {
	if res, err := db.Collection.CountDocuments(ctx, filter); err != nil {
		return 0
	} else {
		return res
	}
}

// Paginate 分页查询
func (db *Db) Paginate(option map[string]interface{}, now int, limit int, callback func(err interface{}, page int, total int, result interface{})) {
	var filter interface{}
	if _, ok := option["filter"]; ok {
		filter = option["filter"]
	} else {
		filter = bson.M{}
	}
	// 获取数据总数
	total, err := db.Collection.CountDocuments(ctx, filter)
	if err != nil {
		callback(err.Error(), 0, 0, nil)
		return
	}
	//获取总页数
	page := int64(math.Ceil(float64(total) / float64(limit)))
	option["skip"] = int64((now - 1) * limit)
	option["limit"] = int64(limit)
	// 查询数据
	db.FindMany(option, func(err interface{}, result interface{}) {
		if err == nil {
			callback(nil, int(page), int(total), result)
		} else {
			callback(err, 0, 0, nil)
		}
	})
}

// InsertOne 插入数据
func (db *Db) InsertOne(data interface{}, callback func(err interface{}, result interface{})) {
	result, err := db.Collection.InsertOne(ctx, data)
	if err != nil {
		callback(err.Error(), nil)
	} else {
		callback(nil, result)
	}
}

// InsertOneIncrease 插入值并自动更新ID
func (db *Db) InsertOneIncrease(data interface{}, key string, callback func(err interface{}, result interface{})) {
	// 最终查询结果
	var result interface{}
	op := options.FindOne()
	// 设置sort
	op.SetSort(bson.M{key: -1})
	op.SetProjection(bson.M{key: 1})
	// 执行查询
	res := db.Collection.FindOne(ctx, bson.M{}, op)
	// 解析数据
	if err := res.Decode(&result); err != nil {
		callback(err.Error(), nil)
		return
	}
	// 如果无误，那么获取一下最大值
	var id int32 = 0
	var ok bool
	// 获取当前最大值id
	if result != nil {
		if id, ok = result.(bson.D).Map()[key].(int32); !ok {
			id = 0
		}
	}
	// 插入数据
	r, err := db.Collection.InsertOne(ctx, data)
	if err != nil {
		callback(err.Error(), nil)
		return
	}
	if r == nil {
		callback("无法获取插入返回值", nil)
		return
	}
	id++
	// 找出新插入的值并自动更新id
	insRes, err := db.Collection.UpdateOne(ctx, bson.M{"_id": r.InsertedID}, bson.M{"$set": bson.M{key: id}})
	if err != nil {
		callback(err.Error(), nil)
	} else {
		callback(nil, insRes)
	}
}

// UpdateOne 更新一条记录
func (db *Db) UpdateOne(option map[string]interface{}, callback func(err interface{}, result interface{})) {
	op := options.Update()
	var filter interface{}
	var update interface{}
	// upsert表示如果这条记录不存在时则插入
	if data, ok := option["upsert"]; ok {
		op.SetUpsert(tools.Interface2Bool(data))
	}
	// 判断过滤条件
	if _, ok := option["filter"]; ok {
		filter = option["filter"]
	} else {
		filter = bson.M{}
	}
	// 判断更新条件
	if _, ok := option["update"]; ok {
		update = option["update"]
	} else {
		update = bson.M{}
	}
	result, err := db.Collection.UpdateOne(ctx, filter, update, op)
	if err != nil {
		callback(err.Error(), nil)
	} else {
		callback(nil, result)
	}
}

// UpdateMany 更新多条记录
func (db *Db) UpdateMany(option map[string]interface{}, callback func(err interface{}, result interface{})) {
	var filter interface{}
	var update interface{}
	// 判断过滤条件
	if _, ok := option["filter"]; ok {
		filter = option["filter"]
	} else {
		filter = bson.M{}
	}
	// 判断更新条件
	if _, ok := option["update"]; ok {
		update = option["update"]
	} else {
		update = bson.M{}
	}
	result, err := db.Collection.UpdateMany(ctx, filter, update)
	if err != nil {
		callback(err.Error(), nil)
	} else {
		callback(nil, result)
	}
}

// DeleteOne 删除一条数据
func (db *Db) DeleteOne(filter map[string]interface{}, callback func(err interface{}, result interface{})) {
	// 判断过滤条件
	if result, err := db.Collection.DeleteOne(ctx, filter); err != nil {
		callback(err.Error(), nil)
	} else {
		callback(nil, result)
	}
}

// DeleteMany 删除多条数据
func (db *Db) DeleteMany(filter map[string]interface{}, callback func(err interface{}, result interface{})) {
	// 判断过滤条件
	if result, err := db.Collection.DeleteMany(ctx, filter); err != nil {
		callback(err.Error(), nil)
	} else {
		callback(nil, result)
	}
}

// DatabaseRegex 数据库正则替换
func DatabaseRegex(regex string) primitive.Regex {
	return db.Regex(regex)
}
