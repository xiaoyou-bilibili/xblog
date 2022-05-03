package com.xiaoyou.library.common.base

import android.content.Context
import androidx.databinding.DataBindingUtil
import androidx.databinding.ViewDataBinding
import com.chad.library.adapter.base.BaseQuickAdapter
import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.net.entity.response.PostDetail

/**
 * @description 基础adapter，继承BaseQuickAdapter，集成了dataBind功能
 * @author 小游
 * @data 2021/02/21
 */
abstract class BaseAdapter<T,B: ViewDataBinding>(layoutId :Int,data: MutableList<T>) :
    BaseQuickAdapter<T, BaseDataBindingHolder<B>>(layoutId,data){


    /**
     *  当ViewHolder创建的时候我们就进行数据绑定
     * @param viewHolder BaseViewHolder
     * @param viewType Int
     */
    override fun onItemViewHolderCreated(viewHolder: BaseDataBindingHolder<B>, viewType: Int) {
        // 开启动画效果
        animationEnable = true
        // 动画一直持续
        isAnimationFirstOnly = false
        // 设置动画效果
        this.setAnimationWithDefault(AnimationType.SlideInLeft)
        // 进行数据绑定
        DataBindingUtil.bind<B>(viewHolder.itemView)
    }

}