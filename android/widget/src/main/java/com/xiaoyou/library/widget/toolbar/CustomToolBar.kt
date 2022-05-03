package com.xiaoyou.library.widget.toolbar

import android.content.Context
import android.util.AttributeSet
import android.view.LayoutInflater
import android.widget.FrameLayout
import androidx.appcompat.widget.AppCompatTextView
import androidx.appcompat.widget.Toolbar
import com.xiaoyou.library.widget.R
import de.hdodenhof.circleimageview.CircleImageView

/**
 * @description 自定义顶部工具栏
 * @author 小游
 * @data 2021/02/20
 */
// 参考https://www.jianshu.com/p/fad3265f7e37
class CustomToolBar : FrameLayout {

    private lateinit var toolBar: Toolbar
    private lateinit var toolBarTitle: AppCompatTextView
    private lateinit var topImg: CircleImageView

    constructor(context: Context) : super(context)

    constructor(context: Context, attrs: AttributeSet) : super(context, attrs) {
        init(context, attrs)
    }

    constructor(context: Context, attrs: AttributeSet, defStyleAttr: Int) : super(
        context,
        attrs,
        defStyleAttr
    ) {
        init(context, attrs)
    }

    private fun init(context: Context, attrs: AttributeSet) {
        val view = LayoutInflater.from(context).inflate(R.layout.toolbar_layout_custom, this)
        toolBar = view.findViewById(R.id.toolBar)
        toolBar.inflateMenu(R.menu.menu_top)
        toolBar.title = ""

        // 自定义标题，因为我们需要文字显示在中央
        toolBarTitle = view.findViewById(R.id.toolbarTitle)
        // 获取顶部图片
        topImg = view.findViewById(R.id.profileImage)
    }

    // 设置是否显示搜索按钮
    fun setSearchIcon(boolean: Boolean){
        toolBar.menu.getItem(0).isVisible = boolean
    }

    // 设置标题
    fun setCenterTitle(titleStr: String) {
        toolBarTitle.text = titleStr
    }
    // 设置标题 id形式
    fun setCenterTitle(titleResId: Int) {
        toolBarTitle.text = context.getString(titleResId)
    }
    // 设置标题颜色
    fun setCenterTitleColor(colorResId: Int) {
        toolBarTitle.setTextColor(colorResId)
    }
    // 返回toobar对象
    fun getBaseToolBar(): Toolbar = toolBar

    // 返回topImage对象
    fun getTopImg() :CircleImageView{
        // 先设置头像为显示状态
        topImg.visibility = VISIBLE
        return topImg
    }
}
