package com.xiaoyou.library.net.entity.param

/**
 * @description 工具板块请求参数
 * @author 小游
 * @data 2021/02/19
 */

/**
 *  用户反馈意见
 * @property concat String 反馈的意见
 * @property content String 联系方式
 * @constructor
 */
data class SubmitAdvice(
        val concat: String,
        val content: String
)