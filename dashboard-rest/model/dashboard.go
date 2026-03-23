package model

// CoreIndicators 核心指标
type CoreIndicators struct {
	RegisteredUsers  Indicator `json:"registeredUsers"`
	AccountCount     Indicator `json:"开户人数"`
	ActiveUsers      Indicator `json:"activeUsers"`
	DepositUsers     Indicator `json:"depositUsers"`
	DepositAmount    Indicator `json:"depositAmount"`
	StockTradeUsers  Indicator `json:"stockTradeUsers"`
	StockTradeAmount Indicator `json:"stockTradeAmount"`
	FundTradeUsers   Indicator `json:"fundTradeUsers"`
	FundTradeAmount  Indicator `json:"fundTradeAmount"`
}

// Indicator 指标结构
type Indicator struct {
	Value      int     `json:"value"`
	Change     int     `json:"change"`
	ChangeRate float64 `json:"changeRate"`
}

// Performance 业绩表现
type Performance struct {
	StockRevenue           int              `json:"stockRevenue"`
	FundRevenue            int              `json:"fundRevenue"`
	RevenueStructure       []RevenueItem    `json:"revenueStructure"`
	RevenueTrend           []RevenueTrend   `json:"revenueTrend"`
	RevenueChange          RevenueChange    `json:"revenueChange"`
	DepositNetGrowth       int              `json:"depositNetGrowth"`
	DepositTrend           []DepositTrend   `json:"depositTrend"`
	DepositWithdrawalRatio float64          `json:"depositWithdrawalRatio"`
}

// RevenueItem 收入项
type RevenueItem struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// RevenueTrend 收入趋势
type RevenueTrend struct {
	Month string `json:"month"`
	Stock int    `json:"stock"`
	Fund  int    `json:"fund"`
}

// RevenueChange 收入变化
type RevenueChange struct {
	Yoy float64 `json:"yoy"`
	Mom float64 `json:"mom"`
}

// DepositTrend 入金趋势
type DepositTrend struct {
	Month string `json:"month"`
	Value int    `json:"value"`
}

// Response 通用响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
