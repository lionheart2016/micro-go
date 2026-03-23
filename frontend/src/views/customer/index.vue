<template>
  <div class="customer">
    <h2>客户分析</h2>
    
    <!-- 客户分布 -->
    <BaseCard title="客户分布">
      <div v-if="loading.distribution" class="loading">加载中...</div>
      <div v-else-if="error.distribution" class="error">{{ error.distribution }}</div>
      <div v-else-if="!data.distribution?.data?.distribution?.length" class="empty">暂无数据</div>
      <div v-else>
        <div class="filter">
          <label for="dimension">维度：</label>
          <select id="dimension" v-model="dimension" @change="fetchCustomerDistribution">
            <option value="age">年龄</option>
            <option value="region">地区</option>
            <option value="job">职业</option>
            <option value="education">教育程度</option>
            <option value="type">客户类型</option>
          </select>
        </div>
        <div class="distribution-chart">
          <PieChart 
            :data="distributionChartData" 
            title="客户分布"
          />
        </div>
        <div class="distribution-table">
          <table>
            <thead>
              <tr>
                <th>名称</th>
                <th>数量</th>
                <th>总资产</th>
                <th>平均资产</th>
                <th>平均交易金额</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in data.distribution.data.distribution" :key="item.name">
                <td>{{ item.name }}</td>
                <td>{{ item.value }}</td>
                <td>{{ item.totalAsset.toLocaleString() }}</td>
                <td>{{ item.avgAsset.toLocaleString() }}</td>
                <td>{{ item.avgTradeAmount.toLocaleString() }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </BaseCard>

    <!-- 客户行为 -->
    <BaseCard title="客户行为" class="mt-4">
      <div v-if="loading.behavior" class="loading">加载中...</div>
      <div v-else-if="error.behavior" class="error">{{ error.behavior }}</div>
      <div v-else-if="!data.behavior?.data?.trend?.length" class="empty">暂无数据</div>
      <div v-else>
        <div class="filter">
          <label for="behaviorType">指标类型：</label>
          <select id="behaviorType" v-model="behaviorType" @change="fetchCustomerBehavior">
            <option value="retention">留存率</option>
            <option value="churn">流失率</option>
            <option value="clv">客户生命周期价值</option>
            <option value="activity">活跃度</option>
          </select>
        </div>
        <div class="behavior-chart">
          <LineChart 
            :x-axis-data="behaviorXAxisData" 
            :series="behaviorSeriesData" 
            title="客户行为趋势"
          />
        </div>
        <div class="behavior-stats">
          <div class="stat-item">
            <span class="stat-label">行业平均：</span>
            <span class="stat-value">{{ data.behavior.data.industryAverage }}%</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">竞争对手平均：</span>
            <span class="stat-value">{{ data.behavior.data.competitorAverage }}%</span>
          </div>
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
import { getCustomerDistribution, getCustomerBehavior } from '../../api/dashboard'

// 状态管理
const dimension = ref('age')
const behaviorType = ref('retention')
const loading = ref({
  distribution: false,
  behavior: false
})
const error = ref({
  distribution: '',
  behavior: ''
})
const data = ref({
  distribution: null,
  behavior: null
})

// 计算属性 - 客户分布图表数据
const distributionChartData = computed(() => {
  if (!data.value.distribution?.data?.distribution) return []
  return data.value.distribution.data.distribution.map(item => ({
    name: item.name,
    value: item.value
  }))
})

// 计算属性 - 客户行为图表数据
const behaviorXAxisData = computed(() => {
  if (!data.value.behavior?.data?.trend) return []
  return data.value.behavior.data.trend.map(item => item.quarter)
})

const behaviorSeriesData = computed(() => {
  if (!data.value.behavior?.data?.trend) return []
  return [
    {
      name: '留存率',
      data: data.value.behavior.data.trend.map(item => item.retentionRate),
      color: '#409eff'
    },
    {
      name: '流失率',
      data: data.value.behavior.data.trend.map(item => item.churnRate),
      color: '#f56c6c'
    }
  ]
})

// 获取客户分布数据
const fetchCustomerDistribution = async () => {
  loading.value.distribution = true
  error.value.distribution = ''
  try {
    const response = await getCustomerDistribution(dimension.value)
    data.value.distribution = response
  } catch (err: any) {
    error.value.distribution = err.message || '获取客户分布数据失败'
  } finally {
    loading.value.distribution = false
  }
}

// 获取客户行为数据
const fetchCustomerBehavior = async () => {
  loading.value.behavior = true
  error.value.behavior = ''
  try {
    const response = await getCustomerBehavior(behaviorType.value)
    data.value.behavior = response
  } catch (err: any) {
    error.value.behavior = err.message || '获取客户行为数据失败'
  } finally {
    loading.value.behavior = false
  }
}

// 组件挂载时获取数据
onMounted(() => {
  fetchCustomerDistribution()
  fetchCustomerBehavior()
})
</script>

<style scoped>
.customer {
  padding: 20px 0;
}

.customer h2 {
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

.distribution-chart,
.behavior-chart {
  margin-bottom: 20px;
}

.distribution-table {
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

.behavior-stats {
  display: flex;
  gap: 30px;
  margin-top: 20px;
}

.stat-item {
  display: flex;
  align-items: center;
}

.stat-label {
  margin-right: 10px;
  color: #606266;
}

.stat-value {
  font-weight: 600;
  color: #409eff;
}
</style>