package com.xiaoyou.library.net.entity.response

/**
 * @description webSocket的响应体
 * @author 小游
 * @data 2021/03/01
 */
data class WsReturn<T>(
    val code :Int,
    val message:String,
    val data: T
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
data class ChatInfo(
    val id:String,
    val user_id:Int,
    val content: String,
    val avatar:String,
    val nickname:String,
    val date:Long,
    val target:Int,
    val message_type:Int,
    val read:Boolean
)

