<!--登录界面的样式-->
<template>
  <client-only>
    <div id="app">
      <el-container>
        <div id="background-slideshow" class="" />
        <el-main>
          <div class="signup-form ">
            <div class="signup-form__logo-box">
              <div class="signup-form__logo" :style="`background-image:url()`" />
              <div class="signup-form__catchphrase" />
            </div>
            <nuxt />
          </div>
        </el-main>
        <el-footer>
          <div id="footer">
            <div class="footer-item">
              <a href="/" target="_blank">©</a>
            </div>
          </div>
        </el-footer>
      </el-container>
    </div>
  </client-only>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
}
</script>

<style scoped>
#footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  width: 100%;
  height: 30px;
  background-color: #464a4d;
  text-align: left;
  color: hsla(0, 0%, 100%, .5);
  font-size: 12px;
}

#footer .footer-item, #footer .footer-item a {
  display: inline-block;
  color: hsla(0, 0%, 100%, .5);
  margin-top: 3px;
}

#footer a {
  text-decoration: none;
}

.user-info a:hover {
  color: #409EFF !important;
}

.footer-item a:hover {
  color: #ffffff !important;
}

/* 背景图片 */
#background-slideshow .slide {
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  z-index: -1;
  position: fixed;
}

#background-slideshow .wallpaper {
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position-x: 50%;
  background-position-y: 0;
  opacity: 1;
  -webkit-transition-property: opacity;
  transition-property: opacity;
  -webkit-transition-duration: 1s;
  transition-duration: 1s;
  -webkit-transition-timing-function: ease;
  transition-timing-function: ease;
}

#background-slideshow .user-info a:link, #background-slideshow .user-info a:visited {
  text-decoration: none;
  color: #fff;
  background-color: transparent;
}
/*底部背景图*/
/*登录注册按钮*/
.signup-form {
  margin-top: 20vh;
  margin-left: auto;
  margin-right: auto;
  text-align: center;
  background: #ffffff;
  width: 363px;
  padding: 10px 0;
  border-radius: 4px;
}

.signup-form .signup-form__logo-box {
  margin-bottom: 40px;
}

.signup-form .signup-form__logo {
  margin-bottom: 6px;
  display: inline-block;
  width: 180px;
  height: 94px;
  background-repeat: no-repeat;
  background-size: cover;
}

.signup-form__catchphrase {
  color: #757c80;
  font-size: 12px;
  font-weight: 700;
}

</style>
