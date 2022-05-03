<template>
  <div class="navi-bar">
    <div class="open-icon" @click="handleClick">
      <font-awesome-icon v-if="sideStatus" icon="indent" />
      <font-awesome-icon v-else icon="outdent" />
    </div>
    <div class="right-side">
      <div class="version-info">
        <span class="version-item" title="程序版本"><font-awesome-icon icon="server" /> {{ settingAdmin.version }}</span>
      </div>
      <div class="full-icon open-icon" @click="fullScreen">
        <font-awesome-icon v-if="!full" icon="expand-arrows-alt" />
        <font-awesome-icon v-else icon="compress-arrows-alt" />
      </div>
      <div class="person-info" @click="visible=!visible">
        <img :src="userInfo.avatar" alt class="person-avatar">
        <span class="person-name">{{ userInfo.nickname }}</span>
        <font-awesome-icon size="sm" icon="sort-down" />
        <div v-show="visible" class="contextmenu">
          <div @click="$router.push('/admin/person/info')">
            <font-awesome-icon icon="user-cog" />
            个人信息
          </div>
          <div @click="showForget=true">
            <font-awesome-icon icon="lock" />
            修改密码
          </div>
          <div @click="loginOut">
            <font-awesome-icon icon="sign-out-alt" />
            退出登录
          </div>
        </div>
      </div>
    </div>
    <!--忘记密码的dialog-->
    <el-dialog
      title="修改密码"
      :visible.sync="showForget"
      width="350px"
    >
      <!-- 这里必须要写model,后面那个rule里面的prop必须和model对象一致，不然会无法使用-->
      <el-form ref="pass" label-width="80px" :rules="rules" :model="pass">
        <el-form-item prop="oldPass" label="原密码">
          <el-input v-model="pass.oldPass" type="password" />
        </el-form-item>
        <el-form-item prop="newPass" label="新密码">
          <el-input v-model="pass.newPass" type="password" />
        </el-form-item>
        <el-form-item prop="repeatPass" label="重复密码">
          <el-input v-model="pass.repeatPass" type="password" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="showForget = false">取 消</el-button>
        <el-button type="primary" @click="changePassword">修改密码</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'NaviBar',
  data () {
    return {
      visible: false, // 显示个人信息设置
      showForget: false, // 显示修改密码弹框
      pass: {
        oldPass: '', // 旧密码
        repeatPass: '', // 重复密码
        newPass: '' // 新密码
      },
      full: false, // 当前是否是全屏状态
      rules: { // 修改密码规则
        oldPass: [{
          required: true,
          message: '请填写密码'
        }],
        repeatPass: [{
          required: true,
          validator: (rule, value, callback) => {
            value !== this.pass.newPass ? callback(new Error('两次输入的密码不一致')) : callback()
          },
          trigger: 'blur'
        }],
        newPass: [{
          required: true,
          message: '请输入新密码'
        }]
      }
    }
  },
  computed: {
    ...mapState('admin', ['sideStatus']),
    ...mapState('user', ['userInfo']),
    ...mapState('setting', ['settingAdmin'])
  },
  mounted () {
    // 自动更新个人信息
    this.$store.dispatch('user/getUserInfo')
      .then((data) => { if (data.identity !== 1) { this.$store.commit('admin/switchNormal') } })
      .catch(() => { this.$message.error('你还有没有登录请先登录!'); window.location = '/' })
    this.authInfo()
  },
  methods: {
    // 侧边栏收缩
    handleClick () {
      this.$store.dispatch('admin/changeSideBar')
    },
    // 修改用户密码
    changePassword () {
      this.$refs.pass.validate((valid) => {
        if (valid) {
          // 发送请求修改密码
          this.$store.dispatch('user/updateInfo', { old_password: this.pass.oldPass, new_password: this.pass.newPass })
            .then(() => {
              this.showForget = false
              this.$message.success('修改密码成功')
              // 清除cookie信息
              this.$store.commit('user/clearAuth')
              // 跳转到登录路由
              window.location.href = '/'
            })
            .catch((msg) => { this.$message.error(msg) })
        } else {
          return false
        }
      })
    },
    // 退出登录
    loginOut () {
      // 清除cookie信息
      this.$store.commit('user/clearAuth')
      // 跳转到主页
      window.location.href = '/'
    },
    fullScreen () { // 设置网页全屏
      if (!document.fullscreenElement) {
        document.documentElement.requestFullscreen()
        this.full = true
      } else if (document.exitFullscreen) {
        document.exitFullscreen()
        this.full = false
      }
    }, // 提示授权失败信息
    showAuthError () {
      this.$prompt('请输入你的用户ID后再进行验证,如果你域名更换了或者没有购买请点击购买系统！', '域名验证失败', {
        confirmButtonText: '重新验证',
        cancelButtonText: '购买系统',
        showClose: false,
        closeOnClickModal: false,
        closeOnPressEscape: false,
        closeOnHashChange: false,
        distinguishCancelAndClose: true
      }).then(({ value }) => {
        this.$store.dispatch('setting/updateAuthUserID', value)
          .then(_ => this.$message.success('验证成功，欢迎使用XBlog博客系统！'))
          .catch(_ => this.$message.error('更新失败'))
          .finally(() => this.authInfo())
      }).catch(() => { window.location = 'https://xblog.xiaoyou66.com/' })
    },
    // 授权认证
    authInfo () {
      // 获取系统设置
      this.$store.dispatch('setting/getAdminSetting', {}).catch((msg) => {
        this.$message.error(msg)
      })
    }
  }
}
</script>

<style scoped lang="scss">
  /*版本信息*/
  .version-info {
    display: flex;
    flex-wrap: wrap;
    color: #666;
    //width: 147px;
    font-size: 15px;
    margin: 5px;
    align-items: center;
  }
  .version-info .version-item {
    justify-content: center;
    flex: 1;
  }
  /*个人头像点击*/
  .contextmenu {
    position: fixed;
    top: 48px;
    right: 3px;
    z-index: 1999;
    border: 1px solid #d4d4d5;
    line-height: 1.4285em;
    max-width: 160px;
    background: #fff;
    font-weight: 400;
    font-style: normal;
    color: rgba(0, 0, 0, 0.87);
    border-radius: 3px;
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
.navi-bar{
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  -webkit-box-shadow: 0 1px 4px rgba(0,21,41,0.08);
  box-shadow: 0 1px 4px rgba(0,21,41,0.08);
  display: flex;
  flex-wrap: wrap;
  align-items: center;
}
.open-icon{
  width: 20px;
  line-height: 50px;
  padding: 0 15px;
  margin: 0 10px 0 0;
  cursor: pointer;
}
.open-icon:hover{
  background: rgba(0,0,0,0.025);
}
.right-side{
  flex: 1;
  display: flex;
  justify-content: flex-end;
}
.person-avatar{
  width: 40px;
  border-radius: 50%;
}
.person-info{
  cursor: pointer;
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  padding: 0 10px 0 0;
}
.person-name{
  font-size: 12px;
  margin: 0 5px;
}
</style>
