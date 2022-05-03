package com.xiaoyou.xblog.ui.fragment.index

import android.os.Bundle
import androidx.recyclerview.widget.LinearLayoutManager
import com.xiaoyou.library.common.base.BaseVmFragment
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.net.entity.base.PageInfo
import com.xiaoyou.library.net.entity.response.DiaryDetail
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.library.net.entity.response.ReturnList
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.ui.adapter.plugins.DiaryAdapter
import com.xiaoyou.xblog.ui.adapter.posts.PostAdapter
import com.xiaoyou.xblog.viewmodel.PluginsVM
import com.xiaoyou.xblog.viewmodel.PostsVM
import kotlinx.android.synthetic.main.fragment_diary.*
import kotlinx.android.synthetic.main.fragment_posts.*
import kotlinx.android.synthetic.main.fragment_posts.refreshLayout

/**
 * @description 日记界面fragment
 * @author 小游
 * @data 2021/02/20
 */
class DiaryFragment(override val layoutId: Int = R.layout.fragment_diary) : BaseVmFragment<PluginsVM>() {
    // 文章列表
    private val diary:MutableList<DiaryDetail> = ArrayList()
    // 页面信息
    private var page = PageInfo()

    /**
     * 初始化view操作
     */
    override fun initView(savedInstanceState: Bundle?) {
        // 初始化adapter
        val adapter = DiaryAdapter(diary)
        diaryList.layoutManager = LinearLayoutManager(appContext)
        diaryList.adapter = adapter

        // 刷新事件
        refreshLayout.apply {
            setOnRefreshListener{
                mViewModel.getDiary(page.reset())
            }
            setOnLoadMoreListener{
                if (page.hasMore()){
                    mViewModel.getDiary(page.next())
                } else {
                    finishLoadMoreWithNoMoreData()
                }
            }
        }

        // 直接监听数据变化
        mViewModel.diary.observe(viewLifecycleOwner){
            refreshLayout.finishRefresh()
            refreshLayout.finishLoadMore()
            if (it is ReturnList<DiaryDetail>){
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

    // 错误界面点击重试界面
    override fun onLoadRetry() {
        // 重新加载界面
        mViewModel.getDiary(page.reset())
    }
}