package com.xiaoyou.library.plugins.project

import android.content.Intent
import android.net.Uri
import android.os.Bundle
import androidx.recyclerview.widget.StaggeredGridLayoutManager
import com.bumptech.glide.Glide
import com.bumptech.glide.load.resource.bitmap.RoundedCorners
import com.bumptech.glide.request.RequestOptions
import com.lxj.xpopup.XPopup
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.util.Common
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.response.Project
import com.xiaoyou.library.net.entity.response.ProjectBottom
import com.xiaoyou.library.net.entity.response.ProjectTop
import com.xiaoyou.library.plugins.databinding.ActivityProjectBinding
import com.xiaoyou.library.plugins.viewmodel.PluginsVM
import com.youth.banner.adapter.BannerImageAdapter
import com.youth.banner.holder.BannerImageHolder
import com.youth.banner.indicator.CircleIndicator
import okhttp3.internal.wait
import java.lang.Exception
import java.util.*
import kotlin.collections.ArrayList


/**
 * @description
 * @author 小游
 * @data 2021/03/10
 */
class ProjectActivity: BaseDbActivity<PluginsVM, ActivityProjectBinding>() {
    // 我的项目
    private val projects:MutableList<ProjectBottom> = ArrayList()

    override fun initView(savedInstanceState: Bundle?) {
        // 初始化adapter
        val adapter = ProjectAdapter(projects)
        mDataBind.projectList.layoutManager =  StaggeredGridLayoutManager(
            1,
            StaggeredGridLayoutManager.VERTICAL
        )
        mDataBind.projectList.adapter = adapter
        // 获取我们的项目
        mViewModel.getProject(object : DataReceive<Project?>() {
            override fun success(data: Project?) {
                if (data is Project) {
                    // 设置轮播图
                    setBanner(data.top_content)
                    // 设置我的项目卡片
                    adapter.setList(data.bottom_content)
                }
            }
        })
        // 设置adapter的点击时间
        adapter.setOnItemClickListener{ _, _, position ->
            val data = adapter.getItem(position)
            // 设置选项还有链接
            val item:MutableList<String> = LinkedList()
            val link:MutableList<String> = LinkedList()
            // 判断是否为空
            if (data.blog.isNotEmpty()){
                item.add("博客")
                link.add(data.blog)
            }
            if (data.code.isNotEmpty()){
                item.add("代码")
                link.add(data.code)
            }
            if (data.video.isNotEmpty()){
                item.add("视频")
                link.add(data.video)
            }
            XPopup.Builder(this).asCenterList("", item.toTypedArray())
            { i, _ ->  Common.openUrl(link[i]) }.show()
        }
    }

    // 设置轮播图数据
    private fun setBanner(data: List<ProjectTop>){
        mDataBind.banner.setAdapter(object : BannerImageAdapter<ProjectTop>(data) {
            override fun onBindView(
                holder: BannerImageHolder,
                data: ProjectTop,
                position: Int,
                size: Int
            ) {
                // 设置点击事件
                holder.itemView.setOnClickListener{
                    if (data.url.isEmpty()){
                        MyToast.waring("无链接")
                    } else {
                        startActivity(Intent(Intent.ACTION_VIEW, Uri.parse(data.url)))
                    }
                }

                // 加载图片
                Glide.with(holder.itemView)
                    .load(data.image)
                    .apply(RequestOptions.bitmapTransform(RoundedCorners(30)))
                    .into(holder.imageView)
            }
        }).addBannerLifecycleObserver(this).indicator = CircleIndicator(this)
    }
}
