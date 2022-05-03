package com.xiaoyou.library.net.entity.base

/**
 * @description 页面信息实体
 * @author 小游
 * @data 2021/02/22
 */
class PageInfo(var total:Int = 0,var current:Int = 1){
    fun next() = ++ current

    fun reset():Int{
        current = 0
        return 1
    }
    fun isFirst() = current == 1
    fun hasMore() = current < total
}