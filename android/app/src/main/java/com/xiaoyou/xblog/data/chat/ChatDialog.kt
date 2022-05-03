package com.xiaoyou.xblog.data.chat

import com.stfalcon.chatkit.commons.models.IDialog
import com.stfalcon.chatkit.commons.models.IUser

/**
 * @description 聊天界面的信息
 * @author 小游
 * @data 2021/02/22
 */
class ChatDialog(val id:Int, private val avatar:String, val name:String, var message: ChatMessage?,val count:Int): IDialog<ChatMessage> {
    // 群组id，这里我们可以直接使用用户id
    override fun getId() = id.toString()
    // 群组的头像
    override fun getDialogPhoto() =  avatar
    // 群组的名字
    override fun getDialogName() = name
    // 用户信息
    override fun getUsers() = listOf(ChatUser(0,"",avatar))
    // 最后一条消息
    override fun getLastMessage() = message
    // 手动设置最后的信息
    override fun setLastMessage(message: ChatMessage?) {
        this.message = message
    }
    // 获取未读数目
    override fun getUnreadCount() = count
}