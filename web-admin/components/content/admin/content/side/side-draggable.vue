<template>
  <draggable
    v-model="lists"
    class="list-group"
    tag="ul"
    :empty-insert-threshold="100"
    v-bind="dragOptions"
    @start="drag = true"
    @end="drag = false"
  >
    <transition-group type="transition" :name="!drag ? 'flip-list' : null">
      <div
        v-for="(item,index) in lists"
        :key="index+0"
        class="list-group-item"
      >
        {{ item.name }}
      </div>
    </transition-group>
  </draggable>
</template>

<script>
import draggable from 'vuedraggable'
export default {
  name: 'SideDraggable',
  components: { draggable },
  props: { list: { type: Array, default: () => [] } },
  data () {
    return {
      lists: this.list,
      drag: false
    }
  },
  computed: {
    dragOptions () {
      return {
        animation: 200,
        group: 'side',
        disabled: false,
        ghostClass: 'ghost'
      }
    }
  },
  watch: {
    list (data) {
      this.lists = data
    }
  }
}
</script>

<style scoped>
.list-group-item {
  position: relative;
  display: block;
  padding: 10px 10px;
  margin-bottom: -1px;
  background-color: #fff;
  border: 1px solid rgba(0,0,0,.125);
}
.list-group {
  min-height: 50px;
  padding: 6px;
  margin: 5px;
}
.flip-list-move {
  transition: transform 0.5s;
}
.no-move {
  transition: transform 0s;
}
.ghost {
  opacity: 0.5;
  background: #c8ebfb;
}
.list-group-item {
  cursor: move;
}
.list-group-item i {
  cursor: pointer;
}
</style>
