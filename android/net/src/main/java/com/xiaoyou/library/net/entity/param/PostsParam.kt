package com.xiaoyou.library.net.entity.param

import com.google.gson.annotations.SerializedName

/**
 * @description 文章板块请求参数
 * @author 小游
 * @data 2021/02/19
 */

/**
 *  文章状态
 * @property good Boolean 点赞
 * @property collection Boolean 收藏
 * @constructor
 */
data class PostsStatus(
        val good:Boolean = false,
        val collection: Boolean = false
)

/**
 *  发表评论
 * @property content String 评论内容
 * @property name String 评论昵称
 * @property userId Int 用户id
 * @property parent Int 父节点
 * @property avatar String 头像
 * @constructor
 */
data class PostComment(
        val content:String,
        val name:String,
        @SerializedName("user_id")
        val userId:Int,
        val parent:Int,
        val avatar:String,
        val email:String
)

/**
 *  加密文章内容
 * @property id Int 文章id
 * @property content String 文章内容
 * @constructor
 */
data class EncryptContent(
        val id:Int,
        val content: String
)