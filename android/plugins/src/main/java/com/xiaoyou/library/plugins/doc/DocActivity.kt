package com.xiaoyou.library.plugins.doc

import android.annotation.SuppressLint
import android.os.Bundle
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.base.WEB
import com.xiaoyou.library.common.ext.isVisible
import com.xiaoyou.library.common.ext.visibleOrGone
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.response.DocContent
import com.xiaoyou.library.net.entity.response.DocList
import com.xiaoyou.library.plugins.databinding.ActivityDocBinding
import com.xiaoyou.library.plugins.viewmodel.PluginsVM
import com.zcolin.treeview.TreeNode
import com.zcolin.treeview.TreeViewAdapter
import java.util.*
import kotlin.collections.ArrayList

/**
 * @description
 * @author 小游
 * @data 2021/03/10
 */
class DocActivity : BaseDbActivity<PluginsVM, ActivityDocBinding>() {
    // 当点击的id
    private var now = 0

    override fun initView(savedInstanceState: Bundle?) {
        mDataBind.docList.layoutManager = LinearLayoutManager(this)
        // 获取文档数据
        mViewModel.getDocs(object: DataReceive<List<DocList>?>(){
            override fun success(data: List<DocList>?) {
                if (data is List<DocList>){
                    setAdapter(data)
                }
            }
        })
        // 侧边栏点击事件
        mDataBind.menu.setOnClickListener{clickMenu()}
        // 点击web时跳转
        mDataBind.postContentWeb.setOnClickListener{clickMenu()}
    }
    // 初始化adapter
    private fun setAdapter(data :List<DocList>){
        val adapter = TreeViewAdapter(parseData(data), listOf(DocBinder()))
        adapter.setOnTreeNodeListener(object : TreeViewAdapter.OnTreeNodeListener {
            override fun onClick(node: TreeNode<*>?, holder: RecyclerView.ViewHolder?): Boolean {
                val content = node!!.content as DocNode
                // 如果是叶子直接获取文章
                if (node.isLeaf){
                    getContent(content.id)
                }
                return false
            }
            override fun onToggle(isExpand: Boolean, holder: RecyclerView.ViewHolder?) {

            }
        })
        mDataBind.docList.adapter = adapter
    }

    // 获取文章内容
    private fun getContent(id:Int){
        mViewModel.getDocContent(id,object :DataReceive<DocContent?>(){
            override fun success(data: DocContent?) {
                if (data is DocContent){
                    setWeb(data.content)
                    mToolbar.setCenterTitle(data.title)
                    clickMenu()
                }
            }
        })
    }

    // 点击按钮事件
    private fun clickMenu(){
        if(mDataBind.docList.isVisible()){
            mDataBind.docList.visibleOrGone(false)
        } else {
            mDataBind.docList.visibleOrGone(true)
        }
    }


    /**
     *  显示文章界面内容
     * @param web String 文章内容
     */
    @SuppressLint("SetJavaScriptEnabled")
    private fun setWeb(web:String){
        // 加入CSS和JS
        val content = "<!DOCTYPE html><html><head><meta charset=utf-8><link href='$WEB/static/css/blog-post.css' rel='stylesheet'></head>" +
                "<body><div class='content'>${web}</div><script src='$WEB/static/js/prism.js'></script></body></html>"
        // webView设置内容
        mDataBind.postContentWeb.apply {
            // 加载内容
            loadDataWithBaseURL(null,content,"text/html","UTF-8",null)
            // 允许js
            settings.javaScriptEnabled = true
        }
    }

    // 解析节点数据
    private fun parseData(data :List<DocList>):List<TreeNode<DocNode>>{
        val row :MutableList<DocList> = data as MutableList<DocList>
        val nodes = LinkedList<TreeNode<DocNode>>()
        for (i in data.count() - 1 downTo 0){
            // 先遍历根节点
            if (row[i].parent==0){
                val parent=TreeNode(DocNode(row[i].title,row[i].id,false))
                val id = row[i].id
                row.removeAt(i)
                // 解析子节点
                parseChild(row,id,parent)
                nodes.add(parent)
            }
        }
        return nodes
    }

    // 解析子节点
    private fun parseChild(data: MutableList<DocList>?,parent:Int,nodes: TreeNode<DocNode>){
        if (data.isNullOrEmpty()){
            return
        }
        for (i in data.count() - 1 downTo 0){
            if(i>=data.size){
                break
            }
            // 先遍历根节点
            if (data[i].parent==parent){
                var node=TreeNode(DocNode(data[i].title,data[i].id,false))
                val tmp = data[i]
                data.removeAt(i)
                // 解析子节点
                parseChild(data,tmp.id,node)
                // 判断当前节点下是否有孩子,没有孩子就是叶子节点
                if (node.childList.isNullOrEmpty()){
                    node = TreeNode(DocNode(tmp.title,tmp.id,true))
                }
                nodes.addChild(node)
            }
        }
    }
}