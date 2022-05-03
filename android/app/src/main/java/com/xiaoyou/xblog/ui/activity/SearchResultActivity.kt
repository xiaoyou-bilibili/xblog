package com.xiaoyou.xblog.ui.activity

import android.content.Intent
import android.os.Bundle
import androidx.recyclerview.widget.LinearLayoutManager
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.ext.initBack
import com.xiaoyou.library.net.entity.base.PageInfo
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.library.net.entity.response.ReturnList
import com.xiaoyou.xblog.databinding.ActivitySearchResultBinding
import com.xiaoyou.xblog.ui.adapter.posts.PostAdapter
import com.xiaoyou.xblog.viewmodel.PostsVM
import kotlinx.android.synthetic.main.fragment_posts.*

/**
 * @description 搜索结果
 * @author 小游
 * @data 2021/02/25
 */
class SearchResultActivity : BaseDbActivity<PostsVM,ActivitySearchResultBinding>() {
    // 搜索关键词
    private var key:String = ""
    // 文章列表
    private val posts:MutableList<PostDetail> = ArrayList()
    // 页面信息
    private var page = PageInfo()

    // view初始化
    override fun initView(savedInstanceState: Bundle?) {
        // 首先获取关键词
        key = intent.getStringExtra("key")?:""
        // 设置标题
        mToolbar.initBack("${key}-搜索结果"){finish()}
        // 初始化adapter
        val adapter = PostAdapter(posts)
        mDataBind.postList.layoutManager = LinearLayoutManager(appContext)
        mDataBind.postList.adapter = adapter
        // 获取文章数据
        mViewModel.getPosts(page.reset(),key)
        // 刷新事件
        refreshLayout.apply {
            setOnRefreshListener{
                mViewModel.getPosts(page.reset(),key)
            }
            setOnLoadMoreListener{
                if (page.hasMore()){
                    mViewModel.getPosts(page.next(),key)
                } else {
                    finishLoadMoreWithNoMoreData()
                }
            }
        }

        // 直接监听文章数据变化
        mViewModel.posts.observe(this){
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

    // 界面点击重试
    override fun onLoadRetry() {
        // 重新加载界面
        mViewModel.getPosts(page.reset(),key)
    }

    override fun showToolBar() = true

}