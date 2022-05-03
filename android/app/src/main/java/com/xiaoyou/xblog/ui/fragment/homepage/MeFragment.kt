package com.xiaoyou.xblog.ui.fragment.homepage

import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.content.IntentFilter
import android.net.Uri
import android.os.Bundle
import androidx.lifecycle.ViewModelProvider
import androidx.localbroadcastmanager.content.LocalBroadcastManager
import com.lxj.xpopup.XPopup
import com.lxj.xpopup.enums.PopupAnimation
import com.xiaoyou.library.common.base.*
import com.xiaoyou.library.common.util.Common
import com.xiaoyou.library.common.util.FileUtil
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.core.Token
import com.xiaoyou.library.net.entity.param.SubmitAdvice
import com.xiaoyou.library.net.entity.response.UserDetail
import com.xiaoyou.xblog.databinding.FragmentMeBinding
import com.xiaoyou.xblog.ui.activity.CollectionActivity
import com.xiaoyou.xblog.ui.activity.EditInfoActivity
import com.xiaoyou.xblog.ui.activity.PersonCenterActivity
import com.xiaoyou.xblog.ui.activity.SignInActivity
import com.xiaoyou.xblog.ui.dialog.ThemeChooseDialog
import com.xiaoyou.xblog.util.Variable
import com.xiaoyou.xblog.viewmodel.ToolsVM
import com.xiaoyou.xblog.viewmodel.UserVM


/**
 * @description 个人中心
 * @author 小游
 * @data 2021/02/21
 */
class MeFragment: BaseDbFragment<UserVM, FragmentMeBinding>(){

    // 工具版本的viewModel
    private val toolsViewModel by lazy { ViewModelProvider(this).get(ToolsVM::class.java) }

    /**
     * 初始化view操作
     */
    override fun initView(savedInstanceState: Bundle?) {
        // 添加数据
        initUI()
    }

    // UI初始化
    private fun initUI() {
        // 第一次使用自动获取用户信息
        if (Token.getToken().userId > 0){
            mViewModel.getUserInfo()
        }
        // 注册广播接受者接收数据
        LocalBroadcastManager.getInstance(appContext).registerReceiver(object :
            BroadcastReceiver() {
            override fun onReceive(context: Context?, intent: Intent?) {
                mViewModel.getUserInfo()
            }
        }, IntentFilter(Variable.LOGIN_SUCCESS))
        // 监听用户数据变化
        mViewModel.info.observe(viewLifecycleOwner){
            mDataBind.info = it
            // 设置用户头像
            ImageUtil.setImageByUrl(it.avatar, mDataBind.userAvatar)
            // 设置顶部头像
            ImageUtil.setImageByUrl(it.avatar, mVmActivity.mToolbar.getTopImg())
        }
        // 点击登录
        mDataBind.clickLogin.setOnClickListener{ startView(SignInActivity::class.java) }
        // 修改个人信息
        mDataBind.editInfo.setOnClickListener{ startView(EditInfoActivity::class.java) }
        // 我的收藏
        mDataBind.myCollection.setOnClickListener{ startView(CollectionActivity::class.java) }
        // 个人中心
        mDataBind.personCenter.setOnClickListener{ startView(PersonCenterActivity::class.java) }
        // 退出登录
        mDataBind.loginOut.setOnClickListener{
            // 删除缓存数据
            Token.clearToken()
            // 清空viewModel
            mViewModel.info.postValue(UserDetail())
        }
        // 软件设置
        mDataBind.setting.setOnClickListener{ MyToast.info("该软件暂无设置！") }
        // 清理缓存
        mDataBind.clearCache.setOnClickListener{
            XPopup.Builder(mActivity)
                    .popupAnimation(PopupAnimation.ScaleAlphaFromCenter)
                    .asConfirm(
                            "提示",
                            "当前缓存大小为${FileUtil.getTotalCacheSize()},确认清除?",
                            "取消",
                            "确定",
                            {
                                FileUtil.clearAllCache()
                                MyToast.success("清除成功!")
                            },
                            null,
                            false
                    ).show()
        }
        // 检查更新
        mDataBind.checkUpdate.setOnClickListener{ Common.checkUpdate(true) }
        // 技术支持
        mDataBind.support.setOnClickListener{
            XPopup.Builder(mActivity)
                .popupAnimation(PopupAnimation.ScaleAlphaFromCenter)
                .asConfirm(
                    "技术支持",
                    "本软件由XBlog系统提供技术支持\n开发作者:小游",
                    "取消",
                    "访问官网",
                    { startActivity(Intent(Intent.ACTION_VIEW, Uri.parse("https://xblog.xiaoyou66.com"))) },
                    null,
                    false
                ).show()
        }
        // 意见反馈
        mDataBind.submitAdvice.setOnClickListener{
            XPopup.Builder(mActivity).popupAnimation(PopupAnimation.ScrollAlphaFromLeft)
                .asInputConfirm("意见反馈", "", "请输入反馈内容")
            {
                if (it.isEmpty()){
                    MyToast.waring("请输入反馈内容")
                } else {
                    toolsViewModel.submitAdvice(it,object :DataReceive<SubmitAdvice?>(){
                        override fun success(data: SubmitAdvice?) {
                            if (data is SubmitAdvice){
                                MyToast.success("反馈成功！")
                            }
                        }
                    })
                }
            }.show()
        }
        // 扫码功能
        mDataBind.myScan.setOnClickListener{ MyToast.info("开发中!") }
        // 切换主题
        mDataBind.myTheme.setOnClickListener{ ThemeChooseDialog(context!!,mActivity).show() }
        // 夜间模式
        mDataBind.myNight.setOnClickListener{ MyToast.info("开发中!") }
    }

    // 启动activity
    private fun <T> startView(clazz: Class<T>){
        startActivity(Intent(appContext, clazz))
    }

    // 当界面恢复的时候我们获取一下用户信息
    override fun onResume() {
        super.onResume()
        if (Token.getToken().userId != 0){
            mViewModel.getUserInfo()
        }
    }

}