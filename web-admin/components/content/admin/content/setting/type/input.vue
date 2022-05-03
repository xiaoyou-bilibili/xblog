<template>
  <div>
    <div class="tool-input-item">
      <h4>{{ title }}</h4><el-input v-model="data" @change="change" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'TypeInput',
  props: {
    title: { type: String, default: '' },
    field: { type: String, default: '' },
    model: { type: String, default: '' }
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
      this.$store.dispatch('admin-settings/updateOption', { key: this.field, type: 'string', value })
        .then(() => { this.$message.success('保存成功'); this.$emit('update:model', this.data) }) // 提示成功，并主动更新store里面的数据
        .catch(() => { this.$message.error('保存设置失败'); this.data = this.model })
    }
  }
}
</script>

<style scoped lang="scss">
/*侧边栏设置的相关设置*/
.tool-input-item{
  display: flex;
  flex-direction: column;
  .item-input{
    display: flex;
    flex-direction: row;
    margin: 5px 0 0 0;
  }
  .item-input .el-input{
    margin-right: 5px;
  }
  h4{
    font-size: 14px;
    font-weight: bold;
    margin: 6px 0;
  }
}
</style>
