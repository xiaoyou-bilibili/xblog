package com.xiaoyou.library.plugins.navigation.data

import com.chad.library.adapter.base.entity.node.BaseNode

/**
 * @description 链接节点，显示链接信息
 * @author 小游
 * @data 2021/03/30
 */
data class LinkNode(
        val name:String,
        val url:String,
        override val childNode: MutableList<BaseNode>? = null
) : BaseNode()