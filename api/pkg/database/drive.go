// Package database @Description MongoDB核心驱动封装
// @Author 小游
// @Date 2021/04/10
package database

import (
	"context"
	"errors"
	"github.com/Unknwon/goconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Db 这个封装了查询需要的信息
type Db struct {
	Collection *mongo.Collection //集合名
	Limit      int64             //限制多少位(option相关)
	Skip       int64             //跳过多少条数据(option相关)
	Sort       interface{}       //排序条件(option相关)
	Filter     bson.M            //条件过滤(查找相关)
	Result     interface{}       //执行某个任务后的结果
	update     bson.M            //更新内容(更新相关)
	Pipeline   []bson.M          //聚合查询有关的参数
	Upsert     bool              //不存在这条记录则插入
	Group      bson.M            // 聚合查询的分组操作
	Projection bson.M            // 设置投影操作符来返回指定的键
}

//------数据库函数初始化

// DB 数据库DB对象
var DB *mongo.Database

//context对象
var ctx = context.Background()

// GetAppPath 获取系统的路径
func GetAppPath() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	// 如果是直接run，需要这里设置一下运行模式
	if os.Getenv("XBLOG_RUNMODE") == "1" {
		dir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	//fmt.Println(strings.Replace(dir, "\\", "/", -1) + "/")
	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1) + "/"
}

// GetConfig 获取配置文件
func GetConfig(module string) map[string]string {
	path := GetAppPath()
	config, err := goconfig.LoadConfigFile(path + "configs/app.ini") //加载配置文件
	if err != nil {
		return nil
	}
	glob, _ := config.GetSection(module) //读取全部mysql配置
	return glob
}

// DbInit 数据库初始化
func DbInit() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() //养成良好的习惯，在调用WithTimeout之后defer cancel()
	//获取MongoDB数据库配置
	config := GetConfig("mongo")
	//fmt.Println(config)
	// 初始mongo连接池
	MongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config["url"])) // change the URI when you need
	if err != nil {
		log.Fatal("数据库连接失败!")
		os.Exit(3)
	}
	// 检查连接
	if err = MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("数据库连接失败！")
		os.Exit(3)
	}
	//选择数据库
	DB = MongoClient.Database(config["name"])
}

// DbClose 关闭数据库
func DbClose() {
	_ = DB.Client().Disconnect(ctx)
}

// DbReconnect 数据库重连函数
func DbReconnect() {
	//获取client对象，并判断是否可用
	if DB != nil {
		MongoClient := DB.Client()
		if MongoClient.Ping(ctx, readpref.Primary()) == nil {
			return
		} else {
			// 关闭旧的连接
			DbClose()
		}
	}
	DbInit()
}

// NewDb find初始化(相当于初始化对象)
func NewDb(collection string) *Db {
	//先重连数据库
	DbReconnect()
	return &Db{Collection: DB.Collection(collection), update: bson.M{}, Pipeline: []bson.M{}}
}

//--------------- option相关

// SetLimit 设置limit
func (db *Db) SetLimit(limit int64) *Db {
	db.Limit = limit
	return db
}

// SetSkip 设置skip数目
func (db *Db) SetSkip(skip int64) *Db {
	db.Skip = skip
	return db
}

// SetSort 设置排序
func (db *Db) SetSort(sort interface{}) *Db {
	db.Sort = sort
	return db
}

// SetUpsert 设置upsert 如果这条记录不存在那么就插入这条记录
func (db *Db) SetUpsert(choose bool) *Db {
	db.Upsert = choose
	return db
}

// SetProjection 设置投影
func (db *Db) SetProjection(projection bson.M) *Db {
	db.Projection = projection
	return db
}

//--------------- filter相关

// SetFilter 设置and条件
func (db *Db) SetFilter(filter bson.M) *Db {
	db.Filter = filter
	return db
}

// AddFilter 添加过滤条件
func (db *Db) AddFilter(key string, value interface{}) *Db {
	if db.Filter == nil {
		db.Filter = bson.M{}
	}
	db.Filter[key] = value
	return db
}

// OR 设置or的条件
func (db *Db) OR(or []bson.M) *Db {
	if db.Filter == nil {
		db.Filter = bson.M{}
	}
	db.Filter["$or"] = or
	return db
}

// ---- update相关的操作

// Set 修改某个字段的值
func (db *Db) Set(set bson.M) *Db {
	db.update["$set"] = set
	return db
}

// Inc 对某个数字字段进行加减
func (db *Db) Inc(inc bson.M) *Db {
	db.update["$inc"] = inc
	return db
}

// Unset 删除某个字段
func (db *Db) Unset(unset bson.M) *Db {
	db.update["$unset"] = unset
	return db
}

// Push 数组类型添加字段
func (db *Db) Push(push bson.M) *Db {
	db.update["$push"] = push
	return db
}

// PushAll 数组类型添加多个字段
func (db *Db) PushAll(pushAll bson.M) *Db {
	db.update["$pushAll"] = pushAll
	return db
}

// AddToSet 加一个值到数组内，而且只有当这个值在数组中不存在时才增加。
func (db *Db) AddToSet(AddToSet bson.M) *Db {
	db.update["$addToSet"] = AddToSet
	return db
}

// Pop 删除数组内某一个值(1是第一个值 -1是最后一个值)
func (db *Db) Pop(pop bson.M) *Db {
	db.update["$pop"] = pop
	return db
}

// Pull 从数组中删除一个值
func (db *Db) Pull(pull bson.M) *Db {
	db.update["$pull"] = pull
	return db
}

// PullAll 从数组中删除多个值
func (db *Db) PullAll(pullAll bson.M) *Db {
	db.update["$pullAll"] = pullAll
	return db
}

// Rename 对字段进行重命名
func (db *Db) Rename(rename bson.M) *Db {
	db.update["$rename"] = rename
	return db
}

// SetUpdate 直接设置
func (db *Db) SetUpdate(update bson.M) *Db {
	db.update = update
	return db
}

//---------数据库增删改查

// InsertOne 插入单条数据
func (db *Db) InsertOne(data interface{}) error {
	result, err := db.Collection.InsertOne(ctx, data)
	db.Result = result
	return err
}

// InsertOneIncrease 某个值id自动递增
func (db *Db) InsertOneIncrease(data interface{}, key string) (int, error) {
	//首先获取最大的id
	db.Sort = bson.M{key: -1}
	var result interface{}
	// 设置投影找出最大的值(查找的时候需要强制清空，要不然会导致找不到对应的值)
	if err := db.SetFilter(bson.M{}).SetProjection(bson.M{key: 1}).FindOne(&result); err != nil {
		if err.Error() != MongoNoResult {
			return 0, err
		}
	}
	var id int32 = 0
	var ok bool
	// 获取当前最大值id
	if result != nil {
		if id, ok = result.(bson.D).Map()[key].(int32); !ok {
			id = 0
		}
	}
	//插入数据
	if err := db.InsertOne(data); err != nil {
		return 0, err
	}
	id++
	// 找出新插入的值并自动更新id
	db.SetFilter(bson.M{"_id": db.Result.(*mongo.InsertOneResult).InsertedID}).Set(bson.M{key: id})
	//自动更新数据，如果更新失败就删除这个数据
	if err := db.UpdateOne(); err != nil {
		_ = db.DeleteOne()
		return 0, err
	}
	return int(id), nil
}

// InsertMany 插入多条数据
func (db *Db) InsertMany(data []interface{}) error {
	result, err := db.Collection.InsertMany(ctx, data)
	db.Result = result
	return err
}

// DeleteOne 删除单条数据
func (db *Db) DeleteOne() error {
	result, err := db.Collection.DeleteOne(ctx, db.Filter)
	db.Result = result
	return err
}

// DeleteMore 删除多条数据
func (db *Db) DeleteMore() error {
	result, err := db.Collection.DeleteMany(ctx, db.Filter)
	db.Result = result
	return err
}

// UpdateOne 更改单条数据
func (db *Db) UpdateOne() error {
	option := options.Update()
	option.SetUpsert(db.Upsert)
	result, err := db.Collection.UpdateOne(ctx, db.Filter, db.update, option)
	db.Result = result
	return err
}

// UpdateMany 更改多个字段
func (db *Db) UpdateMany() error {
	result, err := db.Collection.UpdateMany(ctx, db.Filter, db.update)
	db.Result = result
	return err
}

// FindOne 查询可以参考 https://www.runoob.com/mongodb/mongodb-query.html
//查询单条数据
func (db *Db) FindOne(result interface{}) error {
	option := options.FindOne()
	// 设置投影
	if db.Projection != nil {
		option.SetProjection(db.Projection)
	}
	//设置跳过
	if db.Skip != 0 {
		option.SetSkip(db.Skip)
	}
	//设置排序
	if db.Sort != nil {
		option.SetSort(db.Sort)
	}
	res := db.Collection.FindOne(ctx, db.Filter, option)
	return res.Decode(result)
}

// FindMore 查询多条数据
func (db *Db) FindMore(result interface{}) error {
	option := options.Find()
	// 设置投影
	if db.Projection != nil {
		option.SetProjection(db.Projection)
	}
	if db.Skip != 0 {
		option.SetSkip(db.Skip)
	}
	//设置排序
	if db.Sort != nil {
		option.SetSort(db.Sort)
	}
	//设置limit
	if db.Limit != 0 {
		option.SetLimit(db.Limit)
	}
	//开始查找
	cursor, err := db.Collection.Find(ctx, db.Filter, option)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	//取出cursor所有数据然后导出到result中
	err = cursor.All(ctx, result)
	if err != nil {
		return err
	}
	return nil
}

// GetCount 获取记录的数目
func (db *Db) GetCount() (int64, error) {
	return db.Collection.CountDocuments(ctx, db.Filter)
}

// Paginate 分页查询
//正常情况下分页查询需要自己自定义排序的属性以及查找条件（因为这两个是可以没有的）
//这个函数需要传入当前的页数（string是为了适配echo），每页多少(interface是为了适配设置的配置)，以及传入的值
//函数返回总页数以及相关的错误信息
func (db *Db) Paginate(now int, limit int, result interface{}) (int, error) {
	//获取数据总数
	page, err := db.GetCount()
	if err != nil {
		return 0, errors.New("获取页数失败")
	}
	//获取总页数
	page = int64(math.Ceil(float64(page) / float64(limit)))
	//开始查询
	err = db.SetSkip(int64((now - 1) * limit)).SetLimit(int64(limit)).FindMore(result)
	//是否查询到了数据
	if err != nil {
		return 0, err
	}
	//这里说明查询到了数据
	return int(page), nil
}

// PaginateWithTotal 分页查询带总记录条数版本
// 正常情况下分页查询需要自己自定义排序的属性以及查找条件（因为这两个是可以没有的）
// 这个函数需要传入当前的页数（string是为了适配echo），每页多少(interface是为了适配设置的配置)，以及传入的值
// 函数返回总页数以及相关的错误信息 (第一个是总记录数，第二个是总页数，第三是错误信息)
func (db *Db) PaginateWithTotal(now int, limit int, result interface{}) (int, int, error) {
	//获取数据总数
	total, err := db.GetCount()
	if err != nil {
		return 0, 0, errors.New("获取页数失败")
	}
	//获取总页数
	page := int64(math.Ceil(float64(total) / float64(limit)))
	//开始查询
	err = db.SetSkip(int64((now - 1) * limit)).SetLimit(int64(limit)).FindMore(result)
	//是否查询到了数据
	if err != nil {
		return 0, 0, err
	}
	//这里说明查询到了数据
	return int(total), int(page), nil
}

//---------聚合查询所有的结构函数

// SetLookUp 设置聚合查询的表的对应关系
func (db *Db) SetLookUp(concatCollection string, field string, foreign string, as string) *Db {
	db.Pipeline = append(db.Pipeline, bson.M{"$lookup": bson.M{"from": concatCollection, "localField": field, "foreignField": foreign, "as": as}})
	return db
}

// SetUnwind 聚合查询对数组进行拆分（主要用于拆分聚合的字段）
func (db *Db) SetUnwind(field string) *Db {
	db.Pipeline = append(db.Pipeline, bson.M{"$unwind": "$" + field})
	return db
}

// SetAddFields 设置聚合表中只返回的字段
func (db *Db) SetAddFields(field bson.M) *Db {
	db.Pipeline = append(db.Pipeline, bson.M{"$addFields": field})
	return db
}

// SetProject 聚合查询设置project（返回结果只返回某一个字段）
func (db *Db) SetProject(project bson.M) *Db {
	db.Pipeline = append(db.Pipeline, bson.M{"$project": project})
	return db
}

// 聚合查询设置group
func (db *Db) setGroup(group bson.M) *Db {
	db.Pipeline = append(db.Pipeline, bson.M{"$group": group})
	return db
}

// GetAggregateCount 获取聚合查询的记录数
func (db *Db) GetAggregateCount() (int64, error) {
	pipeline := db.Pipeline
	if db.Filter != nil {
		pipeline = append(pipeline, bson.M{"$match": db.Filter})
	}
	if db.Sort != nil {
		pipeline = append(pipeline, bson.M{"$sort": db.Sort})
	}
	if db.Skip != 0 {
		pipeline = append(pipeline, bson.M{"$skip": db.Skip})
	}
	if db.Limit != 0 {
		pipeline = append(pipeline, bson.M{"$limit": db.Limit})
	}
	pipeline = append(pipeline, bson.M{"$count": "total"})
	//开始连表查询
	cursor, err := db.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)
	cursor.Next(ctx)
	var num interface{}
	err = cursor.Decode(&num)
	if err != nil {
		return 0, err
	}
	total, ok := num.(bson.D).Map()["total"].(int32)
	if !ok {
		return 0, errors.New("获取数据条数失败")
	}
	return int64(total), nil
}

// GetFieldSum 聚合查询统计某个字段的和
func (db *Db) GetFieldSum(id interface{}, field string) ([]FieldSum, error) {
	data := new([]FieldSum)
	// 直接进行聚合查询(我们这里实际上是使用了group来进行分组，id是分组字段，我们这里不同，后面total就是我们自定义的字段，$sum这个用于计算总值)
	db.setGroup(bson.M{"_id": id, "total": bson.M{"$sum": "$" + field}})
	if err := db.Aggregate(data); err == nil {
		return *data, nil
	} else {
		return []FieldSum{}, err
	}
}

// Aggregate 聚合查询(耗时较长，不建议经常使用)
func (db *Db) Aggregate(result interface{}) error {
	pipeline := db.Pipeline
	if db.Filter != nil {
		pipeline = append(pipeline, bson.M{"$match": db.Filter})
	}
	if db.Sort != nil {
		pipeline = append(pipeline, bson.M{"$sort": db.Sort})
	}
	if db.Skip != 0 {
		pipeline = append(pipeline, bson.M{"$skip": db.Skip})
	}
	if db.Limit != 0 {
		pipeline = append(pipeline, bson.M{"$limit": db.Limit})
	}
	if db.Group != nil {
		pipeline = append(pipeline, bson.M{"$group": db.Group})
	}
	//开始连表查询
	cursor, err := db.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, result)
}

// AggregatePaginate 聚合分页查询（这个就是在原有接口的基础上进行封装）
func (db *Db) AggregatePaginate(now int, limit int, result interface{}) (int, error) {
	//获取数据总数
	page, err := db.GetAggregateCount()
	if err != nil {
		return 0, err
	}
	//获取总页数
	page = int64(int(math.Ceil(float64(page) / float64(limit))))
	err = db.
		SetSkip(int64((now - 1) * limit)).
		SetLimit(int64(limit)).
		Aggregate(result)
	if err != nil {
		return 0, err
	}
	return int(page), nil
}

// DropCollection 数据库删除单个集合
func (db *Db) DropCollection() error {
	return db.Collection.Drop(ctx)
}

// DropDatabase 删除整个数据库
func DropDatabase() error {
	//重连数据库
	DbReconnect()
	return DB.Drop(ctx)
}
