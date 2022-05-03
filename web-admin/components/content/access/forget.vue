<!--忘记密码功能-->
<template>
  <a id="login-account-switch" href="javascript:void(0)" @click="forget">忘记密码</a>
</template>

<script>
export default {
  name: 'Forget',
  methods: {
    // 忘记密码
    forget () {
      this.$prompt('请输入邮箱地址，你会收到一封带有重置密码链接的电子邮件', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
        inputErrorMessage: '邮箱格式不正确'
      }).then(({ value }) => {
        // 重置密码
        this.$store.dispatch('user/forgetPassword', { email: value }).then(() => {
          this.$message.success('重置链接已发至你的邮箱')
        }).catch((msg) => {
          this.$message.error(msg)
        })
      }).catch(() => {})
    }
  }
}
</script>

<style scoped>
#login-account-switch{
  margin-left: 10px;
  float: left;
  font: 14px 'PingFangSC-Semibold';
  color: #333;
  cursor: pointer;
  text-decoration: none;
  outline: none;
}
</style>
