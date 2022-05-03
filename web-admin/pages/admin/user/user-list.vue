<template>
  <div>
    <el-card class="box-card">
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
        <!-- 操作模板 -->
        <template v-slot:option="{ row }">
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
          <vxe-button v-if="row.status" icon="fa fa-close" title="拒绝" circle @click="updateComment(row.id,'0')" />
          <vxe-button v-else icon="fa fa-check" title="允许" circle @click="updateComment(row.id,'1')" />
          <!--删除按钮 onConfirm 点击确认按钮激发-->
          <el-popconfirm title="确定要删除这个用户？" @onConfirm="updateComment(row.id,null, true)">
            <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
          </el-popconfirm>
        </template>
      </vxe-grid>
    </el-card>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { getUserList } from '@/api/admin/users'
// 定义用户状态
const userStatus = [{ label: '已激活', value: 1 }, { label: '未激活', value: 0 }]
const userSubscription = [{ label: '未订阅', value: false }, { label: '已订阅', value: true }]
const userIdentity = [{ label: '管理员', value: 1 }, { label: '普通用户', value: 2 }]
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
          query: ({ page, form }) => getUserList({ page: page.currentPage, page_size: page.pageSize, search_type: form.search_type, search_key: form.search_key }),
          // 删除方法被触发
          delete: _ => this.updateComment(this.selects, null, true)
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
                { label: '用户id', value: 'user_id' },
                { label: '用户名', value: 'username' },
                { label: '昵称', value: 'nickname' },
                { label: '用户邮箱', value: 'email' },
                { label: '是否订阅', value: 'subscription' },
                { label: '用户状态', value: 'status' },
                { label: '用户权限', value: 'identity' }
              ]
            }
          },
          // visibleMethod用来控制控件是否显示
          { field: 'search_key', visibleMethod: e => e.data.search_type !== 'status' && e.data.search_type !== 'subscription' && e.data.search_type !== 'identity', title: '关键词', itemRender: { name: '$input', attrs: { placeholder: '请输入文章标题' } } },
          { field: 'search_key', visibleMethod: e => e.data.search_type === 'status', title: '用户状态', itemRender: { name: '$select', options: userStatus } },
          { field: 'search_key', visibleMethod: e => e.data.search_type === 'subscription', title: '是否订阅', itemRender: { name: '$select', options: userSubscription } },
          { field: 'search_key', visibleMethod: e => e.data.search_type === 'identity', title: '用户权限', itemRender: { name: '$select', options: userIdentity } },
          { itemRender: { name: '$button', props: { content: '查询', type: 'submit', status: 'primary' } } },
          { itemRender: { name: '$button', props: { content: '重置', type: 'reset' } } }
        ]
      },
      // 工具栏设置
      tableToolbar: {
        buttons: [
          { code: 'agree', name: '允许', icon: 'fa fa-check', status: 'success' },
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
        { field: 'id', title: '用户id', width: 80 },
        { field: 'username', title: '用户名', editRender: { name: 'input' } },
        { field: 'nickname', title: '昵称', editRender: { name: 'input' } },
        { field: 'email', title: '邮箱', editRender: { name: 'input' } },
        { field: 'registered', title: '注册时间' },
        { field: 'last_login', title: '上次登录时间' },
        { field: 'login_ip', title: '上次登录ip' },
        { field: 'subscription', title: '是否订阅', formatter: 'formatUserSubscription', width: 100, editRender: { name: '$select', options: userSubscription } },
        { field: 'status', title: '状态', formatter: 'formatUserStatus', width: 100, editRender: { name: '$select', options: userStatus } },
        { field: 'identity', title: '权限', formatter: 'formatIdentity', width: 100, editRender: { name: '$select', options: userIdentity } },
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
          this.updateComment(this.selects, '1')
          break
        case 'deny':
          this.updateComment(this.selects, '0')
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
          username: row.username,
          nickname: row.nickname,
          email: row.email,
          status: row.status.toString(),
          subscription: row.subscription.toString(),
          identity: row.identity.toString()
        }
        this.$store.dispatch('admin-users/updateUser', { id: row.id, data }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
      })
    },
    // 提示更新成功并刷新数据
    updateData () {
      this.$message.success('更新成功')
      // 调用commitProxy 手动更新数据
      this.$refs.grid.commitProxy('query')
    },
    // 更新评论，当草稿字段为null的时候就代表我们需要删除
    updateComment (id, status, trash = false) {
      if (id === '') {
        this.$message.warning('请选择内容！')
      } else {
        const data = { status }
        // 判断是更新还是删除
        if (!trash) {
          // 更新评论
          this.$store.dispatch('admin-users/updateUser', { id, data }).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
        } else {
          // 删除评论
          this.$store.dispatch('admin-users/deleteUser', id).then(_ => this.updateData()).catch(msg => this.$message.error(msg))
        }
      }
    }
  }
}
</script>

<style scoped>

.vxe-button+.vxe-button{
  margin-left: 0!important;
}
</style>
