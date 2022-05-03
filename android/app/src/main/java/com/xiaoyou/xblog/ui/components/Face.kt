package com.xiaoyou.xblog.ui.components

import android.content.Context
import android.os.Build
import android.util.AttributeSet
import android.view.LayoutInflater
import android.view.View
import android.widget.LinearLayout
import android.widget.TextView
import androidx.annotation.RequiresApi
import androidx.recyclerview.widget.GridLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.ui.adapter.common.FaceAdapter

/**
 * @description 自定义表情显示组件
 * @author 小游
 * @data 2021/02/18
 */
// 参考： https://www.jianshu.com/p/e287865175d9 来实现视图绑定
@RequiresApi(Build.VERSION_CODES.M)
class Face(context: Context, attrs: AttributeSet) : LinearLayout(context, attrs) {
    // 只能使用这种方式来加载布局文件，不能使用binding
    private val view = LayoutInflater.from(context).inflate(R.layout.component_face, this)
    private var faceAdapter = FaceAdapter(context)
    // 回调对象
    private lateinit var click: (View) -> Unit

    // 设置回调监听
    fun setListener(call: (View) -> Unit) {
        this.click = call
    }


    // 布局映射
    private val faceList : RecyclerView = view.findViewById(R.id.faceList)
    private val faceBiliBili : TextView = view.findViewById(R.id.faceBilibili)
    private val faceTv : TextView = view.findViewById(R.id.faceTv)
    private val faceZhiHu : TextView = view.findViewById(R.id.faceZhihu)
    private val faceTieBa : TextView = view.findViewById(R.id.faceTieba)
    private val faceText : TextView = view.findViewById(R.id.faceText)


    /**
     * 切换表情
     */
    private fun changeFace(item: String) {
        val face = "颜文字"
        faceBiliBili.setBackgroundColor(appContext.getColor(if ("bilibili" == item) R.color.face_back else R.color.white))
        faceTv.setBackgroundColor(appContext.getColor(if ("小电视" == item) R.color.face_back else R.color.white))
        faceZhiHu.setBackgroundColor(appContext.getColor(if ("知乎表情" == item) R.color.face_back else R.color.white))
        faceTieBa.setBackgroundColor(appContext.getColor(if ("贴吧泡泡" == item) R.color.face_back else R.color.white))
        faceText.setBackgroundColor(appContext.getColor(if ("颜文字" == item) R.color.face_back else R.color.white))
//        faceBiliBili.background.alpha = App.getAlpha(context)
//        faceTv.background.alpha = App.getAlpha(context)
//        faceZhiHu.background.alpha = App.getAlpha(context)
//        faceTieBa.background.alpha = App.getAlpha(context)
//        faceText.background.alpha = App.getAlpha(context)

        // 回调函数上报
        faceAdapter.setListener { click(it) }

        //切换内容
        faceAdapter.change(item)
        //颜文字自动切换宫格数目
        if (face == item) {
            faceList.layoutManager = GridLayoutManager(context, 3)
        } else {
            faceList.layoutManager = GridLayoutManager(context, 5)
        }
    }

    init {
        //显示所有的表情（每行显示5个）
        val layoutManager = GridLayoutManager(context, 5)
        layoutManager.orientation = RecyclerView.VERTICAL
        // 设置recycleView的adapter
        faceList.adapter = faceAdapter
        faceList.layoutManager = layoutManager
        // bilibili点击
        changeFace("bilibili")
        //点击事件监听
        faceBiliBili.setOnClickListener { changeFace("bilibili") }
        faceTv.setOnClickListener { changeFace("小电视") }
        faceZhiHu.setOnClickListener { changeFace("知乎表情") }
        faceTieBa.setOnClickListener { changeFace("贴吧泡泡") }
        faceText.setOnClickListener { changeFace("颜文字") }
    }
}