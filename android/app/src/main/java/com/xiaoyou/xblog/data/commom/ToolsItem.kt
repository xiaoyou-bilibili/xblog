package com.xiaoyou.xblog.data.commom

import android.graphics.drawable.Drawable

/**
 * @description 工具界面实体
 * @author 小游
 * @data 2021/02/22
 */

/**
 *  自定义功能布局
 * @property drawable Drawable 图片
 * @property title String 标题
 * @constructor
 */
data class ToolsItem(
    val drawable: Drawable?,
    var title: String,
    val activity: Class<*>? = null
)