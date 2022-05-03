package com.xiaoyou.xblog.ui.activity

import com.xiaoyou.library.net.entity.param.UserGetCodeParam

import android.os.Bundle
import android.view.View
import android.widget.EditText
import androidx.core.content.ContextCompat
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.ext.initBack
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.common.util.Setting
import com.xiaoyou.library.common.util.Validator
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.param.UserGetUsername
import com.xiaoyou.library.net.entity.param.UserLoginParam
import com.xiaoyou.library.net.entity.param.UserRegisterParam
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.databinding.ActivitySignUpBinding
import com.xiaoyou.xblog.ui.components.CustomEdit.setEditTextBackground
import com.xiaoyou.xblog.ui.components.CustomEdit.setError
import com.xiaoyou.xblog.viewmodel.UserVM

class SignUpActivity : BaseDbActivity<UserVM, ActivitySignUpBinding>() {

    // 数据绑定的参数
    private var param = UserRegisterParam()

    /**
     * 初始化view
     */
    override fun initView(savedInstanceState: Bundle?) {
        mToolbar.initBack("注册"){finish()}
        // 数据绑定
        mDataBind.item = param
        initUI()
        // 设置登录背景
        ImageUtil.setImageByUrl(Setting.getSetting().login,mDataBind.loginBackground, R.drawable.sign_in_bk)
    }

    // ui初始化
    private fun initUI(){
        //修复背景切换的Bug
        setEditTextBackground(mDataBind.userEditText)
        setEditTextBackground(mDataBind.eMailEditText)
        setEditTextBackground(mDataBind.passwordEditText)
        setEditTextBackground(mDataBind.verifyEditText)
        setEditTextBackground(mDataBind.nickNameEditText)
        //获取验证码
        mDataBind.verifyButton.setOnClickListener {
            if (param.username.isEmpty()) {
                MyToast.waring("请输入用户名!")
                //设置数据错误背景
                setError(mDataBind.userEditText)
            } else if (!Validator.isEmailOk(param.email)){
                setError(mDataBind.eMailEditText)
                MyToast.waring("邮箱格式不正确!")
            } else  {
                // 先验证用户名是否存在
                mViewModel.isRegister(
                        param.username,
                        param.email,
                        object : DataReceive<UserGetUsername?>() {
                            // 当我们找到这个用户时，提示已注册
                            override fun success(data: UserGetUsername?) {
                                if (data is UserGetUsername) {
                                    setError(mDataBind.userEditText)
                                    setError(mDataBind.eMailEditText)
                                    MyToast.waring("用户名或邮箱已注册!")
                                }
                            }
                            // 没有找到这个用户我们可以获取验证码
                            override fun error(message: String) {
                                mViewModel.getCode(
                                        UserGetCodeParam(param.email, "register"),
                                        object : DataReceive<UserGetCodeParam?>() {
                                            override fun success(data: UserGetCodeParam?) {
                                                if (data is UserGetCodeParam) {
                                                    // 触发获取验证码按钮变化
                                                    mViewModel.editGetCode(mDataBind.verifyButton,this@SignUpActivity)
                                                    MyToast.success("获取验证码成功，请检查邮箱")
                                                }
                                            }
                                        }
                                )
                            }
                        }
                )
            }
        }

        // 用户点击注册按钮
        mDataBind.signUpButton.setOnClickListener {
            // 判断用户输入
            if (param.username.isNotEmpty()
                && param.nickname.isNotEmpty()
                && param.code.isNotEmpty()
                && param.email.isNotEmpty()
                && param.password.isNotEmpty()
            ){
                mViewModel.userRegister(param,object :DataReceive<UserRegisterParam?>(){
                    override fun success(data: UserRegisterParam?) {
                        if (data is UserRegisterParam){
                            MyToast.success("注册成功，请登录！")
                            finish()
                        }
                    }
                })
            }else{
                // 判断用户那部分没有输入
                when {
                    param.username.isEmpty() -> setError(mDataBind.userEditText)
                    param.nickname.isEmpty() -> setError(mDataBind.nickNameEditText)
                    param.code.isEmpty() -> setError(mDataBind.verifyEditText)
                    param.password.isEmpty() -> setError(mDataBind.passwordEditText)
                    param.email.isEmpty() -> setError(mDataBind.eMailEditText)
                    param.password.isEmpty() -> setError(mDataBind.passwordEditText)
                }
                // 提示用户
                MyToast.waring("请填入相关信息")
            }
        }
    }


}