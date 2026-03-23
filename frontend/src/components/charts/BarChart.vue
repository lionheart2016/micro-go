<template>
  <BaseChart :option="chartOption" :height="height" />
</template>

<script setup lang="ts">
import { computed, defineProps } from 'vue'
import BaseChart from '../base/BaseChart.vue'

const props = defineProps({
  title: {
    type: String,
    default: ''
  },
  xAxisData: {
    type: Array as () => string[],
    required: true
  },
  series: {
    type: Array as () => Array<{
      name: string
      data: number[]
      color?: string
    }>,
    required: true
  },
  height: {
    type: String,
    default: '400px'
  }
})

const chartOption = computed(() => {
  return {
    title: {
      text: props.title,
      left: 'center'
    },
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: props.series.map(item => item.name),
      bottom: 0
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '15%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: props.xAxisData
    },
    yAxis: {
      type: 'value'
    },
    series: props.series.map(item => ({
      name: item.name,
      type: 'bar',
      data: item.data,
      itemStyle: {
        color: item.color
      }
    }))
  }
})
</script>