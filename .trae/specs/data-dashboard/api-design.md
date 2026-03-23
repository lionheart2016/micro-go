# 证券公司港美股经纪业务数据看板 - API接口设计文档

## 1. 接口设计规范

### 1.1 基本规范

- **RESTful 风格**：使用 HTTP 标准方法和 URL 资源表示
- **HTTP 方法**：GET（查询）、POST（创建）、PUT（更新）、DELETE（删除）
- **URL 命名**：使用小写字母，单词间用连字符分隔
- **状态码**：使用 HTTP 标准状态码
- **响应格式**：统一 JSON 格式

### 1.2 响应格式

```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

- `code`：HTTP 状态码
- `message`：响应消息
- `data`：响应数据

### 1.3 认证授权

- **认证方式**：JWT（JSON Web Token）
- **请求头**：`Authorization: Bearer {token}`
- **权限控制**：基于角色的访问控制（RBAC）

## 2. 接口设计

### 2.1 首页-关键指标概览

#### 2.1.1 获取核心指标

- **路径**：`/api/dashboard/core-indicators`
- **方法**：GET
- **参数**：
  - `period`：统计周期（可选，默认近1月）
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "registeredUsers": {
        "value": 10000,
        "change": 500,
        "changeRate": 5.26
      },
      "开户人数": {
        "value": 8000,
        "change": 300,
        "changeRate": 3.95
      },
      "activeUsers": {
        "value": 3000,
        "change": -200,
        "changeRate": -6.25
      },
      "depositUsers": {
        "value": 2000,
        "change": 150,
        "changeRate": 8.11
      },
      "depositAmount": {
        "value": 50000000,
        "change": 5000000,
        "changeRate": 11.11
      },
      "stockTradeUsers": {
        "value": 1500,
        "change": 100,
        "changeRate": 7.14
      },
      "stockTradeAmount": {
        "value": 30000000,
        "change": 2000000,
        "changeRate": 7.14
      },
      "fundTradeUsers": {
        "value": 800,
        "change": 50,
        "changeRate": 6.67
      },
      "fundTradeAmount": {
        "value": 10000000,
        "change": 500000,
        "changeRate": 5.26
      }
    }
  }
  ```

#### 2.1.2 获取业绩表现

- **路径**：`/api/dashboard/performance`
- **方法**：GET
- **参数**：
  - `period`：统计周期（可选，默认近1月）
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "stockRevenue": 2000000,
      "fundRevenue": 500000,
      "revenueStructure": [
        {
          "name": "佣金",
          "value": 1200000
        },
        {
          "name": "平台费",
          "value": 400000
        },
        {
          "name": "利息",
          "value": 300000
        },
        {
          "name": "IPO手续费",
          "value": 100000
        }
      ],
      "revenueTrend": [
        {
          "month": "2023-01",
          "stock": 1800000,
          "fund": 450000
        },
        {
          "month": "2023-02",
          "stock": 1900000,
          "fund": 480000
        },
        {
          "month": "2023-03",
          "stock": 2000000,
          "fund": 500000
        }
      ],
      "revenueChange": {
        "yoy": 15.38,
        "mom": 5.26
      },
      "depositNetGrowth": 5000000,
      "depositTrend": [
        {
          "month": "2023-01",
          "value": 4000000
        },
        {
          "month": "2023-02",
          "value": 4500000
        },
        {
          "month": "2023-03",
          "value": 5000000
        }
      ],
      "depositWithdrawalRatio": 1.5
    }
  }
  ```

### 2.2 客户分析

#### 2.2.1 获取客户基础信息分布

- **路径**：`/api/customer/distribution`
- **方法**：GET
- **参数**：
  - `dimension`：维度（age/region/job/education/type）
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "dimension": "age",
      "distribution": [
        {
          "name": "18-25岁",
          "value": 1000,
          "totalAsset": 5000000,
          "avgAsset": 5000,
          "avgTradeAmount": 2000
        },
        {
          "name": "26-35岁",
          "value": 3000,
          "totalAsset": 30000000,
          "avgAsset": 10000,
          "avgTradeAmount": 5000
        },
        {
          "name": "36-45岁",
          "value": 2500,
          "totalAsset": 37500000,
          "avgAsset": 15000,
          "avgTradeAmount": 8000
        },
        {
          "name": "46-55岁",
          "value": 1000,
          "totalAsset": 20000000,
          "avgAsset": 20000,
          "avgTradeAmount": 10000
        },
        {
          "name": "55岁以上",
          "value": 500,
          "totalAsset": 15000000,
          "avgAsset": 30000,
          "avgTradeAmount": 15000
        }
      ]
    }
  }
  ```

#### 2.2.2 获取客户行为指标

- **路径**：`/api/customer/behavior`
- **方法**：GET
- **参数**：
  - `type`：指标类型（retention/churn/clv/activity）
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "type": "retention",
      "trend": [
        {
          "quarter": "2022-Q1",
          "retentionRate": 85,
          "churnRate": 15
        },
        {
          "quarter": "2022-Q2",
          "retentionRate": 82,
          "churnRate": 18
        },
        {
          "quarter": "2022-Q3",
          "retentionRate": 80,
          "churnRate": 20
        },
        {
          "quarter": "2022-Q4",
          "retentionRate": 78,
          "churnRate": 22
        },
        {
          "quarter": "2023-Q1",
          "retentionRate": 80,
          "churnRate": 20
        }
      ],
      "industryAverage": 75,
      "competitorAverage": 78
    }
  }
  ```

### 2.3 交易分析

#### 2.3.1 获取港美股交易详情

- **路径**：`/api/trade/stock`
- **方法**：GET
- **参数**：
  - `type`：详情类型（execution/成交率/时段分布/热门股票/投资分布）
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "type": "execution",
      "executionTime": [
        {
          "time": "9:30-10:30",
          "avgTime": 0.5
        },
        {
          "time": "10:30-11:30",
          "avgTime": 0.3
        },
        {
          "time": "13:30-14:30",
          "avgTime": 0.4
        },
        {
          "time": "14:30-15:30",
          "avgTime": 0.6
        }
      ],
      "marketComparison": [
        {
          "market": "港股",
          "avgTime": 0.4
        },
        {
          "market": "美股",
          "avgTime": 0.5
        }
      ]
    }
  }
  ```

#### 2.3.2 获取基金交易详情

- **路径**：`/api/trade/fund`
- **方法**：GET
- **参数**：
  - `type`：详情类型（销售情况/趋势/投资组合/热门基金）
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "type": "销售情况",
      "salesByType": [
        {
          "type": "股票型",
          "amount": 5000000,
          "count": 200
        },
        {
          "type": "债券型",
          "amount": 3000000,
          "count": 150
        },
        {
          "type": "混合型",
          "amount": 1500000,
          "count": 100
        },
        {
          "type": "货币型",
          "amount": 500000,
          "count": 50
        }
      ]
    }
  }
  ```

### 2.4 PI用户分析

#### 2.4.1 获取PI用户基本信息

- **路径**：`/api/pi/basic`
- **方法**：GET
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "totalUsers": 500,
      "change": 50,
      "changeRate": 11.11,
      "assetDistribution": [
        {
          "name": "港股",
          "value": 60
        },
        {
          "name": "美股",
          "value": 30
        },
        {
          "name": "基金",
          "value": 10
        }
      ],
      "totalAsset": 50000000,
      "demographics": [
        {
          "job": "金融行业",
          "percentage": 40,
          "avgAsset": 150000,
          "avgTradeFrequency": 10,
          "avgFundCount": 5
        },
        {
          "job": "科技行业",
          "percentage": 30,
          "avgAsset": 120000,
          "avgTradeFrequency": 8,
          "avgFundCount": 4
        },
        {
          "job": "其他行业",
          "percentage": 30,
          "avgAsset": 80000,
          "avgTradeFrequency": 5,
          "avgFundCount": 3
        }
      ]
    }
  }
  ```

#### 2.4.2 获取PI用户行为特征

- **路径**：`/api/pi/behavior`
- **方法**：GET
- **参数**：
  - `type`：行为类型（stock/fund）
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "type": "stock",
      "avgTradeAmount": 20000,
      "avgTradeFrequency": 10,
      "avgHoldingPeriod": 30,
      "moneyFlow": [
        {
          "from": "科技行业",
          "to": "金融行业",
          "value": 5000000
        },
        {
          "from": "金融行业",
          "to": "消费行业",
          "value": 3000000
        }
      ]
    }
  }
  ```

### 2.5 开户主题

#### 2.5.1 获取开户流程分析

- **路径**：`/api/account/process`
- **方法**：GET
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "funnel": [
        {
          "stage": "注册",
          "count": 1000,
          "conversion": 100
        },
        {
          "stage": "身份验证",
          "count": 800,
          "conversion": 80
        },
        {
          "stage": "风险评估",
          "count": 700,
          "conversion": 87.5
        },
        {
          "stage": "签署协议",
          "count": 600,
          "conversion": 85.7
        },
        {
          "stage": "开户成功",
          "count": 500,
          "conversion": 83.3
        }
      ],
      "stageTime": [
        {
          "stage": "注册",
          "avgTime": 2,
          "medianTime": 1.5,
          "p25Time": 1,
          "p75Time": 2.5,
          "maxTime": 5
        },
        {
          "stage": "身份验证",
          "avgTime": 5,
          "medianTime": 4,
          "p25Time": 3,
          "p75Time": 6,
          "maxTime": 10
        }
      ]
    }
  }
  ```

#### 2.5.2 获取开户异常分析

- **路径**：`/api/account/exception`
- **方法**：GET
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "exceptionTypes": [
        {
          "type": "身份验证失败",
          "count": 50,
          "percentage": 41.7
        },
        {
          "type": "资料不完整",
          "count": 30,
          "percentage": 25
        },
        {
          "type": "风险评估不通过",
          "count": 20,
          "percentage": 16.7
        },
        {
          "type": "其他",
          "count": 20,
          "percentage": 16.7
        }
      ],
      "exceptionStages": [
        {
          "stage": "身份验证",
          "count": 50
        },
        {
          "stage": "风险评估",
          "count": 30
        },
        {
          "stage": "签署协议",
          "count": 10
        }
      ],
      "exceptionHandling": [
        {
          "type": "身份验证失败",
          "avgHandleTime": 24,
          "maxHandleTime": 72,
          "successRate": 80
        },
        {
          "type": "资料不完整",
          "avgHandleTime": 12,
          "maxHandleTime": 48,
          "successRate": 90
        }
      ]
    }
  }
  ```

### 2.6 IPO主题

#### 2.6.1 获取IPO客户认购情况

- **路径**：`/api/ipo/subscription`
- **方法**：GET
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "totalUsers": 1000,
      "userDistribution": [
        {
          "type": "PI用户",
          "count": 400,
          "percentage": 40
        },
        {
          "type": "普通用户",
          "count": 600,
          "percentage": 60
        }
      ],
      "subscriptionAmount": [
        {
          "project": "项目A",
          "piAmount": 5000000,
          "regularAmount": 3000000,
          "totalAmount": 8000000
        },
        {
          "project": "项目B",
          "piAmount": 3000000,
          "regularAmount": 2000000,
          "totalAmount": 5000000
        }
      ],
      "subscriptionMultiples": [
        {
          "project": "项目A",
          "multiple": 10
        },
        {
          "project": "项目B",
          "multiple": 8
        },
        {
          "project": "项目C",
          "multiple": 12
        }
      ]
    }
  }
  ```

#### 2.6.2 获取IPO相关分析

- **路径**：`/api/ipo/analysis`
- **方法**：GET
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "projectPerformance": [
        {
          "project": "项目A",
          "priceChange": 20,
          "avgReturn": 15
        },
        {
          "project": "项目B",
          "priceChange": -5,
          "avgReturn": -3
        },
        {
          "project": "项目C",
          "priceChange": 30,
          "avgReturn": 25
        }
      ],
      "influenceFactors": [
        {
          "factor": "市场环境",
          "correlation": 0.8
        },
        {
          "factor": "公司基本面",
          "correlation": 0.9
        },
        {
          "factor": "发行价格",
          "correlation": 0.7
        }
      ],
      "competitorComparison": [
        {
          "competitor": "竞争对手A",
          "subscriptionUsers": 800,
          "subscriptionAmount": 6000000,
          "multiple": 8
        },
        {
          "competitor": "竞争对手B",
          "subscriptionUsers": 1200,
          "subscriptionAmount": 9000000,
          "multiple": 11
        }
      ]
    }
  }
  ```

### 2.7 融资主题

#### 2.7.1 获取客户融资情况

- **路径**：`/api/finance/customer`
- **方法**：GET
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "totalUsers": 500,
      "userDistribution": [
        {
          "type": "PI用户",
          "count": 200,
          "percentage": 40
        },
        {
          "type": "普通用户",
          "count": 300,
          "percentage": 60
        }
      ],
      "financingSize": {
        "avgSize": 100000,
        "medianSize": 80000,
        "p25Size": 50000,
        "p75Size": 150000,
        "maxSize": 500000
      },
      "financingPurpose": [
        {
          "purpose": "股票交易",
          "percentage": 70
        },
        {
          "purpose": "基金投资",
          "percentage": 20
        },
        {
          "purpose": "其他",
          "percentage": 10
        }
      ],
      "riskAssessment": [
        {
          "type": "PI用户",
          "defaultRate": 0.5,
          "overdueRate": 1
        },
        {
          "type": "普通用户",
          "defaultRate": 2,
          "overdueRate": 3
        }
      ]
    }
  }
  ```

#### 2.7.2 获取热门股票融资情况

- **路径**：`/api/finance/stock`
- **方法**：GET
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "hotStocks": [
        {
          "stock": "腾讯控股",
          "financingBalance": 100000000,
          "marketCapRatio": 5
        },
        {
          "stock": "阿里巴巴",
          "financingBalance": 80000000,
          "marketCapRatio": 4
        },
        {
          "stock": "美团",
          "financingBalance": 60000000,
          "marketCapRatio": 3
        }
      ],
      "financingTrend": [
        {
          "date": "2023-03-01",
          "buyAmount": 5000000,
          "sellAmount": 3000000
        },
        {
          "date": "2023-03-02",
          "buyAmount": 6000000,
          "sellAmount": 4000000
        },
        {
          "date": "2023-03-03",
          "buyAmount": 7000000,
          "sellAmount": 5000000
        }
      ],
      "marginRatio": [
        {
          "stock": "腾讯控股",
          "ratio": 2,
          "priceVolatility": 2
        },
        {
          "stock": "阿里巴巴",
          "ratio": 1.8,
          "priceVolatility": 2.5
        }
      ]
    }
  }
  ```

## 3. 数据钻取接口

### 3.1 获取明细数据

- **路径**：`/api/drilldown/detail`
- **方法**：GET
- **参数**：
  - `type`：数据类型（customer/trade/ipo/finance）
  - `id`：数据ID
  - `page`：页码（默认1）
  - `pageSize`：每页条数（默认10）
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "total": 100,
      "page": 1,
      "pageSize": 10,
      "list": [
        {
          "id": "1",
          "customerId": "C001",
          "name": "张三",
          "age": 30,
          "region": "北京",
          "asset": 100000,
          "tradeAmount": 50000
        },
        {
          "id": "2",
          "customerId": "C002",
          "name": "李四",
          "age": 35,
          "region": "上海",
          "asset": 150000,
          "tradeAmount": 80000
        }
      ]
    }
  }
  ```

### 3.2 获取历史趋势数据

- **路径**：`/api/drilldown/trend`
- **方法**：GET
- **参数**：
  - `type`：数据类型（indicator/revenue/customer/trade）
  - `id`：数据ID
  - `period`：时间范围（1m/3m/6m/1y）
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "type": "indicator",
      "name": "活跃用户数",
      "trend": [
        {
          "date": "2023-01",
          "value": 2800
        },
        {
          "date": "2023-02",
          "value": 2900
        },
        {
          "date": "2023-03",
          "value": 3000
        }
      ]
    }
  }
  ```

## 4. 认证接口

### 4.1 用户登录

- **路径**：`/api/auth/login`
- **方法**：POST
- **参数**：
  - `username`：用户名
  - `password`：密码
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "user": {
        "id": "U001",
        "username": "admin",
        "role": "admin"
      }
    }
  }
  ```

### 4.2 获取用户信息

- **路径**：`/api/auth/info`
- **方法**：GET
- **响应**：
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "id": "U001",
      "username": "admin",
      "role": "admin",
      "permissions": [
        "dashboard:view",
        "customer:view",
        "trade:view"
      ]
    }
  }
  ```

## 5. 总结

本API接口设计文档严格遵循RESTful API设计规范，为证券公司港美股经纪业务数据看板提供了完整的接口定义。接口覆盖了所有功能模块，包括关键指标概览、客户分析、交易分析、PI用户分析、开户主题、IPO主题和融资主题等。每个接口都定义了详细的路径、方法、参数和响应格式，确保前端与后端的顺畅通信。

在实施过程中，应根据实际业务需求和技术条件，对接口进行适当调整和优化，以达到最佳的系统性能和用户体验。