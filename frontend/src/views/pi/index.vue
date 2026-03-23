<template>
  <div class="pi">
    <h2>PI用户分析</h2>
    
    <!-- PI用户基本信息 -->
    <BaseCard title="PI用户基本信息">
      <div v-if="loading.basic" class="loading">加载中...</div>
      <div v-else-if="error.basic" class="error">{{ error.basic }}</div>
      <div v-else-if="!data.basic?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="basic-stats">
          <div class="stat-card">
            <div class="stat-value">{{ data.basic.data.totalUsers.toLocaleString() }}</div>
            <div class="stat-label">总用户数</div>
            <div class="stat-change" :class="{ positive: data.basic.data.change >= 0, negative: data.basic.data.change < 0 }">
              {{ data.basic.data.change >= 0 ? '+' : '' }}{{ data.basic.data.change }} ({{ data.basic.data.changeRate }}%)
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-value">{{ data.basic.data.totalAsset.toLocaleString() }}</div>
            <div class="stat-label">总资产</div>
          </div>
        </div>
        <div class="basic-charts">
          <div class="chart-item">
            <h3>资产分布</h3>
            <PieChart 
              :data="assetDistributionChartData" 
              title="PI用户资产分布"
            />
          </div>
          <div class="chart-item">
            <h3>人口统计</h3>
            <div class="demographics-table">
              <table>
                <thead>
                  <tr>
                    <th>职业</th>
                    <th>占比</th>
                    <th>平均资产</th>
                    <th>平均交易频率</th>
                    <th>平均基金数量</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in data.basic.data.demographics" :key="item.job">
                    <td>{{ item.job }}</td>
                    <td>{{ item.percentage }}%</td>
                    <td>{{ item.avgAsset.toLocaleString() }}</td>
                    <td>{{ item.avgTradeFrequency }}</td>
                    <td>{{ item.avgFundCount }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </BaseCard>

    <!-- PI用户行为 -->
    <BaseCard title="PI用户行为" class="mt-4">
      <div v-if="loading.behavior" class="loading">加载中...</div>
      <div v-else-if="error.behavior" class="error">{{ error.behavior }}</div>
      <div v-else-if="!data.behavior?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="filter">
          <label for="behaviorType">行为类型：</label>
          <select id="behaviorType" v-model="behaviorType" @change="fetchPIBehavior">
            <option value="trade">交易行为</option>
            <option value="asset">资产配置</option>
            <option value="login">登录行为</option>
            <option value="preference">投资偏好</option>
          </select>
        </div>
        <div class="behavior-stats">
          <div class="stat-card">
            <div class="stat-value">{{ data.behavior.data.avgTradeAmount.toLocaleString() }}</div>
            <div class="stat-label">平均交易金额</div>
          </div>
          <div class="stat-card">
            <div class="stat-value">{{ data.behavior.data.avgTradeFrequency }}</div>
            <div class="stat-label">平均交易频率</div>
          </div>
          <div class="stat-card">
            <div class="stat-value">{{ data.behavior.data.avgHoldingPeriod }}天</div>
            <div class="stat-label">平均持有期</div>
          </div>
        </div>
        <div class="behavior-chart">
          <h3>资金流向</h3>
          <div class="money-flow-table">
            <table>
              <thead>
                <tr>
                  <th>来源</th>
                  <th>去向</th>
                  <th>金额</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(item, index) in data.behavior.data.moneyFlow" :key="index">
                  <td>{{ item.from }}</td>
                  <td>{{ item.to }}</td>
                  <td>{{ item.value.toLocaleString() }}</td>
                </tr>
              </tbody>
            </table>
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
import { getPIBasic, getPIBehavior } from '../../api/dashboard'

// 状态管理
const behaviorType = ref('trade')
const loading = ref({
  basic: false,
  behavior: false
})
const error = ref({
  basic: '',
  behavior: ''
})
const data = ref({
  basic: null,
  behavior: null
})

// 计算属性 - 资产分布图表数据
const assetDistributionChartData = computed(() => {
  if (!data.value.basic?.data?.assetDistribution) return []
  return data.value.basic.data.assetDistribution.map(item => ({
    name: item.name,
    value: item.value
  }))
})

// 获取PI用户基本信息
const fetchPIBasic = async () => {
  loading.value.basic = true
  error.value.basic = ''
  try {
    const response = await getPIBasic()
    data.value.basic = response
  } catch (err: any) {
    error.value.basic = err.message || '获取PI用户基本信息失败'
  } finally {
    loading.value.basic = false
  }
}

// 获取PI用户行为
const fetchPIBehavior = async () => {
  loading.value.behavior = true
  error.value.behavior = ''
  try {
    const response = await getPIBehavior(behaviorType.value)
    data.value.behavior = response
  } catch (err: any) {
    error.value.behavior = err.message || '获取PI用户行为数据失败'
  } finally {
    loading.value.behavior = false
  }
}

// 组件挂载时获取数据
onMounted(() => {
  fetchPIBasic()
  fetchPIBehavior()
})
</script>

<style scoped>
.pi {
  padding: 20px 0;
}

.pi h2 {
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

.basic-stats,
.behavior-stats {
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
  margin-bottom: 8px;
}

.stat-change {
  font-size: 12px;
}

.stat-change.positive {
  color: #67c23a;
}

.stat-change.negative {
  color: #f56c6c;
}

.basic-charts {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.chart-item h3 {
  margin-bottom: 10px;
  font-size: 16px;
  color: #303133;
}

.demographics-table,
.money-flow-table {
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
  .basic-charts {
    grid-template-columns: 1fr;
  }
  
  .basic-stats,
  .behavior-stats {
    flex-direction: column;
  }
}
</style>