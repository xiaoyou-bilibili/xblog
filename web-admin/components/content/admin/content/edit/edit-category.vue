<template>
  <div>
    <el-cascader
      :value="choose"
      :options="category"
      :props="props"
      @change="change"
    />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'EditCategory',
  props: {
    id: { type: Array, default: () => ([]) }
  },
  data () {
    return {
      props: { multiple: true }
    }
  },
  computed: {
    // 把getter混入到computed对象中,我们可以直接调用vex里面的内容即可
    ...mapGetters('admin-posts', ['category']),
    choose () {
      // 手动获取目录
      const choose = []
      // 遍历目录id
      this.id.map((item) => {
        // 遍历分类
        this.category.map((item2) => {
          // 遍历该分类下的子节点，获取父节点
          item2.children.filter((value, index1, array) => { if (value.value === item) { choose.push([item2.value, value.value]) } })
        })
      })
      return choose
    }
  },
  mounted () {
    // 获取文章分类
    this.$store.dispatch('admin-posts/getCategory')
  },
  methods: {
    // 当选择的内容改变的时候，我们自动去修改prop属性，使它可以正常的显示分类id
    change (choose) { this.$emit('update:id', choose.map(item => (item[1]))) }
  }
}
</script>

<style scoped>

</style>
