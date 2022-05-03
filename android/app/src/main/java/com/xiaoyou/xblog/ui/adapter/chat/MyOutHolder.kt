package com.xiaoyou.xblog.ui.adapter.chat

import android.view.View
import android.widget.TextView
import com.stfalcon.chatkit.messages.MessageHolders
import com.xiaoyou.library.common.util.TimeUtil
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.data.chat.ChatMessage
import io.noties.markwon.Markwon
import io.noties.markwon.html.HtmlPlugin
import io.noties.markwon.image.glide.GlideImagesPlugin

/**
 * @description
 * @author 小游
 * @data 2021/03/01
 */
class MyOutHolder (view: View, payload :Any?) : MessageHolders.OutcomingTextMessageViewHolder<ChatMessage>(view,payload){
    // 使用dataBind来绑定布局
    private val messageText : TextView = itemView.findViewById(R.id.messageText)
    private val messageTime : TextView = itemView.findViewById(R.id.messageTime)
    // 加载网络图片
    private val builder = Markwon
           .builder(itemView.context)
           .usePlugin(HtmlPlugin.create())
           .usePlugin(GlideImagesPlugin.create(itemView.context))
           .build()


    override fun onBind(message: ChatMessage?) {
       message?.apply {
           // 设置markdown
           builder.setMarkdown(messageText,message.text)
           messageTime.text = TimeUtil.formatTimeFromUnix("HH:mm",message.date)
       }
    }
}