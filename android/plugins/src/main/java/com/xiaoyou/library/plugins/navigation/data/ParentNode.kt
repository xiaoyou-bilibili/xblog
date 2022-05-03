package com.xiaoyou.library.plugins.navigation.data

import com.chad.library.adapter.base.entity.node.BaseNode

/**
 * @description 父节点，显示分类信息
 * @author 小游
 * @data 2021/03/30
 */
data class ParentNode(
        val id:Int,
        val color:String,
        val title:String,
        override val childNode: MutableList<BaseNode>? = null
) :BaseNode()