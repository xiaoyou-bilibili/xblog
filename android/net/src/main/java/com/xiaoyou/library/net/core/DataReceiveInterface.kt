package com.xiaoyou.library.net.core

/**
 * @description 自定义数据接收接口
 * @author 小游
 * @data 2021/02/25
 */
interface DataReceiveInterface<T>{
    /**
     *  当获取到数据时进行回调
     * @param data T 正确数据
     */
    fun success(data: T)

    /**
     *  失败时的回调
     * @param message String 错误消息
     */
     fun error(message: String){}
}