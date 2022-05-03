package com.xiaoyou.library.net.entity.response

import com.google.gson.annotations.SerializedName

/**
 * @description 用户板块返回的数据
 * @author 小游
 * @data 2021/02/14
 */

/**
 *  token数据
 * @property userId Int 用户id
 * @property token String 用户token
 * @constructor
 */
data class TokenDetail(
        @SerializedName("user_id")
        val userId:Int,
        val token:String
)

/**
 * 用户详细信息
 * @property avatar String 头像
 * @property sign String 个性签名
 * @property level Int 等级
 * @property hang String 头像挂件
 * @property username String 用户名
 * @property nickname String 昵称
 * @property email String 邮箱地址
 * @property user_id Int 用户id
 * @property identity Int 用户身份
 * @property subscription String 用户是否订阅邮件
 * @constructor
 */
data class UserDetail(
        var avatar:String="",
        var sign:String="",
        val level: Int=0,
        val hang:String="",
        val username:String="",
        var nickname:String="",
        var email:String="",
        val user_id:Int=0,
        val identity:Int=0,
        val subscription:Boolean=false
)