package com.xiaoyou.library.plugins.navigation.provider

import android.graphics.Color
import android.widget.LinearLayout
import android.widget.TextView
import com.chad.library.adapter.base.entity.node.BaseNode
import com.chad.library.adapter.base.provider.BaseNodeProvider
import com.chad.library.adapter.base.viewholder.BaseViewHolder
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.navigation.data.ParentNode


/**
 * @description 父节点
 * @author 小游
 * @data 2021/03/30
 */
class RootNodeProvider(override val itemViewType: Int = 0, override val layoutId: Int = R.layout.item_link_root) :BaseNodeProvider(){
    override fun convert(helper: BaseViewHolder, item: BaseNode) {
        // 获取数据
        val data = item as ParentNode
        // 设置文字
        helper.getView<TextView>(R.id.navigationName).text = data.title
        // 设置颜色
        helper.getView<LinearLayout>(R.id.navigationColor).setBackgroundColor(Color.parseColor(data.color))
        helper.getView<LinearLayout>(R.id.navigationLine).setBackgroundColor(Color.parseColor(data.color))
    }

}