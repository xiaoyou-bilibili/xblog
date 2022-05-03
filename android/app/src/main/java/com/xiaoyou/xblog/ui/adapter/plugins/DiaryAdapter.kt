package com.xiaoyou.xblog.ui.adapter.plugins

import android.content.Intent
import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.common.base.BaseAdapter
import com.xiaoyou.library.common.base.appContext
import com.xiaoyou.library.net.entity.response.DiaryDetail
import com.xiaoyou.library.net.entity.response.PostDetail
import com.xiaoyou.xblog.R
import com.xiaoyou.xblog.databinding.ItemDiaryBinding
import com.xiaoyou.xblog.databinding.ItemPostBinding
import com.xiaoyou.xblog.ui.activity.PostActivity

/**
 * @description 日记板块的adapter
 * @author 小游
 * @data 2021/02/24
 */
class DiaryAdapter(data: MutableList<DiaryDetail>) : BaseAdapter<DiaryDetail, ItemDiaryBinding>(R.layout.item_diary,data){

    override fun convert(holder: BaseDataBindingHolder<ItemDiaryBinding>, item: DiaryDetail) {
        val binding = holder.dataBinding
        // 进行数据绑定
        binding?.let {
            it.diary = item
            // 设置点击事件
            it.cardView.setOnClickListener{
                // 设置Intent
                val intent = Intent(appContext, PostActivity::class.java)
                intent.putExtra("id", item.diary_id)
                // 因为我们在非activity中启动，所以我们还需要设置flag
                // https://www.jianshu.com/p/165d88db2124
                intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK)
                // 启动activity
                appContext.startActivity(intent)
            }
        }
    }

}