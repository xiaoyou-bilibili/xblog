package com.xiaoyou.xblog.ui.activity

import android.Manifest
import android.content.Context
import android.content.Intent
import android.os.Bundle
import android.widget.ImageView
import com.hjq.permissions.XXPermissions
import com.lxj.xpopup.XPopup
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.ext.initBack
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.common.util.Validator
import com.xiaoyou.library.common.util.XLog
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.param.UpdateUserInfo
import com.xiaoyou.library.net.entity.response.UploadFile
import com.xiaoyou.library.net.entity.response.UserDetail
import com.xiaoyou.xblog.databinding.ActivityEditInfoBinding
import com.xiaoyou.xblog.viewmodel.UserVM
import com.yuyh.library.imgsel.ISNav
import com.yuyh.library.imgsel.config.ISListConfig
import kotlinx.android.synthetic.main.activity_edit_info.*


/**
 * @description 编辑个人信息
 * @author 小游
 * @data 2021/03/05
 */
class EditInfoActivity :BaseDbActivity<UserVM, ActivityEditInfoBinding>(){
    // 图片回调返回代码
    private val imageResult = 666
    // 个人信息
    private lateinit var info : UserDetail
    // 视图初始化
    override fun initView(savedInstanceState: Bundle?) {
        // 初始化顶部标题
        mToolbar.initBack("修改个人信息") { finish() }
        // 获取个人信息
        mViewModel.getUserInfo()
        // 监听个人信息变化
        mViewModel.info.observe(this){
            if (it is UserDetail){
                // 修改头像
                ImageUtil.setImageByUrl(it.avatar, mDataBind.avatar)
                info = it
                // 绑定布局
                mDataBind.item = it
            }
        }
        // 初始化ui
        initUI()
    }

    private fun initUI(){
        // 修改头像点击时间
        mDataBind.editAvatar.setOnClickListener{
            XXPermissions.with(this)
                .permission(Manifest.permission.WRITE_EXTERNAL_STORAGE)
                .permission(Manifest.permission.READ_EXTERNAL_STORAGE)
                .request { _, _ ->
                    // 打开手机相册
                    chooseImage()
                }
        }
        // 昵称点击事件
        mDataBind.nickname.setOnClickListener{
            XPopup.Builder(this).asInputConfirm("请输入昵称", "", info.nickname, ""){
                // 验证昵称
                if (it.isNullOrEmpty()){
                    MyToast.waring("昵称不能为空")
                } else {
                    updateInfo(UpdateUserInfo(nickname = it))
                }
            }.show()
        }
        // 邮箱点击事件
        mDataBind.email.setOnClickListener{
            XPopup.Builder(this).asInputConfirm("请输入邮箱", "", info.email, ""){
                // 邮箱
                if (!Validator.isEmailOk(it)){
                    MyToast.waring("邮箱格式不合法")
                } else {
                    updateInfo(UpdateUserInfo(email = it))
                }
            }.show()
        }
        // 签名点击事件
        mDataBind.sign.setOnClickListener{
            XPopup.Builder(this).asInputConfirm("请输入签名", "", info.sign, ""){
                updateInfo(UpdateUserInfo(sign = it))
            }.show()
        }
    }

    // 更新信息
    private fun updateInfo(param: UpdateUserInfo){
        // 更新个人信息
        mViewModel.updateUserInfo(param, object : DataReceive<UpdateUserInfo?>() {
            override fun success(data: UpdateUserInfo?) {
                // 更新个人信息
                mViewModel.getUserInfo()
            }
        })
    }

    // activity的回调事件
    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        super.onActivityResult(requestCode, resultCode, data)
        if (requestCode == imageResult && resultCode == RESULT_OK && data != null){
            val path = data.getStringArrayListExtra("result")
            if (!path.isNullOrEmpty()){
                val bitmap = ImageUtil.getBitMapFromPath(path[0])
                if (bitmap != null){
                    // 图片上传
                    mViewModel.uploadFile(ImageUtil.bitmap2Base64(bitmap),object :DataReceive<UploadFile?>(){
                        override fun success(data: UploadFile?) {
                            if (data is UploadFile){
                                // 如果上传成功，那么我们就更新用户头像
                                updateInfo(UpdateUserInfo(avatar=data.url))
                            }
                        }
                    })
                }
            }
        }
    }

    // 选择图片
    private fun chooseImage(){
        ISNav.getInstance().init { _, path, imageView ->
            ImageUtil.setImageByUrl(path,imageView)
        }
        // 设置图片选择器的配置
        val config = ISListConfig.Builder().multiSelect(false).needCrop(true).build()
        //  启动图片选择界面
        ISNav.getInstance().toListActivity(this,config,imageResult)
    }
}