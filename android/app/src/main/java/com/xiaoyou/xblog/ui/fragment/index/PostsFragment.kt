package com.xiaoyou.xblog.ui.fragment.index

import android.content.Intent
import android.os.Bundle
import android.view.View
import android.widget.AdapterView
import androidx.recyclerview.widget.LinearLayoutManager
import com.chad.library.adapter.base.BaseQuickAdapter
import com.chad.library.adapter.base.listener.OnItemClickListener
import com.xiaoyou.library.common.base.BaseVmFragment
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.entity.base.PageInfo
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.library.net.entity.response.ReturnList
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.ui.activity.PostActivity
import com.xiaoyou.xblog.ui.adapter.posts.PostAdapter
import com.xiaoyou.xblog.viewmodel.PostsVM
import kotlinx.android.synthetic.main.fragment_posts.*

/**
 * @description 文章界面fragment
 * @author 小游
 * @data 2021/02/20
 */
class PostsFragment(override val layoutId: Int = R.layout.fragment_posts) : BaseVmFragment<PostsVM>() {
    // 文章列表
    private val posts:MutableList<PostDetail> = ArrayList()
    // 页面信息
    private var page = PageInfo()


    /**
     * 初始化view操作
     */
    override fun initView(savedInstanceState: Bundle?) {
        // 初始化adapter
        val adapter = PostAdapter(posts)
        postList.layoutManager = LinearLayoutManager(appContext)
        postList.adapter = adapter

        // 刷新事件
        refreshLayout.apply {
            setOnRefreshListener{
                mViewModel.getPosts(page.reset())
            }
            setOnLoadMoreListener{
                if (page.hasMore()){
                    mViewModel.getPosts(page.next())
                } else {
                    finishLoadMoreWithNoMoreData()
                }
            }
        }

        // 直接监听文章数据变化
        mViewModel.posts.observe(viewLifecycleOwner){
            refreshLayout.finishRefresh()
            refreshLayout.finishLoadMore()
            if (it is ReturnList<PostDetail>){
                page.current = it.current
                page.total = it.total
                // 判断不同的页数
                if (page.isFirst()){
                    adapter.setList(it.contents)
                } else {
                    adapter.addData(it.contents)
                }
            }
        }
    }

    override fun onLoadRetry() {
        // 重新加载界面
        mViewModel.getPosts(page.reset())
    }
}