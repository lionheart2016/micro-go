# 证券公司港美股经纪业务数据看板 - 前端架构设计

## 1. 技术栈选择

| 技术 | 版本 | 用途 |
|------|------|------|
| Vue | 3.x | 前端框架 |
| TypeScript | 5.x | 类型系统 |
| ECharts | 5.x | 数据可视化 |
| Pinia | 2.x | 状态管理 |
| Vue Router | 4.x | 路由管理 |
| Axios | 1.x | HTTP客户端 |
| Element Plus | 2.x | UI组件库 |
| Vite | 4.x | 构建工具 |

## 2. 架构设计

### 2.1 整体架构

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              浏览器层                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                              路由层                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                            页面组件层                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                          业务组件层                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                          基础组件层                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                           工具/服务层                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                            API 层                                        │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 2.2 核心模块

1. **路由层**：负责页面导航和路由管理
2. **页面组件层**：包含各个主要页面
3. **业务组件层**：包含业务相关的可复用组件
4. **基础组件层**：包含通用的UI组件
5. **工具/服务层**：包含工具函数和服务
6. **API层**：负责与后端API通信

## 3. 目录结构

```
├── public/              # 静态资源
├── src/                 # 源代码
│   ├── assets/          # 资源文件
│   ├── components/      # 组件
│   │   ├── base/        # 基础组件
│   │   ├── business/    # 业务组件
│   │   └── charts/      # 图表组件
│   ├── views/           # 页面
│   │   ├── dashboard/   # 首页-关键指标概览
│   │   ├── customer/    # 客户分析
│   │   ├── trade/       # 交易分析
│   │   ├── pi/          # PI用户分析
│   │   ├── account/     # 开户主题
│   │   ├── ipo/         # IPO主题
│   │   └── finance/     # 融资主题
│   ├── router/          # 路由配置
│   ├── store/           # 状态管理
│   ├── api/             # API请求
│   ├── utils/           # 工具函数
│   ├── services/        # 服务
│   ├── types/           # 类型定义
│   ├── App.vue          # 根组件
│   └── main.ts          # 入口文件
├── vite.config.ts       # Vite配置
├── tsconfig.json        # TypeScript配置
├── package.json         # 项目配置
└── README.md            # 项目说明
```

## 4. 组件设计

### 4.1 基础组件

| 组件名称 | 功能描述 | 路径 |
|---------|---------|------|
| BaseCard | 卡片组件 | components/base/BaseCard.vue |
| BaseChart | 图表基础组件 | components/base/BaseChart.vue |
| BaseTable | 表格组件 | components/base/BaseTable.vue |
| BaseFilter | 筛选组件 | components/base/BaseFilter.vue |
| BasePagination | 分页组件 | components/base/BasePagination.vue |

### 4.2 业务组件

| 组件名称 | 功能描述 | 路径 |
|---------|---------|------|
| KpiCard | 关键指标卡片 | components/business/KpiCard.vue |
| TrendChart | 趋势图表 | components/business/TrendChart.vue |
| DistributionChart | 分布图表 | components/business/DistributionChart.vue |
| HeatmapChart | 热力图 | components/business/HeatmapChart.vue |
| FunnelChart | 漏斗图 | components/business/FunnelChart.vue |

### 4.3 图表组件

| 组件名称 | 功能描述 | 路径 |
|---------|---------|------|
| LineChart | 折线图 | components/charts/LineChart.vue |
| BarChart | 柱状图 | components/charts/BarChart.vue |
| PieChart | 饼图 | components/charts/PieChart.vue |
| ScatterChart | 散点图 | components/charts/ScatterChart.vue |
| RadarChart | 雷达图 | components/charts/RadarChart.vue |
| SankeyChart | 桑基图 | components/charts/SankeyChart.vue |
| MapChart | 地图 | components/charts/MapChart.vue |

## 5. 页面设计

### 5.1 首页-关键指标概览

**路径**：`views/dashboard/index.vue`

**组件结构**：
- KpiCard × 多个（核心指标）
- TrendChart（业绩趋势）
- BarChart（收入结构）
- LineChart（入金趋势）

**功能**：
- 展示核心业务指标
- 展示业绩表现摘要
- 支持指标钻取

### 5.2 客户分析页面

**路径**：`views/customer/index.vue`

**组件结构**：
- DistributionChart（客户分布）
- MapChart（地域分布）
- LineChart（留存率趋势）
- BoxPlotChart（CLV分布）
- HeatmapChart（活跃度热力图）

**功能**：
- 客户基础信息与分布分析
- 客户行为指标分析
- 客户留存率与流失率分析
- 客户生命周期价值分析

### 5.3 交易分析页面

**路径**：`views/trade/index.vue`

**组件结构**：
- LineChart（订单执行时间）
- BarChart（成交率）
- AreaChart（交易时段分布）
- Table（热门股票排行）
- BarChart（投资分布）

**功能**：
- 港美股交易详情分析
- 基金交易详情分析
- 交易时段分布分析
- 热门股票分析

### 5.4 PI用户分析页面

**路径**：`views/pi/index.vue`

**组件结构**：
- PieChart（资产占比）
- RadarChart（用户特征）
- Table（交易行为）
- SankeyChart（资金流动）

**功能**：
- PI用户基本信息分析
- PI用户港美股交易行为分析
- PI用户基金投资行为分析

### 5.5 开户主题页面

**路径**：`views/account/index.vue`

**组件结构**：
- FunnelChart（开户流程）
- BoxPlotChart（各阶段耗时）
- PieChart（异常类型分布）
- BarChart（异常发生阶段）
- Table（异常处理记录）

**功能**：
- 开户流程阶段分析
- 开户异常情况分析

### 5.6 IPO主题页面

**路径**：`views/ipo/index.vue`

**组件结构**：
- BarChart（认购客户分布）
- StackedBarChart（认购金额）
- LineChart（认购倍数趋势）
- Table（项目收益表现）
- ScatterChart（影响因素分析）

**功能**：
- IPO客户认购情况分析
- IPO项目收益表现分析
- 认购决策影响因素分析

### 5.7 融资主题页面

**路径**：`views/finance/index.vue`

**组件结构**：
- PieChart（融资客户分布）
- BoxPlotChart（融资规模）
- BarChart（融资用途）
- Table（融资风险评估）
- BarChart（热门股票融资）
- LineChart（融资买入卖出趋势）

**功能**：
- 客户融资情况分析
- 热门股票融资情况分析
- 融资风险评估

## 6. 状态管理

使用 Pinia 进行状态管理，主要模块包括：

| 模块名称 | 功能描述 | 路径 |
|---------|---------|------|
| dashboardStore | 首页数据状态 | store/dashboard.ts |
| customerStore | 客户分析数据状态 | store/customer.ts |
| tradeStore | 交易分析数据状态 | store/trade.ts |
| piStore | PI用户分析数据状态 | store/pi.ts |
| accountStore | 开户主题数据状态 | store/account.ts |
| ipoStore | IPO主题数据状态 | store/ipo.ts |
| financeStore | 融资主题数据状态 | store/finance.ts |
| commonStore | 通用状态（如加载状态、错误信息） | store/common.ts |

## 7. API设计

### 7.1 API目录结构

```
src/api/
├── index.ts              # API入口
├── dashboard.ts          # 首页相关API
├── customer.ts           # 客户分析相关API
├── trade.ts              # 交易分析相关API
├── pi.ts                 # PI用户分析相关API
├── account.ts            # 开户主题相关API
├── ipo.ts                # IPO主题相关API
├── finance.ts            # 融资主题相关API
└── utils.ts              # API工具函数
```

### 7.2 API请求封装

使用 Axios 封装 API 请求，支持：
- 请求/响应拦截器
- 错误处理
- 超时设置
- 重试机制
- 统一的请求/响应格式

## 8. 路由设计

```typescript
// src/router/index.ts
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: () => import('../views/dashboard/index.vue'),
    meta: { title: '关键指标概览' }
  },
  {
    path: '/customer',
    name: 'Customer',
    component: () => import('../views/customer/index.vue'),
    meta: { title: '客户分析' }
  },
  {
    path: '/trade',
    name: 'Trade',
    component: () => import('../views/trade/index.vue'),
    meta: { title: '交易分析' }
  },
  {
    path: '/pi',
    name: 'PI',
    component: () => import('../views/pi/index.vue'),
    meta: { title: 'PI用户分析' }
  },
  {
    path: '/account',
    name: 'Account',
    component: () => import('../views/account/index.vue'),
    meta: { title: '开户主题' }
  },
  {
    path: '/ipo',
    name: 'IPO',
    component: () => import('../views/ipo/index.vue'),
    meta: { title: 'IPO主题' }
  },
  {
    path: '/finance',
    name: 'Finance',
    component: () => import('../views/finance/index.vue'),
    meta: { title: '融资主题' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
```

## 9. 性能优化

### 9.1 代码优化

- 使用 TypeScript 进行类型检查
- 组件懒加载
- 代码分割
- 按需引入第三方库

### 9.2 渲染优化

- 使用虚拟列表处理大数据表格
- 图表组件按需渲染
- 避免不必要的重渲染
- 使用 `v-memo` 优化列表渲染

### 9.3 网络优化

- API请求缓存
- 批量请求
- 数据预加载
- 压缩传输数据

### 9.4 资源优化

- 图片懒加载
- 静态资源缓存
- CDN加速
- 字体优化

## 10. 响应式设计

- 使用 Element Plus 的响应式布局
- 针对不同屏幕尺寸优化布局
- 图表组件自适应容器大小
- 移动端友好的交互设计

## 11. 安全措施

- XSS防护
- CSRF防护
- 敏感数据加密
- 权限控制
- 输入验证

## 12. 开发与部署

### 12.1 开发环境

- Vite 开发服务器
- 热更新
- 代码检查

### 12.2 构建与部署

- Vite 生产构建
- 静态资源优化
- 部署到 CDN
- 持续集成/持续部署

## 13. 总结

本前端架构设计基于 Vue3 + TypeScript + ECharts 技术栈，采用组件化、模块化的设计思想，确保系统的可维护性、可扩展性和性能。通过合理的目录结构、组件设计和状态管理，实现了数据看板的各项功能需求，为用户提供直观、高效的数据可视化体验。

在实施过程中，应根据实际业务需求和技术条件，对架构进行适当调整和优化，以达到最佳的用户体验和系统性能。