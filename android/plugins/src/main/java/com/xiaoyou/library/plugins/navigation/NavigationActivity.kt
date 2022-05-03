package com.xiaoyou.library.plugins.navigation

import android.os.Bundle
import androidx.recyclerview.widget.StaggeredGridLayoutManager
import com.chad.library.adapter.base.entity.node.BaseNode
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.response.Link
import com.xiaoyou.library.net.entity.response.ProjectBottom
import com.xiaoyou.library.plugins.databinding.ActivityNavigationBinding
import com.xiaoyou.library.plugins.navigation.data.LinkNode
import com.xiaoyou.library.plugins.navigation.data.ParentNode
import com.xiaoyou.library.plugins.project.ProjectAdapter
import com.xiaoyou.library.plugins.viewmodel.PluginsVM

/**
 * @description
 * @author 小游
 * @data 2021/03/10
 */
class NavigationActivity :BaseDbActivity<PluginsVM,ActivityNavigationBinding>() {
    // 我的链接
    private val links:MutableList<BaseNode> = ArrayList()

    override fun initView(savedInstanceState: Bundle?) {
        // 初始化adapter
        val adapter = NodeAdapter()
        mDataBind.linkList.layoutManager =  StaggeredGridLayoutManager(
                3,
                StaggeredGridLayoutManager.VERTICAL
        )
        mDataBind.linkList.adapter = adapter

        // 获取导航链接
        mViewModel.getNavigationLink(object :DataReceive<List<Link>?>(){
            override fun success(data: List<Link>?) {
                if (data is List<Link>){
                    // 遍历节点添加数据
                    for (node in data){
                        val child:MutableList<BaseNode> = ArrayList()
                        // 先遍历子节点
                        for (item in node.child){
                            child.add(LinkNode(item.name,item.url))
                        }
                        links.add(ParentNode(node.id,node.color,node.name,child))
                    }
                    // 然后我们更新一下数据
                    adapter.setList(links)
                }
            }
        })
    }
}