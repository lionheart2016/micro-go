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
  data: {
    type: Array as () => Array<{
      name: string
      value: number
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
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)'
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      data: props.data.map(item => item.name)
    },
    series: [
      {
        name: props.title,
        type: 'pie',
        radius: '50%',
        center: ['50%', '60%'],
        data: props.data,
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
})
</script>