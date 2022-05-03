package com.xiaoyou.library.net.entity.response

/**
 * @description 一些通用的数据结构
 * @author 小游
 * @data 2021/02/13
 */

/**
 *  统一返回的列表数据
 * @param T 列表数据内容
 * @property total Int 总页数
 * @property current Int 当前第几页
 * @property contents List<T> 内容
 * @constructor
 */
data class ReturnList<T>(val total:Int,val current:Int,val contents: List<T>)

/**
 *  返回的错误内容
 * @property message String 错误信息
 * @constructor
 */
data class ReturnError(val message: String)