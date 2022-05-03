<template>
  <div id="page-main-content" style="position: relative">
    <Nuxt />
  </div>
</template>

<style>
/*图片地址为空时自动隐藏img*/
img[src=""],img:not([src]){
  opacity:0;
}
/*清除输入框选中黑色*/
input{outline:none}
</style>
