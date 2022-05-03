package com.xiaoyou.library.plugins.douban

import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.common.base.BaseAdapter
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.net.entity.response.AnimationDetail
import com.xiaoyou.library.net.entity.response.DouBanDetail
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.databinding.ItemAnimationBinding
import com.xiaoyou.library.plugins.databinding.ItemDouBanBinding

/**
 * @description 我的豆瓣的adapter
 * @author 小游
 * @data 2021/02/26
 */
class DouBanAdapter (data: MutableList<DouBanDetail>) : BaseAdapter<DouBanDetail, ItemDouBanBinding>(
    R.layout.item_dou_ban,data){
    override fun convert(holder: BaseDataBindingHolder<ItemDouBanBinding>, item: DouBanDetail) {
        val binding = holder.dataBinding
        // 进行数据绑定
        binding?.let {
            it.item = item
            // 加载头像
            it.animationImg.setImageURI(item.picture)
        }
    }

}