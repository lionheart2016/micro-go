package model

// CoreIndicators 核心指标
type CoreIndicators struct {
	RegisteredUsers  Indicator `json:"registeredUsers"`
	AccountCount     Indicator `json:"accountCount"`
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
	StockRevenue           int            `json:"stockRevenue"`
	FundRevenue            int            `json:"fundRevenue"`
	RevenueStructure       []RevenueItem  `json:"revenueStructure"`
	RevenueTrend           []RevenueTrend `json:"revenueTrend"`
	RevenueChange          RevenueChange  `json:"revenueChange"`
	DepositNetGrowth       int            `json:"depositNetGrowth"`
	DepositTrend           []DepositTrend `json:"depositTrend"`
	DepositWithdrawalRatio float64        `json:"depositWithdrawalRatio"`
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

// CustomerDistribution 客户分布
type CustomerDistribution struct {
	Dimension    string                     `json:"dimension"`
	Distribution []CustomerDistributionItem `json:"distribution"`
}

// CustomerDistributionItem 客户分布项
type CustomerDistributionItem struct {
	Name           string `json:"name"`
	Value          int    `json:"value"`
	TotalAsset     int    `json:"totalAsset"`
	AvgAsset       int    `json:"avgAsset"`
	AvgTradeAmount int    `json:"avgTradeAmount"`
}

// CustomerBehavior 客户行为
type CustomerBehavior struct {
	Type              string                  `json:"type"`
	Trend             []CustomerBehaviorTrend `json:"trend"`
	IndustryAverage   float64                 `json:"industryAverage"`
	CompetitorAverage float64                 `json:"competitorAverage"`
}

// CustomerBehaviorTrend 客户行为趋势
type CustomerBehaviorTrend struct {
	Quarter       string  `json:"quarter"`
	RetentionRate float64 `json:"retentionRate"`
	ChurnRate     float64 `json:"churnRate"`
}

// StockTrade 股票交易
type StockTrade struct {
	Type             string             `json:"type"`
	ExecutionTime    []ExecutionTime    `json:"executionTime"`
	MarketComparison []MarketComparison `json:"marketComparison"`
}

// ExecutionTime 执行时间
type ExecutionTime struct {
	Time    string  `json:"time"`
	AvgTime float64 `json:"avgTime"`
}

// MarketComparison 市场比较
type MarketComparison struct {
	Market  string  `json:"market"`
	AvgTime float64 `json:"avgTime"`
}

// FundTrade 基金交易
type FundTrade struct {
	Type        string        `json:"type"`
	SalesByType []SalesByType `json:"salesByType"`
}

// SalesByType 按类型销售
type SalesByType struct {
	Type   string `json:"type"`
	Amount int    `json:"amount"`
	Count  int    `json:"count"`
}

// PIBasic PI用户基本信息
type PIBasic struct {
	TotalUsers        int                 `json:"totalUsers"`
	Change            int                 `json:"change"`
	ChangeRate        float64             `json:"changeRate"`
	AssetDistribution []AssetDistribution `json:"assetDistribution"`
	TotalAsset        int                 `json:"totalAsset"`
	Demographics      []Demographic       `json:"demographics"`
}

// AssetDistribution 资产分布
type AssetDistribution struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// Demographic 人口统计
type Demographic struct {
	Job               string  `json:"job"`
	Percentage        float64 `json:"percentage"`
	AvgAsset          int     `json:"avgAsset"`
	AvgTradeFrequency int     `json:"avgTradeFrequency"`
	AvgFundCount      int     `json:"avgFundCount"`
}

// PIBehavior PI用户行为特征
type PIBehavior struct {
	Type              string      `json:"type"`
	AvgTradeAmount    int         `json:"avgTradeAmount"`
	AvgTradeFrequency int         `json:"avgTradeFrequency"`
	AvgHoldingPeriod  int         `json:"avgHoldingPeriod"`
	MoneyFlow         []MoneyFlow `json:"moneyFlow"`
}

// MoneyFlow 资金流向
type MoneyFlow struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}

// AccountProcess 开户流程
type AccountProcess struct {
	Funnel    []FunnelStage `json:"funnel"`
	StageTime []StageTime   `json:"stageTime"`
}

// FunnelStage 漏斗阶段
type FunnelStage struct {
	Stage      string  `json:"stage"`
	Count      int     `json:"count"`
	Conversion float64 `json:"conversion"`
}

// StageTime 阶段时间
type StageTime struct {
	Stage      string  `json:"stage"`
	AvgTime    float64 `json:"avgTime"`
	MedianTime float64 `json:"medianTime"`
	P25Time    float64 `json:"p25Time"`
	P75Time    float64 `json:"p75Time"`
	MaxTime    float64 `json:"maxTime"`
}

// AccountException 开户异常
type AccountException struct {
	ExceptionTypes    []ExceptionType     `json:"exceptionTypes"`
	ExceptionStages   []ExceptionStage    `json:"exceptionStages"`
	ExceptionHandling []ExceptionHandling `json:"exceptionHandling"`
}

// ExceptionType 异常类型
type ExceptionType struct {
	Type       string  `json:"type"`
	Count      int     `json:"count"`
	Percentage float64 `json:"percentage"`
}

// ExceptionStage 异常阶段
type ExceptionStage struct {
	Stage string `json:"stage"`
	Count int    `json:"count"`
}

// ExceptionHandling 异常处理
type ExceptionHandling struct {
	Type          string  `json:"type"`
	AvgHandleTime float64 `json:"avgHandleTime"`
	MaxHandleTime float64 `json:"maxHandleTime"`
	SuccessRate   float64 `json:"successRate"`
}

// IPOSubscription IPO客户认购情况
type IPOSubscription struct {
	TotalUsers            int                    `json:"totalUsers"`
	UserDistribution      []UserDistribution     `json:"userDistribution"`
	SubscriptionAmount    []SubscriptionAmount   `json:"subscriptionAmount"`
	SubscriptionMultiples []SubscriptionMultiple `json:"subscriptionMultiples"`
}

// UserDistribution 用户分布
type UserDistribution struct {
	Type       string  `json:"type"`
	Count      int     `json:"count"`
	Percentage float64 `json:"percentage"`
}

// SubscriptionAmount 认购金额
type SubscriptionAmount struct {
	Project       string `json:"project"`
	PiAmount      int    `json:"piAmount"`
	RegularAmount int    `json:"regularAmount"`
	TotalAmount   int    `json:"totalAmount"`
}

// SubscriptionMultiple 认购倍数
type SubscriptionMultiple struct {
	Project  string  `json:"project"`
	Multiple float64 `json:"multiple"`
}

// IPOAAnalysis IPO相关分析
type IPOAAnalysis struct {
	ProjectPerformance   []ProjectPerformance   `json:"projectPerformance"`
	InfluenceFactors     []InfluenceFactor      `json:"influenceFactors"`
	CompetitorComparison []CompetitorComparison `json:"competitorComparison"`
}

// ProjectPerformance 项目表现
type ProjectPerformance struct {
	Project     string  `json:"project"`
	PriceChange float64 `json:"priceChange"`
	AvgReturn   float64 `json:"avgReturn"`
}

// InfluenceFactor 影响因素
type InfluenceFactor struct {
	Factor      string  `json:"factor"`
	Correlation float64 `json:"correlation"`
}

// CompetitorComparison 竞争对手比较
type CompetitorComparison struct {
	Competitor         string  `json:"competitor"`
	SubscriptionUsers  int     `json:"subscriptionUsers"`
	SubscriptionAmount int     `json:"subscriptionAmount"`
	Multiple           float64 `json:"multiple"`
}

// FinanceCustomer 客户融资情况
type FinanceCustomer struct {
	TotalUsers       int                `json:"totalUsers"`
	UserDistribution []UserDistribution `json:"userDistribution"`
	FinancingSize    FinancingSize      `json:"financingSize"`
	FinancingPurpose []FinancingPurpose `json:"financingPurpose"`
	RiskAssessment   []RiskAssessment   `json:"riskAssessment"`
}

// FinancingSize 融资规模
type FinancingSize struct {
	AvgSize    float64 `json:"avgSize"`
	MedianSize float64 `json:"medianSize"`
	P25Size    float64 `json:"p25Size"`
	P75Size    float64 `json:"p75Size"`
	MaxSize    float64 `json:"maxSize"`
}

// FinancingPurpose 融资目的
type FinancingPurpose struct {
	Purpose    string  `json:"purpose"`
	Percentage float64 `json:"percentage"`
}

// RiskAssessment 风险评估
type RiskAssessment struct {
	Type        string  `json:"type"`
	DefaultRate float64 `json:"defaultRate"`
	OverdueRate float64 `json:"overdueRate"`
}

// FinanceStock 热门股票融资情况
type FinanceStock struct {
	HotStocks      []HotStock       `json:"hotStocks"`
	FinancingTrend []FinancingTrend `json:"financingTrend"`
	MarginRatio    []MarginRatio    `json:"marginRatio"`
}

// HotStock 热门股票
type HotStock struct {
	Stock            string  `json:"stock"`
	FinancingBalance int     `json:"financingBalance"`
	MarketCapRatio   float64 `json:"marketCapRatio"`
}

// FinancingTrend 融资趋势
type FinancingTrend struct {
	Date       string `json:"date"`
	BuyAmount  int    `json:"buyAmount"`
	SellAmount int    `json:"sellAmount"`
}

// MarginRatio 保证金比例
type MarginRatio struct {
	Stock           string  `json:"stock"`
	Ratio           float64 `json:"ratio"`
	PriceVolatility float64 `json:"priceVolatility"`
}

// DrilldownDetail 明细数据
type DrilldownDetail struct {
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
	List     []DrilldownItem `json:"list"`
}

// DrilldownItem 明细项
type DrilldownItem struct {
	ID          string `json:"id"`
	CustomerId  string `json:"customerId"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Region      string `json:"region"`
	Asset       int    `json:"asset"`
	TradeAmount int    `json:"tradeAmount"`
}

// DrilldownTrend 历史趋势数据
type DrilldownTrend struct {
	Type  string      `json:"type"`
	Name  string      `json:"name"`
	Trend []TrendItem `json:"trend"`
}

// TrendItem 趋势项
type TrendItem struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// User 用户信息
type User struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions,omitempty"`
}

// Response 通用响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
