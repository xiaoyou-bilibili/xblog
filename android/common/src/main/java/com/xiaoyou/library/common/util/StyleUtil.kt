package com.xiaoyou.library.common.util

import android.content.res.Resources
import android.util.DisplayMetrics
import com.xiaoyou.library.common.base.appContext
import kotlin.math.acos


/**
 * @description 和页面样式有关的工具类
 * @author 小游
 * @data 2021/02/24
 */
object StyleUtil {
    /**
     * 获取系统的屏幕宽度
     * @return 返回当前宽度(单位px)
     */
    fun getWidth(): Int {
        val resources: Resources = appContext.resources
        val dm: DisplayMetrics = resources.displayMetrics
        return dm.widthPixels
    }
}