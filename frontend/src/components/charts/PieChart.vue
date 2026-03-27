<template>
  <BaseChart :option="chartOption" :height="height" :loading="loading" />
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
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const chartOption = computed(() => {
  // 金融行业专业配色
  const colors = [
    '#2563eb', // 主蓝色
    '#10b981', // 绿色
    '#f59e0b', // 橙色
    '#8b5cf6', // 紫色
    '#ec4899', // 粉色
    '#6366f1'  // 靛蓝色
  ]
  
  return {
    title: {
      text: props.title,
      left: 'center',
      textStyle: {
        color: 'var(--text-primary)',
        fontSize: 16,
        fontWeight: 600
      }
    },
    tooltip: {
      trigger: 'item',
      backgroundColor: 'var(--bg-primary)',
      borderColor: 'var(--border-light)',
      borderWidth: 1,
      textStyle: {
        color: 'var(--text-primary)'
      },
      formatter: function(params: any) {
        return `${params.seriesName}<br/>${params.name}: ${params.value} (${params.percent.toFixed(1)}%)`
      }
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      data: props.data.map(item => item.name),
      textStyle: {
        color: 'var(--text-secondary)',
        fontSize: 12
      },
      itemWidth: 12,
      itemHeight: 12,
      itemGap: 12
    },
    series: [
      {
        name: props.title,
        type: 'pie',
        radius: ['40%', '70%'],
        center: ['60%', '55%'],
        data: props.data,
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 8,
          borderColor: 'var(--bg-primary)',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 16,
            fontWeight: 'bold',
            color: 'var(--text-primary)'
          },
          itemStyle: {
            shadowBlur: 15,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0, 0, 0, 0.3)'
          }
        },
        labelLine: {
          show: false
        },
        color: colors
      }
    ]
  }
})
</script>