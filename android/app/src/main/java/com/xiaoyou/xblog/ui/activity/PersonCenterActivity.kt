package com.xiaoyou.xblog.ui.activity

import android.media.MediaRouter
import android.os.Bundle
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.ext.initBack
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.param.UpdateUserInfo
import com.xiaoyou.library.net.entity.response.UserDetail
import com.xiaoyou.xblog.databinding.ActivityPersonCenterBinding
import com.xiaoyou.xblog.viewmodel.UserVM

/**
 * @description 个人中心
 * @author 小游
 * @data 2021/03/06
 */
class PersonCenterActivity: BaseDbActivity<UserVM,ActivityPersonCenterBinding>() {

    // 个人信息
    private val info = UserDetail()

    override fun initView(savedInstanceState: Bundle?) {
        // 初始化标题
        mToolbar.initBack("个人中心") { finish() }
        // 获取个人信息
        mViewModel.getUserInfo()
        // 监听个人信息变化
        mViewModel.info.observe(this){
            if (it is UserDetail){
                XLog.e(it)
                mDataBind.item = it
            }
        }
        // 邮件订阅按钮点击事件
        mDataBind.subscription.setOnCheckedChangeListener { _, isChecked ->
            mViewModel.updateUserInfo(UpdateUserInfo(subscription = isChecked.toString()),object :DataReceive<UpdateUserInfo?>(){
                override fun success(data: UpdateUserInfo?) {
                    if (data is UpdateUserInfo){
                        mViewModel.getUserInfo()
                    }
                }
            })
        }
    }
}