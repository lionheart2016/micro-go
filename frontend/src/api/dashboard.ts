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
    accountCount: Indicator
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

// 客户分布项接口
export interface CustomerDistributionItem {
  name: string
  value: number
  totalAsset: number
  avgAsset: number
  avgTradeAmount: number
}

// 客户分布响应
export interface CustomerDistributionResponse {
  code: number
  message: string
  data: {
    dimension: string
    distribution: CustomerDistributionItem[]
  }
}

// 客户行为趋势接口
export interface CustomerBehaviorTrend {
  quarter: string
  retentionRate: number
  churnRate: number
}

// 客户行为响应
export interface CustomerBehaviorResponse {
  code: number
  message: string
  data: {
    type: string
    trend: CustomerBehaviorTrend[]
    industryAverage: number
    competitorAverage: number
  }
}

// 执行时间接口
export interface ExecutionTime {
  time: string
  avgTime: number
}

// 市场比较接口
export interface MarketComparison {
  market: string
  avgTime: number
}

// 股票交易响应
export interface StockTradeResponse {
  code: number
  message: string
  data: {
    type: string
    executionTime: ExecutionTime[]
    marketComparison: MarketComparison[]
  }
}

// 销售类型接口
export interface SalesByType {
  type: string
  amount: number
  count: number
}

// 基金交易响应
export interface FundTradeResponse {
  code: number
  message: string
  data: {
    type: string
    salesByType: SalesByType[]
  }
}

// 资产分布接口
export interface AssetDistribution {
  name: string
  value: number
}

// 人口统计接口
export interface Demographic {
  job: string
  percentage: number
  avgAsset: number
  avgTradeFrequency: number
  avgFundCount: number
}

// PI用户基本信息响应
export interface PIBasicResponse {
  code: number
  message: string
  data: {
    totalUsers: number
    change: number
    changeRate: number
    assetDistribution: AssetDistribution[]
    totalAsset: number
    demographics: Demographic[]
  }
}

// 资金流向接口
export interface MoneyFlow {
  from: string
  to: string
  value: number
}

// PI用户行为响应
export interface PIBehaviorResponse {
  code: number
  message: string
  data: {
    type: string
    avgTradeAmount: number
    avgTradeFrequency: number
    avgHoldingPeriod: number
    moneyFlow: MoneyFlow[]
  }
}

// 漏斗阶段接口
export interface FunnelStage {
  stage: string
  count: number
  conversion: number
}

// 阶段时间接口
export interface StageTime {
  stage: string
  avgTime: number
  medianTime: number
  p25Time: number
  p75Time: number
  maxTime: number
}

// 开户流程响应
export interface AccountProcessResponse {
  code: number
  message: string
  data: {
    funnel: FunnelStage[]
    stageTime: StageTime[]
  }
}

// 异常类型接口
export interface ExceptionType {
  type: string
  count: number
  percentage: number
}

// 异常阶段接口
export interface ExceptionStage {
  stage: string
  count: number
}

// 异常处理接口
export interface ExceptionHandling {
  type: string
  avgHandleTime: number
  maxHandleTime: number
  successRate: number
}

// 开户异常响应
export interface AccountExceptionResponse {
  code: number
  message: string
  data: {
    exceptionTypes: ExceptionType[]
    exceptionStages: ExceptionStage[]
    exceptionHandling: ExceptionHandling[]
  }
}

// 用户分布接口
export interface UserDistribution {
  type: string
  count: number
  percentage: number
}

// 认购金额接口
export interface SubscriptionAmount {
  project: string
  piAmount: number
  regularAmount: number
  totalAmount: number
}

// 认购倍数接口
export interface SubscriptionMultiple {
  project: string
  multiple: number
}

// IPO认购响应
export interface IPOSubscriptionResponse {
  code: number
  message: string
  data: {
    totalUsers: number
    userDistribution: UserDistribution[]
    subscriptionAmount: SubscriptionAmount[]
    subscriptionMultiples: SubscriptionMultiple[]
  }
}

// 项目表现接口
export interface ProjectPerformance {
  project: string
  priceChange: number
  avgReturn: number
}

// 影响因素接口
export interface InfluenceFactor {
  factor: string
  correlation: number
}

// 竞争对手比较接口
export interface CompetitorComparison {
  competitor: string
  subscriptionUsers: number
  subscriptionAmount: number
  multiple: number
}

// IPO分析响应
export interface IPOAnalysisResponse {
  code: number
  message: string
  data: {
    projectPerformance: ProjectPerformance[]
    influenceFactors: InfluenceFactor[]
    competitorComparison: CompetitorComparison[]
  }
}

// 融资规模接口
export interface FinancingSize {
  avgSize: number
  medianSize: number
  p25Size: number
  p75Size: number
  maxSize: number
}

// 融资目的接口
export interface FinancingPurpose {
  purpose: string
  percentage: number
}

// 风险评估接口
export interface RiskAssessment {
  type: string
  defaultRate: number
  overdueRate: number
}

// 融资客户响应
export interface FinanceCustomerResponse {
  code: number
  message: string
  data: {
    totalUsers: number
    userDistribution: UserDistribution[]
    financingSize: FinancingSize
    financingPurpose: FinancingPurpose[]
    riskAssessment: RiskAssessment[]
  }
}

// 热门股票接口
export interface HotStock {
  stock: string
  financingBalance: number
  marketCapRatio: number
}

// 融资趋势接口
export interface FinancingTrend {
  date: string
  buyAmount: number
  sellAmount: number
}

// 保证金比例接口
export interface MarginRatio {
  stock: string
  ratio: number
  priceVolatility: number
}

// 融资股票响应
export interface FinanceStockResponse {
  code: number
  message: string
  data: {
    hotStocks: HotStock[]
    financingTrend: FinancingTrend[]
    marginRatio: MarginRatio[]
  }
}

// 明细项接口
export interface DrilldownItem {
  id: string
  customerId: string
  name: string
  age: number
  region: string
  asset: number
  tradeAmount: number
}

// 明细数据响应
export interface DrilldownDetailResponse {
  code: number
  message: string
  data: {
    total: number
    page: number
    pageSize: number
    list: DrilldownItem[]
  }
}

// 趋势项接口
export interface TrendItem {
  date: string
  value: number
}

// 趋势数据响应
export interface DrilldownTrendResponse {
  code: number
  message: string
  data: {
    type: string
    name: string
    trend: TrendItem[]
  }
}

// 用户接口
export interface User {
  id: string
  username: string
  role: string
  permissions?: string[]
}

// 登录响应
export interface LoginResponse {
  code: number
  message: string
  data: {
    token: string
    user: User
  }
}

// 认证信息响应
export interface AuthInfoResponse {
  code: number
  message: string
  data: User
}

// 获取核心指标
export const getCoreIndicators = async (period?: string): Promise<CoreIndicatorsResponse> => {
  const params = period ? { period } : {}
  return await apiClient.get('/dashboard/core-indicators', { params })
}

// 获取业绩表现
export const getPerformance = async (period?: string): Promise<PerformanceResponse> => {
  const params = period ? { period } : {}
  return await apiClient.get('/dashboard/performance', { params })
}

// 获取客户分布
export const getCustomerDistribution = async (dimension?: string): Promise<CustomerDistributionResponse> => {
  const params = dimension ? { dimension } : {}
  return await apiClient.get('/customer/distribution', { params })
}

// 获取客户行为
export const getCustomerBehavior = async (type?: string): Promise<CustomerBehaviorResponse> => {
  const params = type ? { type } : {}
  return await apiClient.get('/customer/behavior', { params })
}

// 获取股票交易
export const getStockTrade = async (type?: string): Promise<StockTradeResponse> => {
  const params = type ? { type } : {}
  return await apiClient.get('/trade/stock', { params })
}

// 获取基金交易
export const getFundTrade = async (type?: string): Promise<FundTradeResponse> => {
  const params = type ? { type } : {}
  return await apiClient.get('/trade/fund', { params })
}

// 获取PI用户基本信息
export const getPIBasic = async (): Promise<PIBasicResponse> => {
  return await apiClient.get('/pi/basic')
}

// 获取PI用户行为
export const getPIBehavior = async (type?: string): Promise<PIBehaviorResponse> => {
  const params = type ? { type } : {}
  return await apiClient.get('/pi/behavior', { params })
}

// 获取开户流程
export const getAccountProcess = async (): Promise<AccountProcessResponse> => {
  return await apiClient.get('/account/process')
}

// 获取开户异常
export const getAccountException = async (): Promise<AccountExceptionResponse> => {
  return await apiClient.get('/account/exception')
}

// 获取IPO认购
export const getIPOSubscription = async (): Promise<IPOSubscriptionResponse> => {
  return await apiClient.get('/ipo/subscription')
}

// 获取IPO分析
export const getIPOAnaalysis = async (): Promise<IPOAnalysisResponse> => {
  return await apiClient.get('/ipo/analysis')
}

// 获取融资客户
export const getFinanceCustomer = async (): Promise<FinanceCustomerResponse> => {
  return await apiClient.get('/finance/customer')
}

// 获取融资股票
export const getFinanceStock = async (): Promise<FinanceStockResponse> => {
  return await apiClient.get('/finance/stock')
}

// 获取明细数据
export const getDrilldownDetail = async (type: string, id: string, page: number = 1, pageSize: number = 10): Promise<DrilldownDetailResponse> => {
  return await apiClient.get('/drilldown/detail', {
    params: { type, id, page, pageSize }
  })
}

// 获取趋势数据
export const getDrilldownTrend = async (type: string, id: string, period: string = '1m'): Promise<DrilldownTrendResponse> => {
  return await apiClient.get('/drilldown/trend', {
    params: { type, id, period }
  })
}

// 登录
export const login = async (username: string, password: string): Promise<LoginResponse> => {
  return await apiClient.post('/auth/login', { username, password })
}

// 获取认证信息
export const getAuthInfo = async (): Promise<AuthInfoResponse> => {
  return await apiClient.get('/auth/info')
}
