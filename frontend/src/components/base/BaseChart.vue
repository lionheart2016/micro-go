<template>
  <div ref="chartRef" class="base-chart" :style="{ width, height }"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch, defineProps } from 'vue'
import * as echarts from 'echarts'

const props = defineProps({
  option: {
    type: Object,
    required: true
  },
  width: {
    type: String,
    default: '100%'
  },
  height: {
    type: String,
    default: '400px'
  }
})

const chartRef = ref<HTMLElement | null>(null)
let chart: echarts.ECharts | null = null

const initChart = () => {
  if (chartRef.value) {
    chart = echarts.init(chartRef.value)
    updateChart()
  }
}

const updateChart = () => {
  if (chart) {
    chart.setOption(props.option)
  }
}

const handleResize = () => {
  chart?.resize()
}

onMounted(() => {
  initChart()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  chart?.dispose()
  window.removeEventListener('resize', handleResize)
})

watch(() => props.option, updateChart, { deep: true })
</script>

<style scoped>
.base-chart {
  width: 100%;
  height: 400px;
}
</style>