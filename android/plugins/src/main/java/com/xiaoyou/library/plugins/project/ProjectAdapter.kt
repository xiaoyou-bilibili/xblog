package com.xiaoyou.library.plugins.project

import android.text.method.ScrollingMovementMethod
import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.common.base.BaseAdapter
import com.xiaoyou.library.net.entity.response.ProjectBottom
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.databinding.ItemProjectBinding

/**
 * @description 我的豆瓣的adapter
 * @author 小游
 * @data 2021/03/27
 */
class ProjectAdapter (data: MutableList<ProjectBottom>) : BaseAdapter<ProjectBottom, ItemProjectBinding>(R.layout.item_project,data){
    override fun convert(holder: BaseDataBindingHolder<ItemProjectBinding>, item: ProjectBottom) {
        val binding = holder.dataBinding
        // 进行数据绑定
        binding?.let {
            it.projectCover.setImageURI(item.image)
            // 设置textView可以滚动
            it.projectDec.movementMethod = ScrollingMovementMethod.getInstance()
            // 设置数据
            it.item = item
        }
    }
}