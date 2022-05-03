<template>
  <div style="margin: 5px">
    <el-button size="medium" type="primary" @click="send">
      {{ title }}
    </el-button>
  </div>
</template>

<script>
export default {
  name: 'TypeRowRequest',
  props: {
    title: { type: String, default: '' },
    url: { type: String, default: '' }
  },
  methods: {
    send () {
      const loading = this.$loading({
        lock: true,
        text: '正在处理..',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      this.$store.dispatch('admin-settings/sendRowRequest', { url: this.url, type: 'get', data: {} })
        .then(() => this.$message.success('操作成功'))
        .catch(msg => this.$message.error(msg))
        .finally(() => loading.close())
    }
  }
}
</script>

<style scoped lang="scss">

</style>
