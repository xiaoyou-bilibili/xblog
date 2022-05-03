// Package server @Description 用户板块v3版本
// @Author 小游
// @Date 2021/01/21
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"strconv"
	"time"
	"xBlog/internal/app/common"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// UserGetUser 获取用户信息
func UserGetUser(c *gin.Context) {
	//fmt.Println("获取用户信息")
	if data, err := common.AccessGetTokenV2(c); err == nil {
		user := new(db.User)
		if database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": data.UserID}).FindOne(user) == nil {
			var info model.UserInfo
			info.UserID = user.UserID
			if user.Avatar == "" {
				info.Avatar = db.GetSiteOptionString(db.KeyNoFace)
			} else {
				info.Avatar = user.Avatar
			}
			info.Avatar = user.Avatar
			info.Nickname = user.Nickname
			info.Email = user.Email
			info.Hang = user.Hang
			info.Level = user.Level
			info.Username = user.Username
			info.Sign = user.Sign
			info.Identity = user.Identity
			info.Subscription = user.Subscription
			tools.GlobalResponse.ResponseOk(c, info)
		} else {
			tools.GlobalResponse.ResponseServerError(c)
		}
	} else {
		tools.GlobalResponse.ResponseUnauthorized(c)
	}
}

// UserPutUser 更新用户信息
func UserPutUser(c *gin.Context) {
	// 鉴权
	if user, err := common.AccessGetTokenV2(c); err == nil {
		// 先获取用户信息
		collection := database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": user.UserID})
		info := new(db.User)
		if collection.FindOne(info) != nil {
			tools.GlobalResponse.ResponseNotFound(c, "获取用户信息失败")
			return
		}
		// 参数验证
		param := new(model.UserPutUserParam)
		if tools.ValidatorParam(c, param) {
			return
		}
		// 判断参数
		set := bson.M{}
		if param.Avatar != "" {
			set["avatar"] = param.Avatar
		}
		if param.Hang != "" {
			set["hang"] = param.Hang
		}
		if param.Nickname != "" {
			set["nickname"] = param.Nickname
		}
		if param.Email != "" {
			var tmp db.User
			// 检查一下邮箱是否占用
			if err := database.NewDb(db.CollUser).SetFilter(bson.M{"email": param.Email}).FindOne(&tmp); (err != nil && err.Error() == database.MongoNoResult) || tmp.UserID == user.UserID {
				set["email"] = param.Email
			} else {
				fmt.Println(tmp.UserID == user.UserID)
				tools.GlobalResponse.ResponseBadRequest(c, "该邮箱已被占用，请更换邮箱!")
				return
			}
		}
		if param.Sign != "" {
			set["sign"] = param.Sign
		}
		if param.Subscription != "" {
			set["subscription"] = tools.Str2Bool(param.Subscription)
		}
		// 如果用户想重置密码，那么就必须要确保旧密码和新密码相同
		if param.NewPassword != "" {
			if tools.Encrypt(param.OldPassword) == info.Password {
				set["password"] = tools.Encrypt(param.NewPassword)
			} else {
				tools.GlobalResponse.ResponseUnauthorized(c, "旧密码和原密码不相同")
				return
			}
		}
		// 修改个人信息
		if collection.Set(set).UpdateOne() == nil {
			tools.GlobalResponse.ResponseCreated(c, param)
		} else {
			tools.GlobalResponse.ResponseServerError(c, "更新个人信息失败")
		}
	} else {
		tools.GlobalResponse.ResponseUnauthorized(c)
	}
}

// UserAddUser 用户注册
func UserAddUser(c *gin.Context) {
	// 参数验证
	param := new(model.AddUserParam)
	//fmt.Println(param)
	if tools.ValidatorParam(c, param) {
		return
	}
	//插入数据前先判断用户是否已经注册
	res := new(db.User)
	if err := database.NewDb(db.CollUser).OR([]bson.M{{"email": param.Email}, {"username": param.Username}}).FindOne(res); err == nil {
		tools.GlobalResponse.ResponseBadRequest(c, "该用户已注册，请更换邮箱地址或用户名后注册")
	} else if err.Error() == database.MongoNoResult { //没有找到记录，可以注册
		var user db.User
		//生成随机的token
		user.Token = tools.GetRandomNum(20)
		//密码加密
		user.Password = tools.Encrypt(param.Password)
		user.ID = primitive.NewObjectID()
		user.Username = param.Username
		user.Nickname = param.Nickname
		user.Email = param.Email
		user.LoginInfo = []db.LoginInfo{}
		user.LastTime = time.Now()
		user.Registered = time.Now()
		user.Identity = 2
		//插入数据
		if user.UserID, err = database.NewDb(db.CollUser).InsertOneIncrease(user, "user_id"); err != nil {
			tools.GlobalResponse.ResponseServerError(c, "内部错误，请联系管理员解决")
			return
		}
		sever := db.GetSiteOptionString(db.KeySiteApiServer)
		site := db.GetSiteOptionString(db.KeySiteName)
		//激活的邮件地址
		active := sever + "/api/v3/user/" + tools.Int2Str(user.UserID) + "/status?token=" + user.Token + "&site=" + sever
		body := `
			<div class="content" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; max-width: 600px; display: block; margin: 0 auto; padding: 20px;">
				<table class="main" width="100%" cellpadding="0" cellspacing="0" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; border-radius: 3px; background-color: #fff; margin: 0; border: 1px solid #e9e9e9;" bgcolor="#fff"><tbody><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="alert alert-warning" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 16px; vertical-align: top; color: #fff; font-weight: 500; text-align: center; border-radius: 3px 3px 0 0; background-color: #009688; margin: 0; padding: 20px;" align="center" bgcolor="#FF9F00" valign="top">
							` + site + `用户注册
						</td>
					</tr><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-wrap" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 20px;" valign="top">
							<table width="100%" cellpadding="0" cellspacing="0" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><tbody><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
										亲爱的 <strong style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><span style="border-bottom: 1px dashed rgb(204, 204, 204); z-index: 1; position: static;" t="7" onclick="return false;" data="1589294503" isout="1">` + user.Username + `</span></strong> ：
									</td>
								</tr><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
										感谢您注册` + site + `,请点击下方按钮完成账户激活。
									</td>
								</tr><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
										<a href="` + active + `" class="btn-primary" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; color: #FFF; text-decoration: none; line-height: 2em; font-weight: bold; text-align: center; cursor: pointer; display: inline-block; border-radius: 5px; text-transform: capitalize; background-color: #009688; margin: 0; border-color: #009688; border-style: solid; border-width: 10px 20px;" rel="noopener" target="_blank">激活账户</a>
									</td>
								</tr><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
										欢迎加入我们这个大家庭！
									</td>
								</tr></tbody></table></td>
					</tr></tbody></table><div class="footer" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; width: 100%; clear: both; color: #999; margin: 0; padding: 20px;">
					<table width="100%" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><tbody><tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="aligncenter content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 12px; vertical-align: top; color: #999; text-align: center; margin: 0; padding: 0 0 20px;" align="center" valign="top">此邮件由系统自动发送，请不要直接回复。</td>
						</tr></tbody></table></div></div>
		`
		if tools.SendMail([]string{param.Email}, "新用户注册邮箱验证", body) != nil {
			tools.GlobalResponse.ResponseServerError(c, "激活邮件发送失败请尝试通过忘记密码来激活!")
		} else {
			tools.GlobalResponse.ResponseCreated(c, param)
		}
	} else {
		tools.GlobalResponse.ResponseServerError(c, "内部错误，请稍后重试")
	}
}

// UserActiveUser 用户激活
func UserActiveUser(c *gin.Context) {
	userID := tools.Str2Int(c.Param("id"))
	token := c.Query("token")
	if token == "" || userID == 0 {
		tools.GlobalResponse.ResponseBadRequest(c)
		return
	}
	//判断用户的token是否正确
	user := new(db.User)
	collection := database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": userID})
	if collection.FindOne(user) == nil && user.Token == token {
		buff, err := ioutil.ReadFile(database.GetAppPath() + "assets/template/302.html")
		if err != nil {
			tools.GlobalResponse.ResponseServerError(c, "内部错误")
			return
		}
		//激活用户
		if collection.Set(bson.M{"status": 1, "token": ""}).UpdateOne() == nil {
			//返回激活界面
			tools.GlobalResponse.ResponseHtml(c, string(buff))
		}
	} else {
		tools.GlobalResponse.ResponseServerError(c, "激活失败")
	}
}

// UserToken 用户登录
func UserToken(c *gin.Context) {
	// 参数验证
	param := new(model.UserLoginParam)
	if tools.ValidatorParam(c, param) {
		return
	}
	//密码加密
	password := tools.Encrypt(param.Password)
	user := new(db.User)
	collection := database.NewDb(db.CollUser).OR([]bson.M{{"username": param.Username}, {"email": param.Username}})
	//判断用户是否激活
	if err := collection.FindOne(user); err == nil && user.Status == 0 {
		tools.GlobalResponse.ResponseUnauthorized(c, "你的账户未激活，如果没有收到邮件请直接重置密码")
	} else if err == nil && user.Password == password { //判断用户名和密码是否正确
		//判断账户是否被锁定
		if user.LoginFail > 9 { //超过10次用户密码被锁定
			tools.GlobalResponse.ResponseUnauthorized(c, "你的账户已锁定，请重置密码!")
			return
		}
		//获取用户登录信息
		var info db.LoginInfo
		info.Ip = c.ClientIP()
		info.LoginTime = time.Now()
		info.Agent = c.GetHeader("User-Agent")
		if token, err := tools.JwtCreateToken(); err != nil {
			tools.GlobalResponse.ResponseServerError(c, "内部错误，请重新登录")
		} else {
			info.Token = token
		}
		//存储登录信息到数据库
		// fix: 修复不能登录的问题
		if collection.AddToSet(bson.M{"login_info": info}).UpdateOne() != nil {
			tools.GlobalResponse.ResponseServerError(c, "内部错误，请重新登录")
			return
		}
		//更新上次登录ip和时间,重置密码登录错误次数为0
		_ = collection.Set(bson.M{"login_fail": 0, "last_time": time.Now(), "last_ip": c.ClientIP()}).UpdateOne()
		var token model.UserLogin
		token.UserID = user.UserID
		token.Token = info.Token
		tools.GlobalResponse.ResponseCreated(c, token)
	} else if err == nil && user.Password != password { //用户密码不正确
		//判断账户是否锁定
		if user.LoginFail > 9 {
			tools.GlobalResponse.ResponseUnauthorized(c, "你的账户已锁定，请重置密码!")
			return
		}
		//登录错误次数+1
		_ = collection.Inc(bson.M{"login_fail": 1}).UpdateOne()
		tools.GlobalResponse.ResponseUnauthorized(c, "用户名或密码错误,你还有"+strconv.Itoa(9-user.LoginFail)+"次机会，用完后账户会锁定")
	} else {
		tools.GlobalResponse.ResponseUnauthorized(c, "用户名或密码错误")
	}
}

// UserGetToken 用户是否登录
func UserGetToken(c *gin.Context) {
	if user, err := common.AccessGetTokenV2(c); err != nil {
		tools.GlobalResponse.ResponseUnauthorized(c)
	} else {
		tools.GlobalResponse.ResponseOk(c, user)
	}
}

// UserV3PasswordEmail 发送重置邮件
func UserV3PasswordEmail(c *gin.Context) {
	// 参数验证
	param := new(model.UserResetPasswordEmail)
	//跳转到重置密码的界面
	if c.Request.Method == "GET" && c.Query("option") == "reset" {
		buff, err := ioutil.ReadFile(database.GetAppPath() + "assets/template/reset.html")
		if err != nil {
			tools.GlobalResponse.ResponseUnauthorized(c, "用户名或密码错误")
		} else {
			tools.GlobalResponse.ResponseHtml(c, string(buff))
		}
		return
	} else if tools.ValidatorParam(c, param) {
		return
	}
	//发送重置密码的邮件
	if common.AccessResetPasswordV2(param.Email) {
		tools.GlobalResponse.ResponseCreated(c, param)
	} else {
		tools.GlobalResponse.ResponseServerError(c, "邮件发送失败，请重试")
	}
}

// UserResetPasswordToken 用户通过token来重置密码（这个用户邮件重置密码）
func UserResetPasswordToken(c *gin.Context) {
	// 参数验证
	param := new(model.UserResetPasswordToken)
	if tools.ValidatorParam(c, param) {
		return
	}

	collection := database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": tools.Str2Int(param.UserID)})
	user := new(db.User)
	if collection.FindOne(user) == nil && user.Token == param.Token {
		// 重置密码,同时清空token
		if collection.Set(bson.M{"password": tools.Encrypt(param.Password), "login_fail": 0, "status": 1, "token": ""}).UpdateOne() == nil {
			tools.GlobalResponse.ResponseCreated(c, param)
		}
	} else {
		tools.GlobalResponse.ResponseServerError(c, "密码重置失败")
	}
}

// UserUserGetUserName 判断用户名是否存在
func UserUserGetUserName(c *gin.Context) {
	// 参数验证
	param := new(model.UserGetUserParam)
	param.User = c.Query("user")
	param.Email = c.Query("email")
	user := new(db.User)
	var err error
	if param.User != "" {
		err = database.NewDb(db.CollUser).SetFilter(bson.M{"username": param.User}).FindOne(user)
	} else if param.Email != "" {
		err = database.NewDb(db.CollUser).SetFilter(bson.M{"email": param.Email}).FindOne(user)
	} else {
		tools.GlobalResponse.ResponseBadRequest(c)
		return
	}
	//如果记录不存在说明没有注册
	if err != nil && err.Error() == database.MongoNoResult {
		tools.GlobalResponse.ResponseNotFound(c, "用户名或邮箱未注册")
	} else {
		tools.GlobalResponse.ResponseOk(c, param)
	}
}

// UserPostCode 手机端获取验证码
func UserPostCode(c *gin.Context) {
	// 参数验证
	param := new(model.UserPostCodeParam)
	if tools.ValidatorParam(c, param) {
		return
	}
	//插入验证码
	token := tools.GetRandomNum(4)
	// 我们手动添加验证码到集合里去
	if database.NewDb(db.CollEmailRegistered).SetFilter(bson.M{"email": param.Email}).Set(bson.M{"email": param.Email, "token": token, "send_time": time.Now()}).SetUpsert(true).UpdateOne() == nil {
		//subject 邮件主题 body 发送邮件的内容
		var subject, body string
		site := db.GetSiteOptionString(db.KeySiteName)
		//判断用户是登录还是注册
		if param.Option == "register" {
			// 如果是邮箱注册要先判断这个邮箱是否注册过
			user := new(db.User)
			if database.NewDb(db.CollUser).SetFilter(bson.M{"email": param.Email}).FindOne(user) == nil {
				tools.GlobalResponse.ResponseBadRequest(c, "该邮箱已注册!")
				return
			}
			subject = "新用户注册邮箱验证"
			//内容
			body = `你正在注册` + site + `账号,你的验证码为:<h1>` + token + "</h1><br>如果你没有注册该网站,请忽略该条消息"
		} else if param.Option == "forget" {
			subject = "密码重置"
			//内容
			body = `你正在重置你的` + site + `账号密码,你的验证码为:<h1>` + token + "</h1><br>如果你没有注册该网站,请忽略该条消息"
		} else {
			tools.GlobalResponse.ResponseBadRequest(c, "未知操作")
			return
		}
		if tools.SendMail([]string{param.Email}, subject, body) == nil {
			tools.GlobalResponse.ResponseCreated(c, param)
		}
	} else {
		tools.GlobalResponse.ResponseServerError(c, "邮件发送失败")
	}
}

// UserPostApp 手机端用户注册
func UserPostApp(c *gin.Context) {
	// 参数验证
	param := new(model.UserPostAddParam)
	if tools.ValidatorParam(c, param) {
		return
	}
	//先判断验证码是否正确
	emailRegister := new(db.EmailRegistered)
	if database.NewDb(db.CollEmailRegistered).SetFilter(bson.M{"email": param.Email}).FindOne(emailRegister) == nil && emailRegister.Token == param.Code {
		//判断用户是否存在
		res := new(db.User)
		if err := database.NewDb(db.CollUser).OR([]bson.M{{"email": param.Email}, {"username": param.UserName}}).FindOne(res); err == nil {
			tools.GlobalResponse.ResponseBadRequest(c, "该用户已注册，请更换邮箱地址或用户名后注册")
		} else if err.Error() == database.MongoNoResult { //没有找到记录，可以组成
			var user db.User
			//密码加密
			user.Password = tools.Encrypt(param.Password)
			user.ID = primitive.NewObjectID()
			user.Username = param.UserName
			user.Nickname = param.Nickname
			user.Email = param.Email
			user.Status = 1
			user.LoginInfo = []db.LoginInfo{}
			user.LastTime = time.Now()
			user.Registered = time.Now()
			user.Identity = 2
			//插入数据
			if _, err := database.NewDb(db.CollUser).InsertOneIncrease(user, "user_id"); err == nil {
				// 清除验证码
				_ = database.NewDb(db.CollEmailRegistered).SetFilter(bson.M{"email": param.Email}).Set(bson.M{"token": ""}).UpdateOne()
				tools.GlobalResponse.ResponseCreated(c, param)
			}
		}
	} else {
		tools.GlobalResponse.ResponseServerError(c, "注册失败")
	}
}

// UserPutAppPassword 手机端重置密码
func UserPutAppPassword(c *gin.Context) {
	// 参数验证
	param := new(model.UserPutAppPasswordParam)
	if tools.ValidatorParam(c, param) {
		return
	}
	//判断验证码是否正确
	emailRegister := new(db.EmailRegistered)
	if database.NewDb(db.CollEmailRegistered).SetFilter(bson.M{"email": param.Email}).FindOne(emailRegister) == nil && emailRegister.Token == param.Code {
		//修改用户密码
		if database.NewDb(db.CollUser).SetFilter(bson.M{"email": param.Email}).Set(bson.M{"password": tools.Encrypt(param.Password), "login_fail": 0, "status": 1}).UpdateOne() == nil {
			// 清除验证码
			_ = database.NewDb(db.CollEmailRegistered).SetFilter(bson.M{"email": param.Email}).Set(bson.M{"token": ""}).UpdateOne()
			tools.GlobalResponse.ResponseCreated(c, param)
		}
	} else {
		tools.GlobalResponse.ResponseBadRequest(c, "验证码错误")
	}
}

// UserCollections 获取用户收藏的文章
func UserCollections(c *gin.Context) {
	// 验证用户
	user, err := common.AccessGetTokenV2(c)
	if err != nil {
		tools.GlobalResponse.ResponseUnauthorized(c)
		return
	}
	collections := new([]db.Article)
	var result []model.PostListContent
	// 初始化，避免为空
	result = []model.PostListContent{}
	// 开始查询
	err = database.NewDb(db.CollPostCollection).
		SetUnwind("collection").
		SetLookUp(db.CollArticle, "collection", "post_id", "post").
		SetUnwind("post").
		SetAddFields(bson.M{"post_id": "$post.post_id", "title": "$post.title", "content": "$post.content", "post_time": "$post.post_time", "view": "$post.view", "good": "$post.good", "comment": "$post.comment", "post_type": "$post.post_type", "status": "$post.status"}).
		SetSort(bson.M{"post_id": -1}).
		SetFilter(bson.M{"user_id": user.UserID}).
		Aggregate(collections)
	if err != nil {
		tools.GlobalResponse.ResponseServerError(c, "获取数据失败")
		return
	}
	for _, article := range *collections {
		//获取文章的属性
		post := new(model.PostListContent)
		post.ID = article.PostID
		post.Title = article.Title
		post.Content = common.GetDec(article.Content, db.GetSiteOptionInt(db.KeyPostDecCount))
		post.Date = tools.Time2String(article.PostTime, false)
		post.Image = tools.GetRandomImg()
		post.View = article.View
		post.Good = article.Good
		post.Comment = article.Comment
		result = append(result, *post)
	}
	tools.GlobalResponse.ResponseOk(c, result)
}

// UserWeChatCollections 获取小程序用户收藏的文章
func UserWeChatCollections(c *gin.Context) {
	// 获取openid
	openID := c.Param("openid")
	if openID == "" {
		tools.GlobalResponse.ResponseBadRequest(c)
		return
	}
	// 获取收藏的内容
	collections := new([]db.Article)
	var result []model.PostListContent
	// 初始化，避免为空
	result = []model.PostListContent{}
	// 开始查询
	err := database.NewDb(db.CollPostCollection).
		SetUnwind("collection").
		SetLookUp(db.CollArticle, "collection", "post_id", "post").
		SetUnwind("post").
		SetAddFields(bson.M{"post_id": "$post.post_id", "title": "$post.title", "content": "$post.content", "post_time": "$post.post_time", "view": "$post.view", "good": "$post.good", "comment": "$post.comment", "post_type": "$post.post_type", "status": "$post.status"}).
		SetSort(bson.M{"post_id": -1}).
		SetFilter(bson.M{"openid": openID}).
		Aggregate(collections)
	if err != nil {
		tools.GlobalResponse.ResponseServerError(c, "获取数据失败")
		return
	}
	for _, article := range *collections {
		//获取文章的属性
		post := new(model.PostListContent)
		post.ID = article.PostID
		post.Title = article.Title
		post.Content = common.GetDec(article.Content, db.GetSiteOptionInt(db.KeyPostDecCount))
		post.Date = tools.Time2String(article.PostTime, false)
		post.Image = tools.GetRandomImg()
		post.View = article.View
		post.Good = article.Good
		post.Comment = article.Comment
		result = append(result, *post)
	}
	tools.GlobalResponse.ResponseOk(c, result)
}
