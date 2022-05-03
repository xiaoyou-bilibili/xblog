package com.xiaoyou.library.plugins.animation

import android.content.Intent
import android.os.Bundle
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.StaggeredGridLayoutManager
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.entity.base.PageInfo
import com.xiaoyou.library.net.entity.response.AnimationDetail
import com.xiaoyou.library.net.entity.response.Animations
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.library.net.entity.response.ReturnList
import com.xiaoyou.library.plugins.databinding.ActivityAnimationBinding
import com.xiaoyou.library.plugins.viewmodel.PluginsVM
import kotlinx.android.synthetic.main.activity_animation.*

/**
 * @description 我的追番的activity界面
 * @author 小游
 * @data 2021/02/26
 */
class AnimationActivity : BaseDbActivity<PluginsVM,ActivityAnimationBinding>(){

    // 追番数据
    private val animations:MutableList<AnimationDetail> = ArrayList()

    // 页面信息
    private var page = PageInfo()

    override fun initView(savedInstanceState: Bundle?) {
        // 初始化adapter
        val adapter = AnimationAdapter(animations)
        mDataBind.animationList.layoutManager =  StaggeredGridLayoutManager(
            3,
            StaggeredGridLayoutManager.VERTICAL
        )
        mDataBind.animationList.adapter = adapter

        // 刷新事件
        refreshLayout.apply {
            setOnRefreshListener{
                mViewModel.getAnimation(page.reset())
            }
            setOnLoadMoreListener{
                if (page.hasMore()){
                    mViewModel.getAnimation(page.next())
                } else {
                    finishLoadMoreWithNoMoreData()
                }
            }
        }

        // 直接监听文章数据变化
        mViewModel.animations.observe(this){
            refreshLayout.finishRefresh()
            refreshLayout.finishLoadMore()
            if (it is Animations){
                page.current = it.current
                page.total = it.total
                // 设置追番总数
                mDataBind.animationTotal.text = it.num.toString()
                    // 判断不同的页数
                if (page.isFirst()){
                    adapter.setList(it.contents)
                } else {
                    adapter.addData(it.contents)
                }
            }
        }

        // 设置追番界面的点击事件
        adapter.setOnItemClickListener { _, _, position ->
//            XLog.e("点击事件")
            val data = adapter.getItem(position)
            AnimationDialog(this,data).show()
        }
    }

    override fun onLoadRetry() {
        // 重新加载界面
        mViewModel.getAnimation(page.reset())
    }
}