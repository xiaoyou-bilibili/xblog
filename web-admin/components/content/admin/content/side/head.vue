<template>
  <div>
    <vxe-grid
      ref="grid"
      border
      row-key
      class="sortable-tree-demo"
      :scroll-y="{enabled: false}"
      :columns="tableColumn"
      :data="tableTreeData"
      :tree-config="{children: 'children'}"
      :toolbar-config="tableToolbar"
      :edit-config="editConfig"
      @toolbar-button-click="toolbarClick"
    >
      <!-- 操作模板 -->
      <template v-slot:option="{ row }">
        <vxe-button icon="fa fa-plus" title="添加子节点" circle @click="addChild(row)" />
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
        <!--删除按钮 onConfirm 点击确认按钮激发-->
        <el-popconfirm title="确定删除这个链接？" @onConfirm="deleteNode(row)">
          <vxe-button slot="reference" icon="fa fa-trash" title="删除" circle />
        </el-popconfirm>
      </template>
    </vxe-grid>
  </div>
</template>

<script>
import Sortable from 'sortablejs'
import XEUtils from 'xe-utils'
import { mapGetters } from 'vuex'
export default {
  data () {
    return {
      showHelpTip2: false,
      // 工具栏设置
      tableToolbar: {
        buttons: [
          { code: 'add', name: '添加菜单', icon: 'fa  fa-plus', status: 'success' },
          { code: 'save', name: '保存菜单', icon: 'fa  fa-save', status: 'primary' }
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
      // 行数据
      tableColumn: [
        {
          width: 60,
          slots: {
            default: () => {
              return [
                <span class="drag-btn"><i class="vxe-icon--menu"></i></span>
              ]
            },
            header: () => {
              return [
                <vxe-tooltip v-model={ this.showHelpTip2 } content="按住后可以上下拖动排序！" enterable>
                  <i class="vxe-icon--question" onClick={ () => { this.showHelpTip2 = !this.showHelpTip2 } }></i>
                </vxe-tooltip>
              ]
            }
          }
        },
        { field: 'title', title: '标题', editRender: { name: 'input' }, treeNode: true },
        { field: 'link', title: '地址', editRender: { name: 'input' } },
        { title: '操作', slots: { default: 'option' } }
      ]
    }
  },
  computed: {
    // 把getter混入到computed对象中,我们可以直接调用vex里面的内容即可
    ...mapGetters('admin-settings', ['tableTreeData'])
  },
  created () {
    this.treeDrop()
  },
  beforeDestroy () {
    if (this.sortable2) {
      this.sortable2.destroy()
    }
  },
  methods: {
    // 删除子节点
    deleteNode (row) {
      // console.log(this.tableTreeData)
      this.deleteTableData(this.tableTreeData, row._XID)
    },
    // 工具栏按钮点击事件
    toolbarClick ({ code, button }) {
      // 根据我们的code来判断不同操作
      switch (code) {
        case 'add':
          // 头部插入数据
          this.tableTreeData.push({ title: '新菜单', link: '地址', children: [] })
          break
        case 'save':
          // 保存数据
          this.$store.dispatch('admin-settings/updateNavOption', this.tableTreeData).then(() => {
            this.$message.success('保存成功')
            this.$store.dispatch('admin-settings/getNavOption')
          })
          break
      }
    },
    // 点击编辑事件
    editRowEvent (row) {
      // 我们激活编辑框
      this.$refs.grid.setActiveRow(row)
    },
    // 点击保存事件
    saveRowEvent (row) {
      // 点击保存的时候自动触发代理的保存事件
      this.$refs.grid.clearActived()
    },
    // 添加子菜单
    addChild (row) {
      row.children.push({ title: '新菜单', link: '地址', children: [] })
    },
    // 节点拖动事件
    // 参考 https://xuliangzhan_admin.gitee.io/vxe-table/#/table/other/sortableRow
    treeDrop () {
      this.$nextTick(() => {
        const xTable = this.$refs.grid
        this.sortable2 = Sortable.create(xTable.$el.querySelector('.body--wrapper>.vxe-table--body tbody'), {
          handle: '.drag-btn',
          onEnd: ({ item, oldIndex }) => {
            const options = { children: 'children' }
            const targetTrElem = item
            const wrapperElem = targetTrElem.parentNode
            const prevTrElem = targetTrElem.previousElementSibling
            const tableTreeData = this.tableTreeData
            const selfRow = xTable.getRowNode(targetTrElem).item
            const selfNode = XEUtils.findTree(tableTreeData, row => row === selfRow, options)
            if (prevTrElem) {
              // 移动到节点
              const prevRow = xTable.getRowNode(prevTrElem).item
              const prevNode = XEUtils.findTree(tableTreeData, row => row === prevRow, options)
              if (XEUtils.findTree(selfRow[options.children], row => prevRow === row, options)) {
                // 错误的移动
                const oldTrElem = wrapperElem.children[oldIndex]
                wrapperElem.insertBefore(targetTrElem, oldTrElem)
                return this.$XModal.message({ content: '不允许自己给自己拖动！', status: 'error' })
              }
              const currRow = selfNode.items.splice(selfNode.index, 1)[0]
              if (xTable.isTreeExpandByRow(prevRow)) {
                // 移动到当前的子节点
                prevRow[options.children].splice(0, 0, currRow)
              } else {
                // 移动到相邻节点
                prevNode.items.splice(prevNode.index + (selfNode.index < prevNode.index ? 0 : 1), 0, currRow)
              }
            } else {
              // 移动到第一行
              const currRow = selfNode.items.splice(selfNode.index, 1)[0]
              tableTreeData.unshift(currRow)
            }
            // 如果变动了树层级，需要刷新数据
            this.tableTreeData = [...tableTreeData]
          }
        })
      })
    },
    // 递归删除节点
    deleteTableData (arr, id) {
      for (let i = 0; i < arr.length; i++) {
        if (arr[i]._XID === id) {
          arr.splice(i, 1)
          return
        }
        if (arr[i].children && arr[i].children.length) {
          this.deleteTableData(arr[i].children, id)
        }
      }
    }
  }
}
</script>

<style>
.sortable-tree-demo .drag-btn {
  cursor: move;
  font-size: 12px;
}
.sortable-tree-demo .vxe-body--row.sortable-ghost,
.sortable-tree-demo .vxe-body--row.sortable-chosen {
  background-color: #dfecfb;
}
</style>
