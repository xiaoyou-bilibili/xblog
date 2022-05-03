package com.xiaoyou.library.net.entity.response

import android.graphics.Paint
import com.bin.david.form.annotation.SmartColumn
import com.bin.david.form.annotation.SmartTable


/**
 * @description 插件模块返回的数据结构
 * @author 小游
 * @data 2021/02/12
 */

/**
 * 日记Diary
 * @property diary_id Int 日记id
 * @property content String 日记内容
 * @property date String 日记发布时间
 * @property comment Int 评论数
 * @property good Int 点赞数
 * @property avatar String 头像
 * @property nickname String 昵称
 * @property encryption Boolean 日记是否加密
 * @constructor
 */
data class DiaryDetail(
        val diary_id: Int,
        val content: String,
        val date: String,
        val comment: Int,
        val good: Int,
        val avatar: String,
        val nickname: String,
        val encryption: Boolean
)

/**
 * 友链模型
 * @property name String 友链名称
 * @property url String 友链地址
 * @property avatar String 友链头像
 * @property dec String 友链描述
 * @constructor
 */
data class FriendDetail(
        val name: String,
        val url: String,
        val avatar: String,
        val dec: String
)

/**
 * 追番数据
 * @property current Int 当前为第几页
 * @property total Int 	页数
 * @property num Int 	总共多少部番剧
 * @property contents List<Animation> 番剧内容
 * @constructor
 */
data class Animations(val current: Int, val num: Int, val total: Int, val contents: List<AnimationDetail>)

/**
 * 番剧的详细信息
 * @property title String 番剧标题
 * @property cover String 番剧封面
 * @property dec String 番剧描述
 * @property total String 番剧总集数
 * @property current String 现在在看第几集
 * @property percent Float 番剧观看进度
 * @property url String 番剧的地址
 * @constructor
 */
data class AnimationDetail(
        val title: String,
        val cover: String,
        val dec: String,
        val total: String,
        val current: String,
        val percent: Float,
        val url: String
)

/**
 * 赞助人模型
 * @property nickname String
 * @property donate String
 * @property comment String
 * @constructor
 */
@SmartTable()
data class Sponsors(
        @SmartColumn(id = 1, name = "昵称", align = Paint.Align.LEFT, titleAlign = Paint.Align.LEFT)
        val nickname: String,
        @SmartColumn(id = 2, name = "赞助额", align = Paint.Align.LEFT, titleAlign = Paint.Align.LEFT)
        val donate: String,
        @SmartColumn(id = 3, name = "赞助说明", align = Paint.Align.LEFT, titleAlign = Paint.Align.LEFT)
        val comment: String
)

/**
 * 豆瓣数据
 * @property comment String 评论信息
 * @property name String 名字
 * @property picture String 封面
 * @property pub String 出版信息
 * @property star String 评分
 * @property status String 状态
 * @property url String 链接
 * @constructor
 */
data class DouBanDetail(
        val name: String,
        val picture: String,
        val star: String,
        val pub: String,
        val comment: String,
        val status: String,
        val url: String
)

/**
 *  聊天室
 * @property id Int 聊天室id
 * @property avatar String 聊天室头像
 * @property name String 聊天室名字
 * @property message ChatMessage 聊天信息
 * @property count Int 未读数
 * @constructor
 */
data class ChatRoom(
        val id: Int,
        val avatar: String,
        val name: String,
        val message: ChatMessage,
        val count: Int
)

/**
 *  聊天信息
 * @property id Int id
 * @property user_id Int 用户id
 * @property content String 内容
 * @property date Int 发送时间
 * @property target Int 发送目标
 * @property message_type Int 信息类型
 * @property read Boolean 是否已读
 * @constructor
 */
data class ChatMessage(
        val id:String,
        val user_id:Int,
        val content: String,
        val date:Long,
        val target:Int,
        val message_type:Int,
        val read:Boolean
)

/**
 *  文章列表
 * @property id Int 文档id
 * @property title String 文档标题
 * @property parent Int 文档的父节点
 * @constructor
 */
data class DocList(
        val id:Int,
        val title:String,
        val parent: Int
)

/**
 *  文档内容
 * @property id Int 文档id
 * @property title String 文章标题
 * @property content String 文档内容
 * @constructor
 */
data class DocContent(
        val id:Int,
        val title:String,
        val content:String
)


/**
 *  项目卡片
 * @property top_content List<ProjectTop> 顶部轮播图
 * @property bottom_content List<ProjectBottom> 底部项目地址
 * @constructor
 */
data class Project(
        val top_content:List<ProjectTop>,
        val bottom_content:List<ProjectBottom>
)

/**
 *  顶部轮播图
 * @property image String 轮播图图片
 * @property title String 轮播图标题
 * @property url String 轮播图地址
 * @constructor
 */
data class ProjectTop(
        val image:String,
        val title:String,
        val url: String
)

/**
 *  底部轮播图
 * @property name String 项目名字
 * @property image String 项目图片
 * @property dec String 项目描述
 * @property time String 项目制作时间
 * @property video String 视频地址
 * @property blog String 博客地址
 * @property code String 代码地址
 * @constructor
 */
data class ProjectBottom(
        val name:String,
        val image:String,
        val dec: String,
        val time: String,
        val video: String,
        val blog: String,
        val code: String
)

/**
 *  个人导航
 * @property id Int 导航id
 * @property name String 导航分类
 * @property color String 导航颜色
 * @property child List<LinkItem> 子节点
 * @constructor
 */
data class Link(
        val id: Int,
        val name: String,
        val color: String,
        val child: List<LinkItem>
)

/**
 *  网址信息
 * @property name String 名字
 * @property url String 地址
 * @constructor
 */
data class LinkItem(
        val name:String,
        val url:String
)