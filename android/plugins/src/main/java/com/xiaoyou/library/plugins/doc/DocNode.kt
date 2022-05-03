package com.xiaoyou.library.plugins.doc

import com.xiaoyou.library.plugins.R
import com.zcolin.treeview.Node

/**
 * @description 节点信息
 * @author 小游
 * @data 2021/03/24
 */
class DocNode(val name:String,val id:Int, private val leaf:Boolean): Node {


    override fun getLayoutId() = R.layout.item_doc

    override fun isLeaf() = leaf
}