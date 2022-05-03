<template>
  <div>
    <el-menu-item v-if="item.child === undefined" :index="indexKey.toString()" @click="gotoSite(item.path, item.text)">
      <font-awesome-icon :icon="item.icon" />
      <span slot="title">{{ item.text }}</span>
    </el-menu-item>
    <el-submenu v-else :index="indexKey.toString()">
      <template slot="title">
        <font-awesome-icon :icon="item.icon" />
        <span slot="title">{{ item.text }}</span>
      </template>
      <el-menu-item-group>
        <el-menu-item v-for="(item2,index2) in item.child" :key="index2" :index="`${indexKey.toString()}-${index2.toString()}`" @click="gotoSite(`${item.path}/${item2.path}`, item2.text)">
          <font-awesome-icon :icon="item2.icon" />
          <span slot="title">{{ item2.text }}</span>
        </el-menu-item>
      </el-menu-item-group>
    </el-submenu>
  </div>
</template>

<script>
export default {
  name: 'SideMenu',
  props: { indexKey: { type: Number, default: 0 }, item: { type: Object, default: () => { return {} } } },
  methods: {
    gotoSite (path, name) {
      // 跳转路由
      this.$router.push(path)
      // 添加标签
      this.$store.dispatch('admin/addTag', { name, path })
    }
  }
}
</script>

<style scoped>

</style>
