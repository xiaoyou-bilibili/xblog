package com.xiaoyou.xblog.ui.fragment.homepage

import android.os.Bundle
import androidx.fragment.app.Fragment
import androidx.viewpager2.widget.ViewPager2
import com.flyco.tablayout.listener.CustomTabEntity
import com.flyco.tablayout.listener.OnTabSelectListener
import com.xiaoyou.library.common.base.BaseVmFragment
import com.xiaoyou.xblog.R
import com.xiaoyou.library.plugins.entity.TabEntity
import com.xiaoyou.library.common.base.FragmentAdapter
import com.xiaoyou.xblog.ui.fragment.index.DiaryFragment
import com.xiaoyou.xblog.ui.fragment.index.PostsFragment
import com.xiaoyou.xblog.viewmodel.PostsVM
import kotlinx.android.synthetic.main.fragment_index.*

/**
 * @description 主页界面的fragment
 * @author 小游
 * @data 2021/02/21
 */
class IndexFragment(override val layoutId: Int = R.layout.fragment_index) : BaseVmFragment<PostsVM>() {
    private val choose = ArrayList<CustomTabEntity>().apply {
        add(TabEntity("文章"))
        add(TabEntity("日记"))
    }
    private val pages = ArrayList<Fragment>().apply {
        add(PostsFragment())
        add(DiaryFragment())
    }

    /**
     * 初始化view操作
     */
    override fun initView(savedInstanceState: Bundle?) {
        // 添加数据
        initUI()
    }

    // UI初始化
    private fun initUI() {
        // 设置指示器内容
        topChoose.setTabData(choose)
        // 加载界面
        indexPage.adapter = FragmentAdapter(pages, this.mActivity)
        // viewpage页面切换时进行的操作
        indexPage.registerOnPageChangeCallback(object : ViewPager2.OnPageChangeCallback() {
            override fun onPageSelected(position: Int) {
                topChoose.currentTab = position
            }
        })
        // 顶部选择栏切换时操作
        topChoose.setOnTabSelectListener(object : OnTabSelectListener {
            override fun onTabSelect(p0: Int) {
                indexPage.currentItem = p0
            }
            override fun onTabReselect(p0: Int) {}
        })
    }

}