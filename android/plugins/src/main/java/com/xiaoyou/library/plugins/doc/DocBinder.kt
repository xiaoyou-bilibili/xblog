package com.xiaoyou.library.plugins.doc

import android.view.View
import android.widget.TextView
import com.xiaoyou.library.plugins.R
import com.zcolin.treeview.TreeNode
import com.zcolin.treeview.TreeViewBinder

/**
 * @description 文档binder对象
 * @author 小游
 * @data 2021/03/24
 */
class DocBinder : TreeViewBinder<DocBinder.ViewHolder>() {
    // viewHolder绑定当前布局
    class ViewHolder(root :View): TreeViewBinder.ViewHolder(root){
        var name: TextView = root.findViewById(R.id.tv_name)
    }
    // 获取布局对象
    override fun getLayoutId() = R.layout.item_doc
    // 提供viewHolder对象
    override fun provideViewHolder(itemView: View?) = ViewHolder(itemView!!)
    // 设置布局信息
    override fun bindView(holder: ViewHolder?, position: Int, node: TreeNode<*>?) {
       val nodes = node?.content as DocNode
        holder?.let {
            it.name.text = nodes.name
        }
    }
}