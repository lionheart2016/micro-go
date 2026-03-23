package client

import (
	"context"
	"dashboard-rest/model"
	"dashboard-rest/pkg/logger"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DashboardClient struct {
	logger  *logger.Logger
	conn    *grpc.ClientConn
}

func New(grpcHost string, grpcPort int, logger *logger.Logger) (*DashboardClient, error) {
	addr := fmt.Sprintf("%s:%d", grpcHost, grpcPort)
	// 连接到gRPC服务器
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Failed to connect to gRPC server: %v", err)
		return nil, err
	}

	return &DashboardClient{
		logger:  logger,
		conn:    conn,
	}, nil
}

func (c *DashboardClient) GetCoreIndicators(ctx context.Context, period string) (*model.Response, error) {
	c.logger.Info("Getting core indicators")

	// 模拟数据
	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.CoreIndicators{
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
		},
	}, nil
}

func (c *DashboardClient) GetPerformance(ctx context.Context, period string) (*model.Response, error) {
	c.logger.Info("Getting performance")

	// 模拟数据
	revenueStructure := []model.RevenueItem{
		{Name: "佣金", Value: 1200000},
		{Name: "平台费", Value: 400000},
		{Name: "利息", Value: 300000},
		{Name: "IPO手续费", Value: 100000},
	}

	revenueTrend := []model.RevenueTrend{
		{Month: "2023-01", Stock: 1800000, Fund: 450000},
		{Month: "2023-02", Stock: 1900000, Fund: 480000},
		{Month: "2023-03", Stock: 2000000, Fund: 500000},
	}

	depositTrend := []model.DepositTrend{
		{Month: "2023-01", Value: 4000000},
		{Month: "2023-02", Value: 4500000},
		{Month: "2023-03", Value: 5000000},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.Performance{
			StockRevenue: 2000000,
			FundRevenue:  500000,
			RevenueStructure: revenueStructure,
			RevenueTrend:     revenueTrend,
			RevenueChange: model.RevenueChange{
				Yoy: 15.38,
				Mom: 5.26,
			},
			DepositNetGrowth:       5000000,
			DepositTrend:           depositTrend,
			DepositWithdrawalRatio: 1.5,
		},
	}, nil
}

// 客户分析相关方法
func (c *DashboardClient) GetCustomerDistribution(ctx context.Context, dimension string) (*model.Response, error) {
	c.logger.Info("Getting customer distribution")

	// 模拟数据
	distribution := []model.CustomerDistributionItem{
		{Name: "18-25岁", Value: 1000, TotalAsset: 5000000, AvgAsset: 5000, AvgTradeAmount: 2000},
		{Name: "26-35岁", Value: 3000, TotalAsset: 30000000, AvgAsset: 10000, AvgTradeAmount: 5000},
		{Name: "36-45岁", Value: 2500, TotalAsset: 37500000, AvgAsset: 15000, AvgTradeAmount: 8000},
		{Name: "46-55岁", Value: 1000, TotalAsset: 20000000, AvgAsset: 20000, AvgTradeAmount: 10000},
		{Name: "55岁以上", Value: 500, TotalAsset: 15000000, AvgAsset: 30000, AvgTradeAmount: 15000},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.CustomerDistribution{
			Dimension:   dimension,
			Distribution: distribution,
		},
	}, nil
}

func (c *DashboardClient) GetCustomerBehavior(ctx context.Context, behaviorType string) (*model.Response, error) {
	c.logger.Info("Getting customer behavior")

	// 模拟数据
	trend := []model.CustomerBehaviorTrend{
		{Quarter: "2022-Q1", RetentionRate: 85, ChurnRate: 15},
		{Quarter: "2022-Q2", RetentionRate: 82, ChurnRate: 18},
		{Quarter: "2022-Q3", RetentionRate: 80, ChurnRate: 20},
		{Quarter: "2022-Q4", RetentionRate: 78, ChurnRate: 22},
		{Quarter: "2023-Q1", RetentionRate: 80, ChurnRate: 20},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.CustomerBehavior{
			Type:             behaviorType,
			Trend:            trend,
			IndustryAverage:  75,
			CompetitorAverage: 78,
		},
	}, nil
}

// 交易分析相关方法
func (c *DashboardClient) GetStockTrade(ctx context.Context, tradeType string) (*model.Response, error) {
	c.logger.Info("Getting stock trade")

	// 模拟数据
	executionTime := []model.ExecutionTime{
		{Time: "9:30-10:30", AvgTime: 0.5},
		{Time: "10:30-11:30", AvgTime: 0.3},
		{Time: "13:30-14:30", AvgTime: 0.4},
		{Time: "14:30-15:30", AvgTime: 0.6},
	}

	marketComparison := []model.MarketComparison{
		{Market: "港股", AvgTime: 0.4},
		{Market: "美股", AvgTime: 0.5},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.StockTrade{
			Type:              tradeType,
			ExecutionTime:     executionTime,
			MarketComparison:  marketComparison,
		},
	}, nil
}

func (c *DashboardClient) GetFundTrade(ctx context.Context, tradeType string) (*model.Response, error) {
	c.logger.Info("Getting fund trade")

	// 模拟数据
	salesByType := []model.SalesByType{
		{Type: "股票型", Amount: 5000000, Count: 200},
		{Type: "债券型", Amount: 3000000, Count: 150},
		{Type: "混合型", Amount: 1500000, Count: 100},
		{Type: "货币型", Amount: 500000, Count: 50},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.FundTrade{
			Type:       tradeType,
			SalesByType: salesByType,
		},
	}, nil
}

// PI用户分析相关方法
func (c *DashboardClient) GetPIBasic(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting PI basic information")

	// 模拟数据
	assetDistribution := []model.AssetDistribution{
		{Name: "港股", Value: 60},
		{Name: "美股", Value: 30},
		{Name: "基金", Value: 10},
	}

	demographics := []model.Demographic{
		{Job: "金融行业", Percentage: 40, AvgAsset: 150000, AvgTradeFrequency: 10, AvgFundCount: 5},
		{Job: "科技行业", Percentage: 30, AvgAsset: 120000, AvgTradeFrequency: 8, AvgFundCount: 4},
		{Job: "其他行业", Percentage: 30, AvgAsset: 80000, AvgTradeFrequency: 5, AvgFundCount: 3},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.PIBasic{
			TotalUsers:       500,
			Change:           50,
			ChangeRate:       11.11,
			AssetDistribution: assetDistribution,
			TotalAsset:       50000000,
			Demographics:     demographics,
		},
	}, nil
}

func (c *DashboardClient) GetPIBehavior(ctx context.Context, behaviorType string) (*model.Response, error) {
	c.logger.Info("Getting PI behavior")

	// 模拟数据
	moneyFlow := []model.MoneyFlow{
		{From: "科技行业", To: "金融行业", Value: 5000000},
		{From: "金融行业", To: "消费行业", Value: 3000000},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.PIBehavior{
			Type:              behaviorType,
			AvgTradeAmount:    20000,
			AvgTradeFrequency: 10,
			AvgHoldingPeriod:  30,
			MoneyFlow:         moneyFlow,
		},
	}, nil
}

// 开户主题相关方法
func (c *DashboardClient) GetAccountProcess(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting account process")

	// 模拟数据
	funnel := []model.FunnelStage{
		{Stage: "注册", Count: 1000, Conversion: 100},
		{Stage: "身份验证", Count: 800, Conversion: 80},
		{Stage: "风险评估", Count: 700, Conversion: 87.5},
		{Stage: "签署协议", Count: 600, Conversion: 85.7},
		{Stage: "开户成功", Count: 500, Conversion: 83.3},
	}

	stageTime := []model.StageTime{
		{Stage: "注册", AvgTime: 2, MedianTime: 1.5, P25Time: 1, P75Time: 2.5, MaxTime: 5},
		{Stage: "身份验证", AvgTime: 5, MedianTime: 4, P25Time: 3, P75Time: 6, MaxTime: 10},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.AccountProcess{
			Funnel:     funnel,
			StageTime:  stageTime,
		},
	}, nil
}

func (c *DashboardClient) GetAccountException(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting account exception")

	// 模拟数据
	exceptionTypes := []model.ExceptionType{
		{Type: "身份验证失败", Count: 50, Percentage: 41.7},
		{Type: "资料不完整", Count: 30, Percentage: 25},
		{Type: "风险评估不通过", Count: 20, Percentage: 16.7},
		{Type: "其他", Count: 20, Percentage: 16.7},
	}

	exceptionStages := []model.ExceptionStage{
		{Stage: "身份验证", Count: 50},
		{Stage: "风险评估", Count: 30},
		{Stage: "签署协议", Count: 10},
	}

	exceptionHandling := []model.ExceptionHandling{
		{Type: "身份验证失败", AvgHandleTime: 24, MaxHandleTime: 72, SuccessRate: 80},
		{Type: "资料不完整", AvgHandleTime: 12, MaxHandleTime: 48, SuccessRate: 90},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.AccountException{
			ExceptionTypes:     exceptionTypes,
			ExceptionStages:    exceptionStages,
			ExceptionHandling:  exceptionHandling,
		},
	}, nil
}

// IPO主题相关方法
func (c *DashboardClient) GetIPOSubscription(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting IPO subscription")

	// 模拟数据
	userDistribution := []model.UserDistribution{
		{Type: "PI用户", Count: 400, Percentage: 40},
		{Type: "普通用户", Count: 600, Percentage: 60},
	}

	subscriptionAmount := []model.SubscriptionAmount{
		{Project: "项目A", PiAmount: 5000000, RegularAmount: 3000000, TotalAmount: 8000000},
		{Project: "项目B", PiAmount: 3000000, RegularAmount: 2000000, TotalAmount: 5000000},
	}

	subscriptionMultiples := []model.SubscriptionMultiple{
		{Project: "项目A", Multiple: 10},
		{Project: "项目B", Multiple: 8},
		{Project: "项目C", Multiple: 12},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.IPOSubscription{
			TotalUsers:         1000,
			UserDistribution:   userDistribution,
			SubscriptionAmount: subscriptionAmount,
			SubscriptionMultiples: subscriptionMultiples,
		},
	}, nil
}

func (c *DashboardClient) GetIPOAnaalysis(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting IPO analysis")

	// 模拟数据
	projectPerformance := []model.ProjectPerformance{
		{Project: "项目A", PriceChange: 20, AvgReturn: 15},
		{Project: "项目B", PriceChange: -5, AvgReturn: -3},
		{Project: "项目C", PriceChange: 30, AvgReturn: 25},
	}

	influenceFactors := []model.InfluenceFactor{
		{Factor: "市场环境", Correlation: 0.8},
		{Factor: "公司基本面", Correlation: 0.9},
		{Factor: "发行价格", Correlation: 0.7},
	}

	competitorComparison := []model.CompetitorComparison{
		{Competitor: "竞争对手A", SubscriptionUsers: 800, SubscriptionAmount: 6000000, Multiple: 8},
		{Competitor: "竞争对手B", SubscriptionUsers: 1200, SubscriptionAmount: 9000000, Multiple: 11},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.IPOAAnalysis{
			ProjectPerformance:    projectPerformance,
			InfluenceFactors:      influenceFactors,
			CompetitorComparison:  competitorComparison,
		},
	}, nil
}

// 融资主题相关方法
func (c *DashboardClient) GetFinanceCustomer(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting finance customer")

	// 模拟数据
	userDistribution := []model.UserDistribution{
		{Type: "PI用户", Count: 200, Percentage: 40},
		{Type: "普通用户", Count: 300, Percentage: 60},
	}

	financingPurpose := []model.FinancingPurpose{
		{Purpose: "股票交易", Percentage: 70},
		{Purpose: "基金投资", Percentage: 20},
		{Purpose: "其他", Percentage: 10},
	}

	riskAssessment := []model.RiskAssessment{
		{Type: "PI用户", DefaultRate: 0.5, OverdueRate: 1},
		{Type: "普通用户", DefaultRate: 2, OverdueRate: 3},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.FinanceCustomer{
			TotalUsers:       500,
			UserDistribution: userDistribution,
			FinancingSize: model.FinancingSize{
				AvgSize:    100000,
				MedianSize: 80000,
				P25Size:    50000,
				P75Size:    150000,
				MaxSize:    500000,
			},
			FinancingPurpose: financingPurpose,
			RiskAssessment:   riskAssessment,
		},
	}, nil
}

func (c *DashboardClient) GetFinanceStock(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting finance stock")

	// 模拟数据
	hotStocks := []model.HotStock{
		{Stock: "腾讯控股", FinancingBalance: 100000000, MarketCapRatio: 5},
		{Stock: "阿里巴巴", FinancingBalance: 80000000, MarketCapRatio: 4},
		{Stock: "美团", FinancingBalance: 60000000, MarketCapRatio: 3},
	}

	financingTrend := []model.FinancingTrend{
		{Date: "2023-03-01", BuyAmount: 5000000, SellAmount: 3000000},
		{Date: "2023-03-02", BuyAmount: 6000000, SellAmount: 4000000},
		{Date: "2023-03-03", BuyAmount: 7000000, SellAmount: 5000000},
	}

	marginRatio := []model.MarginRatio{
		{Stock: "腾讯控股", Ratio: 2, PriceVolatility: 2},
		{Stock: "阿里巴巴", Ratio: 1.8, PriceVolatility: 2.5},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.FinanceStock{
			HotStocks:       hotStocks,
			FinancingTrend:  financingTrend,
			MarginRatio:     marginRatio,
		},
	}, nil
}

// 数据钻取相关方法
func (c *DashboardClient) GetDrilldownDetail(ctx context.Context, dataType, id string, page, pageSize int) (*model.Response, error) {
	c.logger.Info("Getting drilldown detail")

	// 模拟数据
	list := []model.DrilldownItem{
		{ID: "1", CustomerId: "C001", Name: "张三", Age: 30, Region: "北京", Asset: 100000, TradeAmount: 50000},
		{ID: "2", CustomerId: "C002", Name: "李四", Age: 35, Region: "上海", Asset: 150000, TradeAmount: 80000},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.DrilldownDetail{
			Total:    100,
			Page:     page,
			PageSize: pageSize,
			List:     list,
		},
	}, nil
}

func (c *DashboardClient) GetDrilldownTrend(ctx context.Context, dataType, id, period string) (*model.Response, error) {
	c.logger.Info("Getting drilldown trend")

	// 模拟数据
	trend := []model.TrendItem{
		{Date: "2023-01", Value: 2800},
		{Date: "2023-02", Value: 2900},
		{Date: "2023-03", Value: 3000},
	}

	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.DrilldownTrend{
			Type:  dataType,
			Name:  "活跃用户数",
			Trend: trend,
		},
	}, nil
}

// 认证相关方法
func (c *DashboardClient) Login(ctx context.Context, username, password string) (*model.Response, error) {
	c.logger.Info("User login")

	// 模拟数据
	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.LoginResponse{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
			User: model.User{
				ID:       "U001",
				Username: "admin",
				Role:     "admin",
			},
		},
	}, nil
}

func (c *DashboardClient) GetAuthInfo(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting auth info")

	// 模拟数据
	return &model.Response{
		Code:    200,
		Message: "success",
		Data: model.User{
			ID:          "U001",
			Username:    "admin",
			Role:        "admin",
			Permissions: []string{"dashboard:view", "customer:view", "trade:view"},
		},
	}, nil
}

// Close 关闭gRPC连接
func (c *DashboardClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
