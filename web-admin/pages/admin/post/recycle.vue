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
          <vxe-button icon="fa fa-refresh" title="恢复" circle @click="updatePost(row.id)" />
          <!--删除文章按钮 onConfirm 点击确认按钮激发-->
          <el-popconfirm title="确定要彻底删除这篇文章？" @onConfirm="updatePost(row.id, true)">
            <vxe-button slot="reference" icon="fa fa-trash" title="彻底删除" circle />
          </el-popconfirm>
        </template>
      </vxe-grid>
    </el-card>
  </div>
</template>

<script>
import admin from '@/components/mixin/admin-seo'
import { getTrash } from '@/api/admin/posts'
// 文章类型
const postType = [{ label: '文章', value: 'post' }, { label: '日记', value: 'diary' }, { label: '文档', value: 'doc' }]
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
          query: ({ page, form }) => getTrash({ page: page.currentPage, page_size: page.pageSize, search_type: form.search_type, search_key: form.search_key }),
          // 删除方法被触发
          delete: _ => this.updatePost(this.selects, true)
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
                { label: '文章id', value: 'post_id' },
                { label: '文章标题', value: 'title' },
                { label: '文章内容', value: 'content' },
                { label: '文章类型', value: 'post_type' }
              ]
            }
          },
          // visibleMethod用来控制控件是否显示
          { field: 'search_key', visibleMethod: e => e.data.search_type !== 'post_type', title: '关键词', itemRender: { name: '$input', attrs: { placeholder: '请输入文章标题' } } },
          { field: 'search_key', visibleMethod: e => e.data.search_type === 'post_type', title: '文章类型', itemRender: { name: '$select', options: postType } },
          { itemRender: { name: '$button', props: { content: '查询', type: 'submit', status: 'primary' } } },
          { itemRender: { name: '$button', props: { content: '重置', type: 'reset' } } }
        ]
      },
      // 工具栏设置
      tableToolbar: {
        buttons: [
          { code: 'restore', name: '恢复选中', icon: 'fa fa-refresh', status: 'success' },
          { code: 'delete', name: '彻底删除选中', icon: 'fa fa-trash', status: 'danger' }
        ],
        zoom: true, // 缩放
        custom: true // 自定义显示内容
      },
      // 表格结构
      tableColumn: [
        { type: 'checkbox', width: 50 },
        { field: 'id', title: '文章id', width: 100 },
        { field: 'title', title: '文章标题' },
        { field: 'content', title: '文章内容' },
        { field: 'type', title: '文章类型', formatter: 'formatPostKinds', width: 100 },
        { title: '操作', slots: { default: 'option' } }
      ]
    }
  },
  methods: {
    // 工具栏按钮点击事件
    toolbarClick ({ code }) {
      // 当我们点击恢复时恢复删除的文章
      if (code === 'restore') { this.updatePost(this.selects) }
    },
    // 多选框改变时触发的事件
    checkboxChangeEvent (data) {
      const ids = []
      // 把我们选中值的id提取出来
      data.records.map(item => ids.push(item.id))
      this.selects = ids.toString()
    },
    // 更新文章 但trash为true的时候表示删除文章
    updatePost (id, trash = false) {
      if (id === '') {
        this.$message.warning('请选择内容！')
      } else {
        // 表示我们恢复
        const data = { delete: 'false' }
        // 更新表格内容
        const update = () => {
          this.$message.success('更新成功')
          // 调用commitProxy 手动更新数据
          this.$refs.grid.commitProxy('query')
        }
        // 判断是更新还是删除
        if (!trash) {
          // 更新评论
          this.$store.dispatch('admin-posts/updatePost', { id, data }).then(_ => update()).catch(msg => this.$message.error(msg))
        } else {
          // 删除评论
          this.$store.dispatch('admin-posts/deletePost', id).then(_ => update()).catch(msg => this.$message.error(msg))
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
<!--<template>-->
<!--  <div>-->
<!--    <el-card class="box-card">-->
<!--      &lt;!&ndash;显示表格&ndash;&gt;-->
<!--      <el-table-->
<!--        :data="deleteList"-->
<!--        border-->
<!--        style="width: 100%;margin: 10px"-->
<!--      >-->
<!--        <el-table-column-->
<!--          prop="title"-->
<!--          label="标题"-->
<!--          width="300"-->
<!--          fixed="left"-->
<!--        />-->
<!--        <el-table-column-->
<!--          prop="content"-->
<!--          label="内容"-->
<!--          width="300"-->
<!--          fixed="left"-->
<!--        />-->
<!--        <el-table-column-->
<!--          prop="type"-->
<!--          label="类型"-->
<!--          width="100"-->
<!--        >-->
<!--          <template slot-scope="scope">-->
<!--            <div v-if="scope.row.type==='post'" style="color: #67C23A">-->
<!--              文章-->
<!--            </div>-->
<!--            <div v-else-if="scope.row.type==='doc'" style="color: #F56C6C">-->
<!--              文档-->
<!--            </div>-->
<!--            <div v-else-if="scope.row.type==='diary'" style="color: #909399">-->
<!--              日记-->
<!--            </div>-->
<!--          </template>-->
<!--        </el-table-column>-->
<!--        <el-table-column-->
<!--          prop="option"-->
<!--          label="操作"-->
<!--          width="220"-->
<!--        >-->
<!--          <template slot-scope="scope">-->
<!--            <el-button-group>-->
<!--              <el-tooltip content="恢复" placement="top">-->
<!--                <el-button type="success" size="mini" icon="el-icon-refresh" @click="updateItem(scope.row.id, 'restore')" />-->
<!--              </el-tooltip>-->
<!--              <el-tooltip content="永久删除" placement="top">-->
<!--                <el-popconfirm-->
<!--                  icon-color="red"-->
<!--                  title="你确定要永久删除（此操作不可逆）？"-->
<!--                  @onConfirm="updateItem(scope.row.id, 'delete')"-->
<!--                >-->
<!--                  <el-button slot="reference" type="danger" size="mini" icon="el-icon-delete" />-->
<!--                </el-popconfirm>-->
<!--              </el-tooltip>-->
<!--            </el-button-group>-->
<!--          </template>-->
<!--        </el-table-column>-->
<!--      </el-table>-->
<!--    </el-card>-->
<!--  </div>-->
<!--</template>-->

<!--<script>-->
<!--import { mapGetters } from 'vuex'-->
<!--import admin from '@/components/mixin/admin-seo'-->

<!--export default {-->
<!--  layout: 'admin',-->
<!--  mixins: [admin],-->
<!--  computed: {-->
<!--    // 把getter混入到computed对象中,我们可以直接调用vex里面的内容即可-->
<!--    ...mapGetters('admin-post', ['deleteList'])-->
<!--  },-->
<!--  mounted () {-->
<!--    this.getData()-->
<!--  },-->
<!--  methods: {-->
<!--    getData () {-->
<!--      // 获取文章内容-->
<!--      this.$store.dispatch('admin-post/getDeleteList')-->
<!--    },-->
<!--    updateItem (id, option) {-->
<!--      this.$store.dispatch('admin-post/updateDeleteList', { id, option }).then(() => {-->
<!--        this.getData()-->
<!--        this.$message.success('操作成功')-->
<!--      }).catch(msg => this.$message.error(msg))-->
<!--    }-->
<!--  }-->
<!--}-->
<!--</script>-->

<!--<style scoped>-->
<!--.table-tag{-->
<!--  margin: 2px 5px;-->
<!--}-->
<!--</style>-->
