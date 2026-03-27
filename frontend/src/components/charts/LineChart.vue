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
      trigger: 'axis',
      backgroundColor: 'var(--bg-primary)',
      borderColor: 'var(--border-light)',
      borderWidth: 1,
      textStyle: {
        color: 'var(--text-primary)'
      },
      formatter: function(params: any) {
        let result = params[0].name + '<br/>'
        params.forEach((item: any) => {
          result += `<span style="display:inline-block;margin-right:5px;border-radius:10px;width:10px;height:10px;background-color:${item.color};"></span>${item.seriesName}: ${item.value}<br/>`
        })
        return result
      }
    },
    legend: {
      data: props.series.map(item => item.name),
      bottom: 0,
      textStyle: {
        color: 'var(--text-secondary)'
      },
      itemWidth: 12,
      itemHeight: 12,
      itemGap: 20
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '15%',
      top: '15%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: props.xAxisData,
      axisLine: {
        lineStyle: {
          color: 'var(--border-light)'
        }
      },
      axisLabel: {
        color: 'var(--text-secondary)',
        fontSize: 12
      },
      axisTick: {
        show: false
      }
    },
    yAxis: {
      type: 'value',
      axisLine: {
        show: false
      },
      axisLabel: {
        color: 'var(--text-secondary)',
        fontSize: 12,
        formatter: function(value: number) {
          if (value >= 100000000) {
            return (value / 100000000).toFixed(1) + '亿'
          } else if (value >= 10000) {
            return (value / 10000).toFixed(1) + '万'
          }
          return value
        }
      },
      axisTick: {
        show: false
      },
      splitLine: {
        lineStyle: {
          color: 'var(--border-light)',
          type: 'dashed'
        }
      }
    },
    series: props.series.map((item, index) => ({
      name: item.name,
      type: 'line',
      data: item.data,
      smooth: true,
      symbol: 'circle',
      symbolSize: 6,
      lineStyle: {
        width: 3,
        color: item.color || colors[index % colors.length]
      },
      itemStyle: {
        color: item.color || colors[index % colors.length],
        borderColor: 'var(--bg-primary)',
        borderWidth: 2
      },
      areaStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [{
            offset: 0,
            color: (item.color || colors[index % colors.length]) + '33'
          }, {
            offset: 1,
            color: (item.color || colors[index % colors.length]) + '00'
          }]
        }
      }
    }))
  }
})
</script>