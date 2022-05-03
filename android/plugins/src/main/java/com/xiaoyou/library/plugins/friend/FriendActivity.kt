package com.xiaoyou.library.plugins.friend

import android.content.Intent
import android.net.Uri
import android.os.Bundle
import androidx.lifecycle.ViewModelProvider
import androidx.recyclerview.widget.StaggeredGridLayoutManager
import com.lxj.xpopup.XPopup
import com.lxj.xpopup.enums.PopupAnimation
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.listener.MyClickListener
import com.xiaoyou.library.common.util.Common
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.param.SubmitFriend
import com.xiaoyou.library.net.entity.response.FriendDetail
import com.xiaoyou.library.net.entity.response.FriendSettings
import com.xiaoyou.library.plugins.databinding.ActivityFriendBinding
import com.xiaoyou.library.plugins.viewmodel.PluginsVM
import com.xiaoyou.library.plugins.viewmodel.SettingVM


/**
 * @description 友人帐activity界面
 * @author 小游
 * @data 2021/02/26
 */
class FriendActivity : BaseDbActivity<PluginsVM, ActivityFriendBinding>(){

    // 友链数据
    private val friends:MutableList<FriendDetail> = ArrayList()
    // 设置界面viewModel
    private val settingVN by lazy { ViewModelProvider(this).get(SettingVM::class.java) }


    override fun initView(savedInstanceState: Bundle?) {
        // 初始化adapter
        val adapter = FriendAdapter(friends)
        mDataBind.friendList.layoutManager = StaggeredGridLayoutManager(
            2,
            StaggeredGridLayoutManager.VERTICAL
        )
        mDataBind.friendList.adapter = adapter

        // 加载友链数据
        mViewModel.getFriend(object : DataReceive<List<FriendDetail>?>() {
            override fun success(data: List<FriendDetail>?) {
                if (data is List<FriendDetail>) {
                    adapter.addData(data)
                }
            }
        })
        // 设置友链点击事件
        adapter.setOnItemClickListener{ _, _, position ->
            // 获取友链数据
            val data = adapter.getItem(position)
            // 显示弹框
            XPopup.Builder(this)
                    .popupAnimation(PopupAnimation.ScaleAlphaFromLeftTop)
                    .asConfirm(
                            data.name,
                            "介绍:${data.dec}\n网站:${data.url}",
                            "取消",
                            "访问网站",
                            { Common.openUrl(data.url) },
                            null,
                            false
                    ).show()
        }
        // 点击添加友链
        mDataBind.addFriend.setOnClickListener{
            // 先获取友链界面设置
            settingVN.getFriend(object :DataReceive<FriendSettings?>(){
                override fun success(data: FriendSettings?) {
                    if (data is FriendSettings){
                        val dialog = FriendDialog(this@FriendActivity).setContent(data)
                        dialog.setOnSubmitListener(object: MyClickListener<SubmitFriend>{
                            override fun onClick(data: SubmitFriend) {
                                // 提交申请
                                mViewModel.submitFriend(data,object :DataReceive<SubmitFriend?>(){
                                    override fun success(data: SubmitFriend?) {
                                        if (data is SubmitFriend){
                                            // 关闭弹窗
                                            dialog.dismiss()
                                            // 显示成功
                                            XPopup.Builder(this@FriendActivity)
                                                    .popupAnimation(PopupAnimation.ScaleAlphaFromLeftTop)
                                                    .asConfirm(
                                                            "提示",
                                                            "提交成功，审核通过后系统会发送邮件通知你！",
                                                            "取消",
                                                            "我知道了",
                                                            {},
                                                            null,
                                                            false
                                                    ).show()
                                        }
                                    }
                                })
                            }
                        })
                        dialog.show()

                    }
                }
            })
        }
    }
}