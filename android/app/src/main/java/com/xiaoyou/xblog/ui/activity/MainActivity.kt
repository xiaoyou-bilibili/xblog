package com.xiaoyou.xblog.ui.activity

import android.content.Intent
import android.os.Bundle
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import androidx.viewpager2.widget.ViewPager2
import com.xiaoyou.library.common.base.BaseVmActivity
import com.xiaoyou.library.common.ext.visibleOrGone
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.library.net.entity.response.ReturnList
import com.xiaoyou.xblog.R
import com.xiaoyou.library.common.base.FragmentAdapter
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.Common
import com.xiaoyou.library.common.util.Setting
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.response.AppSetting
import com.xiaoyou.library.plugins.viewmodel.SettingVM
import com.xiaoyou.xblog.ui.fragment.homepage.ChatFragment
import com.xiaoyou.xblog.ui.fragment.homepage.IndexFragment
import com.xiaoyou.xblog.ui.fragment.homepage.MeFragment
import com.xiaoyou.xblog.ui.fragment.homepage.ToolsFragment
import com.xiaoyou.xblog.viewmodel.PostsVM
import com.xiaoyouProject.searchbox.SearchFragment
import com.xiaoyouProject.searchbox.custom.IOnSearchClickListener
import com.xiaoyouProject.searchbox.entity.CustomLink
import kotlinx.android.synthetic.main.activity_main.*
import nl.joery.animatedbottombar.AnimatedBottomBar

/**
 * @description 主界面
 * @author 小游
 * @data 2021/02/21
 */
class MainActivity(override val layoutId: Int = R.layout.activity_main) :
    BaseVmActivity<PostsVM>() {

    // 设置界面viewModel
    private val settingVN by lazy { ViewModelProvider(this).get(SettingVM::class.java) }


    // 主页显示的标题栏
    val indexList = arrayListOf("主页", "功能", "聊天","个人中心")
    // 自定义searchView
    private val search = SearchFragment<PostDetail>()

    //viewPager对应的Fragment
    private val fragments = ArrayList<Fragment>().apply {
        add(IndexFragment())
        add(ToolsFragment())
        add(ChatFragment())
        add(MeFragment())
    }

    override fun initView(savedInstanceState: Bundle?) {
        mToolbar.setCenterTitle(R.string.main)
        // 获取数据
        //mViewModel.getPostContent(48)
        // 初始化UI
        initUI()
        // 获取设置信息
        settingVN.getAppSetting(object :DataReceive<AppSetting?>(){
            override fun success(data: AppSetting?) {
                if (data is AppSetting){
                    // 保存当前设置信息
                    Setting.setSetting(data)
                }
            }
        })
        // 判断是否为本站链接
        Common.checkLink(this,PostActivity::class.java)
        // 检查更新
        Common.checkUpdate()
    }

    /**
     *  UI界面初始化
     */
    private fun initUI() {
        // 加载界面
        mainPage.adapter = FragmentAdapter(fragments, this)
        // 进行配置 设置应保留在当前可见页面两侧的页面数
        mainPage.offscreenPageLimit = fragments.size
        // viewpage页面切换时进行的操作
        mainPage.registerOnPageChangeCallback(object : ViewPager2.OnPageChangeCallback() {
            // 页面切换的事件
            override fun onPageSelected(position: Int) {
                bottomNavigation.selectTabAt(position)
                // 设置标题
                mToolbar.setCenterTitle(indexList[position])
                // 设置搜索按钮是否显示
                mToolbar.setSearchIcon(position == 0)
                // 设置顶部标题栏是否显示
                mToolbar.visibleOrGone(position != 3)
            }
        })
        // 监听底部操作
        bottomNavigation.setOnTabSelectListener(object : AnimatedBottomBar.OnTabSelectListener {
            override fun onTabSelected(
                lastIndex: Int,
                lastTab: AnimatedBottomBar.Tab?,
                newIndex: Int,
                newTab: AnimatedBottomBar.Tab
            ) {
                mainPage.currentItem = newIndex
            }
            override fun onTabReselected(index: Int, tab: AnimatedBottomBar.Tab) {
                mainPage.currentItem = index
            }
        })
        // 设置搜索按钮点击事件
        mToolbar.getBaseToolBar().setOnMenuItemClickListener{
            if (it.itemId == com.xiaoyou.library.widget.R.id.action_search){
                search.showFragment(supportFragmentManager, SearchFragment.TAG)
            }
            true
        }
        // 设置头像的点击事件,点击头像自动跳转到个人中心
        mToolbar.getTopImg().setOnClickListener{
            mainPage.currentItem = 3
        }
        // 设置searchView的回调事件
        search.setOnSearchClickListener(object : IOnSearchClickListener<PostDetail> {
            /**
             * 点击搜索按钮时触发
             * @param keyword 搜索的关键词
             */
            override fun onSearchClick(keyword: String) {
                val intent = Intent(appContext,SearchResultActivity::class.java)
                intent.putExtra("key",keyword)
                startActivity(intent)
            }

            /**
             * 点击链接时触发
             * @param data 链接携带的数据
             */
            override fun onLinkClick(data: PostDetail) {
                // 点击链接直接启动activity
                val intent = Intent(appContext,PostActivity::class.java)
                intent.putExtra("id",data.id)
                intent.putExtra("img",data.image)
                startActivity(intent)
            }

            /**
             * 搜索框内容改变时触发数据
             * @param key 关键词
             */
            override fun onTextChange(key: String?) {
                // 关键词改变时我们发送请求获取数据
                if (key != null) {
                    mViewModel.getPosts(1,key)
                }
            }

        })
        // 这里我们监听文章变化
        mViewModel.posts.observe(this){
            // 显示我们的文章内容
            if (it is ReturnList<PostDetail>){
                val data = ArrayList<CustomLink<PostDetail>>()
                // 遍历文章内容并显示
                for (item in it.contents){
                    data.add(CustomLink(item.title,item))
                }
                // 设置链接
                search.setLinks(data)
            }
        }
    }
    override fun showToolBar() = true
}