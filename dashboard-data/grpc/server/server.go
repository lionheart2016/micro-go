package server

import (
	"context"
	"dashboard-data/grpc/proto"
	"dashboard-data/pkg/logger"
	"dashboard-data/repository"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// DashboardService 仪表盘服务实现
type DashboardService struct {
	proto.UnimplementedDashboardServiceServer
	dashboardRepo repository.DashboardRepository
}

// NewDashboardService 创建仪表盘服务实例
func NewDashboardService() *DashboardService {
	return &DashboardService{
		dashboardRepo: repository.NewDashboardRepository(),
	}
}

// RegisterDataService 注册数据服务
func RegisterDataService(s *grpc.Server) {
	proto.RegisterDashboardServiceServer(s, NewDashboardService())
	logger.Info("Dashboard Data Service registered")
}

// GetCoreIndicators 获取核心指标
func (s *DashboardService) GetCoreIndicators(ctx context.Context, req *proto.GetCoreIndicatorsRequest) (*proto.GetCoreIndicatorsResponse, error) {
	// 调用仓库获取数据
	coreIndicators, err := s.dashboardRepo.GetCoreIndicators(req.Period)
	if err != nil {
		logger.Error("Failed to get core indicators",
			zap.Error(err),
			zap.String("period", req.Period),
		)
		return &proto.GetCoreIndicatorsResponse{
			Code:    500,
			Message: "Failed to get core indicators",
		}, err
	}

	logger.Info("Core indicators retrieved successfully",
		zap.String("period", req.Period),
	)

	// 转换为 proto 响应
	return &proto.GetCoreIndicatorsResponse{
		Code:    200,
		Message: "success",
		Data: &proto.CoreIndicators{
			RegisteredUsers: &proto.Indicator{
				Value:      int32(coreIndicators.RegisteredUsers.Value),
				Change:     int32(coreIndicators.RegisteredUsers.Change),
				ChangeRate: coreIndicators.RegisteredUsers.ChangeRate,
			},
			AccountCount: &proto.Indicator{
				Value:      int32(coreIndicators.AccountCount.Value),
				Change:     int32(coreIndicators.AccountCount.Change),
				ChangeRate: coreIndicators.AccountCount.ChangeRate,
			},
			ActiveUsers: &proto.Indicator{
				Value:      int32(coreIndicators.ActiveUsers.Value),
				Change:     int32(coreIndicators.ActiveUsers.Change),
				ChangeRate: coreIndicators.ActiveUsers.ChangeRate,
			},
			DepositUsers: &proto.Indicator{
				Value:      int32(coreIndicators.DepositUsers.Value),
				Change:     int32(coreIndicators.DepositUsers.Change),
				ChangeRate: coreIndicators.DepositUsers.ChangeRate,
			},
			DepositAmount: &proto.Indicator{
				Value:      int32(coreIndicators.DepositAmount.Value),
				Change:     int32(coreIndicators.DepositAmount.Change),
				ChangeRate: coreIndicators.DepositAmount.ChangeRate,
			},
			StockTradeUsers: &proto.Indicator{
				Value:      int32(coreIndicators.StockTradeUsers.Value),
				Change:     int32(coreIndicators.StockTradeUsers.Change),
				ChangeRate: coreIndicators.StockTradeUsers.ChangeRate,
			},
			StockTradeAmount: &proto.Indicator{
				Value:      int32(coreIndicators.StockTradeAmount.Value),
				Change:     int32(coreIndicators.StockTradeAmount.Change),
				ChangeRate: coreIndicators.StockTradeAmount.ChangeRate,
			},
			FundTradeUsers: &proto.Indicator{
				Value:      int32(coreIndicators.FundTradeUsers.Value),
				Change:     int32(coreIndicators.FundTradeUsers.Change),
				ChangeRate: coreIndicators.FundTradeUsers.ChangeRate,
			},
			FundTradeAmount: &proto.Indicator{
				Value:      int32(coreIndicators.FundTradeAmount.Value),
				Change:     int32(coreIndicators.FundTradeAmount.Change),
				ChangeRate: coreIndicators.FundTradeAmount.ChangeRate,
			},
		},
	}, nil
}

// GetPerformance 获取业绩表现
func (s *DashboardService) GetPerformance(ctx context.Context, req *proto.GetPerformanceRequest) (*proto.GetPerformanceResponse, error) {
	// 调用仓库获取数据
	performance, err := s.dashboardRepo.GetPerformance(req.Period)
	if err != nil {
		logger.Error("Failed to get performance",
			zap.Error(err),
			zap.String("period", req.Period),
		)
		return &proto.GetPerformanceResponse{
			Code:    500,
			Message: "Failed to get performance",
		}, err
	}

	logger.Info("Performance retrieved successfully",
		zap.String("period", req.Period),
	)

	// 转换为 proto 响应
	revenueStructure := make([]*proto.RevenueItem, len(performance.RevenueStructure))
	for i, item := range performance.RevenueStructure {
		revenueStructure[i] = &proto.RevenueItem{
			Name:  item.Name,
			Value: int32(item.Value),
		}
	}

	revenueTrend := make([]*proto.RevenueTrend, len(performance.RevenueTrend))
	for i, trend := range performance.RevenueTrend {
		revenueTrend[i] = &proto.RevenueTrend{
			Month: trend.Month,
			Stock: int32(trend.Stock),
			Fund:  int32(trend.Fund),
		}
	}

	depositTrend := make([]*proto.DepositTrend, len(performance.DepositTrend))
	for i, trend := range performance.DepositTrend {
		depositTrend[i] = &proto.DepositTrend{
			Month: trend.Month,
			Value: int32(trend.Value),
		}
	}

	return &proto.GetPerformanceResponse{
		Code:    200,
		Message: "success",
		Data: &proto.Performance{
			StockRevenue:     int32(performance.StockRevenue),
			FundRevenue:      int32(performance.FundRevenue),
			RevenueStructure: revenueStructure,
			RevenueTrend:     revenueTrend,
			RevenueChange: &proto.RevenueChange{
				Yoy: performance.RevenueChange.Yoy,
				Mom: performance.RevenueChange.Mom,
			},
			DepositNetGrowth:       int32(performance.DepositNetGrowth),
			DepositTrend:           depositTrend,
			DepositWithdrawalRatio: performance.DepositWithdrawalRatio,
		},
	}, nil
}

// GetCustomerDistribution 获取客户分布
func (s *DashboardService) GetCustomerDistribution(ctx context.Context, req *proto.GetCustomerDistributionRequest) (*proto.GetCustomerDistributionResponse, error) {
	// 调用仓库获取数据
	distribution, err := s.dashboardRepo.GetCustomerDistribution(req.Dimension)
	if err != nil {
		logger.Error("Failed to get customer distribution",
			zap.Error(err),
			zap.String("dimension", req.Dimension),
		)
		return &proto.GetCustomerDistributionResponse{
			Code:    500,
			Message: "Failed to get customer distribution",
		}, err
	}

	logger.Info("Customer distribution retrieved successfully",
		zap.String("dimension", req.Dimension),
	)

	// 转换为 proto 响应
	distributionItems := make([]*proto.CustomerDistributionItem, len(distribution.Distribution))
	for i, item := range distribution.Distribution {
		distributionItems[i] = &proto.CustomerDistributionItem{
			Name:           item.Name,
			Value:          int32(item.Value),
			TotalAsset:     int32(item.TotalAsset),
			AvgAsset:       int32(item.AvgAsset),
			AvgTradeAmount: int32(item.AvgTradeAmount),
		}
	}

	return &proto.GetCustomerDistributionResponse{
		Code:    200,
		Message: "success",
		Data: &proto.CustomerDistribution{
			Dimension:    distribution.Dimension,
			Distribution: distributionItems,
		},
	}, nil
}

// GetCustomerBehavior 获取客户行为
func (s *DashboardService) GetCustomerBehavior(ctx context.Context, req *proto.GetCustomerBehaviorRequest) (*proto.GetCustomerBehaviorResponse, error) {
	// 调用仓库获取数据
	behavior, err := s.dashboardRepo.GetCustomerBehavior(req.Type)
	if err != nil {
		logger.Error("Failed to get customer behavior",
			zap.Error(err),
			zap.String("type", req.Type),
		)
		return &proto.GetCustomerBehaviorResponse{
			Code:    500,
			Message: "Failed to get customer behavior",
		}, err
	}

	logger.Info("Customer behavior retrieved successfully",
		zap.String("type", req.Type),
	)

	// 转换为 proto 响应
	trendItems := make([]*proto.CustomerBehaviorTrend, len(behavior.Trend))
	for i, trend := range behavior.Trend {
		trendItems[i] = &proto.CustomerBehaviorTrend{
			Quarter:       trend.Quarter,
			RetentionRate: trend.RetentionRate,
			ChurnRate:     trend.ChurnRate,
		}
	}

	return &proto.GetCustomerBehaviorResponse{
		Code:    200,
		Message: "success",
		Data: &proto.CustomerBehavior{
			Type:              behavior.Type,
			Trend:             trendItems,
			IndustryAverage:   behavior.IndustryAverage,
			CompetitorAverage: behavior.CompetitorAverage,
		},
	}, nil
}

// GetStockTrade 获取股票交易
func (s *DashboardService) GetStockTrade(ctx context.Context, req *proto.GetStockTradeRequest) (*proto.GetStockTradeResponse, error) {
	// 调用仓库获取数据
	stockTrade, err := s.dashboardRepo.GetStockTrade(req.Type)
	if err != nil {
		logger.Error("Failed to get stock trade",
			zap.Error(err),
			zap.String("type", req.Type),
		)
		return &proto.GetStockTradeResponse{
			Code:    500,
			Message: "Failed to get stock trade",
		}, err
	}

	logger.Info("Stock trade retrieved successfully",
		zap.String("type", req.Type),
	)

	// 转换为 proto 响应
	executionTimes := make([]*proto.ExecutionTime, len(stockTrade.ExecutionTime))
	for i, item := range stockTrade.ExecutionTime {
		executionTimes[i] = &proto.ExecutionTime{
			Time:    item.Time,
			AvgTime: item.AvgTime,
		}
	}

	marketComparisons := make([]*proto.MarketComparison, len(stockTrade.MarketComparison))
	for i, item := range stockTrade.MarketComparison {
		marketComparisons[i] = &proto.MarketComparison{
			Market:  item.Market,
			AvgTime: item.AvgTime,
		}
	}

	return &proto.GetStockTradeResponse{
		Code:    200,
		Message: "success",
		Data: &proto.StockTrade{
			Type:             stockTrade.Type,
			ExecutionTime:    executionTimes,
			MarketComparison: marketComparisons,
		},
	}, nil
}

// GetFundTrade 获取基金交易
func (s *DashboardService) GetFundTrade(ctx context.Context, req *proto.GetFundTradeRequest) (*proto.GetFundTradeResponse, error) {
	// 调用仓库获取数据
	fundTrade, err := s.dashboardRepo.GetFundTrade(req.Type)
	if err != nil {
		logger.Error("Failed to get fund trade",
			zap.Error(err),
			zap.String("type", req.Type),
		)
		return &proto.GetFundTradeResponse{
			Code:    500,
			Message: "Failed to get fund trade",
		}, err
	}

	logger.Info("Fund trade retrieved successfully",
		zap.String("type", req.Type),
	)

	// 转换为 proto 响应
	salesByType := make([]*proto.SalesByType, len(fundTrade.SalesByType))
	for i, item := range fundTrade.SalesByType {
		salesByType[i] = &proto.SalesByType{
			Type:   item.Type,
			Amount: int32(item.Amount),
			Count:  int32(item.Count),
		}
	}

	return &proto.GetFundTradeResponse{
		Code:    200,
		Message: "success",
		Data: &proto.FundTrade{
			Type:        fundTrade.Type,
			SalesByType: salesByType,
		},
	}, nil
}

// GetPIBasic 获取PI用户基本信息
func (s *DashboardService) GetPIBasic(ctx context.Context, req *proto.GetPIBasicRequest) (*proto.GetPIBasicResponse, error) {
	// 调用仓库获取数据
	piBasic, err := s.dashboardRepo.GetPIBasic()
	if err != nil {
		logger.Error("Failed to get PI basic information",
			zap.Error(err),
		)
		return &proto.GetPIBasicResponse{
			Code:    500,
			Message: "Failed to get PI basic information",
		}, err
	}

	logger.Info("PI basic information retrieved successfully")

	// 转换为 proto 响应
	assetDistribution := make([]*proto.AssetDistribution, len(piBasic.AssetDistribution))
	for i, item := range piBasic.AssetDistribution {
		assetDistribution[i] = &proto.AssetDistribution{
			Name:  item.Name,
			Value: item.Value,
		}
	}

	demographics := make([]*proto.Demographic, len(piBasic.Demographics))
	for i, item := range piBasic.Demographics {
		demographics[i] = &proto.Demographic{
			Job:               item.Job,
			Percentage:        item.Percentage,
			AvgAsset:          int32(item.AvgAsset),
			AvgTradeFrequency: int32(item.AvgTradeFrequency),
			AvgFundCount:      int32(item.AvgFundCount),
		}
	}

	return &proto.GetPIBasicResponse{
		Code:    200,
		Message: "success",
		Data: &proto.PIBasic{
			TotalUsers:        int32(piBasic.TotalUsers),
			Change:            int32(piBasic.Change),
			ChangeRate:        piBasic.ChangeRate,
			AssetDistribution: assetDistribution,
			TotalAsset:        int32(piBasic.TotalAsset),
			Demographics:      demographics,
		},
	}, nil
}

// GetPIBehavior 获取PI用户行为
func (s *DashboardService) GetPIBehavior(ctx context.Context, req *proto.GetPIBehaviorRequest) (*proto.GetPIBehaviorResponse, error) {
	// 调用仓库获取数据
	piBehavior, err := s.dashboardRepo.GetPIBehavior(req.Type)
	if err != nil {
		logger.Error("Failed to get PI behavior",
			zap.Error(err),
			zap.String("type", req.Type),
		)
		return &proto.GetPIBehaviorResponse{
			Code:    500,
			Message: "Failed to get PI behavior",
		}, err
	}

	logger.Info("PI behavior retrieved successfully",
		zap.String("type", req.Type),
	)

	// 转换为 proto 响应
	moneyFlow := make([]*proto.MoneyFlow, len(piBehavior.MoneyFlow))
	for i, item := range piBehavior.MoneyFlow {
		moneyFlow[i] = &proto.MoneyFlow{
			From:  item.From,
			To:    item.To,
			Value: int32(item.Value),
		}
	}

	return &proto.GetPIBehaviorResponse{
		Code:    200,
		Message: "success",
		Data: &proto.PIBehavior{
			Type:              piBehavior.Type,
			AvgTradeAmount:    int32(piBehavior.AvgTradeAmount),
			AvgTradeFrequency: int32(piBehavior.AvgTradeFrequency),
			AvgHoldingPeriod:  int32(piBehavior.AvgHoldingPeriod),
			MoneyFlow:         moneyFlow,
		},
	}, nil
}

// GetAccountProcess 获取开户流程
func (s *DashboardService) GetAccountProcess(ctx context.Context, req *proto.GetAccountProcessRequest) (*proto.GetAccountProcessResponse, error) {
	// 调用仓库获取数据
	accountProcess, err := s.dashboardRepo.GetAccountProcess()
	if err != nil {
		logger.Error("Failed to get account process",
			zap.Error(err),
		)
		return &proto.GetAccountProcessResponse{
			Code:    500,
			Message: "Failed to get account process",
		}, err
	}

	logger.Info("Account process retrieved successfully")

	// 转换为 proto 响应
	funnel := make([]*proto.FunnelStage, len(accountProcess.Funnel))
	for i, item := range accountProcess.Funnel {
		funnel[i] = &proto.FunnelStage{
			Stage:      item.Stage,
			Count:      int32(item.Count),
			Conversion: item.Conversion,
		}
	}

	stageTime := make([]*proto.StageTime, len(accountProcess.StageTime))
	for i, item := range accountProcess.StageTime {
		stageTime[i] = &proto.StageTime{
			Stage:      item.Stage,
			AvgTime:    item.AvgTime,
			MedianTime: item.MedianTime,
			P25Time:    item.P25Time,
			P75Time:    item.P75Time,
			MaxTime:    item.MaxTime,
		}
	}

	return &proto.GetAccountProcessResponse{
		Code:    200,
		Message: "success",
		Data: &proto.AccountProcess{
			Funnel:    funnel,
			StageTime: stageTime,
		},
	}, nil
}

// GetAccountException 获取开户异常
func (s *DashboardService) GetAccountException(ctx context.Context, req *proto.GetAccountExceptionRequest) (*proto.GetAccountExceptionResponse, error) {
	// 调用仓库获取数据
	accountException, err := s.dashboardRepo.GetAccountException()
	if err != nil {
		logger.Error("Failed to get account exception",
			zap.Error(err),
		)
		return &proto.GetAccountExceptionResponse{
			Code:    500,
			Message: "Failed to get account exception",
		}, err
	}

	logger.Info("Account exception retrieved successfully")

	// 转换为 proto 响应
	exceptionTypes := make([]*proto.ExceptionType, len(accountException.ExceptionTypes))
	for i, item := range accountException.ExceptionTypes {
		exceptionTypes[i] = &proto.ExceptionType{
			Type:       item.Type,
			Count:      int32(item.Count),
			Percentage: item.Percentage,
		}
	}

	exceptionStages := make([]*proto.ExceptionStage, len(accountException.ExceptionStages))
	for i, item := range accountException.ExceptionStages {
		exceptionStages[i] = &proto.ExceptionStage{
			Stage: item.Stage,
			Count: int32(item.Count),
		}
	}

	exceptionHandling := make([]*proto.ExceptionHandling, len(accountException.ExceptionHandling))
	for i, item := range accountException.ExceptionHandling {
		exceptionHandling[i] = &proto.ExceptionHandling{
			Type:          item.Type,
			AvgHandleTime: item.AvgHandleTime,
			MaxHandleTime: item.MaxHandleTime,
			SuccessRate:   item.SuccessRate,
		}
	}

	return &proto.GetAccountExceptionResponse{
		Code:    200,
		Message: "success",
		Data: &proto.AccountException{
			ExceptionTypes:    exceptionTypes,
			ExceptionStages:   exceptionStages,
			ExceptionHandling: exceptionHandling,
		},
	}, nil
}

// GetIPOSubscription 获取IPO认购情况
func (s *DashboardService) GetIPOSubscription(ctx context.Context, req *proto.GetIPOSubscriptionRequest) (*proto.GetIPOSubscriptionResponse, error) {
	// 调用仓库获取数据
	ipoSubscription, err := s.dashboardRepo.GetIPOSubscription()
	if err != nil {
		logger.Error("Failed to get IPO subscription",
			zap.Error(err),
		)
		return &proto.GetIPOSubscriptionResponse{
			Code:    500,
			Message: "Failed to get IPO subscription",
		}, err
	}

	logger.Info("IPO subscription retrieved successfully")

	// 转换为 proto 响应
	userDistribution := make([]*proto.UserDistribution, len(ipoSubscription.UserDistribution))
	for i, item := range ipoSubscription.UserDistribution {
		userDistribution[i] = &proto.UserDistribution{
			Type:       item.Type,
			Count:      int32(item.Count),
			Percentage: item.Percentage,
		}
	}

	subscriptionAmount := make([]*proto.SubscriptionAmount, len(ipoSubscription.SubscriptionAmount))
	for i, item := range ipoSubscription.SubscriptionAmount {
		subscriptionAmount[i] = &proto.SubscriptionAmount{
			Project:       item.Project,
			PiAmount:      int32(item.PiAmount),
			RegularAmount: int32(item.RegularAmount),
			TotalAmount:   int32(item.TotalAmount),
		}
	}

	subscriptionMultiples := make([]*proto.SubscriptionMultiple, len(ipoSubscription.SubscriptionMultiples))
	for i, item := range ipoSubscription.SubscriptionMultiples {
		subscriptionMultiples[i] = &proto.SubscriptionMultiple{
			Project:  item.Project,
			Multiple: item.Multiple,
		}
	}

	return &proto.GetIPOSubscriptionResponse{
		Code:    200,
		Message: "success",
		Data: &proto.IPOSubscription{
			TotalUsers:            int32(ipoSubscription.TotalUsers),
			UserDistribution:      userDistribution,
			SubscriptionAmount:    subscriptionAmount,
			SubscriptionMultiples: subscriptionMultiples,
		},
	}, nil
}

// GetIPOAnaalysis 获取IPO分析
func (s *DashboardService) GetIPOAnaalysis(ctx context.Context, req *proto.GetIPOAnaalysisRequest) (*proto.GetIPOAnaalysisResponse, error) {
	// 调用仓库获取数据
	ipoAnalysis, err := s.dashboardRepo.GetIPOAnaalysis()
	if err != nil {
		logger.Error("Failed to get IPO analysis",
			zap.Error(err),
		)
		return &proto.GetIPOAnaalysisResponse{
			Code:    500,
			Message: "Failed to get IPO analysis",
		}, err
	}

	logger.Info("IPO analysis retrieved successfully")

	// 转换为 proto 响应
	projectPerformance := make([]*proto.ProjectPerformance, len(ipoAnalysis.ProjectPerformance))
	for i, item := range ipoAnalysis.ProjectPerformance {
		projectPerformance[i] = &proto.ProjectPerformance{
			Project:     item.Project,
			PriceChange: item.PriceChange,
			AvgReturn:   item.AvgReturn,
		}
	}

	influenceFactors := make([]*proto.InfluenceFactor, len(ipoAnalysis.InfluenceFactors))
	for i, item := range ipoAnalysis.InfluenceFactors {
		influenceFactors[i] = &proto.InfluenceFactor{
			Factor:      item.Factor,
			Correlation: item.Correlation,
		}
	}

	competitorComparison := make([]*proto.CompetitorComparison, len(ipoAnalysis.CompetitorComparison))
	for i, item := range ipoAnalysis.CompetitorComparison {
		competitorComparison[i] = &proto.CompetitorComparison{
			Competitor:         item.Competitor,
			SubscriptionUsers:  int32(item.SubscriptionUsers),
			SubscriptionAmount: int32(item.SubscriptionAmount),
			Multiple:           item.Multiple,
		}
	}

	// 创建响应
	response := &proto.GetIPOAnaalysisResponse{
		Code:    200,
		Message: "success",
	}

	// 设置数据
	response.Data = &proto.IPOAnalysis{
		ProjectPerformance:   projectPerformance,
		InfluenceFactors:     influenceFactors,
		CompetitorComparison: competitorComparison,
	}

	return response, nil
}

// GetFinanceCustomer 获取客户融资情况
func (s *DashboardService) GetFinanceCustomer(ctx context.Context, req *proto.GetFinanceCustomerRequest) (*proto.GetFinanceCustomerResponse, error) {
	// 调用仓库获取数据
	financeCustomer, err := s.dashboardRepo.GetFinanceCustomer()
	if err != nil {
		logger.Error("Failed to get finance customer",
			zap.Error(err),
		)
		return &proto.GetFinanceCustomerResponse{
			Code:    500,
			Message: "Failed to get finance customer",
		}, err
	}

	logger.Info("Finance customer retrieved successfully")

	// 转换为 proto 响应
	userDistribution := make([]*proto.UserDistribution, len(financeCustomer.UserDistribution))
	for i, item := range financeCustomer.UserDistribution {
		userDistribution[i] = &proto.UserDistribution{
			Type:       item.Type,
			Count:      int32(item.Count),
			Percentage: item.Percentage,
		}
	}

	financingPurpose := make([]*proto.FinancingPurpose, len(financeCustomer.FinancingPurpose))
	for i, item := range financeCustomer.FinancingPurpose {
		financingPurpose[i] = &proto.FinancingPurpose{
			Purpose:    item.Purpose,
			Percentage: item.Percentage,
		}
	}

	riskAssessment := make([]*proto.RiskAssessment, len(financeCustomer.RiskAssessment))
	for i, item := range financeCustomer.RiskAssessment {
		riskAssessment[i] = &proto.RiskAssessment{
			Type:        item.Type,
			DefaultRate: item.DefaultRate,
			OverdueRate: item.OverdueRate,
		}
	}

	return &proto.GetFinanceCustomerResponse{
		Code:    200,
		Message: "success",
		Data: &proto.FinanceCustomer{
			TotalUsers:       int32(financeCustomer.TotalUsers),
			UserDistribution: userDistribution,
			FinancingSize: &proto.FinancingSize{
				AvgSize:    financeCustomer.FinancingSize.AvgSize,
				MedianSize: financeCustomer.FinancingSize.MedianSize,
				P25Size:    financeCustomer.FinancingSize.P25Size,
				P75Size:    financeCustomer.FinancingSize.P75Size,
				MaxSize:    financeCustomer.FinancingSize.MaxSize,
			},
			FinancingPurpose: financingPurpose,
			RiskAssessment:   riskAssessment,
		},
	}, nil
}

// GetFinanceStock 获取热门股票融资情况
func (s *DashboardService) GetFinanceStock(ctx context.Context, req *proto.GetFinanceStockRequest) (*proto.GetFinanceStockResponse, error) {
	// 调用仓库获取数据
	financeStock, err := s.dashboardRepo.GetFinanceStock()
	if err != nil {
		logger.Error("Failed to get finance stock",
			zap.Error(err),
		)
		return &proto.GetFinanceStockResponse{
			Code:    500,
			Message: "Failed to get finance stock",
		}, err
	}

	logger.Info("Finance stock retrieved successfully")

	// 转换为 proto 响应
	hotStocks := make([]*proto.HotStock, len(financeStock.HotStocks))
	for i, item := range financeStock.HotStocks {
		hotStocks[i] = &proto.HotStock{
			Stock:            item.Stock,
			FinancingBalance: int32(item.FinancingBalance),
			MarketCapRatio:   item.MarketCapRatio,
		}
	}

	financingTrend := make([]*proto.FinancingTrend, len(financeStock.FinancingTrend))
	for i, item := range financeStock.FinancingTrend {
		financingTrend[i] = &proto.FinancingTrend{
			Date:       item.Date,
			BuyAmount:  int32(item.BuyAmount),
			SellAmount: int32(item.SellAmount),
		}
	}

	marginRatio := make([]*proto.MarginRatio, len(financeStock.MarginRatio))
	for i, item := range financeStock.MarginRatio {
		marginRatio[i] = &proto.MarginRatio{
			Stock:           item.Stock,
			Ratio:           item.Ratio,
			PriceVolatility: item.PriceVolatility,
		}
	}

	return &proto.GetFinanceStockResponse{
		Code:    200,
		Message: "success",
		Data: &proto.FinanceStock{
			HotStocks:      hotStocks,
			FinancingTrend: financingTrend,
			MarginRatio:    marginRatio,
		},
	}, nil
}

// GetDrilldownDetail 获取明细数据
func (s *DashboardService) GetDrilldownDetail(ctx context.Context, req *proto.GetDrilldownDetailRequest) (*proto.GetDrilldownDetailResponse, error) {
	// 调用仓库获取数据
	drilldownDetail, err := s.dashboardRepo.GetDrilldownDetail(req.Type, req.Id, int(req.Page), int(req.PageSize))
	if err != nil {
		logger.Error("Failed to get drilldown detail",
			zap.Error(err),
			zap.String("type", req.Type),
			zap.String("id", req.Id),
			zap.Int32("page", req.Page),
			zap.Int32("page_size", req.PageSize),
		)
		return &proto.GetDrilldownDetailResponse{
			Code:    500,
			Message: "Failed to get drilldown detail",
		}, err
	}

	logger.Info("Drilldown detail retrieved successfully",
		zap.String("type", req.Type),
		zap.String("id", req.Id),
		zap.Int32("page", req.Page),
		zap.Int32("page_size", req.PageSize),
	)

	// 转换为 proto 响应
	list := make([]*proto.DrilldownItem, len(drilldownDetail.List))
	for i, item := range drilldownDetail.List {
		list[i] = &proto.DrilldownItem{
			Id:          item.ID,
			CustomerId:  item.CustomerId,
			Name:        item.Name,
			Age:         int32(item.Age),
			Region:      item.Region,
			Asset:       int32(item.Asset),
			TradeAmount: int32(item.TradeAmount),
		}
	}

	return &proto.GetDrilldownDetailResponse{
		Code:    200,
		Message: "success",
		Data: &proto.DrilldownDetail{
			Total:    int32(drilldownDetail.Total),
			Page:     int32(drilldownDetail.Page),
			PageSize: int32(drilldownDetail.PageSize),
			List:     list,
		},
	}, nil
}

// GetDrilldownTrend 获取历史趋势数据
func (s *DashboardService) GetDrilldownTrend(ctx context.Context, req *proto.GetDrilldownTrendRequest) (*proto.GetDrilldownTrendResponse, error) {
	// 调用仓库获取数据
	drilldownTrend, err := s.dashboardRepo.GetDrilldownTrend(req.Type, req.Id, req.Period)
	if err != nil {
		logger.Error("Failed to get drilldown trend",
			zap.Error(err),
			zap.String("type", req.Type),
			zap.String("id", req.Id),
			zap.String("period", req.Period),
		)
		return &proto.GetDrilldownTrendResponse{
			Code:    500,
			Message: "Failed to get drilldown trend",
		}, err
	}

	logger.Info("Drilldown trend retrieved successfully",
		zap.String("type", req.Type),
		zap.String("id", req.Id),
		zap.String("period", req.Period),
	)

	// 转换为 proto 响应
	trend := make([]*proto.TrendItem, len(drilldownTrend.Trend))
	for i, item := range drilldownTrend.Trend {
		trend[i] = &proto.TrendItem{
			Date:  item.Date,
			Value: int32(item.Value),
		}
	}

	return &proto.GetDrilldownTrendResponse{
		Code:    200,
		Message: "success",
		Data: &proto.DrilldownTrend{
			Type:  drilldownTrend.Type,
			Name:  drilldownTrend.Name,
			Trend: trend,
		},
	}, nil
}

// Login 用户登录
func (s *DashboardService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	// 调用仓库获取数据
	loginResponse, err := s.dashboardRepo.Login(req.Username, req.Password)
	if err != nil {
		logger.Error("Failed to login",
			zap.Error(err),
			zap.String("username", req.Username),
		)
		return &proto.LoginResponse{
			Code:    500,
			Message: "Failed to login",
		}, err
	}

	logger.Info("Login successful",
		zap.String("username", req.Username),
		zap.String("user_id", loginResponse.User.ID),
	)

	// 转换为 proto 响应
	return &proto.LoginResponse{
		Code:    200,
		Message: "success",
		Data: &proto.LoginData{
			Token: loginResponse.Token,
			User: &proto.User{
				Id:       loginResponse.User.ID,
				Username: loginResponse.User.Username,
				Role:     loginResponse.User.Role,
			},
		},
	}, nil
}

// GetAuthInfo 获取用户信息
func (s *DashboardService) GetAuthInfo(ctx context.Context, req *proto.GetAuthInfoRequest) (*proto.GetAuthInfoResponse, error) {
	// 调用仓库获取数据
	user, err := s.dashboardRepo.GetAuthInfo()
	if err != nil {
		logger.Error("Failed to get auth info",
			zap.Error(err),
		)
		return &proto.GetAuthInfoResponse{
			Code:    500,
			Message: "Failed to get auth info",
		}, err
	}

	logger.Info("Auth info retrieved successfully",
		zap.String("user_id", user.ID),
		zap.String("username", user.Username),
	)

	// 转换为 proto 响应
	return &proto.GetAuthInfoResponse{
		Code:    200,
		Message: "success",
		Data: &proto.User{
			Id:          user.ID,
			Username:    user.Username,
			Role:        user.Role,
			Permissions: user.Permissions,
		},
	}, nil
}
