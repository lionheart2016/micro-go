package server

import (
	"context"
	"dashboard-data/grpc/proto"
	"dashboard-data/repository"
	"log"

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
	log.Println("Dashboard Data Service registered")
}

// GetCoreIndicators 获取核心指标
func (s *DashboardService) GetCoreIndicators(ctx context.Context, req *proto.GetCoreIndicatorsRequest) (*proto.GetCoreIndicatorsResponse, error) {
	// 调用仓库获取数据
	coreIndicators, err := s.dashboardRepo.GetCoreIndicators(req.Period)
	if err != nil {
		log.Printf("Failed to get core indicators: %v", err)
		return &proto.GetCoreIndicatorsResponse{
			Code:    500,
			Message: "Failed to get core indicators",
		}, err
	}

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
		log.Printf("Failed to get performance: %v", err)
		return &proto.GetPerformanceResponse{
			Code:    500,
			Message: "Failed to get performance",
		}, err
	}

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
			StockRevenue: int32(performance.StockRevenue),
			FundRevenue:  int32(performance.FundRevenue),
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

