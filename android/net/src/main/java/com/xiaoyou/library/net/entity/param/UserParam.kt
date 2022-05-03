package com.xiaoyou.library.net.entity.param

import com.google.gson.annotations.SerializedName

/**
 * @description 用户登录的请求参数
 * @author 小游
 * @data 2021/02/14
 */
data class UserLoginParam(
        var username: String = "",
        var password: String = ""
)

/**
 *  用户注册的参数
 * @property username String 用户名
 * @property nickname String 昵称
 * @property password String 密码
 * @property email String 用户邮箱
 * @property code String 手机验证码
 * @constructor
 */
data class UserRegisterParam(
        var username: String = "",
        var nickname: String = "",
        var password: String = "",
        var email: String = "",
        var code: String = ""
)

/**
 *  获取验证码参数
 * @property email String 邮箱
 * @property option String 用户操作
 * @constructor
 */
data class UserGetCodeParam(
        val email: String,
        val option: String
)

/**
 *  判断用户名是否存在
 * @property user String 用户名
 * @property email String 邮箱
 * @constructor
 */
data class UserGetUsername(
        val user: String,
        val email: String
)

/**
 *  重置密码
 * @property email String 邮箱
 * @property password String 密码
 * @property code String 密码
 * @constructor
 */
data class ResetPassword(
        var email: String = "",
        var password: String = "",
        var code: String = ""
)

/**
 *  更新用户信息
 * @property avatar String 头像
 * @property hang String 挂件
 * @property nickname String 昵称
 * @property email String 邮箱
 * @property sign String 签名
 * @property subscription Boolean 是否订阅
 * @property oldPassword String 旧密码
 * @property newPassword String 新密码
 * @constructor
 */
data class UpdateUserInfo(
        val avatar: String = "",
        val hang: String= "",
        val nickname: String= "",
        val email: String= "",
        val sign: String= "",
        val subscription: String= "",
        @SerializedName("old_password")
        val oldPassword: String= "",
        @SerializedName("new_password")
        val newPassword: String= ""
)