#!/bin/zsh

# 测试配置
BASE_URL="http://localhost"
LOG_FILE="test_results.log"

# 测试结果统计
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# 颜色定义
GREEN="\033[0;32m"
RED="\033[0;31m"
YELLOW="\033[0;33m"
NC="\033[0m" # No Color

# 初始化测试日志
echo "========================================" > $LOG_FILE
echo "HTTP API测试报告 - $(date)" >> $LOG_FILE
echo "========================================" >> $LOG_FILE
echo "" >> $LOG_FILE

# 测试函数
test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    local expected_status=${4:-200}
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    echo "测试: $method $endpoint" | tee -a $LOG_FILE
    
    if [[ $method == "GET" ]]; then
        response=$(curl -s -w "\n%{http_code}" "$BASE_URL$endpoint")
    else
        response=$(curl -s -w "\n%{http_code}" -X $method -H "Content-Type: application/json" -d "$data" "$BASE_URL$endpoint")
    fi
    
    # 分离响应体和状态码
    status_code=$(echo "$response" | tail -n 1)
    response_body=$(echo "$response" | sed '$d')
    
    # 检查状态码
    if [[ "$status_code" == "$expected_status" ]]; then
        echo "${GREEN}✓ 状态码: $status_code (预期: $expected_status)${NC}" | tee -a $LOG_FILE
        
        # 检查响应格式是否为JSON
        if echo "$response_body" | jq . > /dev/null 2>&1; then
            echo "${GREEN}✓ 响应格式: JSON${NC}" | tee -a $LOG_FILE
            
            # 检查响应是否包含数据
            if [[ $(echo "$response_body" | jq 'length') -gt 0 ]]; then
                echo "${GREEN}✓ 响应包含数据${NC}" | tee -a $LOG_FILE
                PASSED_TESTS=$((PASSED_TESTS + 1))
            else
                echo "${YELLOW}⚠ 响应为空${NC}" | tee -a $LOG_FILE
                FAILED_TESTS=$((FAILED_TESTS + 1))
            fi
        else
            echo "${RED}✗ 响应格式不是JSON${NC}" | tee -a $LOG_FILE
            echo "响应体: $response_body" | tee -a $LOG_FILE
            FAILED_TESTS=$((FAILED_TESTS + 1))
        fi
    else
        echo "${RED}✗ 状态码: $status_code (预期: $expected_status)${NC}" | tee -a $LOG_FILE
        echo "响应体: $response_body" | tee -a $LOG_FILE
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
    
    echo "" | tee -a $LOG_FILE
}

# 测试核心指标相关接口
echo "=== 核心指标相关接口测试 ===" | tee -a $LOG_FILE
test_endpoint "GET" "/api/dashboard/core-indicators"
test_endpoint "GET" "/api/dashboard/performance"

# 测试客户分析相关接口
echo "=== 客户分析相关接口测试 ===" | tee -a $LOG_FILE
test_endpoint "GET" "/api/customer/distribution"
test_endpoint "GET" "/api/customer/behavior"

# 测试交易分析相关接口
echo "=== 交易分析相关接口测试 ===" | tee -a $LOG_FILE
test_endpoint "GET" "/api/trade/stock"
test_endpoint "GET" "/api/trade/fund"

# 测试PI用户分析相关接口
echo "=== PI用户分析相关接口测试 ===" | tee -a $LOG_FILE
test_endpoint "GET" "/api/pi/basic"
test_endpoint "GET" "/api/pi/behavior"

# 测试开户主题相关接口
echo "=== 开户主题相关接口测试 ===" | tee -a $LOG_FILE
test_endpoint "GET" "/api/account/process"
test_endpoint "GET" "/api/account/exception"

# 测试IPO主题相关接口
echo "=== IPO主题相关接口测试 ===" | tee -a $LOG_FILE
test_endpoint "GET" "/api/ipo/subscription"
test_endpoint "GET" "/api/ipo/analysis"

# 测试融资主题相关接口
echo "=== 融资主题相关接口测试 ===" | tee -a $LOG_FILE
test_endpoint "GET" "/api/finance/customer"
test_endpoint "GET" "/api/finance/stock"

# 测试数据钻取相关接口
echo "=== 数据钻取相关接口测试 ===" | tee -a $LOG_FILE
test_endpoint "GET" "/api/drilldown/detail"
test_endpoint "GET" "/api/drilldown/trend"

# 测试认证相关接口
echo "=== 认证相关接口测试 ===" | tee -a $LOG_FILE
test_endpoint "POST" "/api/auth/login" '{"username":"test", "password":"test"}'
test_endpoint "GET" "/api/auth/info"

# 生成测试报告
echo "========================================" | tee -a $LOG_FILE
echo "测试报告摘要" | tee -a $LOG_FILE
echo "========================================" | tee -a $LOG_FILE
echo "总测试用例数: $TOTAL_TESTS" | tee -a $LOG_FILE
echo "通过测试数: ${GREEN}$PASSED_TESTS${NC}" | tee -a $LOG_FILE
echo "失败测试数: ${RED}$FAILED_TESTS${NC}" | tee -a $LOG_FILE
echo "" | tee -a $LOG_FILE

# 计算成功率
if [[ $TOTAL_TESTS -gt 0 ]]; then
    SUCCESS_RATE=$((PASSED_TESTS * 100 / TOTAL_TESTS))
    echo "测试成功率: $SUCCESS_RATE%" | tee -a $LOG_FILE
fi

echo "测试报告已保存到: $LOG_FILE"
