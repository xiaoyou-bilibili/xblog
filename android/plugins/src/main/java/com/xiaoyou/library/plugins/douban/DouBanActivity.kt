package com.xiaoyou.library.plugins.douban

import android.os.Bundle
import androidx.recyclerview.widget.StaggeredGridLayoutManager
import com.flyco.tablayout.listener.CustomTabEntity
import com.flyco.tablayout.listener.OnTabSelectListener
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.net.entity.base.PageInfo
import com.xiaoyou.library.net.entity.response.DouBanDetail
import com.xiaoyou.library.net.entity.response.ReturnList
import com.xiaoyou.library.plugins.databinding.ActivityDouBanBinding
import com.xiaoyou.library.plugins.entity.TabEntity
import com.xiaoyou.library.plugins.viewmodel.PluginsVM
import kotlinx.android.synthetic.main.activity_animation.*


/**
 * @description 豆瓣界面
 * @author 小游
 * @data 2021/02/26
 */
class DouBanActivity  : BaseDbActivity<PluginsVM, ActivityDouBanBinding>(){
    // 顶部指示器
    private val choose = ArrayList<CustomTabEntity>().apply {
        add(TabEntity("图书"))
        add(TabEntity("电影"))
        add(TabEntity("音乐"))
    }
    // 豆瓣记录
    private val records:MutableList<DouBanDetail> = ArrayList()
    // 页面信息
    private var page = PageInfo()
    // 类型
    private var typeList = arrayListOf("book", "movie", "music")
    // 当前选中的内容
    private var type = typeList[0]

    // 视图初始化
    override fun initView(savedInstanceState: Bundle?) {
        // 实例化顶部指示器
        mDataBind.douBanTop.apply {
            // 设置顶部指示器内容
            setTabData(choose)
            // 顶部指示器点击事件
            setOnTabSelectListener(object : OnTabSelectListener {
                override fun onTabSelect(p: Int) {
                    // 设置类型
                    type = typeList[p]
                    // 重新获取数据
                    mViewModel.getDouBan(1, type)
                    // 设置数据加载
                    mDataBind.refreshLayout.setNoMoreData(false)
                }
                override fun onTabReselect(p: Int) {}
            })
        }
        // 初始化adapter
        val adapter = DouBanAdapter(records)
        mDataBind.douBanList.layoutManager =  StaggeredGridLayoutManager(
            3,
            StaggeredGridLayoutManager.VERTICAL
        )
        mDataBind.douBanList.adapter = adapter
        // 刷新事件
        refreshLayout.apply {
            setOnRefreshListener{
                mViewModel.getDouBan(page.reset(), type)
            }
            setOnLoadMoreListener{
                if (page.hasMore()){
                    mViewModel.getDouBan(page.next(), type)
                } else {
                    finishLoadMoreWithNoMoreData()
                }
            }
        }
        // 直接监听文章数据变化
        mViewModel.douBans.observe(this){
            refreshLayout.finishRefresh()
            refreshLayout.finishLoadMore()
            if (it is ReturnList<DouBanDetail>){
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
        // 设置豆瓣界面的点击事件
        adapter.setOnItemClickListener { _, _, position ->
            val data = adapter.getItem(position)
            DouBanDialog(this,data).show()
        }
    }

    override fun onLoadRetry() {
        // 重新加载界面
        mViewModel.getDouBan(page.reset(), type)
    }

}