package com.xiaoyou.library.common.base

import androidx.fragment.app.Fragment
import androidx.fragment.app.FragmentActivity
import androidx.viewpager2.adapter.FragmentStateAdapter

/**
 * @description 自己定义的viewpage的adapter
 * @author 小游
 * @data 2021/02/20
 */
class FragmentAdapter(private val fragments: List<Fragment>, fa: FragmentActivity) : FragmentStateAdapter(fa) {
    // 返回fragment的树木
    override fun getItemCount() = fragments.size

   // 创建fragment
    override fun createFragment(position: Int) = fragments[position]
}