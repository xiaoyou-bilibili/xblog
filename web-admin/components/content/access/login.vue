<template>
  <div id="login">
    <el-form ref="ruleForm" class="loginFrom" :model="loginData" :rules="rules">
      <el-form-item prop="userName">
        <el-input
          v-model="loginData.username"
          class="login-inputorbuttom"
          prefix-icon="el-icon-user"
          placeholder="登录名"
        />
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          v-model="loginData.password"
          class="login-inputorbuttom"
          prefix-icon="el-icon-lock"
          placeholder="密码"
          type="password"
        />
      </el-form-item>
      <el-form-item class="login-item">
        <el-checkbox
          v-model="loginData.remember"
          style="margin-top:-7px;float:left;margin-bottom:-19px"
          label="记住我的登录信息"
          name="type"
        />
        <el-button
          v-popover:popover
          class="login-inputorbuttom login-bottom"
          type="primary"
          :loading="loginIng"
          @click="loginButton"
        >
          登 录
        </el-button>
        <div class="login-options clearfix">
          <nuxt-link to="/" class="login-help">
            返回首页
          </nuxt-link>
          <forget />
          <nuxt-link to="/registered" class="register">
            立即注册
          </nuxt-link>
        </div>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import Forget from '@/components/content/access/forget'
export default {
  name: 'Login',
  components: { Forget },
  data () {
    return {
      loginIng: false, // 登录加载提示
      tips: '拖动左边滑块完成上方拼图',
      loginData: {
        username: '',
        password: '',
        remember: ''
      },
      rules: {
        user: [{
          required: true,
          message: '请填写用户名'
        }],
        password: [{
          required: true,
          message: '请填写密码'
        }]
      },
      visible: false // 弹窗开启关闭
    }
  },
  // 在显示界面前先判断用户是否登录
  beforeCreate () {
    // 判断用户是否登录
    this.$store.dispatch('user/userIsLogin').then(() => {
      this.$message.warning('你已登录,请勿重复登录')
      this.$router.push('/')
    }).catch(() => {}) // 这里必须catch一下，要不然会报错
  },
  methods: {
    // 登录按钮
    loginButton () {
      // 表单验证
      this.$refs.ruleForm.validate((valid) => {
        if (valid) {
          this.visible = true
          // 表单验证成功，直接登录
          this.login()
        }
      })
    },
    // 登录
    login () {
      this.loginIng = true
      // 执行登录请求
      this.$store.dispatch('user/userLogin', this.loginData).then(() => {
        // 这里说明登录成功
        this.$message.success('登录成功！正在跳转到主页...')
        setTimeout(() => {
          // this.$router.push('/')
          window.location.href = '/'
        }, 1000)
      }).catch((res) => {
        this.$message.error(res)
      }).finally(() => {
        this.loginIng = false
      })
    }
  }
}
</script>
<style lang="scss">
#login{
  .signup-form {
    margin-top: 13vh !important;
  }

  .slidingPictures {
    padding: 0;
    width: 300px;
    border-radius: 2px;
  }

  /*忘记密码*/
  .clearfix {
    margin-top: 16px;
    zoom: 1;
    height: 20px;
  }

  .login-options .login-help {
    float: left;
    color: #9b9b9b;
    width: 61px;
    margin-right: 107px;
    font: 14px 'PingFangSC-Semibold';
    cursor: pointer;
    text-decoration: none;
    outline: none;
  }

  #login-account-switch, .login-options .register {
    margin-left: 10px;
    float: left;
    font: 14px 'PingFangSC-Semibold';
    color: #333;
    cursor: pointer;
    text-decoration: none;
    outline: none;
  }

  a:hover {
    color: #00a1d6 !important;
  }
}
</style>
<style scoped lang="scss">
#login {
  .loginFrom {
    margin: 0 20px;
    .login-bottom {
      margin-top: 15px;
    }
    .login-bottom:hover {
      background: rgba(28, 136, 188, 0.5);
    }
    .login-bottom:active {
      background: rgba(228, 199, 200, 0.5);
    }
    .login-inputorbuttom {
      height: 40px;
      width: 100%;
      font-size: 14px;
    }
  }
}
</style>
