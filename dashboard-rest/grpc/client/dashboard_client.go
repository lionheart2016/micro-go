package client

import (
	"context"
	"dashboard-rest/model"
	"dashboard-rest/pkg/logger"
	"fmt"

	"dashboard-rest/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DashboardClient struct {
	logger  *logger.Logger
	conn    *grpc.ClientConn
	service proto.DashboardServiceClient
}

func New(grpcHost string, grpcPort int, logger *logger.Logger) (*DashboardClient, error) {
	addr := fmt.Sprintf("%s:%d", grpcHost, grpcPort)
	// 连接到gRPC服务器
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Failed to connect to gRPC server: %v", err)
		return nil, err
	}

	// 创建服务客户端
	service := proto.NewDashboardServiceClient(conn)

	return &DashboardClient{
		logger:  logger,
		conn:    conn,
		service: service,
	}, nil
}

func (c *DashboardClient) GetCoreIndicators(ctx context.Context, period string) (*model.Response, error) {
	c.logger.Info("Getting core indicators")

	// 构建请求
	req := &proto.GetCoreIndicatorsRequest{
		Period: period,
	}

	// 调用gRPC接口
	resp, err := c.service.GetCoreIndicators(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get core indicators: %v", err)
		return nil, err
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.CoreIndicators{
			RegisteredUsers: model.Indicator{
				Value:      int(resp.Data.RegisteredUsers.Value),
				Change:     int(resp.Data.RegisteredUsers.Change),
				ChangeRate: resp.Data.RegisteredUsers.ChangeRate,
			},
			AccountCount: model.Indicator{
				Value:      int(resp.Data.AccountCount.Value),
				Change:     int(resp.Data.AccountCount.Change),
				ChangeRate: resp.Data.AccountCount.ChangeRate,
			},
			ActiveUsers: model.Indicator{
				Value:      int(resp.Data.ActiveUsers.Value),
				Change:     int(resp.Data.ActiveUsers.Change),
				ChangeRate: resp.Data.ActiveUsers.ChangeRate,
			},
			DepositUsers: model.Indicator{
				Value:      int(resp.Data.DepositUsers.Value),
				Change:     int(resp.Data.DepositUsers.Change),
				ChangeRate: resp.Data.DepositUsers.ChangeRate,
			},
			DepositAmount: model.Indicator{
				Value:      int(resp.Data.DepositAmount.Value),
				Change:     int(resp.Data.DepositAmount.Change),
				ChangeRate: resp.Data.DepositAmount.ChangeRate,
			},
			StockTradeUsers: model.Indicator{
				Value:      int(resp.Data.StockTradeUsers.Value),
				Change:     int(resp.Data.StockTradeUsers.Change),
				ChangeRate: resp.Data.StockTradeUsers.ChangeRate,
			},
			StockTradeAmount: model.Indicator{
				Value:      int(resp.Data.StockTradeAmount.Value),
				Change:     int(resp.Data.StockTradeAmount.Change),
				ChangeRate: resp.Data.StockTradeAmount.ChangeRate,
			},
			FundTradeUsers: model.Indicator{
				Value:      int(resp.Data.FundTradeUsers.Value),
				Change:     int(resp.Data.FundTradeUsers.Change),
				ChangeRate: resp.Data.FundTradeUsers.ChangeRate,
			},
			FundTradeAmount: model.Indicator{
				Value:      int(resp.Data.FundTradeAmount.Value),
				Change:     int(resp.Data.FundTradeAmount.Change),
				ChangeRate: resp.Data.FundTradeAmount.ChangeRate,
			},
		},
	}, nil
}

func (c *DashboardClient) GetPerformance(ctx context.Context, period string) (*model.Response, error) {
	c.logger.Info("Getting performance")

	// 构建请求
	req := &proto.GetPerformanceRequest{
		Period: period,
	}

	// 调用gRPC接口
	resp, err := c.service.GetPerformance(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get performance: %v", err)
		return nil, err
	}

	// 转换revenueStructure
	revenueStructure := make([]model.RevenueItem, len(resp.Data.RevenueStructure))
	for i, item := range resp.Data.RevenueStructure {
		revenueStructure[i] = model.RevenueItem{
			Name:  item.Name,
			Value: int(item.Value),
		}
	}

	// 转换revenueTrend
	revenueTrend := make([]model.RevenueTrend, len(resp.Data.RevenueTrend))
	for i, item := range resp.Data.RevenueTrend {
		revenueTrend[i] = model.RevenueTrend{
			Month: item.Month,
			Stock: int(item.Stock),
			Fund:  int(item.Fund),
		}
	}

	// 转换depositTrend
	depositTrend := make([]model.DepositTrend, len(resp.Data.DepositTrend))
	for i, item := range resp.Data.DepositTrend {
		depositTrend[i] = model.DepositTrend{
			Month: item.Month,
			Value: int(item.Value),
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.Performance{
			StockRevenue:     int(resp.Data.StockRevenue),
			FundRevenue:      int(resp.Data.FundRevenue),
			RevenueStructure: revenueStructure,
			RevenueTrend:     revenueTrend,
			RevenueChange: model.RevenueChange{
				Yoy: resp.Data.RevenueChange.Yoy,
				Mom: resp.Data.RevenueChange.Mom,
			},
			DepositNetGrowth:       int(resp.Data.DepositNetGrowth),
			DepositTrend:           depositTrend,
			DepositWithdrawalRatio: resp.Data.DepositWithdrawalRatio,
		},
	}, nil
}

// 客户分析相关方法
func (c *DashboardClient) GetCustomerDistribution(ctx context.Context, dimension string) (*model.Response, error) {
	c.logger.Info("Getting customer distribution")

	// 构建请求
	req := &proto.GetCustomerDistributionRequest{
		Dimension: dimension,
	}

	// 调用gRPC接口
	resp, err := c.service.GetCustomerDistribution(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get customer distribution: %v", err)
		return nil, err
	}

	// 转换distribution
	distribution := make([]model.CustomerDistributionItem, len(resp.Data.Distribution))
	for i, item := range resp.Data.Distribution {
		distribution[i] = model.CustomerDistributionItem{
			Name:           item.Name,
			Value:          int(item.Value),
			TotalAsset:     int(item.TotalAsset),
			AvgAsset:       int(item.AvgAsset),
			AvgTradeAmount: int(item.AvgTradeAmount),
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.CustomerDistribution{
			Dimension:    resp.Data.Dimension,
			Distribution: distribution,
		},
	}, nil
}

func (c *DashboardClient) GetCustomerBehavior(ctx context.Context, behaviorType string) (*model.Response, error) {
	c.logger.Info("Getting customer behavior")

	// 构建请求
	req := &proto.GetCustomerBehaviorRequest{
		Type: behaviorType,
	}

	// 调用gRPC接口
	resp, err := c.service.GetCustomerBehavior(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get customer behavior: %v", err)
		return nil, err
	}

	// 转换trend
	trend := make([]model.CustomerBehaviorTrend, len(resp.Data.Trend))
	for i, item := range resp.Data.Trend {
		trend[i] = model.CustomerBehaviorTrend{
			Quarter:       item.Quarter,
			RetentionRate: item.RetentionRate,
			ChurnRate:     item.ChurnRate,
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.CustomerBehavior{
			Type:              resp.Data.Type,
			Trend:             trend,
			IndustryAverage:   resp.Data.IndustryAverage,
			CompetitorAverage: resp.Data.CompetitorAverage,
		},
	}, nil
}

// 交易分析相关方法
func (c *DashboardClient) GetStockTrade(ctx context.Context, tradeType string) (*model.Response, error) {
	c.logger.Info("Getting stock trade")

	// 构建请求
	req := &proto.GetStockTradeRequest{
		Type: tradeType,
	}

	// 调用gRPC接口
	resp, err := c.service.GetStockTrade(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get stock trade: %v", err)
		return nil, err
	}

	// 转换executionTime
	executionTime := make([]model.ExecutionTime, len(resp.Data.ExecutionTime))
	for i, item := range resp.Data.ExecutionTime {
		executionTime[i] = model.ExecutionTime{
			Time:    item.Time,
			AvgTime: item.AvgTime,
		}
	}

	// 转换marketComparison
	marketComparison := make([]model.MarketComparison, len(resp.Data.MarketComparison))
	for i, item := range resp.Data.MarketComparison {
		marketComparison[i] = model.MarketComparison{
			Market:  item.Market,
			AvgTime: item.AvgTime,
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.StockTrade{
			Type:             resp.Data.Type,
			ExecutionTime:    executionTime,
			MarketComparison: marketComparison,
		},
	}, nil
}

func (c *DashboardClient) GetFundTrade(ctx context.Context, tradeType string) (*model.Response, error) {
	c.logger.Info("Getting fund trade")

	// 构建请求
	req := &proto.GetFundTradeRequest{
		Type: tradeType,
	}

	// 调用gRPC接口
	resp, err := c.service.GetFundTrade(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get fund trade: %v", err)
		return nil, err
	}

	// 转换salesByType
	salesByType := make([]model.SalesByType, len(resp.Data.SalesByType))
	for i, item := range resp.Data.SalesByType {
		salesByType[i] = model.SalesByType{
			Type:   item.Type,
			Amount: int(item.Amount),
			Count:  int(item.Count),
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.FundTrade{
			Type:        resp.Data.Type,
			SalesByType: salesByType,
		},
	}, nil
}

// PI用户分析相关方法
func (c *DashboardClient) GetPIBasic(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting PI basic information")

	// 构建请求
	req := &proto.GetPIBasicRequest{}

	// 调用gRPC接口
	resp, err := c.service.GetPIBasic(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get PI basic: %v", err)
		return nil, err
	}

	// 转换assetDistribution
	assetDistribution := make([]model.AssetDistribution, len(resp.Data.AssetDistribution))
	for i, item := range resp.Data.AssetDistribution {
		assetDistribution[i] = model.AssetDistribution{
			Name:  item.Name,
			Value: item.Value,
		}
	}

	// 转换demographics
	demographics := make([]model.Demographic, len(resp.Data.Demographics))
	for i, item := range resp.Data.Demographics {
		demographics[i] = model.Demographic{
			Job:               item.Job,
			Percentage:        item.Percentage,
			AvgAsset:          int(item.AvgAsset),
			AvgTradeFrequency: int(item.AvgTradeFrequency),
			AvgFundCount:      int(item.AvgFundCount),
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.PIBasic{
			TotalUsers:        int(resp.Data.TotalUsers),
			Change:            int(resp.Data.Change),
			ChangeRate:        resp.Data.ChangeRate,
			AssetDistribution: assetDistribution,
			TotalAsset:        int(resp.Data.TotalAsset),
			Demographics:      demographics,
		},
	}, nil
}

func (c *DashboardClient) GetPIBehavior(ctx context.Context, behaviorType string) (*model.Response, error) {
	c.logger.Info("Getting PI behavior")

	// 构建请求
	req := &proto.GetPIBehaviorRequest{
		Type: behaviorType,
	}

	// 调用gRPC接口
	resp, err := c.service.GetPIBehavior(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get PI behavior: %v", err)
		return nil, err
	}

	// 转换moneyFlow
	moneyFlow := make([]model.MoneyFlow, len(resp.Data.MoneyFlow))
	for i, item := range resp.Data.MoneyFlow {
		moneyFlow[i] = model.MoneyFlow{
			From:  item.From,
			To:    item.To,
			Value: int(item.Value),
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.PIBehavior{
			Type:              resp.Data.Type,
			AvgTradeAmount:    int(resp.Data.AvgTradeAmount),
			AvgTradeFrequency: int(resp.Data.AvgTradeFrequency),
			AvgHoldingPeriod:  int(resp.Data.AvgHoldingPeriod),
			MoneyFlow:         moneyFlow,
		},
	}, nil
}

// 开户主题相关方法
func (c *DashboardClient) GetAccountProcess(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting account process")

	// 构建请求
	req := &proto.GetAccountProcessRequest{}

	// 调用gRPC接口
	resp, err := c.service.GetAccountProcess(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get account process: %v", err)
		return nil, err
	}

	// 转换funnel
	funnel := make([]model.FunnelStage, len(resp.Data.Funnel))
	for i, item := range resp.Data.Funnel {
		funnel[i] = model.FunnelStage{
			Stage:      item.Stage,
			Count:      int(item.Count),
			Conversion: item.Conversion,
		}
	}

	// 转换stageTime
	stageTime := make([]model.StageTime, len(resp.Data.StageTime))
	for i, item := range resp.Data.StageTime {
		stageTime[i] = model.StageTime{
			Stage:      item.Stage,
			AvgTime:    item.AvgTime,
			MedianTime: item.MedianTime,
			P25Time:    item.P25Time,
			P75Time:    item.P75Time,
			MaxTime:    item.MaxTime,
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.AccountProcess{
			Funnel:    funnel,
			StageTime: stageTime,
		},
	}, nil
}

func (c *DashboardClient) GetAccountException(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting account exception")

	// 构建请求
	req := &proto.GetAccountExceptionRequest{}

	// 调用gRPC接口
	resp, err := c.service.GetAccountException(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get account exception: %v", err)
		return nil, err
	}

	// 转换exceptionTypes
	exceptionTypes := make([]model.ExceptionType, len(resp.Data.ExceptionTypes))
	for i, item := range resp.Data.ExceptionTypes {
		exceptionTypes[i] = model.ExceptionType{
			Type:       item.Type,
			Count:      int(item.Count),
			Percentage: item.Percentage,
		}
	}

	// 转换exceptionStages
	exceptionStages := make([]model.ExceptionStage, len(resp.Data.ExceptionStages))
	for i, item := range resp.Data.ExceptionStages {
		exceptionStages[i] = model.ExceptionStage{
			Stage: item.Stage,
			Count: int(item.Count),
		}
	}

	// 转换exceptionHandling
	exceptionHandling := make([]model.ExceptionHandling, len(resp.Data.ExceptionHandling))
	for i, item := range resp.Data.ExceptionHandling {
		exceptionHandling[i] = model.ExceptionHandling{
			Type:          item.Type,
			AvgHandleTime: item.AvgHandleTime,
			MaxHandleTime: item.MaxHandleTime,
			SuccessRate:   item.SuccessRate,
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.AccountException{
			ExceptionTypes:    exceptionTypes,
			ExceptionStages:   exceptionStages,
			ExceptionHandling: exceptionHandling,
		},
	}, nil
}

// IPO主题相关方法
func (c *DashboardClient) GetIPOSubscription(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting IPO subscription")

	// 构建请求
	req := &proto.GetIPOSubscriptionRequest{}

	// 调用gRPC接口
	resp, err := c.service.GetIPOSubscription(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get IPO subscription: %v", err)
		return nil, err
	}

	// 转换userDistribution
	userDistribution := make([]model.UserDistribution, len(resp.Data.UserDistribution))
	for i, item := range resp.Data.UserDistribution {
		userDistribution[i] = model.UserDistribution{
			Type:       item.Type,
			Count:      int(item.Count),
			Percentage: item.Percentage,
		}
	}

	// 转换subscriptionAmount
	subscriptionAmount := make([]model.SubscriptionAmount, len(resp.Data.SubscriptionAmount))
	for i, item := range resp.Data.SubscriptionAmount {
		subscriptionAmount[i] = model.SubscriptionAmount{
			Project:       item.Project,
			PiAmount:      int(item.PiAmount),
			RegularAmount: int(item.RegularAmount),
			TotalAmount:   int(item.TotalAmount),
		}
	}

	// 转换subscriptionMultiples
	subscriptionMultiples := make([]model.SubscriptionMultiple, len(resp.Data.SubscriptionMultiples))
	for i, item := range resp.Data.SubscriptionMultiples {
		subscriptionMultiples[i] = model.SubscriptionMultiple{
			Project:  item.Project,
			Multiple: item.Multiple,
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.IPOSubscription{
			TotalUsers:            int(resp.Data.TotalUsers),
			UserDistribution:      userDistribution,
			SubscriptionAmount:    subscriptionAmount,
			SubscriptionMultiples: subscriptionMultiples,
		},
	}, nil
}

func (c *DashboardClient) GetIPOAnaalysis(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting IPO analysis")

	// 构建请求
	req := &proto.GetIPOAnaalysisRequest{}

	// 调用gRPC接口
	resp, err := c.service.GetIPOAnaalysis(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get IPO analysis: %v", err)
		return nil, err
	}

	// 转换projectPerformance
	projectPerformance := make([]model.ProjectPerformance, len(resp.Data.ProjectPerformance))
	for i, item := range resp.Data.ProjectPerformance {
		projectPerformance[i] = model.ProjectPerformance{
			Project:     item.Project,
			PriceChange: item.PriceChange,
			AvgReturn:   item.AvgReturn,
		}
	}

	// 转换influenceFactors
	influenceFactors := make([]model.InfluenceFactor, len(resp.Data.InfluenceFactors))
	for i, item := range resp.Data.InfluenceFactors {
		influenceFactors[i] = model.InfluenceFactor{
			Factor:      item.Factor,
			Correlation: item.Correlation,
		}
	}

	// 转换competitorComparison
	competitorComparison := make([]model.CompetitorComparison, len(resp.Data.CompetitorComparison))
	for i, item := range resp.Data.CompetitorComparison {
		competitorComparison[i] = model.CompetitorComparison{
			Competitor:         item.Competitor,
			SubscriptionUsers:  int(item.SubscriptionUsers),
			SubscriptionAmount: int(item.SubscriptionAmount),
			Multiple:           item.Multiple,
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.IPOAAnalysis{
			ProjectPerformance:   projectPerformance,
			InfluenceFactors:     influenceFactors,
			CompetitorComparison: competitorComparison,
		},
	}, nil
}

// 融资主题相关方法
func (c *DashboardClient) GetFinanceCustomer(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting finance customer")

	// 构建请求
	req := &proto.GetFinanceCustomerRequest{}

	// 调用gRPC接口
	resp, err := c.service.GetFinanceCustomer(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get finance customer: %v", err)
		return nil, err
	}

	// 转换userDistribution
	userDistribution := make([]model.UserDistribution, len(resp.Data.UserDistribution))
	for i, item := range resp.Data.UserDistribution {
		userDistribution[i] = model.UserDistribution{
			Type:       item.Type,
			Count:      int(item.Count),
			Percentage: item.Percentage,
		}
	}

	// 转换financingPurpose
	financingPurpose := make([]model.FinancingPurpose, len(resp.Data.FinancingPurpose))
	for i, item := range resp.Data.FinancingPurpose {
		financingPurpose[i] = model.FinancingPurpose{
			Purpose:    item.Purpose,
			Percentage: item.Percentage,
		}
	}

	// 转换riskAssessment
	riskAssessment := make([]model.RiskAssessment, len(resp.Data.RiskAssessment))
	for i, item := range resp.Data.RiskAssessment {
		riskAssessment[i] = model.RiskAssessment{
			Type:        item.Type,
			DefaultRate: item.DefaultRate,
			OverdueRate: item.OverdueRate,
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.FinanceCustomer{
			TotalUsers:       int(resp.Data.TotalUsers),
			UserDistribution: userDistribution,
			FinancingSize: model.FinancingSize{
				AvgSize:    resp.Data.FinancingSize.AvgSize,
				MedianSize: resp.Data.FinancingSize.MedianSize,
				P25Size:    resp.Data.FinancingSize.P25Size,
				P75Size:    resp.Data.FinancingSize.P75Size,
				MaxSize:    resp.Data.FinancingSize.MaxSize,
			},
			FinancingPurpose: financingPurpose,
			RiskAssessment:   riskAssessment,
		},
	}, nil
}

func (c *DashboardClient) GetFinanceStock(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting finance stock")

	// 构建请求
	req := &proto.GetFinanceStockRequest{}

	// 调用gRPC接口
	resp, err := c.service.GetFinanceStock(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get finance stock: %v", err)
		return nil, err
	}

	// 转换hotStocks
	hotStocks := make([]model.HotStock, len(resp.Data.HotStocks))
	for i, item := range resp.Data.HotStocks {
		hotStocks[i] = model.HotStock{
			Stock:            item.Stock,
			FinancingBalance: int(item.FinancingBalance),
			MarketCapRatio:   item.MarketCapRatio,
		}
	}

	// 转换financingTrend
	financingTrend := make([]model.FinancingTrend, len(resp.Data.FinancingTrend))
	for i, item := range resp.Data.FinancingTrend {
		financingTrend[i] = model.FinancingTrend{
			Date:       item.Date,
			BuyAmount:  int(item.BuyAmount),
			SellAmount: int(item.SellAmount),
		}
	}

	// 转换marginRatio
	marginRatio := make([]model.MarginRatio, len(resp.Data.MarginRatio))
	for i, item := range resp.Data.MarginRatio {
		marginRatio[i] = model.MarginRatio{
			Stock:           item.Stock,
			Ratio:           item.Ratio,
			PriceVolatility: item.PriceVolatility,
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.FinanceStock{
			HotStocks:      hotStocks,
			FinancingTrend: financingTrend,
			MarginRatio:    marginRatio,
		},
	}, nil
}

// 数据钻取相关方法
func (c *DashboardClient) GetDrilldownDetail(ctx context.Context, dataType, id string, page, pageSize int) (*model.Response, error) {
	c.logger.Info("Getting drilldown detail")

	// 构建请求
	req := &proto.GetDrilldownDetailRequest{
		Type:     dataType,
		Id:       id,
		Page:     int32(page),
		PageSize: int32(pageSize),
	}

	// 调用gRPC接口
	resp, err := c.service.GetDrilldownDetail(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get drilldown detail: %v", err)
		return nil, err
	}

	// 转换list
	list := make([]model.DrilldownItem, len(resp.Data.List))
	for i, item := range resp.Data.List {
		list[i] = model.DrilldownItem{
			ID:          item.Id,
			CustomerId:  item.CustomerId,
			Name:        item.Name,
			Age:         int(item.Age),
			Region:      item.Region,
			Asset:       int(item.Asset),
			TradeAmount: int(item.TradeAmount),
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.DrilldownDetail{
			Total:    int(resp.Data.Total),
			Page:     int(resp.Data.Page),
			PageSize: int(resp.Data.PageSize),
			List:     list,
		},
	}, nil
}

func (c *DashboardClient) GetDrilldownTrend(ctx context.Context, dataType, id, period string) (*model.Response, error) {
	c.logger.Info("Getting drilldown trend")

	// 构建请求
	req := &proto.GetDrilldownTrendRequest{
		Type:   dataType,
		Id:     id,
		Period: period,
	}

	// 调用gRPC接口
	resp, err := c.service.GetDrilldownTrend(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get drilldown trend: %v", err)
		return nil, err
	}

	// 转换trend
	trend := make([]model.TrendItem, len(resp.Data.Trend))
	for i, item := range resp.Data.Trend {
		trend[i] = model.TrendItem{
			Date:  item.Date,
			Value: int(item.Value),
		}
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.DrilldownTrend{
			Type:  resp.Data.Type,
			Name:  resp.Data.Name,
			Trend: trend,
		},
	}, nil
}

// 认证相关方法
func (c *DashboardClient) Login(ctx context.Context, username, password string) (*model.Response, error) {
	c.logger.Info("User login")

	// 构建请求
	req := &proto.LoginRequest{
		Username: username,
		Password: password,
	}

	// 调用gRPC接口
	resp, err := c.service.Login(ctx, req)
	if err != nil {
		c.logger.Error("Failed to login: %v", err)
		return nil, err
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.LoginResponse{
			Token: resp.Data.Token,
			User: model.User{
				ID:       resp.Data.User.Id,
				Username: resp.Data.User.Username,
				Role:     resp.Data.User.Role,
			},
		},
	}, nil
}

func (c *DashboardClient) GetAuthInfo(ctx context.Context) (*model.Response, error) {
	c.logger.Info("Getting auth info")

	// 构建请求
	req := &proto.GetAuthInfoRequest{}

	// 调用gRPC接口
	resp, err := c.service.GetAuthInfo(ctx, req)
	if err != nil {
		c.logger.Error("Failed to get auth info: %v", err)
		return nil, err
	}

	// 转换为模型
	return &model.Response{
		Code:    int(resp.Code),
		Message: resp.Message,
		Data: model.User{
			ID:          resp.Data.Id,
			Username:    resp.Data.Username,
			Role:        resp.Data.Role,
			Permissions: resp.Data.Permissions,
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
