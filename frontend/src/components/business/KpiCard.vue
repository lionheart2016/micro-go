<template>
  <div class="kpi-card">
    <div class="kpi-header">
      <h4>{{ title }}</h4>
    </div>
    <div class="kpi-value">{{ formatValue(value) }}</div>
    <div class="kpi-change" :class="{ positive: changeRate > 0, negative: changeRate < 0 }">
      <span>{{ changeRate > 0 ? '+' : '' }}{{ changeRate }}%</span>
      <span class="change-desc">{{ change > 0 ? '上升' : '下降' }}{{ Math.abs(change) }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  value: {
    type: Number,
    required: true
  },
  change: {
    type: Number,
    required: true
  },
  changeRate: {
    type: Number,
    required: true
  }
})

const formatValue = (value: number): string => {
  if (value >= 100000000) {
    return (value / 100000000).toFixed(2) + '亿'
  } else if (value >= 10000) {
    return (value / 10000).toFixed(2) + '万'
  }
  return value.toString()
}
</script>

<style scoped>
.kpi-card {
  background-color: white;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  text-align: center;
  height: 100%;
}

.kpi-header {
  margin-bottom: 12px;
}

.kpi-header h4 {
  margin: 0;
  font-size: 14px;
  color: #606266;
  font-weight: normal;
}

.kpi-value {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 8px;
}

.kpi-change {
  font-size: 12px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.kpi-change.positive {
  color: #67c23a;
}

.kpi-change.negative {
  color: #f56c6c;
}

.change-desc {
  margin-top: 2px;
}
</style>