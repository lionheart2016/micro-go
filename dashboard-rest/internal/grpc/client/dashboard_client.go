package client

import (
	"context"
	"dashboard-rest/internal/grpc/proto"
	"dashboard-rest/internal/model"
	"dashboard-rest/pkg/logger"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DashboardClient struct {
	logger  *logger.Logger
	client  proto.DashboardServiceClient
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

	// 创建生成的gRPC客户端
	client := proto.NewDashboardServiceClient(conn)

	return &DashboardClient{
		logger:  logger,
		client:  client,
		conn:    conn,
	}, nil
}

func (c *DashboardClient) GetCoreIndicators(ctx context.Context, period string) (*model.Response, error) {
	c.logger.Info("Getting core indicators from gRPC server")

	// 调用生成的gRPC方法
	response, err := c.client.GetCoreIndicators(ctx, &proto.GetCoreIndicatorsRequest{
		Period: period,
	})
	if err != nil {
		c.logger.Error("Failed to get core indicators: %v", err)
		return nil, err
	}

	// 转换为模型结构
	return &model.Response{
		Code:    int(response.Code),
		Message: response.Message,
		Data: model.CoreIndicators{
			RegisteredUsers: model.Indicator{
				Value:      int(response.Data.GetRegisteredUsers().GetValue()),
				Change:     int(response.Data.GetRegisteredUsers().GetChange()),
				ChangeRate: response.Data.GetRegisteredUsers().GetChangeRate(),
			},
			AccountCount: model.Indicator{
				Value:      int(response.Data.GetAccountCount().GetValue()),
				Change:     int(response.Data.GetAccountCount().GetChange()),
				ChangeRate: response.Data.GetAccountCount().GetChangeRate(),
			},
			ActiveUsers: model.Indicator{
				Value:      int(response.Data.GetActiveUsers().GetValue()),
				Change:     int(response.Data.GetActiveUsers().GetChange()),
				ChangeRate: response.Data.GetActiveUsers().GetChangeRate(),
			},
			DepositUsers: model.Indicator{
				Value:      int(response.Data.GetDepositUsers().GetValue()),
				Change:     int(response.Data.GetDepositUsers().GetChange()),
				ChangeRate: response.Data.GetDepositUsers().GetChangeRate(),
			},
			DepositAmount: model.Indicator{
				Value:      int(response.Data.GetDepositAmount().GetValue()),
				Change:     int(response.Data.GetDepositAmount().GetChange()),
				ChangeRate: response.Data.GetDepositAmount().GetChangeRate(),
			},
			StockTradeUsers: model.Indicator{
				Value:      int(response.Data.GetStockTradeUsers().GetValue()),
				Change:     int(response.Data.GetStockTradeUsers().GetChange()),
				ChangeRate: response.Data.GetStockTradeUsers().GetChangeRate(),
			},
			StockTradeAmount: model.Indicator{
				Value:      int(response.Data.GetStockTradeAmount().GetValue()),
				Change:     int(response.Data.GetStockTradeAmount().GetChange()),
				ChangeRate: response.Data.GetStockTradeAmount().GetChangeRate(),
			},
			FundTradeUsers: model.Indicator{
				Value:      int(response.Data.GetFundTradeUsers().GetValue()),
				Change:     int(response.Data.GetFundTradeUsers().GetChange()),
				ChangeRate: response.Data.GetFundTradeUsers().GetChangeRate(),
			},
			FundTradeAmount: model.Indicator{
				Value:      int(response.Data.GetFundTradeAmount().GetValue()),
				Change:     int(response.Data.GetFundTradeAmount().GetChange()),
				ChangeRate: response.Data.GetFundTradeAmount().GetChangeRate(),
			},
		},
	}, nil
}

func (c *DashboardClient) GetPerformance(ctx context.Context, period string) (*model.Response, error) {
	c.logger.Info("Getting performance from gRPC server")

	// 调用生成的gRPC方法
	response, err := c.client.GetPerformance(ctx, &proto.GetPerformanceRequest{
		Period: period,
	})
	if err != nil {
		c.logger.Error("Failed to get performance: %v", err)
		return nil, err
	}

	// 转换为模型结构
	revenueStructure := make([]model.RevenueItem, len(response.Data.GetRevenueStructure()))
	for i, item := range response.Data.GetRevenueStructure() {
		revenueStructure[i] = model.RevenueItem{
			Name:  item.GetName(),
			Value: int(item.GetValue()),
		}
	}

	revenueTrend := make([]model.RevenueTrend, len(response.Data.GetRevenueTrend()))
	for i, trend := range response.Data.GetRevenueTrend() {
		revenueTrend[i] = model.RevenueTrend{
			Month: trend.GetMonth(),
			Stock: int(trend.GetStock()),
			Fund:  int(trend.GetFund()),
		}
	}

	depositTrend := make([]model.DepositTrend, len(response.Data.GetDepositTrend()))
	for i, trend := range response.Data.GetDepositTrend() {
		depositTrend[i] = model.DepositTrend{
			Month: trend.GetMonth(),
			Value: int(trend.GetValue()),
		}
	}

	return &model.Response{
		Code:    int(response.Code),
		Message: response.Message,
		Data: model.Performance{
			StockRevenue: int(response.Data.GetStockRevenue()),
			FundRevenue:  int(response.Data.GetFundRevenue()),
			RevenueStructure: revenueStructure,
			RevenueTrend:     revenueTrend,
			RevenueChange: model.RevenueChange{
				Yoy: response.Data.GetRevenueChange().GetYoy(),
				Mom: response.Data.GetRevenueChange().GetMom(),
			},
			DepositNetGrowth:       int(response.Data.GetDepositNetGrowth()),
			DepositTrend:           depositTrend,
			DepositWithdrawalRatio: response.Data.GetDepositWithdrawalRatio(),
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
