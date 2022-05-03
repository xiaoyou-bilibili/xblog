package com.xiaoyou.xblog.ui.adapter.ui

import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.common.base.BaseAdapter
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.data.commom.ToolsItem
import com.xiaoyou.xblog.databinding.ItemToolsBinding

/**
 * @description 功能界面adapter
 * @author 小游
 * @data 2021/02/22
 */
class ToolsAdapter(data: MutableList<ToolsItem>) : BaseAdapter<ToolsItem, ItemToolsBinding>(R.layout.item_tools,data) {

    /**
     * Implement this method and use the helper to adapt the view to the given item.
     * 实现此方法，并使用 helper 完成 item 视图的操作
     * @param helper A fully initialized helper.
     * @param item   The item that needs to be displayed.
     */
    override fun convert(holder: BaseDataBindingHolder<ItemToolsBinding>, item: ToolsItem) {
        val binding = holder.dataBinding
        // 进行数据绑定
        binding?.item = item
    }
}