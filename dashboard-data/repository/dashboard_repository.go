package repository

import (
	"dashboard-data/model"
)

// DashboardRepository 仪表盘数据仓库接口
type DashboardRepository interface {
	GetCoreIndicators(period string) (*model.CoreIndicators, error)
	GetPerformance(period string) (*model.Performance, error)
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
