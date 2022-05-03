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
        <!-- 分类模板 -->
        <template v-slot:category="{ row }">
          <el-tag v-for="(item,index) in row.category" :key="index" size="small" class="table-tag">
            {{ item }}
          </el-tag>
        </template>
        <!-- 标签模板 -->
        <template v-slot:tag="{ row }">
          <el-tag v-for="(item,index) in row.tags" :key="index" type="success" size="mini" class="table-tag">
            {{ item }}
          </el-tag>
        </template>
        <!-- 操作模板 -->
        <template v-slot:option="{ row }">
          <vxe-button icon="fa fa-edit" title="编辑" circle @click="gotoEdit(row.id)" />
          <vxe-button v-if="row.is_draft" icon="fa fa-send" title="发布" circle @click="updatePost(row.id,'false')" />
          <vxe-button v-else icon="fa fa-file-o" title="草稿" circle @click="updatePost(row.id,'true')" />
          <!-- 发送订阅通知 -->
          <vxe-button icon="fa fa-envelope" title="发布邮件订阅通知" circle @click="noticeUser(row.id)" />
          <!--删除文章按钮 onConfirm 点击确认按钮激发-->
          <el-popconfirm title="确定要删除这篇文章？" @onConfirm="updatePost(row.id,null)">
            <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
          </el-popconfirm>
        </template>
      </vxe-grid>
    </el-card>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { getArticles } from '@/api/admin/posts'
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
          query: ({ page, form }) => getArticles({ page: page.currentPage, page_size: page.pageSize, search_type: form.search_type, search_key: form.search_key }),
          // 删除方法被触发
          delete: _ => this.updatePost(this.selects, null)
        }
      },
      // 搜索表单设置
      tableForm: {
        items: [
          { field: 'search_type', title: '搜索类型', itemRender: { name: '$select', options: [{ label: '文章id', value: 'post_id' }, { label: '文章标题', value: 'title' }] } },
          { field: 'search_key', title: '关键词', itemRender: { name: '$input', attrs: { placeholder: '请输入文章标题', name: 'sex' } } },
          { itemRender: { name: '$button', props: { content: '查询', type: 'submit', status: 'primary' } } },
          { itemRender: { name: '$button', props: { content: '重置', type: 'reset' } } }
        ]
      },
      // 工具栏设置
      tableToolbar: {
        buttons: [
          { code: 'write', name: '写文章', icon: 'fa fa-plus', status: 'primary' },
          { code: 'publish', name: '发布', icon: 'fa fa-send', status: 'success' },
          { code: 'cancel', name: '取消发布', status: 'warning' },
          { code: 'delete', name: '删除选中', icon: 'fa fa-trash', status: 'danger' }
        ],
        zoom: true, // 缩放
        custom: true // 自定义显示内容
      },
      // 表格结构
      tableColumn: [
        { type: 'checkbox', width: 50 },
        { field: 'id', title: '文章id', width: 90 },
        { field: 'title', title: '文章标题' },
        { title: '分类', slots: { default: 'category' } },
        { title: '标签', slots: { default: 'tag' } },
        { field: 'view', title: '阅读数', width: 100 },
        { field: 'good', title: '点赞数', width: 100 },
        { field: 'comment', title: '评论数', width: 100 },
        { field: 'status', title: '类型', formatter: 'formatPostType', width: 100 },
        { field: 'is_draft', title: '状态', formatter: 'formatPostStatus', width: 100 },
        { field: 'date', title: '发布时间' },
        { title: '操作', slots: { default: 'option' } }
      ]
    }
  },
  methods: {
    // 工具栏按钮点击事件
    toolbarClick ({ code, button }) {
      // 根据我们的code来判断不同操作
      switch (code) {
        case 'write':
          this.gotoEdit(0)
          break
        case 'publish':
          this.updatePost(this.selects, 'false')
          break
        case 'cancel':
          this.updatePost(this.selects, 'true')
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
    // 更新文章 因为我们有草稿和删除操作，当草稿字段为null的时候就代表我们需要删除
    updatePost (id, draft) {
      if (id === '') {
        this.$message.warning('请选择内容！')
      } else {
        let data = {}
        // 先判断用户操作
        if (draft != null) {
          data = { is_draft: draft }
        } else {
          data = { delete: 'true' }
        }
        // 更新文章内容
        this.$store.dispatch('admin-posts/updatePost', { id, data }).then((_) => {
          this.$message.success('更新成功')
          // 调用commitProxy 手动更新数据
          this.$refs.grid.commitProxy('query')
        }).catch(msg => this.$message.error(msg))
      }
    },
    // 发布订阅信息
    noticeUser (id) {
      this.$store.dispatch('admin-posts/noticeUser', id).then(_ => this.$message.success('发布订阅信息成功')).catch(msg => this.$message.error(msg))
    },
    // 跳转到文章编辑器
    gotoEdit (id) {
      let path = '/admin/post/post-edit'
      if (id !== 0) {
        path += '/' + id
      }
      // 跳转路由
      this.$router.push(path)
      // 添加标签
      this.$store.dispatch('admin/editChangeId', { title: '文章编辑器', path })
    }
  }
}
</script>

<style scoped>
.table-tag{
  margin: 2px 5px;
}
.vxe-button+.vxe-button{
  margin-left: 0!important;
}
</style>
