package com.xiaoyou.xblog.ui.adapter.common

import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.common.base.BaseAdapter
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.data.commom.Theme
import com.xiaoyou.xblog.databinding.ItemThemeBinding

/**
 * @description 主题选择的adapter
 * @author 小游
 * @data 2021/03/08
 */
class ThemeAdapter(data: MutableList<Theme>)  : BaseAdapter<Theme, ItemThemeBinding>(R.layout.item_theme,data){
    override fun convert(holder: BaseDataBindingHolder<ItemThemeBinding>, item: Theme) {
        val binding = holder.dataBinding
        // 进行数据绑定
        binding?.let {
            it.item = item
        }
    }
}