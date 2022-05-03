package com.xiaoyou.xblog.ui.adapter.posts

import android.content.Intent
import androidx.core.content.ContextCompat.startActivity
import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.common.base.BaseAdapter
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.databinding.ItemPostBinding
import com.xiaoyou.xblog.ui.activity.PostActivity

/**
 * @description 文章板块adapter,基础了baseAdapter后会自动绑定布局
 * @author 小游
 * @data 2021/02/21
 */
// data bind 参考：https://blog.csdn.net/weixin_44613063/article/details/103553014
class PostAdapter(data: MutableList<PostDetail>) : BaseAdapter<PostDetail,ItemPostBinding>(R.layout.item_post,data){

    /**
     * @param holder A fully initialized helper.
     * @param item   The item that needs to be displayed.
     */
    override fun convert(holder: BaseDataBindingHolder<ItemPostBinding>, item: PostDetail) {
        val binding = holder.dataBinding
        // 进行数据绑定
        binding?.let {
            it.post = item
            // 设置图片
            it.postImg.setImageURI(item.image)
            // 设置点击事件
            it.cardView.setOnClickListener{
                // 设置Intent
                val intent = Intent(appContext, PostActivity::class.java)
                intent.putExtra("id", item.id)
                intent.putExtra("img",item.image)
                // 因为我们在非activity中启动，所以我们还需要设置flag
                // https://www.jianshu.com/p/165d88db2124
                intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK)
                // 启动activity
                appContext.startActivity(intent)
            }
        }
    }
}