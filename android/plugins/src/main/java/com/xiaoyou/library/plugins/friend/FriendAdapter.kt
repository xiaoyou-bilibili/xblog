package com.xiaoyou.library.plugins.friend

import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.common.base.BaseAdapter
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.net.entity.response.FriendDetail
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.databinding.ItemFriendBinding

/**
 * @description
 * @author 小游
 * @data 2021/02/26
 */
class FriendAdapter(data: MutableList<FriendDetail>) : BaseAdapter<FriendDetail, ItemFriendBinding>(R.layout.item_friend,data){

    override fun convert(holder: BaseDataBindingHolder<ItemFriendBinding>, item: FriendDetail) {
        val binding = holder.dataBinding
        // 进行数据绑定
        binding?.let {
            it.item = item
            // 加载头像
            ImageUtil.setImageByUrl(item.avatar,it.userAvatar)
        }
    }
}