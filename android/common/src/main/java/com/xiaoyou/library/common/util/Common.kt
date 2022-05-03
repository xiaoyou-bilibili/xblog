package com.xiaoyou.library.common.util

import android.content.ClipData
import android.content.ClipboardManager
import android.content.Context
import android.content.Context.CLIPBOARD_SERVICE
import android.content.Intent
import android.content.pm.PackageInfo
import android.content.pm.PackageManager
import android.net.Uri
import com.allenliu.versionchecklib.v2.AllenVersionChecker
import com.allenliu.versionchecklib.v2.builder.UIData
import com.lxj.xpopup.XPopup
import com.lxj.xpopup.enums.PopupAnimation
import com.xiaoyou.library.common.base.SERVER
import com.xiaoyou.library.common.base.WEB
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.net.api.ToolsService
import com.xiaoyou.library.net.core.CallHandle
import com.xiaoyou.library.net.core.DataReceive
import com.xiaoyou.library.net.core.Repository
import com.xiaoyou.library.net.entity.response.AppVersion
import java.util.regex.Matcher
import java.util.regex.Pattern


/**
 * @description 常用的不好分类的工具类
 * @author 小游
 * @data 2021/03/03
 */
object Common {
    private val toolsService: ToolsService by lazy { Repository.toolsService(appContext) }

    // 获取当前应用版本号
    fun getAppVersion():String{
        var versionName = ""
        try {
            // ---get the package info---
            val pm: PackageManager = appContext.packageManager
            val pi: PackageInfo = pm.getPackageInfo(appContext.packageName, 0)
            versionName = pi.versionName
            if (versionName == null || versionName.isEmpty()) {
                return ""
            }
        } catch (e: Exception) {

        }
        return versionName
    }

    // 检查版本更新,是否提示用户，默认不提示
    fun checkUpdate(show:Boolean = false){
        // 首先检查更新
        CallHandle.handleCall(toolsService.getAppVersion(), object : DataReceive<AppVersion?>() {
            override fun success(data: AppVersion?) {
                if (data is AppVersion) {
                    // 先获取当前版本信息
                    XLog.e(getAppVersion())
                    if (data.version != getAppVersion()) {
                        AllenVersionChecker
                                .getInstance()
                                .downloadOnly(
                                        UIData
                                                .create()
                                                .setDownloadUrl(data.download)
                                                .setTitle("检测到新版本 v${data.version}")
                                                .setContent(data.dec.replace("%%", "\n"))
                                )
                                .executeMission(appContext)
                    } else {
                        if (show){
                            MyToast.success("您当前的版本(${getAppVersion()})已是最新！")
                        }
                    }
                }
            }
        })
    }

    // 判断是否为本站链接,这里需要传入activity
    fun checkLink(context: Context,activity :Class<*>){
        val link = getClipboard()
        val p = Pattern.compile("$WEB/archives/([0-9|/]+)")
        if (p.matcher(link).matches()) {
            //说明匹配，我们可以吧id给打印出来
            val type: Matcher = p.matcher(link)
            if (type.find()) {
                var id: String = type.group(1)!!
                id = id.replace("/", "")
                // 提示用户发现文章
                XPopup.Builder(context)
                    .popupAnimation(PopupAnimation.ScaleAlphaFromCenter)
                    .asConfirm(
                        "提示",
                        "检测到你复制了本站链接，是否立即跳转？$id",
                        "取消",
                        "确定",
                        {
                            // 清除剪贴板
                            setClipboard("")
                            // 打开文章
                            val intent = Intent(appContext,activity)
                            intent.putExtra("id",id.toInt())
                            context.startActivity(intent)
                        },
                        null,
                        false
                    ).show()
            }
        }
    }

    /**
     *  获取剪贴板内容
     * @return String 剪贴板内容
     */
    fun getClipboard():String{
        return try {
            //获取剪贴板信息
            val cm = appContext.getSystemService(CLIPBOARD_SERVICE) as ClipboardManager
            val data = cm.primaryClip;
            val item = data?.getItemAt(0)
            item?.text.toString()
        }catch (e: Exception){
            ""
        }
    }

    /**
     *  设置剪贴板内容
     * @param data String 剪贴板信息
     * @return Boolean 设置是否成功
     */
    fun setClipboard(data: String):Boolean{
        return try {
            val cm = appContext.getSystemService(CLIPBOARD_SERVICE) as ClipboardManager
            val clip = ClipData.newPlainText("label", data)
            cm.setPrimaryClip(clip)
            true
        } catch (e: Exception){
            false
        }
    }

    /**
     *  打开应用链接
     * @param url String
     */
    fun openUrl(url: String){
        val intent = Intent(Intent.ACTION_VIEW, Uri.parse(url))
        // 如果要启动的话需要加上这个flag
        intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK)
        appContext.startActivity(intent)
    }

}