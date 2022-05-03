<template>
  <div>
    <!--    顶部数据显示-->
    <el-row>
      <el-col :xs="24" :sm="12" :md="6">
        <data-cube title="文章总数" color="#337ab7" icon="book-open" :total="total.post" />
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <data-cube title="用户总数" color="#f96868" icon="user" :total="total.user" />
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <data-cube title="浏览量" color="#15c377" icon="eye" :total="total.view" />
      </el-col>
      <el-col :xs="24" :sm="12" :md="6">
        <data-cube title="评论数" color="#926dde" icon="comments" :total="total.comment" />
      </el-col>
    </el-row>
    <!--    下方图表展示-->
    <el-row>
      <el-col :xs="24" :sm="24" :md="12">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文章类型分布</span>
          </div>
          <div>
            <div id="post-chart" style="width:100%;height:400px" />
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="24" :md="12">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文章相关数据分布</span>
          </div>
          <div>
            <div id="post-detail-chart" style="width:100%;height:400px" />
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import echarts from 'echarts'
import DataCube from '@/components/content/admin/content/data-cube'
import { mapGetters } from 'vuex'
import admin from '@/components/mixin/admin-seo'
export default {
  components: { DataCube },
  layout: 'admin',
  mixins: [admin],
  computed: {
    // 把getter混入到computed对象中,我们可以直接调用vex里面的内容即可
    ...mapGetters('admin-visual', ['total'])
  },
  mounted () {
    this.$store.dispatch('admin-visual/getTotal')
    // 设置文章分布数据
    this.$store.dispatch('admin-visual/getDistributed').then((data) => {
      const myChart = echarts.init(document.getElementById('post-chart'))
      const option = {
        title: {
          text: '文章类型分布',
          left: 'center'
        },
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        legend: {
          orient: 'vertical',
          left: 'left',
          data: ['文章', '日记', '文档']
        },
        series: [
          {
            name: '总数',
            type: 'pie',
            radius: '55%',
            center: ['50%', '60%'],
            data: [
              { value: data.post, name: '文章' },
              { value: data.diary, name: '日记' },
              { value: data.doc, name: '文档' }
            ],
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ]
      }
      myChart.setOption(option)
    })
    // 获取文章的详细分布数据
    this.$store.dispatch('admin-visual/getPostDetail').then((data) => {
      const myChart = echarts.init(document.getElementById('post-detail-chart'))
      const option = {
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'shadow',
            label: {
              show: true
            }
          }
        },
        toolbox: {
          show: true,
          feature: {
            mark: { show: true },
            dataView: { show: true, readOnly: false },
            magicType: { show: true, type: ['line', 'bar'] },
            restore: { show: true },
            saveAsImage: { show: true }
          }
        },
        calculable: true,
        legend: {
          data: ['浏览量', '点赞数', '评论数'],
          itemGap: 5
        },
        grid: {
          top: '12%',
          left: '1%',
          right: '10%',
          containLabel: true
        },
        xAxis: [
          {
            type: 'category',
            data: data.title
          }
        ],
        yAxis: [
          {
            type: 'value',
            name: '数目'
          }
        ],
        dataZoom: [
          {
            show: true,
            start: 94,
            end: 100
          },
          {
            type: 'inside',
            start: 94,
            end: 100
          },
          {
            show: true,
            yAxisIndex: 0,
            filterMode: 'empty',
            width: 30,
            height: '80%',
            showDataShadow: false,
            left: '93%'
          }
        ],
        series: [
          {
            name: '浏览量',
            type: 'bar',
            data: data.view
          },
          {
            name: '点赞数',
            type: 'bar',
            data: data.good
          },
          {
            name: '评论数',
            type: 'bar',
            data: data.comment
          }
        ]
      }
      myChart.setOption(option)
    })
  }
}
</script>

<style scoped>
.box-card{
  margin: 10px;
}
</style>
