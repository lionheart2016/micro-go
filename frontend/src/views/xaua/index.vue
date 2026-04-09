<template>
  <div class="xaua-dashboard">
    <div class="page-header">
      <h1 class="page-title">概览</h1>
      <p class="page-subtitle">XAUa 黄金代币发行管理</p>
    </div>

    <div class="kpi-grid">
      <div class="kpi-card">
        <div class="kpi-card-header">
          <span class="kpi-label">待审批批次</span>
          <div class="kpi-icon-wrapper pending-icon">
            <svg class="kpi-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="10"/>
              <polyline points="12,6 12,12 16,14"/>
            </svg>
          </div>
        </div>
        <div class="kpi-value">2</div>
      </div>

      <div class="kpi-card">
        <div class="kpi-card-header">
          <span class="kpi-label">今日认购 (USDT)</span>
          <div class="kpi-icon-wrapper subscribe-icon">
            <svg class="kpi-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
              <polyline points="7,10 12,15 17,10"/>
              <line x1="12" y1="15" x2="12" y2="3"/>
            </svg>
          </div>
        </div>
        <div class="kpi-value">285,000</div>
      </div>

      <div class="kpi-card">
        <div class="kpi-card-header">
          <span class="kpi-label">累计发行 XAUa</span>
          <div class="kpi-icon-wrapper issue-icon">
            <svg class="kpi-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/>
              <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>
            </svg>
          </div>
        </div>
        <div class="kpi-value">137.22</div>
      </div>

      <div class="kpi-card">
        <div class="kpi-card-header">
          <span class="kpi-label">累计赎回 (USDT)</span>
          <div class="kpi-icon-wrapper redeem-icon">
            <svg class="kpi-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="23,6 13.5,15.5 8.5,10.5 1,18"/>
              <polyline points="17,6 23,6 23,12"/>
            </svg>
          </div>
        </div>
        <div class="kpi-value">221,880</div>
      </div>
    </div>

    <div class="alert-banner">
      <div class="alert-icon">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/>
          <polyline points="12,6 12,12 16,14"/>
        </svg>
      </div>
      <div class="alert-content">
        <p class="alert-title">有 <span class="alert-highlight">2</span> 个批次订单等待审批</p>
        <p class="alert-desc">来自 YunFeng Securities 的订单已汇总，请前往审批中心处理。</p>
        <button class="alert-action">前往审批 →</button>
      </div>
    </div>

    <div class="table-section">
      <h2 class="table-title">最近批次</h2>
      <div class="table-wrapper">
        <table class="batch-table">
          <thead>
            <tr>
              <th>批次编号</th>
              <th>方向</th>
              <th>交易日</th>
              <th>收取</th>
              <th>发送</th>
              <th>子订单</th>
              <th>状态</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="batch in batches" :key="batch.id">
              <td class="batch-id">{{ batch.batchNo }}</td>
              <td>
                <div class="direction-cell">
                  <svg v-if="batch.direction === '认购'" class="direction-icon subscribe" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
                    <polyline points="7,10 12,15 17,10"/>
                    <line x1="12" y1="15" x2="12" y2="3"/>
                  </svg>
                  <svg v-else class="direction-icon redeem" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
                    <polyline points="17,8 12,3 7,8"/>
                    <line x1="12" y1="3" x2="12" y2="15"/>
                  </svg>
                  <span>{{ batch.direction }}</span>
                </div>
              </td>
              <td>{{ batch.tradeDate }}</td>
              <td>{{ batch.receive }}</td>
              <td>{{ batch.send }}</td>
              <td>{{ batch.subOrders }}</td>
              <td>
                <span class="status-badge" :class="getStatusClass(batch.status)">
                  {{ batch.status }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Batch {
  id: string
  batchNo: string
  direction: string
  tradeDate: string
  receive: string
  send: string
  subOrders: number
  status: string
}

const batches = ref<Batch[]>([
  {
    id: '1',
    batchNo: 'BATCH-S-20260325-001',
    direction: '认购',
    tradeDate: '2026-03-25',
    receive: '285,000 USDT',
    send: '102.54 XAUa',
    subOrders: 5,
    status: '待审批'
  },
  {
    id: '2',
    batchNo: 'BATCH-R-20260325-001',
    direction: '赎回',
    tradeDate: '2026-03-25',
    receive: '45.20 XAUa',
    send: '125,622.82 USDT',
    subOrders: 3,
    status: '待审批'
  },
  {
    id: '3',
    batchNo: 'BATCH-S-20260324-001',
    direction: '认购',
    tradeDate: '2026-03-24',
    receive: '520,000 USDT',
    send: '187.42 XAUa',
    subOrders: 3,
    status: '已审批'
  },
  {
    id: '4',
    batchNo: 'BATCH-R-20260324-001',
    direction: '赎回',
    tradeDate: '2026-03-24',
    receive: '80.00 XAUa',
    send: '221,880 USDT',
    subOrders: 2,
    status: '已结算'
  },
  {
    id: '5',
    batchNo: 'BATCH-S-20260323-001',
    direction: '认购',
    tradeDate: '2026-03-23',
    receive: '380,000 USDT',
    send: '137.22 XAUa',
    subOrders: 2,
    status: '已结算'
  },
  {
    id: '6',
    batchNo: 'BATCH-S-20260322-001',
    direction: '认购',
    tradeDate: '2026-03-22',
    receive: '150,000 USDT',
    send: '54.15 XAUa',
    subOrders: 1,
    status: '已拒绝'
  }
])

const getStatusClass = (status: string): string => {
  const statusMap: Record<string, string> = {
    '待审批': 'status-pending',
    '已审批': 'status-approved',
    '已结算': 'status-settled',
    '已拒绝': 'status-rejected'
  }
  return statusMap[status] || ''
}
</script>

<style scoped>
.xaua-dashboard {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: 32px;
}

.page-title {
  font-size: 24px;
  font-weight: 600;
  color: #1a1a2e;
  margin: 0 0 8px 0;
}

.page-subtitle {
  font-size: 14px;
  color: #6b7280;
  margin: 0;
}

.kpi-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
  margin-bottom: 24px;
}

.kpi-card {
  background: #ffffff;
  border-radius: 12px;
  padding: 20px;
  border: 1px solid #e5e7eb;
  transition: all 0.2s ease;
}

.kpi-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  border-color: #d1d5db;
}

.kpi-card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.kpi-label {
  font-size: 14px;
  color: #6b7280;
  font-weight: 500;
}

.kpi-icon-wrapper {
  width: 36px;
  height: 36px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.kpi-icon {
  width: 20px;
  height: 20px;
}

.pending-icon {
  background-color: #fef3c7;
  color: #d97706;
}

.subscribe-icon {
  background-color: #dbeafe;
  color: #2563eb;
}

.issue-icon {
  background-color: #fef3c7;
  color: #d97706;
}

.redeem-icon {
  background-color: #d1fae5;
  color: #059669;
}

.kpi-value {
  font-size: 28px;
  font-weight: 600;
  color: #1a1a2e;
  font-family: 'SF Mono', 'Monaco', 'Inconsolata', 'Roboto Mono', monospace;
}

.alert-banner {
  background: linear-gradient(135deg, #fffbeb 0%, #fef3c7 100%);
  border: 1px solid #fcd34d;
  border-radius: 12px;
  padding: 20px 24px;
  display: flex;
  align-items: flex-start;
  gap: 16px;
  margin-bottom: 32px;
}

.alert-icon {
  width: 24px;
  height: 24px;
  color: #d97706;
  flex-shrink: 0;
  margin-top: 2px;
}

.alert-content {
  flex: 1;
}

.alert-title {
  font-size: 14px;
  font-weight: 600;
  color: #92400e;
  margin: 0 0 4px 0;
}

.alert-highlight {
  color: #d97706;
  font-weight: 700;
}

.alert-desc {
  font-size: 13px;
  color: #b45309;
  margin: 0 0 12px 0;
}

.alert-action {
  background: none;
  border: none;
  color: #d97706;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  padding: 0;
  transition: color 0.2s ease;
}

.alert-action:hover {
  color: #b45309;
}

.table-section {
  background: #ffffff;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  overflow: hidden;
}

.table-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a2e;
  padding: 20px 24px;
  margin: 0;
  border-bottom: 1px solid #e5e7eb;
}

.table-wrapper {
  overflow-x: auto;
}

.batch-table {
  width: 100%;
  border-collapse: collapse;
}

.batch-table thead {
  background-color: #f9fafb;
}

.batch-table th {
  padding: 12px 24px;
  text-align: left;
  font-size: 13px;
  font-weight: 600;
  color: #6b7280;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  border-bottom: 1px solid #e5e7eb;
}

.batch-table td {
  padding: 16px 24px;
  font-size: 14px;
  color: #374151;
  border-bottom: 1px solid #f3f4f6;
}

.batch-table tbody tr:last-child td {
  border-bottom: none;
}

.batch-table tbody tr:hover {
  background-color: #f9fafb;
}

.batch-id {
  font-family: 'SF Mono', 'Monaco', 'Inconsolata', 'Roboto Mono', monospace;
  font-size: 13px;
  color: #1a1a2e;
  font-weight: 500;
}

.direction-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.direction-icon {
  width: 16px;
  height: 16px;
}

.direction-icon.subscribe {
  color: #2563eb;
}

.direction-icon.redeem {
  color: #059669;
}

.status-badge {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 9999px;
  font-size: 12px;
  font-weight: 500;
}

.status-pending {
  background-color: #fef3c7;
  color: #d97706;
}

.status-approved {
  background-color: #dbeafe;
  color: #2563eb;
}

.status-settled {
  background-color: #d1fae5;
  color: #059669;
}

.status-rejected {
  background-color: #fee2e2;
  color: #dc2626;
}

@media (max-width: 1024px) {
  .kpi-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .xaua-dashboard {
    padding: 16px;
  }

  .kpi-grid {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  .batch-table th,
  .batch-table td {
    padding: 12px 16px;
  }
}
</style>
