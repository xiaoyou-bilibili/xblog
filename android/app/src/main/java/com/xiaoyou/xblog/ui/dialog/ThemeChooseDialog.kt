package com.xiaoyou.xblog.ui.dialog

import android.app.Activity
import android.content.Context
import androidx.recyclerview.widget.GridLayoutManager
import com.jaredrummler.cyanea.Cyanea
import com.lxj.xpopup.XPopup
import com.xiaoyou.library.common.base.BaseDialog
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.core.ThemeManage
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.data.commom.Theme
import com.xiaoyou.xblog.databinding.DialogThemeBinding
import com.xiaoyou.xblog.ui.adapter.common.ThemeAdapter

/**
 * @description 主题选择
 * @author 小游
 * @data 2021/03/08
 */
class ThemeChooseDialog(context: Context,activity: Activity) : BaseDialog<DialogThemeBinding>(context,0.7f) {
    // 所有的主题
    private val themes = ArrayList<Theme>()
    // 当前的主题
    private val theme = ThemeManage.getTheme()

    init {
        // 互补配色参考 https://www.shejidaren.com/examples/tools/color-scheme/
        addColor(R.color.theme_pink,R.color.theme_pink_accent,R.string.theme_pink)
        addColor(R.color.theme_red,R.color.theme_red_accent,R.string.theme_red)
        addColor(R.color.theme_yellow,R.color.theme_yellow_accent,R.string.theme_yellow)
        addColor(R.color.theme_green,R.color.theme_green_accent,R.string.theme_green)
        addColor(R.color.theme_blue,R.color.theme_blue_accent,R.string.theme_blue)
        addColor(R.color.theme_purple,R.color.theme_purple_accent,R.string.theme_purple)
        addColor(R.color.theme_default,R.color.theme_default_accent,R.string.theme_default)
        addColor(R.color.theme_dark,R.color.theme_dark_accent,R.string.theme_dark)
        // 初始化recycleView
        mDataBind.chooseList.layoutManager = GridLayoutManager(context,1)
        // 初始化adapter
        val adapter = ThemeAdapter(themes)
        mDataBind.chooseList.adapter = adapter
        // 设置adapter的点击时间
        adapter.setOnItemClickListener { _, _, position ->
            val data = adapter.getItem(position)
            // 保存主题
            ThemeManage.setTheme(data.primary,data.accent)
            // 设置颜色
            Cyanea.instance.edit{
                primary(data.color)
                accent(appContext.getColor(data.accent))
            }.recreate(activity)
            // 设置弹窗颜色
            XPopup.setPrimaryColor(appContext.getColor(data.primary))
        }
    }

    // 添加颜色
    private fun addColor(color:Int,accent:Int,title:Int){
        // 判断当前选中的主题
        val choose = theme.primary == color
        themes.add(Theme(context.getColor(color),color,accent,context.getString(title),choose))
    }

}