package com.xiaoyou.library.plugins.sponsor

import android.Manifest
import android.graphics.BitmapFactory
import android.graphics.drawable.GradientDrawable
import android.os.Bundle
import androidx.lifecycle.ViewModelProvider
import com.facebook.binaryresource.FileBinaryResource
import com.facebook.cache.common.SimpleCacheKey
import com.facebook.drawee.backends.pipeline.Fresco
import com.hjq.permissions.XXPermissions
import com.xiaoyou.library.common.base.BaseDbActivity
import com.xiaoyou.library.common.util.Common
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.common.util.MyToast
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.entity.response.DonateSettings
import com.xiaoyou.library.net.entity.response.Sponsors
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.databinding.ActivitySponsorBinding
import com.xiaoyou.library.plugins.viewmodel.PluginsVM
import com.xiaoyou.library.plugins.viewmodel.SettingVM
import java.io.InputStream

/**
 * @description 赞助博主
 * @author 小游
 * @data 2021/02/26
 */
class SponsorActivity  : BaseDbActivity<PluginsVM, ActivitySponsorBinding>(){
    private val settingVN by lazy { ViewModelProvider(this).get(SettingVM::class.java) }

    private var alipay: String = ""
    private var wechat: String = ""

    /**
     *  当前URL地址
     */
    private var current: String = ""

    // 设置图片选中效果
    private val drawable by lazy {
        GradientDrawable().apply {
            setStroke(4, getColor(R.color.theme_pink))
            cornerRadius = 15f
        }
    }

    override fun initView(savedInstanceState: Bundle?) {
        // 获取赞助记录
        mViewModel.getSponsor(object :DataReceive<List<Sponsors>?>(){
            override fun success(data: List<Sponsors>?) {
                if (data is List<Sponsors>){
                    mDataBind.table.apply {
                        // 显示数据
                        setData(data)
                        // 统计行
                        config.isFixedTitle = true
                        // 隐藏顶部序号列
                        config.isShowXSequence = false
                    }

                }
            }
        })
        // 获取赞助设置
        settingVN.getSponsor(object :DataReceive<DonateSettings?>(){
            override fun success(data: DonateSettings?) {
                if (data is DonateSettings){
                    alipay = data.aliPay
                    wechat = data.weChat
                    // 默认设置支付宝
                    mDataBind.payImage.setImageURI(alipay)
                    // 设置url
                    current = alipay
                    // 设置选中效果
                    mDataBind.loveALiPayChoose.foreground = drawable
                }
            }
        })
        // 切换支付方式，第一个是支付宝的，第二个是微信支付
        mDataBind.loveALiPayChoose.setOnClickListener {
            mDataBind.apply {
                loveALiPayChoose.foreground = drawable
                loveWeChatChoose.foreground = null
                payImage.setImageURI(alipay)
                current = alipay
            }
        }
        mDataBind.loveWeChatChoose.setOnClickListener {
            mDataBind.apply {
                loveALiPayChoose.foreground = null
                loveWeChatChoose.foreground = drawable
                payImage.setImageURI(wechat)
                current = wechat
            }
        }
        // 保存图片按钮点击
        mDataBind.loveSave.setOnClickListener{
            // 先动态申请图片权限
            XXPermissions.with(this)
                    .permission(Manifest.permission.WRITE_EXTERNAL_STORAGE)
                    .permission(Manifest.permission.READ_EXTERNAL_STORAGE)
                    .request { _, _ ->
                        var ins:InputStream? = null
                        val result = try {
                            // 打开手机相册
                            val img = Fresco.getImagePipelineFactory().mainFileCache.getResource(SimpleCacheKey(current)) as FileBinaryResource
                            ins = img.openStream()
                            ImageUtil.saveImage(BitmapFactory.decodeStream(ins),"pay")
                        } catch (e: Exception){
                            false
                        }finally {
                            ins?.close()
                        }
                        if (result) {
                            MyToast.success("保存成功")
                        } else {
                            MyToast.error("保存成功")
                        }
                    }
        }
    }
}