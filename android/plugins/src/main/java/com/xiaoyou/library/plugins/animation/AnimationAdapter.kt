package com.xiaoyou.library.plugins.animation

import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.common.base.BaseAdapter
import com.xiaoyou.library.net.entity.response.AnimationDetail
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.databinding.ItemAnimationBinding

/**
 * @description 我的追番的adapter
 * @author 小游
 * @data 2021/02/26
 */
class AnimationAdapter (data: MutableList<AnimationDetail>) : BaseAdapter<AnimationDetail, ItemAnimationBinding>(
    R.layout.item_animation,data){
    override fun convert(
        holder: BaseDataBindingHolder<ItemAnimationBinding>,
        item: AnimationDetail
    ) {
        val binding = holder.dataBinding
        // 进行数据绑定
        binding?.let {
            it.item = item
            // 加载头像
            it.animationImg.setImageURI(item.cover)
        }
    }
}