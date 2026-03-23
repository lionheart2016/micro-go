<template>
  <div class="account">
    <h2>开户主题</h2>
    
    <!-- 开户流程 -->
    <BaseCard title="开户流程分析">
      <div v-if="loading.process" class="loading">加载中...</div>
      <div v-else-if="error.process" class="error">{{ error.process }}</div>
      <div v-else-if="!data.process?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="process-charts">
          <div class="chart-item">
            <h3>开户漏斗</h3>
            <div class="funnel-container">
              <div class="funnel">
                <div 
                  v-for="(stage, index) in data.process.data.funnel" 
                  :key="stage.stage"
                  class="funnel-stage"
                  :style="{
                    width: `${100 - index * 20}%`,
                    opacity: 1 - index * 0.2
                  }"
                >
                  <div class="stage-info">
                    <div class="stage-name">{{ stage.stage }}</div>
                    <div class="stage-count">{{ stage.count }}人</div>
                    <div class="stage-conversion">转化率: {{ stage.conversion }}%</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="chart-item">
            <h3>阶段时间分析</h3>
            <div class="stage-time-table">
              <table>
                <thead>
                  <tr>
                    <th>阶段</th>
                    <th>平均时间</th>
                    <th>中位数时间</th>
                    <th>25%分位</th>
                    <th>75%分位</th>
                    <th>最大时间</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in data.process.data.stageTime" :key="item.stage">
                    <td>{{ item.stage }}</td>
                    <td>{{ item.avgTime }}分钟</td>
                    <td>{{ item.medianTime }}分钟</td>
                    <td>{{ item.p25Time }}分钟</td>
                    <td>{{ item.p75Time }}分钟</td>
                    <td>{{ item.maxTime }}分钟</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </BaseCard>

    <!-- 开户异常 -->
    <BaseCard title="开户异常分析" class="mt-4">
      <div v-if="loading.exception" class="loading">加载中...</div>
      <div v-else-if="error.exception" class="error">{{ error.exception }}</div>
      <div v-else-if="!data.exception?.data" class="empty">暂无数据</div>
      <div v-else>
        <div class="exception-charts">
          <div class="chart-item">
            <h3>异常类型分布</h3>
            <PieChart 
              :data="exceptionTypeChartData" 
              title="异常类型分布"
            />
          </div>
          <div class="chart-item">
            <h3>异常阶段分布</h3>
            <BarChart 
              :x-axis-data="exceptionStageXAxisData" 
              :series="exceptionStageSeriesData" 
              title="异常阶段分布"
            />
          </div>
        </div>
        <div class="exception-handling">
          <h3>异常处理情况</h3>
          <table>
            <thead>
              <tr>
                <th>异常类型</th>
                <th>平均处理时间</th>
                <th>最长处理时间</th>
                <th>成功率</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in data.exception.data.exceptionHandling" :key="item.type">
                <td>{{ item.type }}</td>
                <td>{{ item.avgHandleTime }}分钟</td>
                <td>{{ item.maxHandleTime }}分钟</td>
                <td>{{ item.successRate }}%</td>
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
import { getAccountProcess, getAccountException } from '../../api/dashboard'

// 状态管理
const loading = ref({
  process: false,
  exception: false
})
const error = ref({
  process: '',
  exception: ''
})
const data = ref({
  process: null,
  exception: null
})

// 计算属性 - 异常类型图表数据
const exceptionTypeChartData = computed(() => {
  if (!data.value.exception?.data?.exceptionTypes) return []
  return data.value.exception.data.exceptionTypes.map(item => ({
    name: item.type,
    value: item.count
  }))
})

// 计算属性 - 异常阶段图表数据
const exceptionStageXAxisData = computed(() => {
  if (!data.value.exception?.data?.exceptionStages) return []
  return data.value.exception.data.exceptionStages.map(item => item.stage)
})

const exceptionStageSeriesData = computed(() => {
  if (!data.value.exception?.data?.exceptionStages) return []
  return [{
    name: '异常数量',
    data: data.value.exception.data.exceptionStages.map(item => item.count),
    color: '#f56c6c'
  }]
})

// 获取开户流程数据
const fetchAccountProcess = async () => {
  loading.value.process = true
  error.value.process = ''
  try {
    const response = await getAccountProcess()
    data.value.process = response
  } catch (err: any) {
    error.value.process = err.message || '获取开户流程数据失败'
  } finally {
    loading.value.process = false
  }
}

// 获取开户异常数据
const fetchAccountException = async () => {
  loading.value.exception = true
  error.value.exception = ''
  try {
    const response = await getAccountException()
    data.value.exception = response
  } catch (err: any) {
    error.value.exception = err.message || '获取开户异常数据失败'
  } finally {
    loading.value.exception = false
  }
}

// 组件挂载时获取数据
onMounted(() => {
  fetchAccountProcess()
  fetchAccountException()
})
</script>

<style scoped>
.account {
  padding: 20px 0;
}

.account h2 {
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

.process-charts,
.exception-charts {
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

.funnel-container {
  height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.funnel {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  padding: 20px 0;
}

.funnel-stage {
  width: 100%;
  background-color: #409eff;
  border-radius: 4px;
  padding: 15px;
  margin-bottom: 10px;
  color: white;
  text-align: center;
  transition: all 0.3s ease;
}

.stage-info {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.stage-name {
  font-weight: 600;
  font-size: 14px;
}

.stage-count {
  font-size: 16px;
  font-weight: 600;
}

.stage-conversion {
  font-size: 12px;
  opacity: 0.8;
}

.stage-time-table,
.exception-handling {
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
  .process-charts,
  .exception-charts {
    grid-template-columns: 1fr;
  }
}
</style>