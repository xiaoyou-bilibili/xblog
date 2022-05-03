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
        @toolbar-button-click="toolbarClick"
        @checkbox-change="checkboxChangeEvent"
        @checkbox-all="checkboxChangeEvent"
      >
        <!-- 操作模板 -->
        <template v-slot:option="{ row }">
          <vxe-button v-if="row.status" icon="fa fa-close" title="拒绝" circle @click="updateComment(row.id,'0')" />
          <vxe-button v-else icon="fa fa-check" title="允许" circle @click="updateComment(row.id,'1')" />
          <!--删除文章按钮 onConfirm 点击确认按钮激发-->
          <el-popconfirm title="确定要删除这条评论？" @onConfirm="updateComment(row.id,null, true)">
            <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
          </el-popconfirm>
        </template>
      </vxe-grid>
    </el-card>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { getComments } from '@/api/admin/comments'
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
          query: ({ page, form }) => getComments({ page: page.currentPage, page_size: page.pageSize, search_type: form.search_type, search_key: form.search_key }),
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
                { label: '昵称', value: 'nickname' },
                { label: '评论内容', value: 'content' },
                { label: '用户邮箱', value: 'email' },
                { label: '用户ip', value: 'ip' },
                { label: '评论状态', value: 'agree' }
              ]
            }
          },
          // visibleMethod用来控制控件是否显示
          { field: 'search_key', visibleMethod: e => e.data.search_type !== 'agree', title: '关键词', itemRender: { name: '$input', attrs: { placeholder: '请输入文章标题' } } },
          { field: 'search_key', visibleMethod: e => e.data.search_type === 'agree', title: '评论状态', itemRender: { name: '$select', options: [{ label: '已通过', value: '1' }, { label: '已拒绝', value: '0' }] } },
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
      // 表格结构
      tableColumn: [
        { type: 'checkbox', width: 50 },
        { field: 'author', title: '昵称' },
        { field: 'content', title: '评论内容' },
        { field: 'ip', title: '用户ip' },
        { field: 'email', title: '邮箱' },
        { field: 'status', title: '状态', formatter: 'formatReview', width: 100 },
        { field: 'date', title: '评论时间' },
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
    // 更新评论，当草稿字段为null的时候就代表我们需要删除
    updateComment (id, agree, trash = false) {
      if (id === '') {
        this.$message.warning('请选择内容！')
      } else {
        const data = { agree }
        // 更新表格内容
        const update = () => {
          this.$message.success('更新成功')
          // 调用commitProxy 手动更新数据
          this.$refs.grid.commitProxy('query')
        }
        // 判断是更新还是删除
        if (!trash) {
          // 更新评论
          this.$store.dispatch('admin-comments/updateComments', { id, data }).then(_ => update()).catch(msg => this.$message.error(msg))
        } else {
          // 删除评论
          this.$store.dispatch('admin-comments/deleteComments', id).then(_ => update()).catch(msg => this.$message.error(msg))
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
