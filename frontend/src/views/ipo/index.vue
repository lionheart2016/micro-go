<template>
  <div class="ipo">
    <h2>IPO主题</h2>
    
    <!-- IPO认购情况 -->
    <BaseCard title="IPO认购情况">
      <div v-if="loading.subscription" class="loading">加载中...</div>
      <div v-else-if="error.subscription" class="error">{{ error.subscription }}</div>
      <div v-else-if="!data.subscription?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="subscription-stats">
          <div class="stat-card">
            <div class="stat-value">{{ data.subscription.data.totalUsers.toLocaleString() }}</div>
            <div class="stat-label">总认购用户数</div>
          </div>
        </div>
        <div class="subscription-charts">
          <div class="chart-item">
            <h3>用户分布</h3>
            <PieChart 
              :data="userDistributionChartData" 
              title="IPO认购用户分布"
            />
          </div>
          <div class="chart-item">
            <h3>认购金额</h3>
            <BarChart 
              :x-axis-data="subscriptionAmountXAxisData" 
              :series="subscriptionAmountSeriesData" 
              title="IPO认购金额对比"
            />
          </div>
        </div>
        <div class="subscription-multiples">
          <h3>认购倍数</h3>
          <BarChart 
            :x-axis-data="subscriptionMultiplesXAxisData" 
            :series="subscriptionMultiplesSeriesData" 
            title="IPO认购倍数"
          />
        </div>
      </div>
    </BaseCard>

    <!-- IPO分析 -->
    <BaseCard title="IPO分析" class="mt-4">
      <div v-if="loading.analysis" class="loading">加载中...</div>
      <div v-else-if="error.analysis" class="error">{{ error.analysis }}</div>
      <div v-else-if="!data.analysis?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="analysis-charts">
          <div class="chart-item">
            <h3>项目表现</h3>
            <BarChart 
              :x-axis-data="projectPerformanceXAxisData" 
              :series="projectPerformanceSeriesData" 
              title="IPO项目表现"
            />
          </div>
          <div class="chart-item">
            <h3>影响因素</h3>
            <div class="influence-table">
              <table>
                <thead>
                  <tr>
                    <th>因素</th>
                    <th>相关系数</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in data.analysis.data.influenceFactors" :key="item.factor">
                    <td>{{ item.factor }}</td>
                    <td>{{ item.correlation }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
        <div class="competitor-comparison">
          <h3>竞争对手比较</h3>
          <table>
            <thead>
              <tr>
                <th>竞争对手</th>
                <th>认购用户数</th>
                <th>认购金额</th>
                <th>认购倍数</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in data.analysis.data.competitorComparison" :key="item.competitor">
                <td>{{ item.competitor }}</td>
                <td>{{ item.subscriptionUsers.toLocaleString() }}</td>
                <td>{{ item.subscriptionAmount.toLocaleString() }}</td>
                <td>{{ item.multiple }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </BaseCard>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import BaseCard from '../../components/base/BaseCard.vue'
import PieChart from '../../components/charts/PieChart.vue'
import BarChart from '../../components/charts/BarChart.vue'
import { getIPOSubscription, getIPOAnaalysis } from '../../api/dashboard'

// 状态管理
const loading = ref({
  subscription: false,
  analysis: false
})
const error = ref({
  subscription: '',
  analysis: ''
})
const data = ref({
  subscription: null,
  analysis: null
})

// 计算属性 - 用户分布图表数据
const userDistributionChartData = computed(() => {
  if (!data.value.subscription?.data?.userDistribution) return []
  return data.value.subscription.data.userDistribution.map(item => ({
    name: item.type,
    value: item.count
  }))
})

// 计算属性 - 认购金额图表数据
const subscriptionAmountXAxisData = computed(() => {
  if (!data.value.subscription?.data?.subscriptionAmount) return []
  return data.value.subscription.data.subscriptionAmount.map(item => item.project)
})

const subscriptionAmountSeriesData = computed(() => {
  if (!data.value.subscription?.data?.subscriptionAmount) return []
  return [
    {
      name: 'PI用户认购',
      data: data.value.subscription.data.subscriptionAmount.map(item => item.piAmount),
      color: '#409eff'
    },
    {
      name: '普通用户认购',
      data: data.value.subscription.data.subscriptionAmount.map(item => item.regularAmount),
      color: '#67c23a'
    }
  ]
})

// 计算属性 - 认购倍数图表数据
const subscriptionMultiplesXAxisData = computed(() => {
  if (!data.value.subscription?.data?.subscriptionMultiples) return []
  return data.value.subscription.data.subscriptionMultiples.map(item => item.project)
})

const subscriptionMultiplesSeriesData = computed(() => {
  if (!data.value.subscription?.data?.subscriptionMultiples) return []
  return [{
    name: '认购倍数',
    data: data.value.subscription.data.subscriptionMultiples.map(item => item.multiple),
    color: '#e6a23c'
  }]
})

// 计算属性 - 项目表现图表数据
const projectPerformanceXAxisData = computed(() => {
  if (!data.value.analysis?.data?.projectPerformance) return []
  return data.value.analysis.data.projectPerformance.map(item => item.project)
})

const projectPerformanceSeriesData = computed(() => {
  if (!data.value.analysis?.data?.projectPerformance) return []
  return [
    {
      name: '价格变动',
      data: data.value.analysis.data.projectPerformance.map(item => item.priceChange),
      color: '#409eff'
    },
    {
      name: '平均收益',
      data: data.value.analysis.data.projectPerformance.map(item => item.avgReturn),
      color: '#67c23a'
    }
  ]
})

// 获取IPO认购数据
const fetchIPOSubscription = async () => {
  loading.value.subscription = true
  error.value.subscription = ''
  try {
    const response = await getIPOSubscription()
    data.value.subscription = response
  } catch (err: any) {
    error.value.subscription = err.message || '获取IPO认购数据失败'
  } finally {
    loading.value.subscription = false
  }
}

// 获取IPO分析数据
const fetchIPOAnaalysis = async () => {
  loading.value.analysis = true
  error.value.analysis = ''
  try {
    const response = await getIPOAnaalysis()
    data.value.analysis = response
  } catch (err: any) {
    error.value.analysis = err.message || '获取IPO分析数据失败'
  } finally {
    loading.value.analysis = false
  }
}

// 组件挂载时获取数据
onMounted(() => {
  fetchIPOSubscription()
  fetchIPOAnaalysis()
})
</script>

<style scoped>
.ipo {
  padding: 20px 0;
}

.ipo h2 {
  margin-bottom: 20px;
  color: #303133;
}

.mt-4 {
  margin-top: 20px;
}

.loading, .error, .empty {
  padding: 40px 0;
  text-align: center;
  color: #909399;
}

.error {
  color: #f56c6c;
}

.subscription-stats {
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
}

.stat-card {
  flex: 1;
  padding: 20px;
  background-color: #f5f7fa;
  border-radius: 8px;
  text-align: center;
}

.stat-value {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: #606266;
}

.subscription-charts,
.analysis-charts {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-bottom: 20px;
}

.chart-item h3 {
  margin-bottom: 10px;
  font-size: 16px;
  color: #303133;
}

.subscription-multiples,
.competitor-comparison {
  margin-top: 20px;
}

.subscription-multiples h3,
.competitor-comparison h3 {
  margin-bottom: 10px;
  font-size: 16px;
  color: #303133;
}

.influence-table,
.competitor-comparison {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

th, td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ebeef5;
}

th {
  background-color: #f5f7fa;
  font-weight: 600;
  color: #303133;
}

@media (max-width: 768px) {
  .subscription-charts,
  .analysis-charts {
    grid-template-columns: 1fr;
  }
  
  .subscription-stats {
    flex-direction: column;
  }
}
</style>