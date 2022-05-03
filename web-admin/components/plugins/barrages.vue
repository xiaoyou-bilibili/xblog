<template>
  <div>
    <vxe-grid
      ref="grid"
      border
      resizable
      height="780"
      row-id="id"
      :pager-config="{pageSize: 10}"
      :proxy-config="tableProxy"
      :form-config="tableForm"
      :toolbar-config="tableToolbar"
      :columns="tableColumn"
      @checkbox-change="checkboxChangeEvent"
      @checkbox-all="checkboxChangeEvent"
    >
      <!-- 友链头像 -->
      <template v-slot:avatar="{ row }">
        <div style="text-align: center">
          <img class="friend-avatar" :src="row.avatar">
        </div>
      </template>
      <template v-slot:content="{ row }">
        <span :style="{'color': row.color}" style="text-align: center;background: rgba(0,0,0,.3);padding: 3px;border-radius: 3px">{{ row.content }}</span>
      </template>
      <!-- 操作模板 -->
      <template v-slot:option="{ row }">
        <!--删除按钮 onConfirm 点击确认按钮激发-->
        <el-popconfirm title="确定要删除这条友链？" @onConfirm="deleteFriend(row.id)">
          <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
        </el-popconfirm>
      </template>
    </vxe-grid>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { pluginRequest } from '@/api/admin/plugins'
export default {
  layout: 'admin',
  mixins: [admin], // 混入对象用于批量修改文章状态
  data () {
    return {
      // 多选框选中的值
      selects: '',
      // 使用动态代理数据
      tableProxy: {
        ajax: {
          // 任何支持 Promise API 的库都可以对接（fetch、jquery、axios、xe-ajax）
          query: ({ page, form }) => pluginRequest({ url: '/barrages', type: 'get', data: { page: page.currentPage, page_size: page.pageSize, search_type: form.search_type, search_key: form.search_key } }),
          // 删除方法被触发
          delete: _ => this.deleteFriend(this.selects)
        }
      },
      // 搜索表单设置
      tableForm: {
        items: [
          {
            field: 'search_type',
            title: '搜索类型',
            itemRender: {
              name: '$select',
              options: [
                { label: '昵称', value: 'nickname' },
                { label: '弹幕内容', value: 'content' }
              ]
            }
          },
          // visibleMethod用来控制控件是否显示
          { field: 'search_key', visibleMethod: e => e.data.search_type !== 'status', title: '关键词', itemRender: { name: '$input', attrs: { placeholder: '请输入关键词' } } },
          { itemRender: { name: '$button', props: { content: '查询', type: 'submit', status: 'primary' } } },
          { itemRender: { name: '$button', props: { content: '重置', type: 'reset' } } }
        ]
      },
      // 工具栏设置
      tableToolbar: {
        buttons: [
          { code: 'delete', name: '删除选中', icon: 'fa fa-trash', status: 'danger' }
        ],
        zoom: true, // 缩放
        custom: true // 自定义显示内容
      },
      // 表格结构
      tableColumn: [
        { type: 'checkbox', width: 50 },
        { field: 'nickname', title: '用户昵称', editRender: { name: 'input' } },
        { field: 'avatar', title: '弹幕头像', slots: { default: 'avatar' }, editRender: { name: 'input' } },
        { field: 'content', title: '弹幕内容', slots: { default: 'content' }, editRender: { name: 'input' } },
        { field: 'send', title: '发送时间', editRender: { name: 'input' } },
        { title: '操作', slots: { default: 'option' } }
      ]
    }
  },
  methods: {
    // 多选框改变时触发的事件
    checkboxChangeEvent (data) {
      const ids = []
      // 把我们选中值的id提取出来
      data.records.map(item => ids.push(item.id))
      this.selects = ids.toString()
    },
    // 提示更新成功并刷新数据
    updateData (show = true) {
      // 是否需要提示更新信息
      if (show) { this.$message.success('更新成功') }
      // 调用commitProxy 手动更新数据
      this.$refs.grid.commitProxy('query')
    },
    // 更新友链，
    deleteFriend (id) {
      if (id === '') {
        this.$message.warning('请选择内容！')
      } else {
        this.$store.dispatch('admin-plugins/pluginRequest', { url: `/barrages/${id}`, type: 'delete', data: null }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
      }
    }
  }
}
</script>

<style scoped>

img.friend-avatar{
  width: 45px;
  height: 45px;
  border-radius: 50%;
}

.vxe-button+.vxe-button{
  margin-left: 0!important;
}
</style>
