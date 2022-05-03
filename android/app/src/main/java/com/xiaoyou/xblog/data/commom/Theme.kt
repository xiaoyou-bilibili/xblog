package com.xiaoyou.xblog.data.commom

/**
 * @description 主题样式data
 * @author 小游
 * @data 2021/03/08
 */

/**
 *  自定义主题信息
 * @property color Int 当前主题颜色(主要是为了选择)
 * @property primary Int  主题颜色
 * @property accent Int 次级颜色
 * @property name String 主题名字
 * @property choose Boolean 是否选中
 * @constructor
 */
data class Theme (
    val color:Int,
    val primary:Int,
    val accent:Int,
    val name:String,
    val choose:Boolean
)