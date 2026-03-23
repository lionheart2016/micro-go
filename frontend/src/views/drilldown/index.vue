<template>
  <div class="drilldown">
    <h2>数据钻取</h2>
    
    <!-- 明细数据 -->
    <BaseCard title="明细数据">
      <div v-if="loading.detail" class="loading">加载中...</div>
      <div v-else-if="error.detail" class="error">{{ error.detail }}</div>
      <div v-else-if="!data.detail?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="filter">
          <div class="filter-item">
            <label for="detailType">类型：</label>
            <select id="detailType" v-model="detailType" @change="fetchDrilldownDetail">
              <option value="customer">客户</option>
              <option value="trade">交易</option>
              <option value="account">开户</option>
              <option value="finance">融资</option>
            </select>
          </div>
          <div class="filter-item">
            <label for="detailId">ID：</label>
            <input type="text" id="detailId" v-model="detailId" placeholder="请输入ID" @change="fetchDrilldownDetail" />
          </div>
          <div class="filter-item">
            <label for="pageSize">每页条数：</label>
            <select id="pageSize" v-model="pageSize" @change="fetchDrilldownDetail">
              <option value="10">10</option>
              <option value="20">20</option>
              <option value="50">50</option>
              <option value="100">100</option>
            </select>
          </div>
        </div>
        <div class="detail-table">
          <table>
            <thead>
              <tr>
                <th>ID</th>
                <th>客户ID</th>
                <th>姓名</th>
                <th>年龄</th>
                <th>地区</th>
                <th>资产</th>
                <th>交易金额</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in data.detail.data.list" :key="item.id">
                <td>{{ item.id }}</td>
                <td>{{ item.customerId }}</td>
                <td>{{ item.name }}</td>
                <td>{{ item.age }}</td>
                <td>{{ item.region }}</td>
                <td>{{ item.asset.toLocaleString() }}</td>
                <td>{{ item.tradeAmount.toLocaleString() }}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="pagination">
          <button 
            :disabled="data.detail.data.page <= 1" 
            @click="changePage(data.detail.data.page - 1)"
          >
            上一页
          </button>
          <span>{{ data.detail.data.page }} / {{ Math.ceil(data.detail.data.total / data.detail.data.pageSize) }}</span>
          <button 
            :disabled="data.detail.data.page >= Math.ceil(data.detail.data.total / data.detail.data.pageSize)" 
            @click="changePage(data.detail.data.page + 1)"
          >
            下一页
          </button>
        </div>
      </div>
    </BaseCard>

    <!-- 趋势数据 -->
    <BaseCard title="趋势数据" class="mt-4">
      <div v-if="loading.trend" class="loading">加载中...</div>
      <div v-else-if="error.trend" class="error">{{ error.trend }}</div>
      <div v-else-if="!data.trend?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="filter">
          <div class="filter-item">
            <label for="trendType">类型：</label>
            <select id="trendType" v-model="trendType" @change="fetchDrilldownTrend">
              <option value="customer">客户</option>
              <option value="trade">交易</option>
              <option value="account">开户</option>
              <option value="finance">融资</option>
            </select>
          </div>
          <div class="filter-item">
            <label for="trendId">ID：</label>
            <input type="text" id="trendId" v-model="trendId" placeholder="请输入ID" @change="fetchDrilldownTrend" />
          </div>
          <div class="filter-item">
            <label for="period">周期：</label>
            <select id="period" v-model="period" @change="fetchDrilldownTrend">
              <option value="7d">7天</option>
              <option value="30d">30天</option>
              <option value="90d">90天</option>
              <option value="1y">1年</option>
            </select>
          </div>
        </div>
        <div class="trend-chart">
          <LineChart 
            :x-axis-data="trendXAxisData" 
            :series="trendSeriesData" 
            title="趋势数据"
          />
        </div>
        <div class="trend-info">
          <div class="info-item">
            <span class="info-label">类型：</span>
            <span class="info-value">{{ data.trend.data.type }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">名称：</span>
            <span class="info-value">{{ data.trend.data.name }}</span>
          </div>
        </div>
      </div>
    </BaseCard>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import BaseCard from '../../components/base/BaseCard.vue'
import LineChart from '../../components/charts/LineChart.vue'
import { getDrilldownDetail, getDrilldownTrend } from '../../api/dashboard'

// 状态管理
const detailType = ref('customer')
const detailId = ref('1')
const page = ref(1)
const pageSize = ref(10)
const trendType = ref('customer')
const trendId = ref('1')
const period = ref('30d')

const loading = ref({
  detail: false,
  trend: false
})
const error = ref({
  detail: '',
  trend: ''
})
const data = ref({
  detail: null,
  trend: null
})

// 计算属性 - 趋势图表数据
const trendXAxisData = computed(() => {
  if (!data.value.trend?.data?.trend) return []
  return data.value.trend.data.trend.map(item => item.date)
})

const trendSeriesData = computed(() => {
  if (!data.value.trend?.data?.trend) return []
  return [{
    name: '数值',
    data: data.value.trend.data.trend.map(item => item.value),
    color: '#409eff'
  }]
})

// 获取明细数据
const fetchDrilldownDetail = async () => {
  loading.value.detail = true
  error.value.detail = ''
  try {
    const response = await getDrilldownDetail(detailType.value, detailId.value, page.value, pageSize.value)
    data.value.detail = response
  } catch (err: any) {
    error.value.detail = err.message || '获取明细数据失败'
  } finally {
    loading.value.detail = false
  }
}

// 获取趋势数据
const fetchDrilldownTrend = async () => {
  loading.value.trend = true
  error.value.trend = ''
  try {
    const response = await getDrilldownTrend(trendType.value, trendId.value, period.value)
    data.value.trend = response
  } catch (err: any) {
    error.value.trend = err.message || '获取趋势数据失败'
  } finally {
    loading.value.trend = false
  }
}

// 切换页码
const changePage = (newPage: number) => {
  page.value = newPage
  fetchDrilldownDetail()
}

// 组件挂载时获取数据
onMounted(() => {
  fetchDrilldownDetail()
  fetchDrilldownTrend()
})
</script>

<style scoped>
.drilldown {
  padding: 20px 0;
}

.drilldown h2 {
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
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.filter-item {
  display: flex;
  align-items: center;
  gap: 10px;
}

.filter-item label {
  font-weight: 500;
  white-space: nowrap;
}

.filter-item select,
.filter-item input {
  padding: 8px 12px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  font-size: 14px;
}

.detail-table {
  overflow-x: auto;
  margin-bottom: 20px;
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

.pagination {
  display: flex;
  align-items: center;
  gap: 10px;
  justify-content: center;
  margin-top: 20px;
}

.pagination button {
  padding: 6px 12px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  background-color: white;
  cursor: pointer;
  font-size: 14px;
}

.pagination button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.trend-chart {
  margin-bottom: 20px;
}

.trend-info {
  display: flex;
  gap: 30px;
  margin-top: 20px;
}

.info-item {
  display: flex;
  align-items: center;
}

.info-label {
  margin-right: 10px;
  color: #606266;
}

.info-value {
  font-weight: 600;
  color: #303133;
}

@media (max-width: 768px) {
  .filter {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .trend-info {
    flex-direction: column;
  }
}
</style>