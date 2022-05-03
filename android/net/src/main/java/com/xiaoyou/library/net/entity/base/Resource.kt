package com.xiaoyou.library.net.entity.base

/**
 * @description 自己封装的一个结果返回类
 * @author 小游
 * @data 2021/02/13
 */
class Resource<out T>(val status: Status, val data: T?, val message: String) {
    companion object {
        // 这里我们确定了三种状态
        fun <T> success(data: T?) = Resource(Status.SUCCESS, data, "")
        fun <T> error(msg: String, data: T? ) = Resource(Status.ERROR, data, msg)
        fun <T> loading(data: T?) = Resource(Status.LOADING, data, "")
    }
}