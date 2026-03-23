package repository

import (
	"dashboard-data/model"
)

// DashboardRepository 仪表盘数据仓库接口
type DashboardRepository interface {
	GetCoreIndicators(period string) (*model.CoreIndicators, error)
	GetPerformance(period string) (*model.Performance, error)
	GetCustomerDistribution(dimension string) (*model.CustomerDistribution, error)
	GetCustomerBehavior(behaviorType string) (*model.CustomerBehavior, error)
	GetStockTrade(tradeType string) (*model.StockTrade, error)
	GetFundTrade(tradeType string) (*model.FundTrade, error)
	GetPIBasic() (*model.PIBasic, error)
	GetPIBehavior(behaviorType string) (*model.PIBehavior, error)
	GetAccountProcess() (*model.AccountProcess, error)
	GetAccountException() (*model.AccountException, error)
	GetIPOSubscription() (*model.IPOSubscription, error)
	GetIPOAnaalysis() (*model.IPOAnalysis, error)
	GetFinanceCustomer() (*model.FinanceCustomer, error)
	GetFinanceStock() (*model.FinanceStock, error)
	GetDrilldownDetail(dataType, id string, page, pageSize int) (*model.DrilldownDetail, error)
	GetDrilldownTrend(dataType, id, period string) (*model.DrilldownTrend, error)
	Login(username, password string) (*model.LoginResponse, error)
	GetAuthInfo() (*model.User, error)
}

// DashboardRepositoryImpl 仪表盘数据仓库实现
type DashboardRepositoryImpl struct {}

// NewDashboardRepository 创建仪表盘数据仓库实例
func NewDashboardRepository() DashboardRepository {
	return &DashboardRepositoryImpl{}
}

// GetCoreIndicators 获取核心指标（mock数据）
func (r *DashboardRepositoryImpl) GetCoreIndicators(period string) (*model.CoreIndicators, error) {
	// 这里可以根据 period 参数返回不同的数据
	// 暂时返回固定的 mock 数据
	return &model.CoreIndicators{
		RegisteredUsers: model.Indicator{
			Value:      10000,
			Change:     500,
			ChangeRate: 5.26,
		},
		AccountCount: model.Indicator{
			Value:      8000,
			Change:     300,
			ChangeRate: 3.95,
		},
		ActiveUsers: model.Indicator{
			Value:      3000,
			Change:     -200,
			ChangeRate: -6.25,
		},
		DepositUsers: model.Indicator{
			Value:      2000,
			Change:     150,
			ChangeRate: 8.11,
		},
		DepositAmount: model.Indicator{
			Value:      50000000,
			Change:     5000000,
			ChangeRate: 11.11,
		},
		StockTradeUsers: model.Indicator{
			Value:      1500,
			Change:     100,
			ChangeRate: 7.14,
		},
		StockTradeAmount: model.Indicator{
			Value:      30000000,
			Change:     2000000,
			ChangeRate: 7.14,
		},
		FundTradeUsers: model.Indicator{
			Value:      800,
			Change:     50,
			ChangeRate: 6.67,
		},
		FundTradeAmount: model.Indicator{
			Value:      10000000,
			Change:     500000,
			ChangeRate: 5.26,
		},
	}, nil
}

// GetPerformance 获取业绩表现（mock数据）
func (r *DashboardRepositoryImpl) GetPerformance(period string) (*model.Performance, error) {
	// 这里可以根据 period 参数返回不同的数据
	// 暂时返回固定的 mock 数据
	return &model.Performance{
		StockRevenue: 2000000,
		FundRevenue:  500000,
		RevenueStructure: []model.RevenueItem{
			{Name: "佣金", Value: 1200000},
			{Name: "平台费", Value: 400000},
			{Name: "利息", Value: 300000},
			{Name: "IPO手续费", Value: 100000},
		},
		RevenueTrend: []model.RevenueTrend{
			{Month: "2023-01", Stock: 1800000, Fund: 450000},
			{Month: "2023-02", Stock: 1900000, Fund: 480000},
			{Month: "2023-03", Stock: 2000000, Fund: 500000},
		},
		RevenueChange: model.RevenueChange{
			Yoy: 15.38,
			Mom: 5.26,
		},
		DepositNetGrowth: 5000000,
		DepositTrend: []model.DepositTrend{
			{Month: "2023-01", Value: 4000000},
			{Month: "2023-02", Value: 4500000},
			{Month: "2023-03", Value: 5000000},
		},
		DepositWithdrawalRatio: 1.5,
	}, nil
}

// GetCustomerDistribution 获取客户分布（mock数据）
func (r *DashboardRepositoryImpl) GetCustomerDistribution(dimension string) (*model.CustomerDistribution, error) {
	return &model.CustomerDistribution{
		Dimension: dimension,
		Distribution: []model.CustomerDistributionItem{
			{Name: "18-25岁", Value: 1000, TotalAsset: 5000000, AvgAsset: 5000, AvgTradeAmount: 2000},
			{Name: "26-35岁", Value: 3000, TotalAsset: 30000000, AvgAsset: 10000, AvgTradeAmount: 5000},
			{Name: "36-45岁", Value: 2500, TotalAsset: 37500000, AvgAsset: 15000, AvgTradeAmount: 8000},
			{Name: "46-55岁", Value: 1000, TotalAsset: 20000000, AvgAsset: 20000, AvgTradeAmount: 10000},
			{Name: "55岁以上", Value: 500, TotalAsset: 15000000, AvgAsset: 30000, AvgTradeAmount: 15000},
		},
	}, nil
}

// GetCustomerBehavior 获取客户行为（mock数据）
func (r *DashboardRepositoryImpl) GetCustomerBehavior(behaviorType string) (*model.CustomerBehavior, error) {
	return &model.CustomerBehavior{
		Type: behaviorType,
		Trend: []model.CustomerBehaviorTrend{
			{Quarter: "2022-Q1", RetentionRate: 85, ChurnRate: 15},
			{Quarter: "2022-Q2", RetentionRate: 82, ChurnRate: 18},
			{Quarter: "2022-Q3", RetentionRate: 80, ChurnRate: 20},
			{Quarter: "2022-Q4", RetentionRate: 78, ChurnRate: 22},
			{Quarter: "2023-Q1", RetentionRate: 80, ChurnRate: 20},
		},
		IndustryAverage:  75,
		CompetitorAverage: 78,
	}, nil
}

// GetStockTrade 获取股票交易（mock数据）
func (r *DashboardRepositoryImpl) GetStockTrade(tradeType string) (*model.StockTrade, error) {
	return &model.StockTrade{
		Type: tradeType,
		ExecutionTime: []model.ExecutionTime{
			{Time: "9:30-10:30", AvgTime: 0.5},
			{Time: "10:30-11:30", AvgTime: 0.3},
			{Time: "13:30-14:30", AvgTime: 0.4},
			{Time: "14:30-15:30", AvgTime: 0.6},
		},
		MarketComparison: []model.MarketComparison{
			{Market: "港股", AvgTime: 0.4},
			{Market: "美股", AvgTime: 0.5},
		},
	}, nil
}

// GetFundTrade 获取基金交易（mock数据）
func (r *DashboardRepositoryImpl) GetFundTrade(tradeType string) (*model.FundTrade, error) {
	return &model.FundTrade{
		Type: tradeType,
		SalesByType: []model.SalesByType{
			{Type: "股票型", Amount: 5000000, Count: 200},
			{Type: "债券型", Amount: 3000000, Count: 150},
			{Type: "混合型", Amount: 1500000, Count: 100},
			{Type: "货币型", Amount: 500000, Count: 50},
		},
	}, nil
}

// GetPIBasic 获取PI用户基本信息（mock数据）
func (r *DashboardRepositoryImpl) GetPIBasic() (*model.PIBasic, error) {
	return &model.PIBasic{
		TotalUsers: 500,
		Change:     50,
		ChangeRate: 11.11,
		AssetDistribution: []model.AssetDistribution{
			{Name: "港股", Value: 60},
			{Name: "美股", Value: 30},
			{Name: "基金", Value: 10},
		},
		TotalAsset: 50000000,
		Demographics: []model.Demographic{
			{Job: "金融行业", Percentage: 40, AvgAsset: 150000, AvgTradeFrequency: 10, AvgFundCount: 5},
			{Job: "科技行业", Percentage: 30, AvgAsset: 120000, AvgTradeFrequency: 8, AvgFundCount: 4},
			{Job: "其他行业", Percentage: 30, AvgAsset: 80000, AvgTradeFrequency: 5, AvgFundCount: 3},
		},
	}, nil
}

// GetPIBehavior 获取PI用户行为（mock数据）
func (r *DashboardRepositoryImpl) GetPIBehavior(behaviorType string) (*model.PIBehavior, error) {
	return &model.PIBehavior{
		Type:              behaviorType,
		AvgTradeAmount:    20000,
		AvgTradeFrequency: 10,
		AvgHoldingPeriod:  30,
		MoneyFlow: []model.MoneyFlow{
			{From: "科技行业", To: "金融行业", Value: 5000000},
			{From: "金融行业", To: "消费行业", Value: 3000000},
		},
	}, nil
}

// GetAccountProcess 获取开户流程（mock数据）
func (r *DashboardRepositoryImpl) GetAccountProcess() (*model.AccountProcess, error) {
	return &model.AccountProcess{
		Funnel: []model.FunnelStage{
			{Stage: "注册", Count: 1000, Conversion: 100},
			{Stage: "身份验证", Count: 800, Conversion: 80},
			{Stage: "风险评估", Count: 700, Conversion: 87.5},
			{Stage: "签署协议", Count: 600, Conversion: 85.7},
			{Stage: "开户成功", Count: 500, Conversion: 83.3},
		},
		StageTime: []model.StageTime{
			{Stage: "注册", AvgTime: 2, MedianTime: 1.5, P25Time: 1, P75Time: 2.5, MaxTime: 5},
			{Stage: "身份验证", AvgTime: 5, MedianTime: 4, P25Time: 3, P75Time: 6, MaxTime: 10},
		},
	}, nil
}

// GetAccountException 获取开户异常（mock数据）
func (r *DashboardRepositoryImpl) GetAccountException() (*model.AccountException, error) {
	return &model.AccountException{
		ExceptionTypes: []model.ExceptionType{
			{Type: "身份验证失败", Count: 50, Percentage: 41.7},
			{Type: "资料不完整", Count: 30, Percentage: 25},
			{Type: "风险评估不通过", Count: 20, Percentage: 16.7},
			{Type: "其他", Count: 20, Percentage: 16.7},
		},
		ExceptionStages: []model.ExceptionStage{
			{Stage: "身份验证", Count: 50},
			{Stage: "风险评估", Count: 30},
			{Stage: "签署协议", Count: 10},
		},
		ExceptionHandling: []model.ExceptionHandling{
			{Type: "身份验证失败", AvgHandleTime: 24, MaxHandleTime: 72, SuccessRate: 80},
			{Type: "资料不完整", AvgHandleTime: 12, MaxHandleTime: 48, SuccessRate: 90},
		},
	}, nil
}

// GetIPOSubscription 获取IPO认购情况（mock数据）
func (r *DashboardRepositoryImpl) GetIPOSubscription() (*model.IPOSubscription, error) {
	return &model.IPOSubscription{
		TotalUsers: 1000,
		UserDistribution: []model.UserDistribution{
			{Type: "PI用户", Count: 400, Percentage: 40},
			{Type: "普通用户", Count: 600, Percentage: 60},
		},
		SubscriptionAmount: []model.SubscriptionAmount{
			{Project: "项目A", PiAmount: 5000000, RegularAmount: 3000000, TotalAmount: 8000000},
			{Project: "项目B", PiAmount: 3000000, RegularAmount: 2000000, TotalAmount: 5000000},
		},
		SubscriptionMultiples: []model.SubscriptionMultiple{
			{Project: "项目A", Multiple: 10},
			{Project: "项目B", Multiple: 8},
			{Project: "项目C", Multiple: 12},
		},
	}, nil
}

// GetIPOAnaalysis 获取IPO分析（mock数据）
func (r *DashboardRepositoryImpl) GetIPOAnaalysis() (*model.IPOAnalysis, error) {
	return &model.IPOAnalysis{
		ProjectPerformance: []model.ProjectPerformance{
			{Project: "项目A", PriceChange: 20, AvgReturn: 15},
			{Project: "项目B", PriceChange: -5, AvgReturn: -3},
			{Project: "项目C", PriceChange: 30, AvgReturn: 25},
		},
		InfluenceFactors: []model.InfluenceFactor{
			{Factor: "市场环境", Correlation: 0.8},
			{Factor: "公司基本面", Correlation: 0.9},
			{Factor: "发行价格", Correlation: 0.7},
		},
		CompetitorComparison: []model.CompetitorComparison{
			{Competitor: "竞争对手A", SubscriptionUsers: 800, SubscriptionAmount: 6000000, Multiple: 8},
			{Competitor: "竞争对手B", SubscriptionUsers: 1200, SubscriptionAmount: 9000000, Multiple: 11},
		},
	}, nil
}

// GetFinanceCustomer 获取客户融资情况（mock数据）
func (r *DashboardRepositoryImpl) GetFinanceCustomer() (*model.FinanceCustomer, error) {
	return &model.FinanceCustomer{
		TotalUsers: 500,
		UserDistribution: []model.UserDistribution{
			{Type: "PI用户", Count: 200, Percentage: 40},
			{Type: "普通用户", Count: 300, Percentage: 60},
		},
		FinancingSize: model.FinancingSize{
			AvgSize:    100000,
			MedianSize: 80000,
			P25Size:    50000,
			P75Size:    150000,
			MaxSize:    500000,
		},
		FinancingPurpose: []model.FinancingPurpose{
			{Purpose: "股票交易", Percentage: 70},
			{Purpose: "基金投资", Percentage: 20},
			{Purpose: "其他", Percentage: 10},
		},
		RiskAssessment: []model.RiskAssessment{
			{Type: "PI用户", DefaultRate: 0.5, OverdueRate: 1},
			{Type: "普通用户", DefaultRate: 2, OverdueRate: 3},
		},
	}, nil
}

// GetFinanceStock 获取热门股票融资情况（mock数据）
func (r *DashboardRepositoryImpl) GetFinanceStock() (*model.FinanceStock, error) {
	return &model.FinanceStock{
		HotStocks: []model.HotStock{
			{Stock: "腾讯控股", FinancingBalance: 100000000, MarketCapRatio: 5},
			{Stock: "阿里巴巴", FinancingBalance: 80000000, MarketCapRatio: 4},
			{Stock: "美团", FinancingBalance: 60000000, MarketCapRatio: 3},
		},
		FinancingTrend: []model.FinancingTrend{
			{Date: "2023-03-01", BuyAmount: 5000000, SellAmount: 3000000},
			{Date: "2023-03-02", BuyAmount: 6000000, SellAmount: 4000000},
			{Date: "2023-03-03", BuyAmount: 7000000, SellAmount: 5000000},
		},
		MarginRatio: []model.MarginRatio{
			{Stock: "腾讯控股", Ratio: 2, PriceVolatility: 2},
			{Stock: "阿里巴巴", Ratio: 1.8, PriceVolatility: 2.5},
		},
	}, nil
}

// GetDrilldownDetail 获取明细数据（mock数据）
func (r *DashboardRepositoryImpl) GetDrilldownDetail(dataType, id string, page, pageSize int) (*model.DrilldownDetail, error) {
	return &model.DrilldownDetail{
		Total:    100,
		Page:     page,
		PageSize: pageSize,
		List: []model.DrilldownItem{
			{ID: "1", CustomerId: "C001", Name: "张三", Age: 30, Region: "北京", Asset: 100000, TradeAmount: 50000},
			{ID: "2", CustomerId: "C002", Name: "李四", Age: 35, Region: "上海", Asset: 150000, TradeAmount: 80000},
		},
	}, nil
}

// GetDrilldownTrend 获取历史趋势数据（mock数据）
func (r *DashboardRepositoryImpl) GetDrilldownTrend(dataType, id, period string) (*model.DrilldownTrend, error) {
	return &model.DrilldownTrend{
		Type: dataType,
		Name: "活跃用户数",
		Trend: []model.TrendItem{
			{Date: "2023-01", Value: 2800},
			{Date: "2023-02", Value: 2900},
			{Date: "2023-03", Value: 3000},
		},
	}, nil
}

// Login 用户登录（mock数据）
func (r *DashboardRepositoryImpl) Login(username, password string) (*model.LoginResponse, error) {
	return &model.LoginResponse{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
		User: model.User{
			ID:       "U001",
			Username: "admin",
			Role:     "admin",
		},
	}, nil
}

// GetAuthInfo 获取用户信息（mock数据）
func (r *DashboardRepositoryImpl) GetAuthInfo() (*model.User, error) {
	return &model.User{
		ID:          "U001",
		Username:    "admin",
		Role:        "admin",
		Permissions: []string{"dashboard:view", "customer:view", "trade:view"},
	}, nil
}
