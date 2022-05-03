package com.xiaoyou.xblog.viewmodel

import android.annotation.SuppressLint
import android.widget.Button
import androidx.lifecycle.LifecycleOwner
import androidx.lifecycle.MutableLiveData
import com.xiaoyou.library.common.base.BaseViewModel
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.ext.handleCall
import com.xiaoyou.library.net.api.ToolsService
import com.xiaoyou.library.net.api.UserService
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.core.Repository
import com.xiaoyou.library.net.entity.base.LoadingType
import com.xiaoyou.library.net.entity.param.*
import com.xiaoyou.library.net.entity.response.*
import java.util.*

/**
 * @description 和用户有关的viewmodel
 * @author 小游
 * @data 2021/02/25
 */
class UserVM (private val service: UserService = Repository.userService(appContext)) : BaseViewModel(){

    private val toolsService: ToolsService  by lazy { Repository.toolsService(appContext) }

    // 用户信息
    val info = MutableLiveData<UserDetail>()
    // 剩余时间
    private val time = MutableLiveData<Int>()
    // 用户收藏的文章
    val collections = MutableLiveData<List<PostDetail>>()

    /**
     *  用户登录
     * @param param UserLoginParam 登录所需要的的参数
     * @param handle OnDataReceive<TokenDetail?>
     */
    fun login(param: UserLoginParam,handle: DataReceive<TokenDetail?>) = handleCall<TokenDetail> {
        call = service.postToken(param)
        onSuccess = { handle.success(it) }
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "登录中..."
    }

    /**
     *  获取用户个人信息
     */
    fun getUserInfo() = handleCall<UserDetail> {
        call = service.getUser()
        onSuccess = {info.postValue(it)}
    }


    /**
     *  获取验证码
     * @param param UserGetCodeParam 验证参数
     * @param handle DataReceive<UserGetCodeParam?> 返回验证码信息
     */
    fun getCode(param: UserGetCodeParam,handle: DataReceive<UserGetCodeParam?>) = handleCall<UserGetCodeParam> {
        call = service.postCode(param)
        onSuccess = {handle.success(it)}

    }

    /**
     *  判断用户是否注册过了
     * @param user String
     * @param email String
     * @param handle DataReceive<UserGetUsername?>
     */
    fun isRegister(user:String,email:String,handle: DataReceive<UserGetUsername?>) = handleCall<UserGetUsername> {
        call = service.getUsername(user,email)
        onSuccess = {handle.success(it)}
        onError = {handle.error(it.message?:"")}
    }

    /**
     *  用户注册功能
     * @param param UserRegisterParam 用户注册所需要的的内容
     * @param handle DataReceive<UserRegisterParam?> 用户注册回调函数
     */
    fun userRegister(param: UserRegisterParam,handle: DataReceive<UserRegisterParam?>) = handleCall<UserRegisterParam> {
        call = service.postRegister(param)
        onSuccess = {handle.success(it)}
    }

    /**
     *  用户重置密码
     * @param param ResetPassword 重置密码所需的参数
     * @param handle DataReceive<ResetPassword?> 回调函数
     */
    fun forgetPassword(param: ResetPassword,handle: DataReceive<ResetPassword?>) = handleCall<ResetPassword> {
        call = service.resetPassword(param)
        onSuccess = {handle.success(it)}
    }

    /**
    *  更新用户信息
    * @param param UpdateUserInfo 更新用户信息参数
    * @param handel DataReceive<UpdateUserInfo?> 回调函数
    */
    fun updateUserInfo(param: UpdateUserInfo,handel: DataReceive<UpdateUserInfo?>) = handleCall<UpdateUserInfo> {
        call = service.updateInfo(param)
        onSuccess = {handel.success(it)}
    }

    /**
     *  用户上传图片
     * @param data String base64的数据
     * @param handle DataReceive<UploadFile?> 回调函数
     */
    fun uploadFile(data: String,handle: DataReceive<UploadFile?>) = handleCall<UploadFile> {
        call = toolsService.uploadImage(data)
        onSuccess = {handle.success(it)}
    }

    /**
     *  获取我收藏的文章
     */
    fun getUserCollection() = handleCall<List<PostDetail>> {
        call = service.getCollection()
        onSuccess = {collections.postValue(it)}
        loadingType = LoadingType.LOADING_DIALOG
        loadingMessage = "正在获取我的收藏..."
        errorType = LoadingType.LOADING_XML
    }


    /**
     *  点击按钮获取验证码
     * @param button Button
     */
    @SuppressLint("SetTextI18n")
    fun editGetCode(button: Button, owner: LifecycleOwner){
        button.isEnabled = false
        time.postValue(50)
        time.observe(owner){
            if (it <= 0){
                button.text = "获取验证码"
                button.isEnabled = true
            } else {
                button.text = "${it}秒后重新获取"
            }
        }
        // 设置定时任务
        Timer().schedule(object : TimerTask(){
            override fun run() {
                // 定时任务自动减1
                val t = time.value?:0
                if ( t > 0){
                    time.postValue(t - 1)
                }
            }
        }, Date(),1000)
    }


}