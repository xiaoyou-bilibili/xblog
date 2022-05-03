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
      :edit-config="editConfig"
      @toolbar-button-click="toolbarClick"
      @checkbox-change="checkboxChangeEvent"
      @checkbox-all="checkboxChangeEvent"
    >
      <!-- 友链头像 -->
      <template v-slot:avatar="{ row }">
        <div style="text-align: center">
          <img class="friend-avatar" :src="row.avatar">
        </div>
      </template>
      <!-- 操作模板 -->
      <template v-slot:option="{ row }">
        <!-- 邮件通知按钮 -->
        <vxe-button icon="fa fa-envelope" title="发布通过通知" circle @click="sendNotification(row.id)" />
        <!-- 编辑保存按钮 -->
        <vxe-button
          v-if="$refs.grid.isActiveByRow(row)"
          icon="fa fa-save"
          status="primary"
          title="保存"
          circle
          @click="saveRowEvent(row)"
        />
        <vxe-button v-else icon="fa fa-edit" title="编辑" circle @click="editRowEvent(row)" />
        <!-- 允许拒绝 -->
        <vxe-button v-if="row.status" icon="fa fa-close" title="拒绝" circle @click="updateFriends(row.id,'0')" />
        <vxe-button v-else icon="fa fa-check" title="允许" circle @click="updateFriends(row.id,'1')" />
        <!--删除按钮 onConfirm 点击确认按钮激发-->
        <el-popconfirm title="确定要删除这条友链？" @onConfirm="updateFriends(row.id,null, true)">
          <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
        </el-popconfirm>
      </template>
    </vxe-grid>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { pluginRequest } from '@/api/admin/plugins'
// 定义友链状态
const friendReview = [{ label: '已通过', value: 1 }, { label: '已拒绝', value: 0 }]
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
          query: ({ page, form }) => pluginRequest({ url: '/friends', type: 'get', data: { page: page.currentPage, page_size: page.pageSize, search_type: form.search_type, search_key: form.search_key } }),
          // 删除方法被触发
          delete: _ => this.updateFriends(this.selects, null, true)
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
                { label: '友链名称', value: 'name' },
                { label: '友链地址', value: 'url' },
                { label: '友链邮箱', value: 'email' },
                { label: '友链描述', value: 'description' },
                { label: '友链状态', value: 'status' }
              ]
            }
          },
          // visibleMethod用来控制控件是否显示
          { field: 'search_key', visibleMethod: e => e.data.search_type !== 'status', title: '关键词', itemRender: { name: '$input', attrs: { placeholder: '请输入关键词' } } },
          { field: 'search_key', visibleMethod: e => e.data.search_type === 'status', title: '友链状态', itemRender: { name: '$select', options: friendReview } },
          { itemRender: { name: '$button', props: { content: '查询', type: 'submit', status: 'primary' } } },
          { itemRender: { name: '$button', props: { content: '重置', type: 'reset' } } }
        ]
      },
      // 工具栏设置
      tableToolbar: {
        buttons: [
          { code: 'insert_actived', name: '添加', icon: 'fa fa-plus', status: 'primary' },
          { code: 'agree', name: '通过', icon: 'fa fa-check', status: 'success' },
          { code: 'deny', name: '拒绝', icon: 'fa fa-close', status: 'warning' },
          { code: 'delete', name: '删除选中', icon: 'fa fa-trash', status: 'danger' }
        ],
        zoom: true, // 缩放
        custom: true // 自定义显示内容
      },
      // 开启行编辑功能
      editConfig: {
        // 手动触发
        trigger: 'manual',
        // 单元格编辑模式
        mode: 'row',
        // 点击其他地方不清除激活状态
        autoClear: false,
        // 不显示按钮
        icon: 'none'
      },
      // 表格结构
      tableColumn: [
        { type: 'checkbox', width: 50 },
        { field: 'name', title: '友链名称', editRender: { name: 'input' } },
        { field: 'url', title: '友链地址', editRender: { name: 'input' } },
        { field: 'avatar', title: '友链头像', slots: { default: 'avatar' }, editRender: { name: 'input' } },
        { field: 'email', title: '友链邮箱', editRender: { name: 'input' } },
        { field: 'description', title: '友链描述', editRender: { name: 'input' } },
        { field: 'status', title: '友链状态', formatter: 'formatReview', width: 100, editRender: { name: '$select', options: friendReview } },
        { title: '操作', slots: { default: 'option' } }
      ]
    }
  },
  methods: {
    // 工具栏按钮点击事件
    toolbarClick ({ code, button }) {
      // 根据我们的code来判断不同操作
      switch (code) {
        case 'agree':
          this.updateFriends(this.selects, '1')
          break
        case 'deny':
          this.updateFriends(this.selects, '0')
          break
      }
    },
    // 多选框改变时触发的事件
    checkboxChangeEvent (data) {
      const ids = []
      // 把我们选中值的id提取出来
      data.records.map(item => ids.push(item.id))
      this.selects = ids.toString()
    },
    // 点击编辑事件
    editRowEvent (row) {
      // 我们激活编辑框
      this.$refs.grid.setActiveRow(row)
    },
    // 点击保存事件
    saveRowEvent (row) {
      // 点击保存的时候自动触发代理的保存事件
      this.$refs.grid.clearActived().then(() => {
        // 设置更新的值
        const data = {
          name: row.name,
          url: row.url,
          avatar: row.avatar,
          description: row.description,
          email: row.email,
          status: row.status === 0 ? '0' : '1'
        }
        // 判断是新增还是保存
        if (row.id.includes('row')) {
          this.$store.dispatch('admin-plugins/pluginRequest', { url: '/friends', type: 'post', data }).then(_ => this.updateData()).catch((msg) => { this.$message.error(msg); this.updateData(false) })
        } else {
          this.$store.dispatch('admin-plugins/pluginRequest', { url: `/friends/${row.id}`, type: 'put', data }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
        }
      })
    },
    // 提示更新成功并刷新数据
    updateData (show = true) {
      // 是否需要提示更新信息
      if (show) { this.$message.success('更新成功') }
      // 调用commitProxy 手动更新数据
      this.$refs.grid.commitProxy('query')
    },
    // 更新友链，
    updateFriends (id, status, trash = false) {
      if (id === '') {
        this.$message.warning('请选择内容！')
      } else {
        const data = { status }
        // 判断是更新还是删除
        if (!trash) {
          // 更新评论
          this.$store.dispatch('admin-plugins/pluginRequest', { url: `/friends/${id}`, type: 'put', data }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
        } else {
          // 删除评论
          this.$store.dispatch('admin-plugins/pluginRequest', { url: `/friends/${id}`, type: 'delete', data: null }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
        }
      }
    },
    // 发送通过通知
    sendNotification (id) {
      this.$store.dispatch('admin-plugins/pluginRequest', { url: `/friends/${id}/notification`, type: 'put', data: null }).then(_ => this.$message.success('已发布邮件通知')).catch(msg => this.$message.error(msg))
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
