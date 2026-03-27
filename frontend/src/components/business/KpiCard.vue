<template>
  <div class="kpi-card" :class="{ 'kpi-card-loading': loading }">
    <div class="kpi-header">
      <h4>{{ title }}</h4>
      <div class="kpi-icon" :class="getIconClass()"></div>
    </div>
    <div class="kpi-value">{{ formatValue(value) }}</div>
    <div class="kpi-change" :class="{ positive: changeRate > 0, negative: changeRate < 0, neutral: changeRate === 0 }">
      <span class="change-rate">{{ changeRate > 0 ? '+' : '' }}{{ changeRate }}%</span>
      <span class="change-desc">{{ change > 0 ? '上升' : change < 0 ? '下降' : '持平' }}{{ Math.abs(change) }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, computed } from 'vue'

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
  },
  loading: {
    type: Boolean,
    default: false
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

const getIconClass = () => {
  // 根据标题返回不同的图标类名
  const titleLower = props.title.toLowerCase()
  if (titleLower.includes('注册')) return 'icon-users'
  if (titleLower.includes('开户')) return 'icon-account'
  if (titleLower.includes('活跃')) return 'icon-active'
  if (titleLower.includes('入金')) return 'icon-deposit'
  if (titleLower.includes('交易')) return 'icon-trade'
  if (titleLower.includes('金额')) return 'icon-amount'
  return 'icon-default'
}
</script>

<style scoped>
.kpi-card {
  background-color: var(--bg-primary);
  border-radius: 12px;
  padding: 24px;
  box-shadow: var(--shadow-md);
  height: 100%;
  transition: all 0.3s ease;
  border: 1px solid var(--border-light);
  position: relative;
  overflow: hidden;
}

.kpi-card:hover {
  box-shadow: var(--shadow-lg);
  transform: translateY(-2px);
  border-color: var(--primary-light);
}

.kpi-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 4px;
  height: 100%;
  background: linear-gradient(180deg, var(--primary) 0%, var(--secondary) 100%);
  border-radius: 12px 0 0 12px;
}

.kpi-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.kpi-header h4 {
  margin: 0;
  font-size: 14px;
  color: var(--text-secondary);
  font-weight: 500;
  letter-spacing: 0.5px;
}

.kpi-icon {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  background-color: rgba(37, 99, 235, 0.1);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  color: var(--primary);
}

.kpi-value {
  font-size: 28px;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 12px;
  font-family: var(--mono);
  letter-spacing: -0.5px;
}

.kpi-change {
  font-size: 13px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.change-rate {
  font-weight: 600;
  font-family: var(--mono);
}

.change-desc {
  font-size: 12px;
  opacity: 0.8;
}

.kpi-change.positive {
  color: var(--success);
}

.kpi-change.negative {
  color: var(--danger);
}

.kpi-change.neutral {
  color: var(--text-tertiary);
}

.kpi-card-loading {
  opacity: 0.7;
  pointer-events: none;
}

.kpi-card-loading::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  animation: loading 1.5s infinite;
}

@keyframes loading {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

/* 图标样式 */
.icon-users::before { content: '👥'; }
.icon-account::before { content: '📋'; }
.icon-active::before { content: '⚡'; }
.icon-deposit::before { content: '💰'; }
.icon-trade::before { content: '📈'; }
.icon-amount::before { content: '💹'; }
.icon-default::before { content: '📊'; }

@media (max-width: 768px) {
  .kpi-card {
    padding: 20px;
  }
  
  .kpi-value {
    font-size: 24px;
  }
  
  .kpi-header h4 {
    font-size: 13px;
  }
}
</style>