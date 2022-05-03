// Package common
// @Description 爬虫相关的
// @Author 小游
// @Date 2021/04/13
package common

import (
	"encoding/json"
	"errors"
	"strings"
	"xBlog/internal/app/model"
	"xBlog/internal/db"
	"xBlog/tools"
)

// BiliGetPersonInfo 获取B站的个人信息（V2版本接口）
func BiliGetPersonInfo(uid string) (model.BiliPersonInfo, error) {
	//避免被B站识别为爬虫
	head := tools.HttpNewHead()
	head["origin"] = "https://space.bilibili.com"
	head["Referer"] = "https://space.bilibili.com"
	head["Cookie"] = tools.Interface2String(db.GetSiteOption(db.KeyBiliCookie))
	var info model.BiliPersonInfo
	info.Uid = uid
	//获取用户基本信息
	content, _ := tools.HttpGetHead("https://api.bilibili.com/x/space/acc/info?mid="+uid, head)
	var row interface{}
	if json.Unmarshal([]byte(content), &row) != nil {
		return model.BiliPersonInfo{}, errors.New("获取个人信息失败")
	}
	infos, ok := row.(map[string]interface{})["data"].(map[string]interface{})
	if ok {
		info.Nickname = tools.Interface2String(infos["name"])
		info.Avatar = tools.Interface2String(infos["face"])
		info.Level = tools.InterfaceFloat2Int(infos["level"])
		info.Sign = tools.Interface2String(infos["sign"])
		info.Sex = tools.Interface2String(infos["sex"])
		info.TopImage = tools.Interface2String(infos["top_photo"])
		if vip, ok := infos["vip"].(map[string]interface{})["status"]; ok {
			info.IsVip = tools.InterfaceFloat2Int(vip)
		}
		//替换头像https链接
		info.Avatar = tools.HttpReplaceHttp(info.Avatar)
	}
	//获取头像挂件，背景卡片等信息
	content, _ = tools.HttpGetHead("https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?host_uid="+uid, head)
	//解析用户数据
	if json.Unmarshal([]byte(content), &row) != nil {
		return model.BiliPersonInfo{}, errors.New("获取背景卡片失败")
	}
	infos, ok = row.(map[string]interface{})["data"].(map[string]interface{})
	if ok && tools.InterfaceFloat2Int(infos["has_more"]) == 1 { //发布过动态，可以获取卡片信息
		infos, ok = infos["cards"].([]interface{})[0].(map[string]interface{})["desc"].(map[string]interface{})["user_profile"].(map[string]interface{})
		if ok {
			info.Hang = tools.Interface2String(infos["pendant"].(map[string]interface{})["image"])
			//装饰卡片可能为空
			if infos["decorate_card"] != nil {
				info.Card = tools.Interface2String(infos["decorate_card"].(map[string]interface{})["card_url"])
			}
			//替换https链接
			info.Hang = tools.HttpReplaceHttp(info.Hang)
			info.Card = tools.HttpReplaceHttp(info.Card)
		}
	}
	//获取浏览量等内容
	content, _ = tools.HttpGetHead("https://api.bilibili.com/x/space/upstat?mid="+uid, head)
	if json.Unmarshal([]byte(content), &row) != nil {
		return model.BiliPersonInfo{}, errors.New("获取浏览量等数据失败")
	}
	infos, ok = row.(map[string]interface{})["data"].(map[string]interface{})
	if ok {
		if infos["archive"] != nil {
			info.View = tools.InterfaceFloat2Int(infos["archive"].(map[string]interface{})["view"])
		}
		if infos["article"] != nil {
			info.Article = tools.InterfaceFloat2Int(infos["article"].(map[string]interface{})["view"])
		}
		info.Good = tools.InterfaceFloat2Int(infos["likes"])
	}
	//获取播放数和粉丝数
	content, _ = tools.HttpGetHead("https://api.bilibili.com/x/relation/stat?vmid="+uid, head)
	if json.Unmarshal([]byte(content), &row) != nil {
		return model.BiliPersonInfo{}, errors.New("获取播放数粉丝数失败")
	}
	if infos, ok = row.(map[string]interface{})["data"].(map[string]interface{}); ok {
		info.Fans = tools.InterfaceFloat2Int(infos["follower"])
		info.Watch = tools.InterfaceFloat2Int(infos["following"])
	}
	//获取B站个人认证
	return info, nil
}

// BiliGetBaseInfo 只获取等级，挂件，头像，uid
func BiliGetBaseInfo(uid string) (model.BiliBaseInfo, error) {
	//避免被B站识别为爬虫
	head := tools.HttpNewHead()
	head["origin"] = "https://space.bilibili.com"
	head["Referer"] = "https://space.bilibili.com"
	head["Cookie"] = tools.Interface2String(db.GetSiteOption(db.KeyBiliCookie))
	var info model.BiliBaseInfo
	//获取用户基本信息
	content, _ := tools.HttpGetHead("https://api.bilibili.com/x/space/acc/info?mid="+uid, head)
	var row interface{}
	if json.Unmarshal([]byte(content), &row) != nil {
		return info, errors.New("获取个人信息失败")
	}
	if data, ok := row.(map[string]interface{})["data"].(map[string]interface{}); ok {
		info.Nickname = tools.Interface2String(data["name"])
		info.Level = tools.InterfaceFloat2Int(data["level"])
		info.Avatar = tools.Interface2String(data["face"])
		if pendant, ok := data["pendant"].(map[string]interface{}); ok {
			info.Hang = tools.Interface2String(pendant["image"])
		}
	}
	return info, nil
}

// GetBiBoDynamicV2 获取B站动态内容V2版本
func GetBiBoDynamicV2(uid string, id string) ([]model.BiBoDynamic, error) {
	//避免被B站识别为爬虫
	head := tools.HttpNewHead()
	head["origin"] = "https://space.bilibili.com"
	head["Referer"] = "https://space.bilibili.com"
	head["Cookie"] = tools.Interface2String(db.GetSiteOption(db.KeyBiliCookie))
	content, _ := tools.HttpGetHead("https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?host_uid="+uid+"&offset_dynamic_id="+id+"&need_top=1", head)
	var v interface{}
	if json.Unmarshal([]byte(content), &v) != nil {
		return nil, errors.New("获取数据失败")
	}
	//开始解析
	if data, ok := v.(map[string]interface{})["data"].(map[string]interface{})["cards"].([]interface{}); ok {
		var dynamics []model.BiBoDynamic
		//遍历动态列表
		for _, v := range data {
			//获取动态的内容
			card := v.(map[string]interface{})
			var dynamic model.BiBoDynamic
			dynamic.Images = []string{}
			dynamic.OImages = []string{}
			//获取动态的基本信息
			if desc, ok := card["desc"].(map[string]interface{}); ok {
				dynamic.View = tools.InterfaceFloat2Int(desc["view"])
				dynamic.RePost = tools.InterfaceFloat2Int(desc["repost"])
				dynamic.Like = tools.InterfaceFloat2Int(desc["like"])
				dynamic.Id = tools.Interface2String(desc["dynamic_id_str"])
				dynamic.Time = tools.Unix2Time(tools.InterfaceFloat2Int(desc["timestamp"]))
				//获取Bv号
				if desc["bvid"] != nil {
					dynamic.Bid = tools.Interface2String(desc["bvid"])
				}
				//获取头像挂件
				if pendent, ok := desc["user_profile"].(map[string]interface{})["pendant"].(map[string]interface{})["image"].(string); ok {
					dynamic.Pendant = pendent
				}
				//获取卡片挂件
				if card, ok := desc["user_profile"].(map[string]interface{})["decorate_card"].(map[string]interface{}); ok && card != nil {
					var decorate model.BiBoDecorate
					decorate.Card = tools.Interface2String(card["card_url"])
					decorate.Number = tools.Interface2String(card["fan"].(map[string]interface{})["num_desc"])
					decorate.Color = tools.Interface2String(card["fan"].(map[string]interface{})["color"])
					decorate.Name = tools.Interface2String(card["name"])
					dynamic.Decorate = decorate
				}
			}
			//获取动态中的表情数据(用于后面内容替换)
			smiles := make(map[string]string)
			//为了避免解析出现错误，需要进行多次判断
			if card["display"] != nil {
				if data, ok := card["display"].(map[string]interface{})["emoji_info"]; ok && data != nil {
					if data2, ok := data.(map[string]interface{})["emoji_details"].([]interface{}); ok {
						for _, v := range data2 {
							if v2, ok := v.(map[string]interface{}); ok {
								smiles[tools.Interface2String(v2["text"])] = tools.Interface2String(v2["url"])
							}
						}
					}
				}
			}
			//开始解析动态内容
			cardContent := tools.Interface2String(card["card"])
			//动态内容为json字符串所以我们需要解析
			var row interface{}
			if json.Unmarshal([]byte(cardContent), &row) != nil {
				//解析时发生错误跳出这个循环
				break
			}
			//动态内容进一步解析
			if content, ok := row.(map[string]interface{}); ok {
				//判断动态的类型
				if strings.HasPrefix(cardContent, `{"aid":`) {
					var video model.BiBoVideoContent
					//这里说明动态为视频
					dynamic.Types = "video"
					//这里解析动态的数据
					video.Aid = tools.InterfaceFloat2String(content["aid"])
					video.Cid = tools.InterfaceFloat2String(content["cid"])
					video.Dec = tools.Interface2String(content["desc"])
					video.Dynamic = tools.Interface2String(content["dynamic"])
					video.Pic = tools.Interface2String(content["pic"])
					video.Title = tools.Interface2String(content["title"])
					//获取硬币等数据
					if content["stat"] != nil {
						if data, ok := content["stat"].(map[string]interface{}); ok {
							video.Coin = tools.InterfaceFloat2Int(data["coin"])
							video.Barrage = tools.InterfaceFloat2Int(data["danmaku"])
							video.View = tools.InterfaceFloat2Int(data["view"])
							video.Share = tools.InterfaceFloat2Int(data["share"])
							video.Comment = tools.InterfaceFloat2Int(data["reply"])
							dynamic.Comment = tools.InterfaceFloat2Int(data["reply"])
							video.Favorite = tools.InterfaceFloat2Int(data["favorite"])
							video.Like = tools.InterfaceFloat2Int(data["like"])
						}
					}
					dynamic.Content = video
				} else if strings.HasPrefix(cardContent, `{"item":`) {
					//这里说明是动态数据
					dynamic.Types = "dynamic"
					//获取动态内容
					if content["item"] != nil {
						if data, ok := content["item"].(map[string]interface{}); ok {
							description := tools.Interface2String(data["description"])
							//遍历替换表情
							for k, v := range smiles {
								description = strings.ReplaceAll(description, k, `<img src="`+v+`" alt="`+k+`" style="margin: -1px 1px 0 1px;display: inline-block; width: 20px; height: 20px; vertical-align: text-bottom;">`)
							}
							dynamic.Content = model.BiBoDynamicContent{Content: description}
							//获取评论数
							dynamic.Comment = tools.InterfaceFloat2Int(data["reply"])
							//判断是否有图片
							if data["pictures"] != nil {
								if images, ok := data["pictures"].([]interface{}); ok {
									for _, v := range images {
										dynamic.Images = append(dynamic.Images, tools.Interface2String(v.(map[string]interface{})["img_src"]))
									}
								}
							}
						}
					}
				} else if strings.HasPrefix(cardContent, `{ "user":`) {
					//这里说明是转发
					dynamic.Types = "share"
					//获取分享的内容
					if content["item"] != nil {
						if data, ok := content["item"].(map[string]interface{}); ok {
							var shareContent model.BiBoShareContent
							description := tools.Interface2String(data["content"])
							//替换表情
							for k, v := range smiles {
								description = strings.ReplaceAll(description, k, `<img src="`+v+`" alt="`+k+`" style="margin: -1px 1px 0 1px;display: inline-block; width: 20px; height: 20px; vertical-align: text-bottom;">`)
							}
							shareContent.Content = description
							//获取原始数据(转发的数据)
							if content["origin"] != nil {
								origin := tools.Interface2String(content["origin"])
								//开始解析原始数据
								var row interface{}
								if json.Unmarshal([]byte(origin), &row) != nil {
									//解析失败跳出循环
									break
								}
								//这里进行解析判断
								if strings.HasPrefix(origin, `{"item":`) {
									shareContent.Type = "text"
									//转发的是普通动态
									if data, ok := row.(map[string]interface{})["user"]; ok && data != nil {
										if user, ok := data.(map[string]interface{}); ok {
											shareContent.Uid = tools.InterfaceFloat2String(user["uid"])
											shareContent.Name = tools.InterfaceFloat2String(user["name"])
											shareContent.Avatar = tools.Interface2String(user["avatar"])
										}
									}
									//获取原动态内容
									if data, ok := row.(map[string]interface{})["item"]; ok && data != nil {
										if data2, ok := data.(map[string]interface{}); ok {
											shareContent.Origin = tools.Interface2String(data2["description"])
											//判断是否有图片
											if data2["pictures"] != nil {
												if data3, ok := data2["pictures"].([]interface{}); ok {
													for _, k := range data3 {
														if data4, ok := k.(map[string]interface{})["img_src"]; ok && data4 != nil {
															dynamic.OImages = append(dynamic.OImages, tools.Interface2String(data4))
														}
													}
												}
											}
										}
									}
								} else if strings.HasPrefix(origin, `{ "user":`) {
									//这里是普通的没有图片的动态
									shareContent.Type = "text"
									//获取用户信息
									if data, ok := row.(map[string]interface{})["user"]; ok && data != nil {
										if user, ok := data.(map[string]interface{}); ok {
											shareContent.Name = tools.InterfaceFloat2String(user["uname"])
											shareContent.Avatar = tools.Interface2String(user["face"])
										}
									}
									//获取原动态内容
									if data, ok := row.(map[string]interface{})["item"]; ok && data != nil {
										if data2, ok := data.(map[string]interface{}); ok {
											shareContent.Origin = tools.Interface2String(data2["content"])
										}
									}
								} else if strings.HasPrefix(origin, `{"aid":`) {
									shareContent.Type = "video"
									if data, ok := row.(map[string]interface{}); ok {
										shareContent.Aid = tools.InterfaceFloat2String(data["aid"])
										shareContent.Cid = tools.InterfaceFloat2String(data["cid"])
										shareContent.Desc = tools.Interface2String(data["desc"])
										shareContent.Dynamic = tools.Interface2String(data["dynamic"])
										shareContent.Pic = tools.Interface2String(data["pic"])
										shareContent.Title = tools.Interface2String(data["title"])
										//获取视频的原主人信息
										if data["owner"] != nil {
											if owner, ok := data["owner"].(map[string]interface{}); ok {
												shareContent.Avatar = tools.Interface2String(owner["face"])
												shareContent.Name = tools.Interface2String(owner["name"])
												shareContent.Uid = tools.InterfaceFloat2String(owner["mid"])
											}
										}
										//获取视频的播放数据
										if data["stat"] != nil {
											if stat, ok := data["stat"].(map[string]interface{}); ok {
												shareContent.View = tools.InterfaceFloat2Int(stat["view"])
												shareContent.Like = tools.InterfaceFloat2Int(stat["like"])
											}
										}

									}

								}
							}
							dynamic.Content = shareContent
						}
					}
				}
			}
			dynamics = append(dynamics, dynamic)
		}
		return dynamics, nil
	}
	return nil, errors.New("获取数据失败")
}

// ToolsSyncMusic163 同步网易云歌单
func ToolsSyncMusic163() ([]model.AdminMusic163, error) {
	//获取网站配置
	server := db.GetSiteOptionString(db.KeySiteApiServer)
	//获取网站的歌单
	head := tools.HttpNewHead()
	head["Cookie"] = "MUSIC_U=" + db.GetSiteOptionString(db.KeyMusicU)
	rowMusic, _ := tools.HttpGetHead("https://music.163.com/api/playlist/detail?id="+db.GetSiteOptionString(db.KeyMusicId), head)
	var row interface{}
	if json.Unmarshal([]byte(rowMusic), &row) != nil {
		return nil, errors.New("解析json数据失败")
	}
	row = row.(map[string]interface{})["result"].(map[string]interface{})["tracks"]
	//遍历获取到的列表数据
	var dataS []model.AdminMusic163
	for _, v := range row.([]interface{}) {
		id := tools.Float2String(v.(map[string]interface{})["id"].(float64))
		var data model.AdminMusic163
		data.Name = v.(map[string]interface{})["name"].(string)
		data.Artist = v.(map[string]interface{})["artists"].([]interface{})[0].(map[string]interface{})["name"].(string)
		data.Url = "https://music.163.com/song/media/outer/url?id=" + id
		data.Cover = v.(map[string]interface{})["album"].(map[string]interface{})["picUrl"].(string)
		// 替换http
		data.Cover = strings.Replace(data.Cover, "http:", "", -1)
		data.Irc = server + "/api/v3/settings/side/music/" + id + "/irc"
		dataS = append(dataS, data)
	}
	// 保存到设置里去
	db.SetSiteOption(db.KeyMusicContent, dataS)
	return dataS, nil
}
