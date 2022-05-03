package com.xiaoyou.library.common.data

/**
 * @description 主题实体类
 * @author 小游
 * @data 2021/03/09
 */
/**
 *  存储着主题相关的信息
 * @property primary Int 主题颜色
 * @property accent Int 次要颜色
 * @constructor
 */
data class ThemeEntity (
        val primary: Int,
        val accent: Int
)