package com.xiaoyou.library.plugins.music

import com.chad.library.adapter.base.viewholder.BaseDataBindingHolder
import com.xiaoyou.library.common.base.BaseAdapter
import com.xiaoyou.library.net.entity.response.MusicDetail
import com.xiaoyou.library.plugins.R
import com.xiaoyou.library.plugins.databinding.ItemMusicBinding

/**
 * @description 音乐播放器的adapter界面
 * @author 小游
 * @data 2021/03/10
 */
class MusicAdapter(data: MutableList<MusicDetail>): BaseAdapter<MusicDetail, ItemMusicBinding>(R.layout.item_music,data) {
    override fun convert(holder: BaseDataBindingHolder<ItemMusicBinding>, item: MusicDetail) {
        val binding = holder.dataBinding
        // 进行数据绑定
        binding?.let {
            it.item = item
            // 加载头像（因为图片可能不带http）
            var cover = item.cover
            if (item.cover.indexOf("http")==-1){
                cover = "https:${cover}"
            }
            it.musicCover.setImageURI(cover)
        }
    }

}