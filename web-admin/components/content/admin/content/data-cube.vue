<template>
  <div class="cube-data" :style="{background:color}">
    <div class="data-icon">
      <font-awesome-icon size="lg" :icon="icon" />
    </div>
    <div class="data-container">
      <div class="data-text-cube">
        <div class="data-text">
          {{ title }}
        </div>
        <div class="data-num">
          {{ total }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DataCube',
  props: {
    title: { type: String, default: '' },
    icon: { type: String, default: '' },
    color: { type: String, default: '' },
    total: { type: Number, default: 0 }
  }
}
</script>

<style scoped>
.cube-data{
  display: flex;
  height: 70px;
  margin: 10px;
  align-items: center;
}
.data-container{
  flex: 1;
  margin-right: 20px;
  display: flex;
  justify-content: flex-end;
}
.data-icon{
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 5px;
  width: 45px;
  height: 45px;
}
.data-text-cube{
  color: #ffffff;
}
.data-text{
  font-size: 12px;
}
.data-num{
  font-size: 24px
}
</style>
