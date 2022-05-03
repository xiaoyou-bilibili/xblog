package com.xiaoyou.xblog.ui.activity

import android.os.Bundle
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.ext.initBack
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.common.util.Setting
import com.xiaoyou.library.common.util.Validator
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.param.ResetPassword
import com.xiaoyou.library.net.entity.param.UserGetCodeParam
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.databinding.ActivitySignForgetBinding
import com.xiaoyou.xblog.ui.components.CustomEdit.setEditTextBackground
import com.xiaoyou.xblog.ui.components.CustomEdit.setError
import com.xiaoyou.xblog.viewmodel.UserVM

/**
 * @description 忘记密码板块
 * @author 小游
 * @data 2021/03/04
 */
class SignForgetActivity : BaseDbActivity<UserVM, ActivitySignForgetBinding>() {
    override fun showToolBar() = true
    private var param = ResetPassword()
    override fun initView(savedInstanceState: Bundle?) {
        mToolbar.initBack("忘记密码"){finish()}
        // 设置数据绑定
        mDataBind.reset = param
        initUI()
        // 设置登录背景
        ImageUtil.setImageByUrl(Setting.getSetting().login,mDataBind.loginBackground, R.drawable.sign_in_bk)
    }

    // 初始化UI
    private fun initUI(){
        //修复背景切换的Bug
        setEditTextBackground(mDataBind.eMailEditText)
        setEditTextBackground(mDataBind.passwordEditText)
        setEditTextBackground(mDataBind.verifyEditText)
        //获取验证码
        mDataBind.verifyButton.setOnClickListener {
             if (!Validator.isEmailOk(param.email)){
                setError(mDataBind.eMailEditText)
                MyToast.waring("邮箱格式不正确!")
            } else  {
                mViewModel.getCode(
                    UserGetCodeParam(param.email, "forget"),
                    object : DataReceive<UserGetCodeParam?>() {
                        override fun success(data: UserGetCodeParam?) {
                            if (data is UserGetCodeParam) {
                                // 触发获取验证码按钮变化
                                mViewModel.editGetCode(mDataBind.verifyButton,this@SignForgetActivity)
                                MyToast.success("获取验证码成功，请检查邮箱")
                            }
                        }
                    }
                )
            }
        }
        // 重置密码点击事件
        mDataBind.forgetButton.setOnClickListener{
            // 先验证一下用户是否输入
            if (param.email.isEmpty() || param.password.isEmpty() || param.code.isEmpty()){
                when{
                    param.email.isEmpty() -> setError(mDataBind.eMailEditText)
                    param.password.isEmpty() -> setError(mDataBind.passwordEditText)
                    param.code.isEmpty() -> setError(mDataBind.verifyEditText)
                }
                MyToast.waring("请输入相关信息")
            } else {
                // 重置密码
                mViewModel.forgetPassword(param,object :DataReceive<ResetPassword?>(){
                    override fun success(data: ResetPassword?) {
                        if (data is ResetPassword){
                            MyToast.success("重置密码成功，请登录！")
                            finish()
                        }
                    }
                })
            }
        }
    }
}