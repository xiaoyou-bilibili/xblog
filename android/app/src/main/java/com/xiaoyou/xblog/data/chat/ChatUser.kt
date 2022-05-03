package com.xiaoyou.xblog.data.chat

import com.stfalcon.chatkit.commons.models.IUser

/**
 * @description 聊天的用户信息
 * @author 小游
 * @data 2021/02/22
 */
class ChatUser(val id:Int, private val nickname:String, private val avatar:String) : IUser {
    // 用户id
    override fun getId() = id.toString()
    // 用户昵称
    override fun getName() = nickname
    // 用户头像
    override fun getAvatar() = avatar
}