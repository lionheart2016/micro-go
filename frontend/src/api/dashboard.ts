import apiClient from './utils'

// 核心指标接口
export interface Indicator {
  value: number
  change: number
  changeRate: number
}

// 核心指标响应
export interface CoreIndicatorsResponse {
  code: number
  message: string
  data: {
    registeredUsers: Indicator
    开户人数: Indicator
    activeUsers: Indicator
    depositUsers: Indicator
    depositAmount: Indicator
    stockTradeUsers: Indicator
    stockTradeAmount: Indicator
    fundTradeUsers: Indicator
    fundTradeAmount: Indicator
  }
}

// 收入项接口
export interface RevenueItem {
  name: string
  value: number
}

// 收入趋势接口
export interface RevenueTrend {
  month: string
  stock: number
  fund: number
}

// 收入变化接口
export interface RevenueChange {
  yoy: number
  mom: number
}

// 入金趋势接口
export interface DepositTrend {
  month: string
  value: number
}

// 业绩表现响应
export interface PerformanceResponse {
  code: number
  message: string
  data: {
    stockRevenue: number
    fundRevenue: number
    revenueStructure: RevenueItem[]
    revenueTrend: RevenueTrend[]
    revenueChange: RevenueChange
    depositNetGrowth: number
    depositTrend: DepositTrend[]
    depositWithdrawalRatio: number
  }
}

// 获取核心指标
export const getCoreIndicators = async (period?: string): Promise<CoreIndicatorsResponse> => {
  const params = period ? { period } : {}
  return await apiClient.get('/core-indicators', { params })
}

// 获取业绩表现
export const getPerformance = async (period?: string): Promise<PerformanceResponse> => {
  const params = period ? { period } : {}
  return await apiClient.get('/performance', { params })
}
