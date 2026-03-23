<template>
  <div class="finance">
    <h2>融资主题</h2>
    
    <!-- 融资客户分析 -->
    <BaseCard title="融资客户分析">
      <div v-if="loading.customer" class="loading">加载中...</div>
      <div v-else-if="error.customer" class="error">{{ error.customer }}</div>
      <div v-else-if="!data.customer?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="customer-stats">
          <div class="stat-card">
            <div class="stat-value">{{ data.customer.data.totalUsers.toLocaleString() }}</div>
            <div class="stat-label">总融资用户数</div>
          </div>
        </div>
        <div class="customer-charts">
          <div class="chart-item">
            <h3>用户分布</h3>
            <PieChart 
              :data="userDistributionChartData" 
              title="融资用户分布"
            />
          </div>
          <div class="chart-item">
            <h3>融资目的</h3>
            <PieChart 
              :data="financingPurposeChartData" 
              title="融资目的分布"
            />
          </div>
        </div>
        <div class="financing-size">
          <h3>融资规模</h3>
          <div class="size-stats">
            <div class="size-item">
              <span class="size-label">平均规模：</span>
              <span class="size-value">{{ data.customer.data.financingSize.avgSize.toLocaleString() }}</span>
            </div>
            <div class="size-item">
              <span class="size-label">中位数：</span>
              <span class="size-value">{{ data.customer.data.financingSize.medianSize.toLocaleString() }}</span>
            </div>
            <div class="size-item">
              <span class="size-label">25%分位：</span>
              <span class="size-value">{{ data.customer.data.financingSize.p25Size.toLocaleString() }}</span>
            </div>
            <div class="size-item">
              <span class="size-label">75%分位：</span>
              <span class="size-value">{{ data.customer.data.financingSize.p75Size.toLocaleString() }}</span>
            </div>
            <div class="size-item">
              <span class="size-label">最大规模：</span>
              <span class="size-value">{{ data.customer.data.financingSize.maxSize.toLocaleString() }}</span>
            </div>
          </div>
        </div>
        <div class="risk-assessment">
          <h3>风险评估</h3>
          <table>
            <thead>
              <tr>
                <th>类型</th>
                <th>违约率</th>
                <th>逾期率</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in data.customer.data.riskAssessment" :key="item.type">
                <td>{{ item.type }}</td>
                <td>{{ item.defaultRate }}%</td>
                <td>{{ item.overdueRate }}%</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </BaseCard>

    <!-- 融资股票分析 -->
    <BaseCard title="融资股票分析" class="mt-4">
      <div v-if="loading.stock" class="loading">加载中...</div>
      <div v-else-if="error.stock" class="error">{{ error.stock }}</div>
      <div v-else-if="!data.stock?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="stock-charts">
          <div class="chart-item">
            <h3>热门融资股票</h3>
            <table>
              <thead>
                <tr>
                  <th>股票</th>
                  <th>融资余额</th>
                  <th>市值占比</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in data.stock.data.hotStocks" :key="item.stock">
                  <td>{{ item.stock }}</td>
                  <td>{{ item.financingBalance.toLocaleString() }}</td>
                  <td>{{ item.marketCapRatio }}%</td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="chart-item">
            <h3>融资趋势</h3>
            <LineChart 
              :x-axis-data="financingTrendXAxisData" 
              :series="financingTrendSeriesData" 
              title="融资买入卖出趋势"
            />
          </div>
        </div>
        <div class="margin-ratio">
          <h3>保证金比例</h3>
          <table>
            <thead>
              <tr>
                <th>股票</th>
                <th>保证金比例</th>
                <th>价格波动率</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in data.stock.data.marginRatio" :key="item.stock">
                <td>{{ item.stock }}</td>
                <td>{{ item.ratio }}%</td>
                <td>{{ item.priceVolatility }}%</td>
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
import LineChart from '../../components/charts/LineChart.vue'
import { getFinanceCustomer, getFinanceStock } from '../../api/dashboard'

// 状态管理
const loading = ref({
  customer: false,
  stock: false
})
const error = ref({
  customer: '',
  stock: ''
})
const data = ref({
  customer: null,
  stock: null
})

// 计算属性 - 用户分布图表数据
const userDistributionChartData = computed(() => {
  if (!data.value.customer?.data?.userDistribution) return []
  return data.value.customer.data.userDistribution.map(item => ({
    name: item.type,
    value: item.count
  }))
})

// 计算属性 - 融资目的图表数据
const financingPurposeChartData = computed(() => {
  if (!data.value.customer?.data?.financingPurpose) return []
  return data.value.customer.data.financingPurpose.map(item => ({
    name: item.purpose,
    value: item.percentage
  }))
})

// 计算属性 - 融资趋势图表数据
const financingTrendXAxisData = computed(() => {
  if (!data.value.stock?.data?.financingTrend) return []
  return data.value.stock.data.financingTrend.map(item => item.date)
})

const financingTrendSeriesData = computed(() => {
  if (!data.value.stock?.data?.financingTrend) return []
  return [
    {
      name: '买入金额',
      data: data.value.stock.data.financingTrend.map(item => item.buyAmount),
      color: '#67c23a'
    },
    {
      name: '卖出金额',
      data: data.value.stock.data.financingTrend.map(item => item.sellAmount),
      color: '#f56c6c'
    }
  ]
})

// 获取融资客户数据
const fetchFinanceCustomer = async () => {
  loading.value.customer = true
  error.value.customer = ''
  try {
    const response = await getFinanceCustomer()
    data.value.customer = response
  } catch (err: any) {
    error.value.customer = err.message || '获取融资客户数据失败'
  } finally {
    loading.value.customer = false
  }
}

// 获取融资股票数据
const fetchFinanceStock = async () => {
  loading.value.stock = true
  error.value.stock = ''
  try {
    const response = await getFinanceStock()
    data.value.stock = response
  } catch (err: any) {
    error.value.stock = err.message || '获取融资股票数据失败'
  } finally {
    loading.value.stock = false
  }
}

// 组件挂载时获取数据
onMounted(() => {
  fetchFinanceCustomer()
  fetchFinanceStock()
})
</script>

<style scoped>
.finance {
  padding: 20px 0;
}

.finance h2 {
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

.customer-stats {
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

.customer-charts,
.stock-charts {
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

.financing-size {
  margin-top: 20px;
}

.financing-size h3 {
  margin-bottom: 10px;
  font-size: 16px;
  color: #303133;
}

.size-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 15px;
  padding: 20px;
  background-color: #f5f7fa;
  border-radius: 8px;
}

.size-item {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.size-label {
  font-size: 14px;
  color: #606266;
}

.size-value {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.risk-assessment,
.margin-ratio {
  margin-top: 20px;
  overflow-x: auto;
}

.risk-assessment h3,
.margin-ratio h3 {
  margin-bottom: 10px;
  font-size: 16px;
  color: #303133;
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
  .customer-charts,
  .stock-charts {
    grid-template-columns: 1fr;
  }
  
  .customer-stats {
    flex-direction: column;
  }
  
  .size-stats {
    grid-template-columns: 1fr;
  }
}
</style>