package com.xiaoyou.library.net.entity.param

/**
 * @description
 * @author 小游
 * @data 2021/03/07
 */

/**
 *  提交友链申请
 * @property name String 友链名字
 * @property site String 友链地址
 * @property dec String 友链描述
 * @property avatar String 友链头像
 * @property email String 友链邮箱
 * @constructor
 */
data class SubmitFriend(
        var name:String="",
        var site:String="",
        var dec:String="",
        var avatar:String="",
        var email:String="",
)