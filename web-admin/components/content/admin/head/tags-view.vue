<template>
  <div id="tags-view-container" class="tags-view-container">
    <div class="tags-view-wrapper">
      <router-link
        v-for="(item,index) in tags"
        :key="index"
        ref="tag"
        tag="span"
        class="tags-view-item"
        :class="{'active': item.active}"
        :to="item.path"
        @contextmenu.prevent.native="openMenu(index,$event)"
      >
        {{ item.title }}
        <!--这里加prevent.stop是为了避免跳转路由-->
        <span v-if="item.close" class="el-icon-close" @click.prevent.stop="closeTag(index)" />
      </router-link>
    </div>
    <ul v-if="visible" :style="{left:left+'px',top:top+'px'}" class="contextmenu">
      <li @click="refresh">
        刷新
      </li>
      <li @click="closeTag(selectedTag)">
        关闭
      </li>
      <li @click="closeOther(selectedTag)">
        关闭其他
      </li>
      <li @click="closeAll">
        全部关闭
      </li>
    </ul>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'TagsView',
  data () {
    return {
      visible: false,
      top: 0,
      left: 0,
      selectedTag: 0
    }
  },
  computed: {
    ...mapState('admin', ['tags'])
  },
  watch: {
    $route () {
      this.changeTag()
    },
    visible (value) { // 监听右键菜单显示变化
      if (value) {
        document.body.addEventListener('click', this.closeMenu)
      } else {
        document.body.removeEventListener('click', this.closeMenu)
      }
    }
  },
  methods: {
    // 监听路由变化修改标签
    changeTag () {
      // 判断新加的标签是否为当前页
      const { path } = this.$route
      // 路径不为空则调用vuex方法来处理数据
      if (path) {
        this.$store.dispatch('admin/changeTag', path)
      }
    },
    // 关闭某一个标签页
    closeTag (index) {
      // 把当前标签页放入tags中
      this.$store.dispatch('admin/closeTag', index)
        .then(res => this.$router.push({ path: res }))
        .catch(() => {})
    },
    // 打开右键菜单
    openMenu (index, e) {
      this.left = e.clientX
      this.top = e.clientY
      this.visible = true
      this.selectedTag = index
    },
    // 关闭右键菜单
    closeMenu () {
      this.visible = false
    },
    // 刷新页面
    refresh () {
      window.location.reload()
    },
    // 关闭其他
    closeOther (index) {
      this.$store.dispatch('admin/closeOther', index)
      // 手动获取最后一个标签并跳转路由
      const path = this.tags[this.tags.length - 1].path
      if (this.$route.path !== path) {
        this.$router.push(path)
      }
    },
    // 关闭所有
    closeAll () {
      this.$store.dispatch('admin/closeAll')
      const path = this.tags[0].path
      if (this.$route.path !== path) {
        this.$router.push(path)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.tags-view-container {
  height: 34px;
  width: 100%;
  background: #fff;
  border-bottom: 1px solid #d8dce5;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .12), 0 0 3px 0 rgba(0, 0, 0, .04);
  .tags-view-wrapper {
    overflow: auto;
    height: 52px;
    white-space: nowrap;
    .tags-view-item {
      display: inline-block;
      position: relative;
      cursor: pointer;
      height: 26px;
      line-height: 26px;
      border: 1px solid #d8dce5;
      color: #495060;
      background: #fff;
      padding: 0 8px;
      font-size: 12px;
      margin-left: 5px;
      margin-top: 4px;
      &:first-of-type {
        margin-left: 15px;
      }
      &:last-of-type {
        margin-right: 15px;
      }
      &.active {
        background-color: #42b983;
        color: #fff;
        border-color: #42b983;
        &::before {
          content: '';
          background: #fff;
          display: inline-block;
          width: 8px;
          height: 8px;
          border-radius: 50%;
          position: relative;
          margin-right: 2px;
        }
      }
    }
  }
  .contextmenu {
    margin: 0;
    background: #fff;
    z-index: 3000;
    position: absolute;
    list-style-type: none;
    padding: 5px 0;
    border-radius: 4px;
    font-size: 12px;
    font-weight: 400;
    color: #333;
    box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, .3);
    li {
      margin: 0;
      padding: 7px 16px;
      cursor: pointer;
      &:hover {
        background: #eee;
      }
    }
  }
}
</style>
<style lang="scss">
//reset element css of el-icon-close
.tags-view-wrapper {
  .tags-view-item {
    .el-icon-close {
      width: 16px;
      height: 16px;
      vertical-align: 2px;
      border-radius: 50%;
      text-align: center;
      transition: all .3s cubic-bezier(.645, .045, .355, 1);
      transform-origin: 100% 50%;
      &:before {
        transform: scale(.6);
        display: inline-block;
        vertical-align: -3px;
      }
      &:hover {
        background-color: #b4bccc;
        color: #fff;
      }
    }
  }
}
</style>
