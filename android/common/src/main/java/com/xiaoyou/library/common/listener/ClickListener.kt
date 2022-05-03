package com.xiaoyou.library.common.listener


/**
 * @description 自己定义的点击回调事件
 * @author 小游
 * @data 2021/03/07
 */
interface MyClickListener<T> {
    /**
     *  点击触发事件
     * @param data T
     */
    fun onClick(data:T)
}