<template>
  <div class="tool-number-item">
    <h4>{{ title }}</h4>
    <el-switch
      v-model="data"
      active-color="#13ce66"
      inactive-color="#ff4949"
      @change="change"
    />
  </div>
</template>

<script>
export default {
  name: 'TypeSwitch',
  props: {
    title: { type: String, default: '' },
    field: { type: String, default: '' },
    model: { type: Boolean, default: false }
  },
  data () {
    return {
      // 因为vue默认不能直接修改设置，所以我们需要自己创建一个副本
      data: this.model
    }
  },
  methods: {
    change (value) {
      // 自动保存改设置
      this.$store.dispatch('admin-settings/updateOption', {
        key: this.field,
        type: 'bool',
        value: value.toString()
      })
        .then(() => {
          this.$message.success('保存成功')
          this.$emit('update:model', this.data)
        }) // 提示成功，并主动更新store里面的数据
        .catch(() => {
          this.$message.error('保存设置失败')
          this.data = this.model
        })
    }
  }
}
</script>

<style scoped lang="scss">
.tool-number-item{
  display: flex;
  flex-direction: column;
  h4{
    font-size: 14px;
    font-weight: bold;
    margin: 6px 0;
  }
}
</style>
