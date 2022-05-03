package com.xiaoyou.library.common.core

import com.tencent.mmkv.MMKV
import com.xiaoyou.library.common.R
import com.xiaoyou.library.common.data.ThemeEntity

/**
 * @description 自定义主题管理
 * @author 小游
 * @data 2021/03/09
 */
object ThemeManage {
    private val kv: MMKV by lazy {  MMKV.defaultMMKV() }
    // 用户id
    private const val PRIMARY = "primary"
    // 用户token信息
    private const val ACCENT = "accent"
    // 获取主题信息
    fun getTheme() = ThemeEntity(kv.getInt(PRIMARY, R.color.theme_pink), kv.getInt(ACCENT,R.color.theme_pink_accent))
    // 设置主题信息
    fun setTheme(primary:Int,accent:Int){
        kv.encode(PRIMARY,primary)
        kv.encode(ACCENT,accent)
    }
    // 删除主题信息
    fun clearTheme() {
        kv.removeValueForKey(PRIMARY)
        kv.removeValueForKey(ACCENT)
    }
}