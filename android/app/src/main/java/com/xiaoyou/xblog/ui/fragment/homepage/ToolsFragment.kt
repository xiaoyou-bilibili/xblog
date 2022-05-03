package com.xiaoyou.xblog.ui.fragment.homepage

import android.content.Intent
import android.graphics.drawable.Drawable
import android.os.Bundle
import android.view.View
import androidx.appcompat.content.res.AppCompatResources.getDrawable
import androidx.recyclerview.widget.StaggeredGridLayoutManager
import com.chad.library.adapter.base.BaseQuickAdapter
import com.xiaoyou.library.common.base.BaseFragment
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.Setting
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.plugins.animation.AnimationActivity
import com.xiaoyou.library.plugins.doc.DocActivity
import com.xiaoyou.library.plugins.douban.DouBanActivity
import com.xiaoyou.library.plugins.friend.FriendActivity
import com.xiaoyou.library.plugins.music.MusicActivity
import com.xiaoyou.library.plugins.navigation.NavigationActivity
import com.xiaoyou.library.plugins.project.ProjectActivity
import com.xiaoyou.library.plugins.sponsor.SponsorActivity
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.data.commom.ToolsItem
import com.xiaoyou.xblog.ui.adapter.ui.ToolsAdapter
import kotlinx.android.synthetic.main.fragment_tools.*

/**
 * @description 功能界面的fragment
 * @author 小游
 * @data 2021/02/21
 */
class ToolsFragment(override val layoutId: Int = R.layout.fragment_tools) : BaseFragment() {

    // 功能能区图标
    private val drawable = ArrayList<Drawable?>()

    // 功能区标题
    private val titles = ArrayList<String>()

    // 功能区activity
    private val activity = ArrayList<Class<*>>()

    // 手动把所有的工具类全部添加进去
    private val functions = ArrayList<ToolsItem>()

    // 视图初始化
    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        val setting = Setting.getSetting()
        // 获取设置数据
        if(setting.friend){
            drawable.add(getDrawable(appContext,R.drawable.ic_tools_friend))
            titles.add(appContext.getString(R.string.function_friend_text))
            activity.add(FriendActivity::class.java)
        }
        if(setting.animation){
            drawable.add(getDrawable(appContext,R.drawable.ic_tools_animation))
            titles.add(appContext.getString(R.string.function_animation_text))
            activity.add(AnimationActivity::class.java)
        }
        if(setting.donate){
            drawable.add(getDrawable(appContext,R.drawable.ic_tools_reward))
            titles.add(appContext.getString(R.string.function_love_text))
            activity.add( SponsorActivity::class.java)
        }
        if (setting.dou_ban){
            drawable.add(getDrawable(appContext,R.drawable.ic_tools_douban))
            titles.add(appContext.getString(R.string.function_douBan_text))
            activity.add(DouBanActivity::class.java)
        }
        if (setting.music){
            drawable.add(getDrawable(appContext,R.drawable.ic_tools_music))
            titles.add(appContext.getString(R.string.function_music_text))
            activity.add(MusicActivity::class.java)
        }
        if (setting.doc){
            drawable.add(getDrawable(appContext,R.drawable.ic_tools_doc))
            titles.add(appContext.getString(R.string.function_doc_text))
            activity.add(DocActivity::class.java)
        }
        if (setting.project){
            drawable.add(getDrawable(appContext,R.drawable.ic_tools_project))
            titles.add(appContext.getString(R.string.function_project_text))
            activity.add(ProjectActivity::class.java)
        }
        if (setting.navigation){
            drawable.add(getDrawable(appContext,R.drawable.ic_tools_navigation))
            titles.add(appContext.getString(R.string.function_navigation_text))
            activity.add(NavigationActivity::class.java)
        }
        // 添加数据
        for (i in 0 until drawable.size){
            functions.add(ToolsItem(drawable[i],titles[i],activity[i]))
        }
        // 绑定布局
       val toolsAdapter = ToolsAdapter(functions)
        // 设置布局类型
        toolsList.layoutManager = StaggeredGridLayoutManager(3, StaggeredGridLayoutManager.VERTICAL)
        // 设置按键监听事件
        toolsAdapter.setOnItemClickListener{ adapter: BaseQuickAdapter<*, *>, v: View, i: Int ->
            // 获取activity并启动
            val item = adapter.getItem(i) as ToolsItem
            val intent = Intent(appContext,item.activity)
            // 放入一些参数
            intent.putExtra("title",item.title)
            if (item.activity != null) startActivity(intent)
        }
        toolsList.adapter = toolsAdapter
    }
}