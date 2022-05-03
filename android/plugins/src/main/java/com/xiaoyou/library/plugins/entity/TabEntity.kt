package com.xiaoyou.library.plugins.entity

import com.flyco.tablayout.listener.CustomTabEntity

/**
 * @description 顶部指示器实体
 * @author 小游
 * @data 2021/02/21
 */
class TabEntity(private val title:String, private val select: Int = 0, private val unSelect:Int = 0) : CustomTabEntity {

    override fun getTabTitle() = title

    override fun getTabSelectedIcon() = select

    override fun getTabUnselectedIcon() = unSelect
}