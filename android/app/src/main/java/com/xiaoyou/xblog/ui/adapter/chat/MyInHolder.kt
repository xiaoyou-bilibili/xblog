package com.xiaoyou.xblog.ui.adapter.chat

import android.content.Intent
import android.view.View
import androidx.databinding.DataBindingUtil
import com.stfalcon.chatkit.messages.MessageHolders
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.common.util.ImageUtil
import com.xiaoyou.library.common.util.TimeUtil
import com.xiaoyou.xblog.data.chat.ChatMessage
import com.xiaoyou.xblog.databinding.ItemChatBinding
import com.xiaoyou.xblog.ui.activity.ChatActivity
import io.noties.markwon.Markwon
import io.noties.markwon.html.HtmlPlugin
import io.noties.markwon.image.glide.GlideImagesPlugin

/**
 * @description 自定义聊天布局，这里是别人发送消息的布局
 * @author 小游
 * @data 2021/03/01
 */
// 图片加载参考： https://noties.io/Markwon/docs/v4/html/
class MyInHolder(view: View, payload :Any?) : MessageHolders.IncomingTextMessageViewHolder<ChatMessage>(view,payload){
    // 使用dataBind来绑定布局
    private val bind = DataBindingUtil.bind<ItemChatBinding>(itemView)
    // 加载网络图片
    private val builder = Markwon
            .builder(itemView.context)
            .usePlugin(HtmlPlugin.create())
            .usePlugin(GlideImagesPlugin.create(itemView.context))
            .build()

    override fun onBind(message: ChatMessage?) {
        bind?.apply {
            message?.apply {
                // 设置头像
                ImageUtil.setImageByUrl(user.avatar, bind.messageUserAvatar)
                // 设置dataBind
                bind.chat = message
                // 设置markdown
                builder.setMarkdown(bind.messageText,message.text)
                // 设置时间
                bind.messageTime.text = TimeUtil.formatTimeFromUnix("HH:mm",message.date)
                // 如果用户头像被点击了
                bind.messageUserAvatar.setOnClickListener{
                    val intent = Intent(appContext, ChatActivity::class.java)
                    // 添加参数
                    intent.putExtra("title",message.user.name)
                    intent.putExtra("target",message.user.id)
                    // 启动activity
                    itemView.context.startActivity(intent)
                }
            }
        }
    }
}