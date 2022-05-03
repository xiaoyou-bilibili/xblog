package com.xiaoyou.library.net.core

/**
 * @description 不使用liveData，使用回调函数来获取数据
 * @author 小游
 * @data 2021/02/25
 */
abstract class DataReceive<T> : DataReceiveInterface<T>{
    /**
     *  当获取到数据时进行回调
     * @param data T 正确数据
     */
    override fun success(data: T) {}
    /**
     *  失败时的回调
     * @param message String
     */
    override fun error(message: String){}
}