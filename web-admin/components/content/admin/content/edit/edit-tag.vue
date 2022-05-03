<template>
  <div class="tag-box">
    <el-tag
      v-for="(edit,index) in tags"
      :key="index"
      closable
      class="tags"
      :disable-transitions="false"
      @close="handleClose(edit)"
    >
      {{ edit }}
    </el-tag>
    <el-input
      v-if="inputVisible"
      ref="saveTagInput"
      v-model="inputValue"
      class="input-new-tag"
      size="small"
      @keyup.enter.native="handleInputConfirm"
      @blur="handleInputConfirm"
    />
    <el-button v-else class="button-new-tag" size="small" @click="showInput">
      添加
    </el-button>
  </div>
</template>

<script>
export default {
  name: 'EditTag',
  props: {
    // tag测试文章id477
    tags: { type: Array, default: () => ([]) }
  },
  data () {
    return {
      inputVisible: false,
      inputValue: ''
    }
  },
  methods: {
    // 删除某个标签
    handleClose (tag) {
      this.tags.splice(this.tags.indexOf(tag), 1)
    },
    // 显示输入框，用于输入标签信息
    showInput () {
      this.inputVisible = true
      // 设置焦点
      this.$nextTick(() => {
        this.$refs.saveTagInput.$refs.input.focus()
      })
    },
    // 输入标签
    handleInputConfirm () {
      // 获取输入框的值
      const inputValue = this.inputValue
      if (inputValue) {
        this.tags.push(inputValue)
      }
      // 隐藏输入框
      this.inputVisible = false
      this.inputValue = ''
    }
  }
}
</script>

<style scoped>
.tag-box{
  display: flex;
  flex-wrap: wrap;
}
span.tags {
  margin: 2px;
}
</style>
