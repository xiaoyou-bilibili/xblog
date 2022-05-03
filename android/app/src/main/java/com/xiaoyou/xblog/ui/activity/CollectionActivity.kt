package com.xiaoyou.xblog.ui.activity

import android.os.Bundle
import androidx.recyclerview.widget.LinearLayoutManager
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.ext.initBack
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.xblog.databinding.ActivityCollectionBinding
import com.xiaoyou.xblog.ui.adapter.posts.PostAdapter
import com.xiaoyou.xblog.viewmodel.UserVM
import kotlinx.android.synthetic.main.fragment_posts.*

/**
 * @description 我的收藏
 * @author 小游
 * @data 2021/03/06
 */
class CollectionActivity: BaseDbActivity<UserVM,ActivityCollectionBinding>() {

    // 我的收藏
    private val posts:MutableList<PostDetail> = ArrayList()


    override fun initView(savedInstanceState: Bundle?) {
        // 初始化标题
        mToolbar.initBack("我的收藏") { finish() }
        // 获取我收藏的文章
        mViewModel.getUserCollection()
        // 初始化ui
        initUI()
    }

    // 初始化ui
    private fun initUI(){
        // 初始化adapter
        val adapter = PostAdapter(posts)
        mDataBind.collectionList.layoutManager = LinearLayoutManager(appContext)
        mDataBind.collectionList.adapter = adapter

        mViewModel.collections.observe(this){
            // 更新adapter
            if (it is List<PostDetail>){
                adapter.setList(it)
            }
        }
    }
}