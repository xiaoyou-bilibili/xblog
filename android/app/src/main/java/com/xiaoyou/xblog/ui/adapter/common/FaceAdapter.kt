package com.xiaoyou.xblog.ui.adapter.common

import android.annotation.SuppressLint
import android.content.Context
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.facebook.drawee.view.SimpleDraweeView
import com.google.gson.Gson
import com.google.gson.reflect.TypeToken
import com.xiaoyou.library.common.base.SERVER
import com.xiaoyou.library.common.ext.readAssets
import com.xiaoyou.library.net.entity.response.Face
import com.xiaoyou.library.net.entity.response.FaceDetail
import com.xiaoyou.xblog.R

import java.util.*

/**
 *作者:created by HP on 2021/2/14 12:52
 *邮箱:sakurajimamai2020@qq.com
 */
class FaceAdapter(val context: Context) : RecyclerView.Adapter<FaceAdapter.ViewHolder?>() {
    private var faces: Map<String, Face>
    private var face: List<FaceDetail>? = null

    private lateinit var click: (View) -> Unit

    /**
     *  设置回调
     * @param call Function1<View, Unit>
     */
    fun setListener(call: (View) -> Unit) {
        this.click = call
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        val view = View.inflate(parent.context, R.layout.item_face, null)
        return ViewHolder(view)
    }

    @SuppressLint("SetTextI18n")
    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
        val src = "assets/images/smilies"
        //判断是文字还是图片
        if (face!![position].icon?.contains(src) == true) {
            holder.faceText.visibility = View.GONE
            holder.faceImg.visibility = View.VISIBLE
            holder.faceImg.setImageURI(SERVER+"/"+face!![position].icon)
            holder.faceImg.tag = face!![position].desc
            // 设置点击事件，我们通过回调上报事件
            holder.faceImg.setOnClickListener{ click(it) }
        } else {
            holder.faceImg.visibility = View.GONE
            holder.faceText.visibility = View.VISIBLE
            holder.faceText.text = face!![position].icon
            holder.faceText.tag = face!![position].desc
            // 设置点击事件，我们通过回调上报事件
            holder.faceText.setOnClickListener{ click(it) }
        }
    }

    override fun getItemCount(): Int {
        return face!!.size
    }

    inner class ViewHolder(item: View) : RecyclerView.ViewHolder(item) {
        //绑定控件
        var faceImg: SimpleDraweeView = item.findViewById(R.id.faceImg)
        var faceText: TextView = item.findViewById(R.id.faceText)
    }

    /**
     * 切换表情
     * @param item 切换的参数
     */
    fun change(item: String) {
        face = Objects.requireNonNull<Face>(faces[item]).container
        notifyDataSetChanged()
    }


    init {
        //这里获取所有的表情数据
        val wow: String = readAssets(context, "json/owo.json")
        val type = object : TypeToken<Map<String?, Face?>?>() {}.type
        val data: Map<String, Face> = Gson().fromJson(wow, type)
        faces = data
    }
}