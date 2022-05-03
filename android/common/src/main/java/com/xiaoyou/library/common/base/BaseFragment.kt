package com.xiaoyou.library.common.base

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment

/**
 * @description 基础fragment
 * @author 小游
 * @data 2021/02/20
 */
abstract class BaseFragment : Fragment() {
    // 界面布局
    abstract val layoutId: Int
    var dataBindView : View? = null
    // 自动完成布局映射
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(layoutId, container, false)
    }


}