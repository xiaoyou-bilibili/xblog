<template>
  <div>
    <el-row :gutter="20">
      <!--功能选择区-->
      <el-col :xs="24" :sm="9" :md="4">
        <el-card class="option-card" shadow="hover">
          <div slot="header" style="background: #996600;" class="card-head">
            <font-awesome-icon icon="list" />
            <span>文档目录</span>
            <span v-show="saveStatus" style="font-size: 12px"><i class="el-icon-circle-check" />已自动保存</span>
          </div>
          <el-tree
            ref="tree"
            class="post-tree"
            :data="docList"
            draggable
            :highlight-current="true"
            :expand-on-click-node="false"
            @node-contextmenu="nodeRightClick"
            @node-click="nodeClick"
            @node-drop="nodeDrop"
          />
          <div class="add_chapter">
            <el-button plain @click="addNewChapter(0)">
              新增章节
            </el-button>
          </div>
        </el-card>
      </el-col>
      <!--文章编辑器-->
      <el-col :xs="24" :sm="15" :md="20">
        <!--使用sync的目的是为了让子组件改变父主键的值-->
        <post-edit ref="edit" :html.sync="doc.content" :md.sync="doc.md" />
      </el-col>
    </el-row>
    <div v-show="visible" :style="{left:left+'px',top:top+'px'}" class="contextmenu">
      <div @click="addNewChapter(clickRightItem.id, clickRightItem)">
        新建子章节
      </div>
      <div @click="changeChapterTitle(clickRightItem)">
        编辑
      </div>
      <div @click="delChapter(clickRightItem, clickRightNode)">
        删除
      </div>
    </div>
  </div>
</template>

<script>
import _ from 'lodash'
import { mapGetters } from 'vuex'
import PostEdit from '@/components/content/admin/content/post-edit'
import admin from '@/components/mixin/admin-seo'

export default {
  components: { PostEdit },
  layout: 'admin',
  mixins: [admin],
  data () {
    return {
      left: 0, // 右键菜单的左边距
      top: 0, // 右键菜单的上边距
      visible: false, // 右键菜单是否可见
      clickRightItem: {}, // 右键选中的对象
      clickRightNode: {}, // 右键选中对象的node信息
      id: 0, // 当前选中的文档id
      doc: { // 文档的内容
        content: '',
        md: ''
      },
      flag: false, // 这个flag用于跳过watch监听，避免自动更新死循环
      saveStatus: false // 当前文档的保存状态
    }
  },
  computed: {
    // 把getter混入到computed对象中,我们可以直接调用vex里面的内容即可
    ...mapGetters('admin-posts', ['docList', 'docContent'])
  },
  watch: {
    visible (value) { // 监听右键菜单显示变化
      if (value) {
        document.body.addEventListener('click', this.closeMenu)
      } else {
        document.body.removeEventListener('click', this.closeMenu)
      }
    },
    doc: { // 监听文章内容变化
      deep: true,
      handler () {
        if (this.flag) {
          // 设置状态未保存
          this.saveStatus = false
          // 自动保存
          this.debouncedAutoSave()
        } else {
          this.flag = true
        }
      }
    }
  },
  created () {
    // 创建一个防反跳函数
    this.debouncedAutoSave = _.debounce(this.saveChapter, 1000)
  },
  mounted () {
    this.$store.dispatch('admin-posts/getDocs')
  },
  methods: {
    nodeRightClick (event, data, node, self) { // 节点点击事件
      this.clickRightItem = data
      this.clickRightNode = node
      // 设置右键菜单位置
      this.left = event.clientX
      this.top = event.clientY
      this.visible = true
    },
    closeMenu () { // 关闭右键菜单
      this.visible = false
    },
    addNewChapter (parent, choose = null) { // 新建章节信息
      this.$prompt('请输入标题', '提示', {
        closeOnClickModal: false, // 点击其他地方不关闭
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      }).then(({ value }) => {
        this.$store.dispatch('admin-posts/addDocs', { parent, title: value }).then((data) => {
          // 自动向节点添加数据
          const newChild = { id: data.id, label: value, children: [] }
          // 判断是否是右键选中的对象
          if (choose !== null) {
            // 判断当前选中的节点是否有子节点
            if (!choose.children) {
              // 手动给当前choose添加子节点（这里需要用到set，如果不用这个那么对象就无法刷新）
              this.$set(choose, 'children', [])
            }
            choose.children.push(newChild)
          } else {
            // 没有传值说明是根节点插入值
            this.$store.commit('admin-posts/docListPush', newChild)
          }
        }).catch((msg) => { this.$message.error(msg) })
      }).catch(() => {})
    },
    updateChapter (data) { // 更新章节信息
      return this.$store.dispatch('admin-posts/updatePost', data)
    },
    delChapter (data, node) { // 删除某个章节
      this.$confirm('删除该文档(该文档下所有子文档都会消失)?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.updateChapter({ id: data.id, data: { delete: 'true' } }).then(() => {
          // 我们手动从文档列表里删除这个文档
          // 获取当前文档的父节点node信息
          const parent = node.parent
          // ||说明: 只要第一个值的布尔值为false，那么永远返回第二个值。逻辑或属于短路操作，第一个值为true时，不再操作第二个值，且返回第一个值。
          // 通过简单测试可以看到如果当前节点为根节点的话，那么就没有children。data里面就是一个数组，如果不是根节点，那么就data里面就是children
          // （这个children里面就是和自己同级的一些节点）
          const children = parent.data.children || parent.data
          // findIndex方法是数组中满足提供的测试函数的第一个元素的索引。若没有找到对应元素 （我们这里目的是为了找到文档id对应的子节在数组中的位置）
          const index = children.findIndex(d => d.id === data.id)
          // 从子节点中删除该文档
          children.splice(index, 1)
          this.$message.success('删除成功')
        }).catch(() => { this.$message.error('删除失败') })
      }).catch(() => {})
    },
    changeChapterTitle (data) { // 修改章节标题
      this.$prompt('新标题', '提示', {
        closeOnClickModal: false, // 点击其他地方不关闭
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputValue: data.label
      }).then(({ value }) => {
        this.updateChapter({ id: data.id, data: { title: value } }).then(() => {
          data.label = value
        }).catch((msg) => { this.$message.error(msg) })
      }).catch(() => {})
    },
    nodeDrop (node, node2, position) { // 节点拖拽事件
      const id = node.data.id // 获取当前节点id
      let parent = node2.data.id // 获取拖动到的父节点的id
      // 判断父节点是否为根节点
      if (position !== 'inner') {
        parent = node2.parent.data.id
        if (parent === undefined) { parent = 0 }
      }
      // 发送请求更改父节点
      this.updateChapter({ id, data: { parent: parent.toString() } }).catch(() => {
        // 拖动失败，重新获取一下节点数据
        this.$store.dispatch('admin-posts/getDocs')
      })
    },
    nodeClick (data) { // 右键菜单点击事件
      this.id = data.id
      // 获取文档内容
      this.$store.dispatch('admin-posts/getDocsContent', data.id).then((data) => {
        this.$refs.edit.setContent(data.md ? data.md : data.content)
      }).catch((msg) => { this.$message.error(msg) })
    },
    saveChapter () { // 更新文档内容
      if (this.id === 0) {
        this.$message.error('请选择要编辑文档!')
        return
      }
      this.updateChapter({ id: this.id, data: { html: this.doc.content, md: this.doc.md } })
        .then(() => { this.saveStatus = true })
        .catch(() => { this.$message.error('保存失败') })
    }
  }
}
</script>

<style>
/*设置card的边距为0，同时设置分割线的间距*/
.option-card .el-card__header{
  padding: 0!important;
}
.option-card .el-divider--horizontal{
  margin: 10px 0!important;
}
</style>

<style scoped lang="scss">
.post-tree{
  max-height: 500px;
  overflow-y: auto;
}
.card-head{
  padding: 18px 20px;
  color: #ffffff;
  font-size: 16px;
}
.option-card{
  margin-bottom: 10px;
}
/*添加章节按钮*/
.add_chapter {
  margin-top: 10px;
  width: 100%;
  display: flex;
  justify-content: center;
}
/*多选框样式*/
.contextmenu {
  position: fixed;
  min-width: min-content;
  z-index: 1900;
  border: 1px solid #d4d4d5;
  line-height: 1.4285em;
  max-width: 150px;
  background: #fff;
  font-weight: 400;
  font-style: normal;
  color: rgba(0,0,0,.87);
  border-radius: .28571429rem;
  box-shadow: 0 2px 4px 0 rgba(34,36,38,.12), 0 2px 10px 0 rgba(34,36,38,.15);
  div {
    position: relative;
    vertical-align: middle;
    line-height: 1;
    -webkit-tap-highlight-color: transparent;
    padding: 10px 15px;
    color: rgba(0,0,0,.87);
    font-size: 14px;
    cursor: pointer;
    &:hover {
       background: #eee;
     }
  }
}
</style>
