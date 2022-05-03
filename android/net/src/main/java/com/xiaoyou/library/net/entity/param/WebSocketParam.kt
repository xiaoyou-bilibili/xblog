package com.xiaoyou.library.net.entity.param

/**
 * @description webSocket提交的参数
 * @author 小游
 * @data 2021/03/01
 */

/**
 *  用户发送的信息
 * @param T 发送的数据类型
 * @property user_id Int 用户id
 * @property token String 用户token数据
 * @property to Int 发送的对象
 * @property option String 发送的类型
 * @property data T 发送的数据
 * @constructor
 */
data class SendMessage<T>(
    val user_id :Int,
    val token :String,
    val to :Int,
    val option:String,
    val data: T
)

/**
 *  用户发送的数据
 * @property message_type Int 数据类型
 * @property content String 数据内容
 * @constructor
 */
data class UserSend(
    val message_type: Int,
    val content: String
)

/**
 * 用户获取数据
 * @property date Long 如果是历史数据，设置历史数据时间
 * @property size Int 一次获取多少数据
 * @constructor
 */
data class UserGet(
    val date: Long,
    val size: Int
)
