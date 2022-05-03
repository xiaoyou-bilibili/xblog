<template>
  <div id="main" class="show">
    <router-link to="/admin">
      管理员
    </router-link>
    <router-link to="/login">
      登录
    </router-link>
  </div>
</template>

<script>

export default {
  layout: 'page'
}
</script>

<style scoped>
#main {
  /*width: 100%;*/
  min-height: 750px;
  /*display: flex;*/
  /*justify-content: center;*/
  /*margin-bottom: 78px;*/
  margin-top: 25px;
}

#bg {
  background-size: cover;
  background-position: right bottom;
  background-repeat: no-repeat;
  position: fixed;
  z-index: -1;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}
</style>
