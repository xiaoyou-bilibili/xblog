<template>
  <div class="sidebar-logo-container" :class="{'collapse':collapse}">
    <transition name="sidebarLogoFade">
      <a v-if="collapse" key="collapse" class="sidebar-logo-link" href="/">
        <img v-if="logo" alt :src="logo" class="sidebar-logo">
        <h1 v-else class="sidebar-title">
          {{ title }}
        </h1>
      </a>
      <a v-else key="expand" class="sidebar-logo-link" href="/">
        <img v-if="logo" alt :src="logo" class="sidebar-logo">
        <h1 class="sidebar-title">
          {{ title }}
        </h1>
      </a>
    </transition>
  </div>
</template>

<script>
import logo from '@/assets/image/admin/logo.png'
export default {
  name: 'SidebarLogo',
  props: {
    collapse: {
      type: Boolean,
      required: true
    }
  },
  data () {
    return {
      title: 'XBlog博客系统',
      logo
    }
  }
}
</script>

<style lang="scss" scoped>
.sidebarLogoFade-enter-active {
  transition: opacity 1.5s;
}

.sidebarLogoFade-enter,
.sidebarLogoFade-leave-to {
  opacity: 0;
}

.sidebar-logo-container {
  position: relative;
  border-right: solid 1px #e6e6e6;
  border-bottom: solid 1px #e6e6e6;
  width: 100%;
  height: 50px;
  line-height: 50px;
  background: #ffffff;
  text-align: center;
  overflow: hidden;

  & .sidebar-logo-link {
    height: 100%;
    width: 100%;

    & .sidebar-logo {
      width: 32px;
      height: 32px;
      vertical-align: middle;
      margin-right: 12px;
      border-radius: 50%;
    }

    & .sidebar-title {
      display: inline-block;
      margin: 0;
      color: #2b2f3a;
      font-weight: 600;
      line-height: 50px;
      font-size: 14px;
      font-family: Avenir, Helvetica Neue, Arial, Helvetica, sans-serif;
      vertical-align: middle;
    }
  }

  &.collapse {
    .sidebar-logo {
      margin-right: 0;
      border-radius: 50%;
    }
  }
}
</style>
