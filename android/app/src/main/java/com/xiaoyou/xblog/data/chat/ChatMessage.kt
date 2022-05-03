package com.xiaoyou.xblog.data.chat

import com.stfalcon.chatkit.commons.models.IMessage
import com.stfalcon.chatkit.commons.models.IUser
import java.util.*

/**
 * @description 聊天信息
 * @author 小游
 * @data 2021/02/22
 */
class ChatMessage( private val id :String,private val user: ChatUser,private val content:String,private val time:Date,val date:Long = 0): IMessage {
    // 聊天的唯一id
    override fun getId() = id
    // 聊天内容
    override fun getText() = content
    // 聊天的用户信息
    override fun getUser() = user
    // 发送时间
    override fun getCreatedAt() = time
}