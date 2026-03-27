<template>
  <div class="dashboard">
    <div class="dashboard-header">
      <div class="header-left">
        <h2>关键指标概览</h2>
        <p class="header-subtitle">实时监控港美股经纪业务核心数据</p>
      </div>
      <div class="header-right">
        <el-button 
          type="primary" 
          @click="loadData" 
          :loading="loading"
          :icon="loading ? 'el-icon-loading' : 'el-icon-refresh'"
          class="refresh-button"
        >
          {{ loading ? '加载中...' : '刷新数据' }}
        </el-button>
        <el-date-picker
          v-model="dateRange"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
          @change="handleDateChange"
          class="date-picker"
        />
      </div>
    </div>
    
    <!-- 核心指标卡片 -->
    <div class="kpi-grid">
      <div class="kpi-item">
        <KpiCard 
          title="累计注册人数" 
          :value="coreIndicators.registeredUsers.value" 
          :change="coreIndicators.registeredUsers.change" 
          :changeRate="coreIndicators.registeredUsers.changeRate" 
          :loading="loading"
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="开户人数" 
          :value="coreIndicators.开户人数.value" 
          :change="coreIndicators.开户人数.change" 
          :changeRate="coreIndicators.开户人数.changeRate" 
          :loading="loading"
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="近1月活跃人数" 
          :value="coreIndicators.activeUsers.value" 
          :change="coreIndicators.activeUsers.change" 
          :changeRate="coreIndicators.activeUsers.changeRate" 
          :loading="loading"
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="入金人数" 
          :value="coreIndicators.depositUsers.value" 
          :change="coreIndicators.depositUsers.change" 
          :changeRate="coreIndicators.depositUsers.changeRate" 
          :loading="loading"
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="入金金额" 
          :value="coreIndicators.depositAmount.value" 
          :change="coreIndicators.depositAmount.change" 
          :changeRate="coreIndicators.depositAmount.changeRate" 
          :loading="loading"
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="股票交易人数" 
          :value="coreIndicators.stockTradeUsers.value" 
          :change="coreIndicators.stockTradeUsers.change" 
          :changeRate="coreIndicators.stockTradeUsers.changeRate" 
          :loading="loading"
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="股票交易金额" 
          :value="coreIndicators.stockTradeAmount.value" 
          :change="coreIndicators.stockTradeAmount.change" 
          :changeRate="coreIndicators.stockTradeAmount.changeRate" 
          :loading="loading"
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="基金交易人数" 
          :value="coreIndicators.fundTradeUsers.value" 
          :change="coreIndicators.fundTradeUsers.change" 
          :changeRate="coreIndicators.fundTradeUsers.changeRate" 
          :loading="loading"
        />
      </div>
    </div>

    <!-- 业绩表现 -->
    <div class="performance-section">
      <h3 class="section-title">业绩表现分析</h3>
      <div class="performance-charts">
        <div class="chart-item">
          <BaseCard title="收入结构">
            <PieChart 
              title="收入结构" 
              :data="performance.revenueStructure" 
              height="350px" 
              :loading="loading"
            />
          </BaseCard>
        </div>
        <div class="chart-item">
          <BaseCard title="收入趋势">
            <LineChart 
              title="收入趋势" 
              :xAxisData="performance.revenueTrend.map(item => item.month)" 
              :series="[
                { name: '股票收入', data: performance.revenueTrend.map(item => item.stock) },
                { name: '基金收入', data: performance.revenueTrend.map(item => item.fund) }
              ]" 
              height="350px" 
              :loading="loading"
            />
          </BaseCard>
        </div>
        <div class="chart-item chart-item-full">
          <BaseCard title="入金趋势">
            <LineChart 
              title="入金趋势" 
              :xAxisData="performance.depositTrend.map(item => item.month)" 
              :series="[
                { name: '入金金额', data: performance.depositTrend.map(item => item.value) }
              ]" 
              height="350px" 
              :loading="loading"
            />
          </BaseCard>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import BaseCard from '../../components/base/BaseCard.vue'
import KpiCard from '../../components/business/KpiCard.vue'
import LineChart from '../../components/charts/LineChart.vue'
import PieChart from '../../components/charts/PieChart.vue'
import { getCoreIndicators, getPerformance } from '../../api/dashboard'

// 核心指标数据
const coreIndicators = ref({
  registeredUsers: {
    value: 0,
    change: 0,
    changeRate: 0
  },
  开户人数: {
    value: 0,
    change: 0,
    changeRate: 0
  },
  activeUsers: {
    value: 0,
    change: 0,
    changeRate: 0
  },
  depositUsers: {
    value: 0,
    change: 0,
    changeRate: 0
  },
  depositAmount: {
    value: 0,
    change: 0,
    changeRate: 0
  },
  stockTradeUsers: {
    value: 0,
    change: 0,
    changeRate: 0
  },
  stockTradeAmount: {
    value: 0,
    change: 0,
    changeRate: 0
  },
  fundTradeUsers: {
    value: 0,
    change: 0,
    changeRate: 0
  },
  fundTradeAmount: {
    value: 0,
    change: 0,
    changeRate: 0
  }
})

// 业绩表现数据
const performance = ref({
  stockRevenue: 0,
  fundRevenue: 0,
  revenueStructure: [],
  revenueTrend: [],
  revenueChange: {
    yoy: 0,
    mom: 0
  },
  depositNetGrowth: 0,
  depositTrend: [],
  depositWithdrawalRatio: 0
})

// 加载状态
const loading = ref(false)

// 日期范围
const dateRange = ref<[string, string]>([
  new Date(new Date().setMonth(new Date().getMonth() - 1)).toISOString().split('T')[0],
  new Date().toISOString().split('T')[0]
])

// 加载数据
const loadData = async () => {
  console.log('开始加载数据')
  loading.value = true
  try {
    // 调用API获取真实数据
    console.log('调用API获取数据')
    const [coreIndicatorsRes, performanceRes] = await Promise.all([
      getCoreIndicators('month'),
      getPerformance('month')
    ])
    
    console.log('核心指标响应:', coreIndicatorsRes)
    console.log('业绩表现响应:', performanceRes)
    
    // 更新数据
    coreIndicators.value = {
      registeredUsers: coreIndicatorsRes.data.registeredUsers,
      开户人数: coreIndicatorsRes.data.accountCount,
      activeUsers: coreIndicatorsRes.data.activeUsers,
      depositUsers: coreIndicatorsRes.data.depositUsers,
      depositAmount: coreIndicatorsRes.data.depositAmount,
      stockTradeUsers: coreIndicatorsRes.data.stockTradeUsers,
      stockTradeAmount: coreIndicatorsRes.data.stockTradeAmount,
      fundTradeUsers: coreIndicatorsRes.data.fundTradeUsers,
      fundTradeAmount: coreIndicatorsRes.data.fundTradeAmount
    }
    performance.value = performanceRes.data
  } catch (error) {
    console.error('Failed to load data:', error)
  } finally {
    loading.value = false
    console.log('数据加载完成')
  }
}

// 日期变化处理
const handleDateChange = (val: [string, string] | null) => {
  if (val) {
    dateRange.value = val
    // 这里可以根据日期范围重新加载数据
    loadData()
  }
}

onMounted(() => {
  console.log('组件挂载，开始加载数据')
  loadData()
})
</script>

<style scoped>
.dashboard {
  padding: 0;
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 32px;
  padding-bottom: 24px;
  border-bottom: 1px solid var(--border-light);
}

.header-left h2 {
  margin: 0 0 8px 0;
  color: var(--text-primary);
  font-size: 24px;
  font-weight: 700;
  letter-spacing: 0.5px;
}

.header-subtitle {
  margin: 0;
  color: var(--text-secondary);
  font-size: 14px;
  line-height: 1.4;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.refresh-button {
  background-color: var(--primary) !important;
  border-color: var(--primary) !important;
  border-radius: 8px !important;
  padding: 8px 16px !important;
  font-size: 14px !important;
  font-weight: 500 !important;
  transition: all 0.3s ease !important;
}

.refresh-button:hover {
  background-color: var(--primary-light) !important;
  border-color: var(--primary-light) !important;
  transform: translateY(-1px) !important;
}

.date-picker {
  border-radius: 8px !important;
  border-color: var(--border-light) !important;
  font-size: 14px !important;
}

.kpi-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
  margin-bottom: 40px;
}

.kpi-item {
  height: 180px;
  transition: all 0.3s ease;
}

.kpi-item:hover {
  transform: translateY(-2px);
}

.performance-section {
  margin-top: 40px;
}

.section-title {
  color: var(--text-primary);
  font-size: 20px;
  font-weight: 600;
  margin: 0 0 24px 0;
  letter-spacing: 0.5px;
}

.performance-charts {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 24px;
}

.chart-item {
  width: 100%;
  transition: all 0.3s ease;
}

.chart-item:hover {
  transform: translateY(-2px);
}

.chart-item-full {
  grid-column: 1 / -1;
}

@media (max-width: 1200px) {
  .kpi-grid {
    grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
    gap: 20px;
  }
  
  .performance-charts {
    grid-template-columns: 1fr;
  }
  
  .chart-item-full {
    grid-column: 1;
  }
}

@media (max-width: 768px) {
  .dashboard-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }
  
  .header-right {
    width: 100%;
    justify-content: space-between;
  }
  
  .kpi-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 16px;
  }
  
  .kpi-item {
    height: 160px;
  }
  
  .header-left h2 {
    font-size: 20px;
  }
  
  .section-title {
    font-size: 18px;
  }
  
  .performance-charts {
    gap: 20px;
  }
}

@media (max-width: 480px) {
  .kpi-grid {
    grid-template-columns: 1fr;
  }
  
  .header-right {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
  
  .date-picker {
    width: 100%;
  }
}
</style>