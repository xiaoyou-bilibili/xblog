package com.xiaoyou.library.plugins.navigation.provider

import android.net.Uri
import android.view.View
import android.widget.TextView
import com.chad.library.adapter.base.entity.node.BaseNode
import com.chad.library.adapter.base.provider.BaseNodeProvider
import com.chad.library.adapter.base.viewholder.BaseViewHolder
import com.facebook.drawee.view.SimpleDraweeView
import com.xiaoyou.library.common.util.Common
import com.xiaoyou.library.net.entity.response.Link
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.navigation.data.LinkNode

/**
 * @description
 * @author 小游
 * @data 2021/03/30
 */
class LinkNodeProvider(override val itemViewType: Int = 1, override val layoutId: Int = R.layout.item_link) : BaseNodeProvider(){

    override fun convert(helper: BaseViewHolder, item: BaseNode) {
        // 获取数据
        val data = item as LinkNode
        // 设置网站名字
        helper.getView<TextView>(R.id.linkName).text = data.name
        // 设置网站图片
        helper.getView<SimpleDraweeView>(R.id.linkIcon).setImageURI("https://statics.dnspod.cn/proxy_favicon/_/favicon?domain="+(Uri.parse(data.url)).host)
    }
    // 设置点击事件
    override fun onClick(helper: BaseViewHolder, view: View, data: BaseNode, position: Int) {
        // 点击事件
        Common.openUrl((data as LinkNode).url)
    }
}