<!--主页面相关的布局文件-->
<template>
  <div id="app">
    <div id="page-main-content" style="position: relative">
      <nuxt />
    </div>
  </div>
</template>

<script>
import $ from 'jquery'

export default {
  mounted () {
    this.$nextTick(() => {
      // 底部自适应
      const height = $(window).height() - $('#navigation').height() - $('#footers').height() - 10
      $('#page-main-content').css('min-height', height + 'px')
    })
  }
}
</script>
<style>
/*图片地址为空时自动隐藏img*/
img[src=""],img:not([src]){
  opacity:0;
}
</style>

<style scoped>

</style>
