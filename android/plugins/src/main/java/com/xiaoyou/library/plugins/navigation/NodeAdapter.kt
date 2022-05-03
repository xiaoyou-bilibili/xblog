package com.xiaoyou.library.plugins.navigation

import com.chad.library.adapter.base.BaseNodeAdapter
import com.chad.library.adapter.base.entity.node.BaseNode
import com.xiaoyou.library.plugins.navigation.data.ParentNode
import com.xiaoyou.library.plugins.navigation.provider.LinkNodeProvider
import com.xiaoyou.library.plugins.navigation.provider.RootNodeProvider

/**
 * @description
 * @author 小游
 * @data 2021/03/30
 */
class NodeAdapter : BaseNodeAdapter() {
    init {
        // 这里我们把父节点还有子节点加进去
        addFullSpanNodeProvider(RootNodeProvider())
        addNodeProvider(LinkNodeProvider())
    }
    override fun getItemType(data: List<BaseNode>, position: Int): Int {
        val node = data[position]
        // 设置节点状态
        return if(node is ParentNode)0 else 1
    }
}