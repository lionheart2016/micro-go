<template>
  <div class="dashboard">
    <div class="dashboard-header">
      <h2>关键指标概览</h2>
      <el-button type="primary" @click="loadData">刷新数据</el-button>
    </div>
    
    <!-- 核心指标卡片 -->
    <div class="kpi-grid">
      <div class="kpi-item">
        <KpiCard 
          title="累计注册人数" 
          :value="coreIndicators.registeredUsers.value" 
          :change="coreIndicators.registeredUsers.change" 
          :changeRate="coreIndicators.registeredUsers.changeRate" 
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="开户人数" 
          :value="coreIndicators.开户人数.value" 
          :change="coreIndicators.开户人数.change" 
          :changeRate="coreIndicators.开户人数.changeRate" 
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="近1月活跃人数" 
          :value="coreIndicators.activeUsers.value" 
          :change="coreIndicators.activeUsers.change" 
          :changeRate="coreIndicators.activeUsers.changeRate" 
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="入金人数" 
          :value="coreIndicators.depositUsers.value" 
          :change="coreIndicators.depositUsers.change" 
          :changeRate="coreIndicators.depositUsers.changeRate" 
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="入金金额" 
          :value="coreIndicators.depositAmount.value" 
          :change="coreIndicators.depositAmount.change" 
          :changeRate="coreIndicators.depositAmount.changeRate" 
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="股票交易人数" 
          :value="coreIndicators.stockTradeUsers.value" 
          :change="coreIndicators.stockTradeUsers.change" 
          :changeRate="coreIndicators.stockTradeUsers.changeRate" 
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="股票交易金额" 
          :value="coreIndicators.stockTradeAmount.value" 
          :change="coreIndicators.stockTradeAmount.change" 
          :changeRate="coreIndicators.stockTradeAmount.changeRate" 
        />
      </div>
      <div class="kpi-item">
        <KpiCard 
          title="基金交易人数" 
          :value="coreIndicators.fundTradeUsers.value" 
          :change="coreIndicators.fundTradeUsers.change" 
          :changeRate="coreIndicators.fundTradeUsers.changeRate" 
        />
      </div>
    </div>

    <!-- 业绩表现 -->
    <div class="performance-section">
      <div class="performance-charts">
        <div class="chart-item">
          <BaseCard title="收入结构">
            <PieChart 
              title="收入结构" 
              :data="performance.revenueStructure" 
              height="300px" 
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
              height="300px" 
            />
          </BaseCard>
        </div>
        <div class="chart-item">
          <BaseCard title="入金趋势">
            <LineChart 
              title="入金趋势" 
              :xAxisData="performance.depositTrend.map(item => item.month)" 
              :series="[
                { name: '入金金额', data: performance.depositTrend.map(item => item.value) }
              ]" 
              height="300px" 
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

onMounted(() => {
  console.log('组件挂载，开始加载数据')
  loadData()
})
</script>

<style scoped>
.dashboard {
  padding: 20px 0;
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.dashboard-header h2 {
  margin: 0;
  color: #303133;
}

.kpi-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.kpi-item {
  height: 160px;
}

.performance-section {
  margin-top: 30px;
}

.performance-charts {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 20px;
}

.chart-item {
  width: 100%;
}

@media (max-width: 768px) {
  .kpi-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .performance-charts {
    grid-template-columns: 1fr;
  }
}
</style>