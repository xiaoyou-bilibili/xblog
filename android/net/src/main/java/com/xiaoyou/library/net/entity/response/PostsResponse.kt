package com.xiaoyou.library.net.entity.response

import com.google.gson.annotations.SerializedName
import java.util.*

/**
 * @description 文章部分板块返回结构
 * @author 小游
 * @data 2021/02/10
 */


/**
 * 文章Object
 * @property comment Int 文章评论数
 * @property content String 文章内容
 * @property date String 文章发布时间
 * @property encryption Boolean 文章是否加密
 * @property good Int 文章点赞数
 * @property id Int 文章id
 * @property image String 文章图片
 * @property is_top Boolean 文章是否置顶
 * @property title String 文章标题
 * @property view Int 文章浏览量
 * @constructor
 */
data class PostDetail(
        val id:Int,
        val title:String,
        val content:String,
        val date:String,
        val good:Int,
        val view:Int,
        val image:String,
        val comment:Int,
        val is_top:Boolean,
        val encryption:Boolean
)

/**
 *  文章内容数据
 * @property id Int 文章id
 * @property title String 文章标题
 * @property date String 文章发布时间
 * @property view Int 浏览量
 * @property comment Int 评论数
 * @property good Int 点赞数
 * @property commentStatus String 评论状态
 * @property tag List<TagDetail> 文章标签
 * @property category List<TagDetail> 文章分类
 * @property content String 文章内容
 * @property modify String 最后修改时间
 * @property image String 图片地址
 * @property aliPay String 支付宝二维码
 * @property weChat String 微信支付二维码
 * @property encrypt Boolean 是否加密
 * @constructor
 */
data class PostContentDetail(
        val id: Int,
        val title: String,
        var date: String,
        var view: Int,
        var comment: Int,
        var good: Int,
        @SerializedName("comment_status")
        val commentStatus: String,
        val tag: List<TagDetail>,
        val category: List<TagDetail>,
        val content: String,
        val modify: String,
        val image: String,
        @SerializedName("alipay")
        val aliPay: String,
        @SerializedName("wechat")
        val weChat: String,
        val encrypt: Boolean
)

/**
 *  标签详细信息
 * @property name String 名字
 * @property link String 链接地址
 * @constructor
 */
data class TagDetail(
        val name: String,
        val link: String
)

/**
 *  评论数据
 * @property id Int 评论id
 * @property userId Int 用户id
 * @property nickname String 昵称
 * @property avatar String 头像
 * @property content String 评论内容
 * @property date String 评论时间
 * @property url String 评论地址
 * @property post_id Int 文章id
 * @property parent Int 父节点
 * @property hang String 头像挂件
 * @property level Int 等级
 * @property uid String B站uid
 * @constructor
 */
data class CommentDetail(
        val id:Int,
        @SerializedName("user_id")
        val userId:Int,
        val nickname:String,
        val avatar:String,
        val content:String,
        val date:String,
        val url:String,
        @SerializedName("post_id")
        val postId:Int,
        val parent:Int,
        val hang:String,
        val level:Int,
        val uid:String
)