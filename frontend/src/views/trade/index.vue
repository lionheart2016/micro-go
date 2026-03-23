<template>
  <div class="trade">
    <h2>交易分析</h2>
    
    <!-- 股票交易 -->
    <BaseCard title="股票交易">
      <div v-if="loading.stock" class="loading">加载中...</div>
      <div v-else-if="error.stock" class="error">{{ error.stock }}</div>
      <div v-else-if="!data.stock?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="filter">
          <label for="stockType">交易类型：</label>
          <select id="stockType" v-model="stockType" @change="fetchStockTrade">
            <option value="execution">执行时间</option>
            <option value="volume">交易量</option>
            <option value="price">价格波动</option>
            <option value="profit">盈利分析</option>
          </select>
        </div>
        <div class="stock-charts">
          <div class="chart-item">
            <h3>执行时间分析</h3>
            <LineChart 
              :x-axis-data="stockExecutionXAxisData" 
              :series="stockExecutionSeriesData" 
              title="执行时间趋势"
            />
          </div>
          <div class="chart-item">
            <h3>市场比较</h3>
            <BarChart 
              :x-axis-data="stockMarketXAxisData" 
              :series="stockMarketSeriesData" 
              title="市场执行时间比较"
            />
          </div>
        </div>
      </div>
    </BaseCard>

    <!-- 基金交易 -->
    <BaseCard title="基金交易" class="mt-4">
      <div v-if="loading.fund" class="loading">加载中...</div>
      <div v-else-if="error.fund" class="error">{{ error.fund }}</div>
      <div v-else-if="!data.fund?.data?.salesByType?.length" class="empty">暂无数据</div>
      <div v-else>
        <div class="filter">
          <label for="fundType">基金类型：</label>
          <select id="fundType" v-model="fundType" @change="fetchFundTrade">
            <option value="equity">股票型</option>
            <option value="bond">债券型</option>
            <option value="mixed">混合型</option>
            <option value="money">货币型</option>
            <option value="index">指数型</option>
          </select>
        </div>
        <div class="fund-charts">
          <div class="chart-item">
            <h3>销售金额分布</h3>
            <PieChart 
              :data="fundAmountChartData" 
              title="基金销售金额分布"
            />
          </div>
          <div class="chart-item">
            <h3>销售数量分布</h3>
            <PieChart 
              :data="fundCountChartData" 
              title="基金销售数量分布"
            />
          </div>
        </div>
        <div class="fund-table">
          <table>
            <thead>
              <tr>
                <th>基金类型</th>
                <th>销售金额</th>
                <th>销售数量</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in data.fund.data.salesByType" :key="item.type">
                <td>{{ item.type }}</td>
                <td>{{ item.amount.toLocaleString() }}</td>
                <td>{{ item.count }}</td>
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
import BarChart from '../../components/charts/BarChart.vue'
import { getStockTrade, getFundTrade } from '../../api/dashboard'

// 状态管理
const stockType = ref('execution')
const fundType = ref('equity')
const loading = ref({
  stock: false,
  fund: false
})
const error = ref({
  stock: '',
  fund: ''
})
const data = ref({
  stock: null,
  fund: null
})

// 计算属性 - 股票执行时间图表数据
const stockExecutionXAxisData = computed(() => {
  if (!data.value.stock?.data?.executionTime) return []
  return data.value.stock.data.executionTime.map(item => item.time)
})

const stockExecutionSeriesData = computed(() => {
  if (!data.value.stock?.data?.executionTime) return []
  return [{
    name: '执行时间',
    data: data.value.stock.data.executionTime.map(item => item.avgTime),
    color: '#409eff'
  }]
})

// 计算属性 - 股票市场比较图表数据
const stockMarketXAxisData = computed(() => {
  if (!data.value.stock?.data?.marketComparison) return []
  return data.value.stock.data.marketComparison.map(item => item.market)
})

const stockMarketSeriesData = computed(() => {
  if (!data.value.stock?.data?.marketComparison) return []
  return [{
    name: '平均执行时间',
    data: data.value.stock.data.marketComparison.map(item => item.avgTime),
    color: '#67c23a'
  }]
})

// 计算属性 - 基金销售金额图表数据
const fundAmountChartData = computed(() => {
  if (!data.value.fund?.data?.salesByType) return []
  return data.value.fund.data.salesByType.map(item => ({
    name: item.type,
    value: item.amount
  }))
})

// 计算属性 - 基金销售数量图表数据
const fundCountChartData = computed(() => {
  if (!data.value.fund?.data?.salesByType) return []
  return data.value.fund.data.salesByType.map(item => ({
    name: item.type,
    value: item.count
  }))
})

// 获取股票交易数据
const fetchStockTrade = async () => {
  loading.value.stock = true
  error.value.stock = ''
  try {
    const response = await getStockTrade(stockType.value)
    data.value.stock = response
  } catch (err: any) {
    error.value.stock = err.message || '获取股票交易数据失败'
  } finally {
    loading.value.stock = false
  }
}

// 获取基金交易数据
const fetchFundTrade = async () => {
  loading.value.fund = true
  error.value.fund = ''
  try {
    const response = await getFundTrade(fundType.value)
    data.value.fund = response
  } catch (err: any) {
    error.value.fund = err.message || '获取基金交易数据失败'
  } finally {
    loading.value.fund = false
  }
}

// 组件挂载时获取数据
onMounted(() => {
  fetchStockTrade()
  fetchFundTrade()
})
</script>

<style scoped>
.trade {
  padding: 20px 0;
}

.trade h2 {
  margin-bottom: 20px;
  color: #303133;
}

.mt-4 {
  margin-top: 20px;
}

.filter {
  margin-bottom: 20px;
}

.filter label {
  margin-right: 10px;
  font-weight: 500;
}

.filter select {
  padding: 8px 12px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  font-size: 14px;
}

.loading, .error, .empty {
  padding: 40px 0;
  text-align: center;
  color: #909399;
}

.error {
  color: #f56c6c;
}

.stock-charts,
.fund-charts {
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

.fund-table {
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
  .stock-charts,
  .fund-charts {
    grid-template-columns: 1fr;
  }
}
</style>