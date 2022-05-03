// 这个地方专门用来表格的各种操作
export default {
  data () {
    return {
      currentSelect: [],
      funcName: ''
    }
  },
  methods: {
    // 单个项目操作操作
    itemOption (id, option) {
      // 执行操作
      this.$store.dispatch(this.funcName, { id, option }).then((res) => {
        this.$message.success('操作成功')
        // 这里判断当前页面是日记还是文章
        // 重新获取文章
        this.getData(this.now)
      }).catch(msg => this.$message.error(msg))
    },
    // 多个项目操作
    itemsOption (option) {
      // 先把表格选中的id获取一下
      const ids = []
      this.currentSelect.map((value, index, array) => ids.push(value.id))
      // 执行操作
      this.itemOption(ids.join(','), option)
    },
    change (row) { // 表格多选框改变的事件，就相当于把当前选中的内容保存起来
      this.currentSelect = row
    }
  }
}
