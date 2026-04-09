# 组件路由映射表

> 本文档建立了界面组件与路由路径、文件位置之间的明确对应关系，用于快速定位需要修改的组件位置，提高开发效率和需求理解准确性。

---

## 一、路由配置概览

| 配置项 | 值 |
|--------|-----|
| 路由模式 | `createWebHistory` (HTML5 History 模式) |
| 路由配置文件 | `frontend/src/router/index.ts` |
| 加载方式 | 懒加载 (`import()` 动态导入) |
| 路由总数 | 10 条 |

---

## 二、页面视图组件映射表

### 2.1 业务页面组件

| 序号 | 组件名称 | 路由路径 | 路由名称 | 页面标题 | 组件文件路径 | 组件类型 | 功能描述 |
|------|---------|---------|---------|---------|-------------|---------|---------|
| 1 | Dashboard | `/` | `Dashboard` | 关键指标概览 | `src/views/dashboard/index.vue` | 业务页面 | 港美股经纪业务核心数据实时监控，包含注册人数、开户人数、活跃人数、入金、交易等 KPI 指标卡片和业绩表现图表 |
| 2 | Customer | `/customer` | `Customer` | 客户分析 | `src/views/customer/index.vue` | 业务页面 | 客户分布分析（年龄/地区/职业/教育/类型维度）和客户行为分析（留存率/流失率/生命周期价值/活跃度） |
| 3 | Trade | `/trade` | `Trade` | 交易分析 | `src/views/trade/index.vue` | 业务页面 | 股票交易分析（执行时间/交易量/价格波动/盈利）和基金交易分析（销售分布/类型统计） |
| 4 | PI | `/pi` | `PI` | PI用户分析 | `src/views/pi/index.vue` | 业务页面 | PI用户基本信息（用户数/资产/人口统计）和PI用户行为分析（交易行为/资产配置/登录行为/投资偏好/资金流向） |
| 5 | Account | `/account` | `Account` | 开户主题 | `src/views/account/index.vue` | 业务页面 | 开户流程分析（漏斗转化/阶段时间）和开户异常分析（异常类型分布/异常阶段/异常处理情况） |
| 6 | IPO | `/ipo` | `IPO` | IPO主题 | `src/views/ipo/index.vue` | 业务页面 | IPO认购情况（用户分布/认购金额/认购倍数）和IPO分析（项目表现/影响因素/竞争对手比较） |
| 7 | Finance | `/finance` | `Finance` | 融资主题 | `src/views/finance/index.vue` | 业务页面 | 融资客户分析（用户分布/融资目的/融资规模/风险评估）和融资股票分析（热门股票/融资趋势/保证金比例） |
| 8 | XAUa | `/xaua` | `XAUa` | XAUa 黄金代币 | `src/views/xaua/index.vue` | 业务页面 | XAUa 黄金代币发行管理，包含批次订单审批、认购/赎回统计、批次列表展示 |
| 9 | Drilldown | `/drilldown` | `Drilldown` | 数据钻取 | `src/views/drilldown/index.vue` | 业务页面 | 明细数据查询（客户/交易/开户/融资）和趋势数据分析，支持分页和多维度筛选 |
| 10 | Login | `/login` | `Login` | 登录 | `src/views/auth/index.vue` | 业务页面 | 用户登录认证，包含用户名密码表单、Token 存储和路由跳转 |

### 2.2 视图组件详细结构

#### Dashboard (关键指标概览)
- **文件路径**: `src/views/dashboard/index.vue`
- **使用的子组件**:
  - `KpiCard` - 8 个 KPI 指标卡片（注册人数、开户人数、活跃人数、入金人数/金额、股票交易人数/金额、基金交易人数）
  - `BaseCard` - 卡片容器（收入结构、收入趋势、入金趋势）
  - `PieChart` - 收入结构饼图
  - `LineChart` - 收入趋势和入金趋势折线图
- **API 调用**: `getCoreIndicators()`, `getPerformance()`
- **特色功能**: 日期范围筛选、数据刷新

#### Customer (客户分析)
- **文件路径**: `src/views/customer/index.vue`
- **使用的子组件**:
  - `BaseCard` - 卡片容器（客户分布、客户行为）
  - `PieChart` - 客户分布饼图
  - `LineChart` - 客户行为趋势折线图
- **API 调用**: `getCustomerDistribution()`, `getCustomerBehavior()`
- **特色功能**: 多维度筛选（年龄/地区/职业/教育/类型）、行为指标切换

#### Trade (交易分析)
- **文件路径**: `src/views/trade/index.vue`
- **使用的子组件**:
  - `BaseCard` - 卡片容器（股票交易、基金交易）
  - `LineChart` - 执行时间趋势折线图
  - `BarChart` - 市场执行时间比较柱状图
  - `PieChart` - 基金销售金额/数量分布饼图
- **API 调用**: `getStockTrade()`, `getFundTrade()`
- **特色功能**: 交易类型筛选、基金类型筛选

#### PI (PI用户分析)
- **文件路径**: `src/views/pi/index.vue`
- **使用的子组件**:
  - `BaseCard` - 卡片容器（PI用户基本信息、PI用户行为）
  - `PieChart` - 资产分布饼图
- **API 调用**: `getPIBasic()`, `getPIBehavior()`
- **特色功能**: 人口统计表格、资金流向表格、行为类型筛选

#### Account (开户主题)
- **文件路径**: `src/views/account/index.vue`
- **使用的子组件**:
  - `BaseCard` - 卡片容器（开户流程分析、开户异常分析）
  - `PieChart` - 异常类型分布饼图
  - `BarChart` - 异常阶段分布柱状图
- **API 调用**: `getAccountProcess()`, `getAccountException()`
- **特色功能**: 开户漏斗可视化、阶段时间统计分析

#### IPO (IPO主题)
- **文件路径**: `src/views/ipo/index.vue`
- **使用的子组件**:
  - `BaseCard` - 卡片容器（IPO认购情况、IPO分析）
  - `PieChart` - 用户分布饼图
  - `BarChart` - 认购金额对比、认购倍数、项目表现柱状图
- **API 调用**: `getIPOSubscription()`, `getIPOAnaalysis()`
- **特色功能**: PI/普通用户认购对比、竞争对手比较

#### Finance (融资主题)
- **文件路径**: `src/views/finance/index.vue`
- **使用的子组件**:
  - `BaseCard` - 卡片容器（融资客户分析、融资股票分析）
  - `PieChart` - 用户分布、融资目的饼图
  - `LineChart` - 融资买入卖出趋势折线图
- **API 调用**: `getFinanceCustomer()`, `getFinanceStock()`
- **特色功能**: 融资规模统计、风险评估表格、保证金比例分析

#### XAUa (XAUa 黄金代币)
- **文件路径**: `src/views/xaua/index.vue`
- **使用的子组件**: 无（独立实现）
- **API 调用**: 无（当前使用静态数据）
- **特色功能**: KPI 指标展示、审批提醒横幅、批次订单表格、状态标签

#### Drilldown (数据钻取)
- **文件路径**: `src/views/drilldown/index.vue`
- **使用的子组件**:
  - `BaseCard` - 卡片容器（明细数据、趋势数据）
  - `LineChart` - 趋势数据折线图
- **API 调用**: `getDrilldownDetail()`, `getDrilldownTrend()`
- **特色功能**: 多维度筛选、分页查询、趋势分析

#### Login (登录)
- **文件路径**: `src/views/auth/index.vue`
- **使用的子组件**: 无（独立实现）
- **API 调用**: `login()`
- **特色功能**: 表单验证、Token 存储、路由跳转

---

## 三、公共组件映射表

### 3.1 基础组件 (Base Components)

| 序号 | 组件名称 | 文件路径 | 组件类型 | 功能描述 | 使用场景 | Props |
|------|---------|---------|---------|---------|---------|-------|
| 1 | BaseCard | `src/components/base/BaseCard.vue` | 公共组件 | 基础卡片容器，提供标题、内容插槽、页脚插槽 | 所有页面的内容区块容器 | `title` (标题), `showFooter` (是否显示页脚) |
| 2 | BaseChart | `src/components/base/BaseChart.vue` | 公共组件 | ECharts 图表基础封装，支持响应式、加载状态 | 所有图表组件的底层实现 | `option` (图表配置), `width` (宽度), `height` (高度), `loading` (加载状态) |

### 3.2 业务组件 (Business Components)

| 序号 | 组件名称 | 文件路径 | 组件类型 | 功能描述 | 使用场景 | Props |
|------|---------|---------|---------|---------|---------|-------|
| 1 | KpiCard | `src/components/business/KpiCard.vue` | 业务组件 | KPI 指标卡片，展示数值、变化量、变化率 | Dashboard 页面的核心指标展示 | `title` (标题), `value` (数值), `change` (变化量), `changeRate` (变化率), `loading` (加载状态) |
| 2 | TrendChart | `src/components/business/TrendChart.vue` | 业务组件 | 趋势图组件，基于 BaseChart 封装的折线图 | 多系列趋势对比展示 | `title` (标题), `xAxisData` (X轴数据), `seriesData` (系列数据), `height` (高度) |

### 3.3 图表组件 (Chart Components)

| 序号 | 组件名称 | 文件路径 | 组件类型 | 功能描述 | 使用场景 | Props |
|------|---------|---------|---------|---------|---------|-------|
| 1 | BarChart | `src/components/charts/BarChart.vue` | 公共组件 | 柱状图组件，支持多系列、自定义颜色 | 市场比较、异常阶段分布等 | `title` (标题), `xAxisData` (X轴数据), `series` (系列数据), `height` (高度) |
| 2 | LineChart | `src/components/charts/LineChart.vue` | 公共组件 | 折线图组件，支持平滑曲线、面积填充、金融行业配色 | 趋势分析、时间序列数据 | `title` (标题), `xAxisData` (X轴数据), `series` (系列数据), `height` (高度), `loading` (加载状态) |
| 3 | PieChart | `src/components/charts/PieChart.vue` | 公共组件 | 环形饼图组件，支持百分比显示、交互高亮 | 分布分析、占比展示 | `title` (标题), `data` (数据), `height` (高度), `loading` (加载状态) |

### 3.4 其他组件

| 序号 | 组件名称 | 文件路径 | 组件类型 | 功能描述 | 使用场景 |
|------|---------|---------|---------|---------|---------|
| 1 | HelloWorld | `src/components/HelloWorld.vue` | 示例组件 | Vue 示例组件（开发参考） | 开发参考 |

---

## 四、组件依赖关系图

```
视图页面 (Views)
├── dashboard/index.vue
│   ├── KpiCard (业务组件)
│   ├── BaseCard (基础组件)
│   ├── PieChart (图表组件)
│   └── LineChart (图表组件)
│
├── customer/index.vue
│   ├── BaseCard (基础组件)
│   ├── PieChart (图表组件)
│   └── LineChart (图表组件)
│
├── trade/index.vue
│   ├── BaseCard (基础组件)
│   ├── LineChart (图表组件)
│   ├── BarChart (图表组件)
│   └── PieChart (图表组件)
│
├── pi/index.vue
│   ├── BaseCard (基础组件)
│   └── PieChart (图表组件)
│
├── account/index.vue
│   ├── BaseCard (基础组件)
│   ├── PieChart (图表组件)
│   └── BarChart (图表组件)
│
├── ipo/index.vue
│   ├── BaseCard (基础组件)
│   ├── PieChart (图表组件)
│   └── BarChart (图表组件)
│
├── finance/index.vue
│   ├── BaseCard (基础组件)
│   ├── PieChart (图表组件)
│   └── LineChart (图表组件)
│
├── xaua/index.vue
│   └── (独立实现，无子组件)
│
├── drilldown/index.vue
│   ├── BaseCard (基础组件)
│   └── LineChart (图表组件)
│
└── auth/index.vue
    └── (独立实现，无子组件)

图表组件 (Charts)
├── BarChart.vue
│   └── BaseChart (基础组件)
├── LineChart.vue
│   └── BaseChart (基础组件)
└── PieChart.vue
    └── BaseChart (基础组件)

业务组件 (Business)
├── KpiCard.vue
│   └── (独立实现)
└── TrendChart.vue
    └── BaseChart (基础组件)

基础组件 (Base)
├── BaseCard.vue
│   └── (独立实现)
└── BaseChart.vue
    └── ECharts (第三方库)
```

---

## 五、API 接口映射表

| 页面组件 | API 函数 | 后端路由 | 请求方法 | 功能描述 |
|---------|---------|---------|---------|---------|
| Dashboard | `getCoreIndicators()` | `/api/dashboard/core-indicators` | GET | 获取核心指标数据 |
| Dashboard | `getPerformance()` | `/api/dashboard/performance` | GET | 获取业绩表现数据 |
| Customer | `getCustomerDistribution()` | `/api/customer/distribution` | GET | 获取客户分布数据 |
| Customer | `getCustomerBehavior()` | `/api/customer/behavior` | GET | 获取客户行为数据 |
| Trade | `getStockTrade()` | `/api/trade/stock` | GET | 获取股票交易数据 |
| Trade | `getFundTrade()` | `/api/trade/fund` | GET | 获取基金交易数据 |
| PI | `getPIBasic()` | `/api/pi/basic` | GET | 获取PI用户基本信息 |
| PI | `getPIBehavior()` | `/api/pi/behavior` | GET | 获取PI用户行为数据 |
| Account | `getAccountProcess()` | `/api/account/process` | GET | 获取开户流程数据 |
| Account | `getAccountException()` | `/api/account/exception` | GET | 获取开户异常数据 |
| IPO | `getIPOSubscription()` | `/api/ipo/subscription` | GET | 获取IPO认购数据 |
| IPO | `getIPOAnaalysis()` | `/api/ipo/analysis` | GET | 获取IPO分析数据 |
| Finance | `getFinanceCustomer()` | `/api/finance/customer` | GET | 获取融资客户数据 |
| Finance | `getFinanceStock()` | `/api/finance/stock` | GET | 获取融资股票数据 |
| Drilldown | `getDrilldownDetail()` | `/api/drilldown/detail` | GET | 获取钻取明细数据 |
| Drilldown | `getDrilldownTrend()` | `/api/drilldown/trend` | GET | 获取钻取趋势数据 |
| Login | `login()` | `/api/auth/login` | POST | 用户登录认证 |

---

## 六、快速定位指南

### 6.1 按功能模块定位

| 需求类型 | 需要修改的文件 |
|---------|---------------|
| 修改首页 KPI 指标 | `src/views/dashboard/index.vue` + `src/components/business/KpiCard.vue` |
| 修改客户分析页面 | `src/views/customer/index.vue` |
| 修改交易分析页面 | `src/views/trade/index.vue` |
| 修改 PI 用户分析 | `src/views/pi/index.vue` |
| 修改开户主题 | `src/views/account/index.vue` |
| 修改 IPO 主题 | `src/views/ipo/index.vue` |
| 修改融资主题 | `src/views/finance/index.vue` |
| 修改 XAUa 页面 | `src/views/xaua/index.vue` |
| 修改数据钻取 | `src/views/drilldown/index.vue` |
| 修改登录页面 | `src/views/auth/index.vue` |
| 修改路由配置 | `src/router/index.ts` |
| 修改卡片样式 | `src/components/base/BaseCard.vue` |
| 修改图表样式 | `src/components/charts/*.vue` |
| 修改 KPI 卡片 | `src/components/business/KpiCard.vue` |

### 6.2 按组件类型定位

| 组件类型 | 目录路径 | 说明 |
|---------|---------|------|
| 页面视图 | `src/views/` | 每个业务模块一个目录，包含 `index.vue` 入口文件 |
| 基础组件 | `src/components/base/` | 通用基础组件，可在多个页面复用 |
| 业务组件 | `src/components/business/` | 包含业务逻辑的组件 |
| 图表组件 | `src/components/charts/` | 数据可视化图表组件 |
| 路由配置 | `src/router/` | Vue Router 路由定义 |
| API 接口 | `src/api/` | 后端 API 调用封装 |

---

## 七、技术栈说明

| 技术 | 版本/说明 |
|------|----------|
| 前端框架 | Vue 3 (Composition API) |
| 路由管理 | Vue Router 4 (createWebHistory 模式) |
| 图表库 | ECharts |
| 语言 | TypeScript |
| 构建工具 | Vite (推测) |
| UI 组件库 | Element Plus (部分使用) |

---

## 八、注意事项

1. **懒加载**: 所有路由均采用 `import()` 动态导入，实现代码分割
2. **组件复用**: `BaseCard`、`BaseChart` 及图表组件在多个页面中复用
3. **数据流**: 页面组件通过 `src/api/dashboard` 调用后端 API 获取数据
4. **状态管理**: 当前使用 `ref` 和 `computed` 进行组件内状态管理，未使用 Vuex/Pinia
5. **样式规范**: 使用 `scoped` 样式，避免样式污染
6. **响应式设计**: 所有组件均包含 `@media` 响应式断点适配

---

> 文档生成时间: 2026-04-09
> 维护建议: 每次新增/修改路由或组件时，同步更新此文档
