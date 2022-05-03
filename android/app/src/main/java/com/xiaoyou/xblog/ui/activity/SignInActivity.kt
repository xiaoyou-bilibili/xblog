package com.xiaoyou.xblog.ui.activity

import android.content.Intent
import android.os.Bundle
import android.view.View
import androidx.localbroadcastmanager.content.LocalBroadcastManager
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.base.kv
import com.xiaoyou.library.common.ext.initBack
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.common.util.Setting
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.core.Token
import com.xiaoyou.library.net.entity.param.UserLoginParam
import com.xiaoyou.library.net.entity.response.TokenDetail
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.databinding.ActivitySignInBinding
import com.xiaoyou.xblog.util.Variable
import com.xiaoyou.xblog.viewmodel.UserVM

/**
 * @description 登录界面
 * @author 小游
 * @data 2021/02/21
 */
class SignInActivity: BaseDbActivity<UserVM,ActivitySignInBinding>() {

    override fun showToolBar() = true
    // 数据绑定的参数
    private var param = UserLoginParam()

    /**
     * 初始化view
     */
    override fun initView(savedInstanceState: Bundle?) {
        mToolbar.initBack("登录"){finish()}
        // 设置数据绑定
        mDataBind.param = param
        // 用户注册
        mDataBind.register.setOnClickListener{ startActivity(Intent(this, SignUpActivity::class.java)) }
        // 忘记密码
        mDataBind.forgetTitle.setOnClickListener{ startActivity(Intent(this, SignForgetActivity::class.java)) }
        // 设置登录背景
        ImageUtil.setImageByUrl(Setting.getSetting().login,mDataBind.loginBackground, R.drawable.sign_in_bk)
    }

    // 用户登录
    fun userLogin(view: View){
        // 判断参数
        if (param.username !="" && param.password != ""){
           mViewModel.login(param,object :DataReceive<TokenDetail?>(){
               /**
                * 当获取到数据时进行回调
                * @param data T
                */
               override fun success(data: TokenDetail?) {
                   if (data is TokenDetail){
                       // 保存token数据
                       kv.encode(Token.USER_ID,data.userId)
                       kv.encode(Token.TOKEN,data.token)
                       // 发送广播,通知登录成功
                       LocalBroadcastManager.getInstance(appContext).sendBroadcast(Intent(Variable.LOGIN_SUCCESS))
                       // 结束当前activity
                       finish()
                   } else {
                       MyToast.error("登录出现异常！")
                   }
               }
           })
        } else {
            MyToast.waring("请输入用户名或密码！")
        }
    }
}