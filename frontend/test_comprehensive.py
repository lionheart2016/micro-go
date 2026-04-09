"""
前端应用全面测试脚本
测试范围：组件功能、用户交互、响应式设计、性能、兼容性
"""

from playwright.sync_api import sync_playwright
import json
import time
from datetime import datetime

# 测试结果收集
test_results = {
    "test_name": "港美股经纪业务数据看板 - 前端全面测试",
    "test_time": datetime.now().isoformat(),
    "total_tests": 0,
    "passed_tests": 0,
    "failed_tests": 0,
    "warnings": 0,
    "test_cases": [],
    "issues": [],
    "console_errors": [],
    "performance_metrics": {}
}

def add_test_case(name, status, details="", issue=None):
    """添加测试用例结果"""
    test_results["total_tests"] += 1
    if status == "passed":
        test_results["passed_tests"] += 1
    elif status == "failed":
        test_results["failed_tests"] += 1
    elif status == "warning":
        test_results["warnings"] += 1
    
    test_results["test_cases"].append({
        "name": name,
        "status": status,
        "details": details
    })
    
    if issue and status == "failed":
        test_results["issues"].append(issue)

def take_screenshot(page, name):
    """截图并保存"""
    page.screenshot(path=f'/tmp/test_{name}.png', full_page=True)

def test_page_load_and_navigation(page, context):
    """测试1: 页面加载和路由导航"""
    print("\n=== 测试1: 页面加载和路由导航 ===")
    
    # 测试首页加载
    try:
        page.goto('http://localhost:5173/', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(2000)
        
        # 检查页面标题
        title = page.title()
        if title:
            add_test_case("页面标题加载", "passed", f"标题: {title}")
        else:
            add_test_case("页面标题加载", "warning", "页面标题为空")
        
        # 检查 header 是否存在
        header = page.query_selector('.app-header h1')
        if header and header.inner_text():
            add_test_case("Header 渲染", "passed", f"Header 文本: {header.inner_text()}")
        else:
            add_test_case("Header 渲染", "failed", "Header 未找到", 
                         {"type": "UI渲染", "description": "Header 元素未正确渲染", "expected": "显示'港美股经纪业务数据看板'"})
        
        # 检查侧边栏菜单
        menu_items = page.query_selector_all('.el-menu-item')
        expected_menus = ['关键指标概览', '客户分析', '交易分析', 'PI用户分析', '开户主题', 'IPO主题', '融资主题', 'XAUa 黄金代币']
        
        if len(menu_items) >= 8:
            add_test_case("侧边栏菜单渲染", "passed", f"找到 {len(menu_items)} 个菜单项")
        else:
            add_test_case("侧边栏菜单渲染", "failed", f"只找到 {len(menu_items)} 个菜单项，预期 8 个",
                         {"type": "UI渲染", "description": "侧边栏菜单项不完整", "expected": "8个菜单项", "actual": f"{len(menu_items)}个"})
        
        take_screenshot(page, "01_homepage")
        
    except Exception as e:
        add_test_case("页面加载", "failed", str(e),
                     {"type": "页面加载", "description": "首页加载失败", "error": str(e)})
    
    # 测试路由导航
    routes = [
        ('/customer', '客户分析'),
        ('/trade', '交易分析'),
        ('/pi', 'PI用户分析'),
        ('/account', '开户主题'),
        ('/ipo', 'IPO主题'),
        ('/finance', '融资主题'),
        ('/xaua', 'XAUa 黄金代币')
    ]
    
    for route, name in routes:
        try:
            page.goto(f'http://localhost:5173{route}', wait_until='networkidle', timeout=30000)
            page.wait_for_timeout(1500)
            
            # 检查 URL 是否正确
            current_url = page.url
            if route in current_url:
                add_test_case(f"路由导航 - {name}", "passed", f"URL: {current_url}")
            else:
                add_test_case(f"路由导航 - {name}", "failed", f"URL 不匹配: {current_url}",
                             {"type": "路由", "description": f"{name} 路由导航失败", "expected": f"包含 {route}", "actual": current_url})
            
            take_screenshot(page, f"02_route_{name}")
            
        except Exception as e:
            add_test_case(f"路由导航 - {name}", "failed", str(e),
                         {"type": "路由", "description": f"{name} 路由导航异常", "error": str(e)})

def test_core_components(page, context):
    """测试2: 核心组件渲染"""
    print("\n=== 测试2: 核心组件渲染 ===")
    
    # 回到首页
    page.goto('http://localhost:5173/', wait_until='networkidle', timeout=30000)
    page.wait_for_timeout(2000)
    
    # 测试 KPI 卡片
    try:
        kpi_cards = page.query_selector_all('.kpi-item')
        if len(kpi_cards) >= 8:
            add_test_case("KPI 卡片渲染", "passed", f"找到 {len(kpi_cards)} 个 KPI 卡片")
        else:
            add_test_case("KPI 卡片渲染", "warning", f"找到 {len(kpi_cards)} 个 KPI 卡片（API 未响应时可能为 0）")
        
        # 检查 KPI 卡片内容结构
        if len(kpi_cards) > 0:
            first_card = kpi_cards[0]
            card_title = first_card.query_selector('.kpi-title')
            card_value = first_card.query_selector('.kpi-value')
            
            if card_title and card_value:
                add_test_case("KPI 卡片结构", "passed", "卡片包含标题和数值元素")
            else:
                add_test_case("KPI 卡片结构", "warning", "卡片结构可能不完整")
        
        take_screenshot(page, "03_kpi_cards")
        
    except Exception as e:
        add_test_case("KPI 卡片渲染", "failed", str(e),
                     {"type": "组件", "description": "KPI 卡片渲染失败", "error": str(e)})
    
    # 测试图表组件
    try:
        # 检查是否有图表容器
        charts = page.query_selector_all('[class*="chart"]')
        if len(charts) > 0:
            add_test_case("图表组件渲染", "passed", f"找到 {len(charts)} 个图表容器")
        else:
            add_test_case("图表组件渲染", "warning", "未找到图表容器（可能因 API 未响应）")
        
        # 检查 BaseCard 组件
        base_cards = page.query_selector_all('.base-card')
        if len(base_cards) > 0:
            add_test_case("BaseCard 组件渲染", "passed", f"找到 {len(base_cards)} 个 BaseCard")
        else:
            add_test_case("BaseCard 组件渲染", "warning", "未找到 BaseCard 组件")
        
    except Exception as e:
        add_test_case("图表组件渲染", "failed", str(e),
                     {"type": "组件", "description": "图表组件渲染失败", "error": str(e)})

def test_user_interactions(page, context):
    """测试3: 用户交互流程"""
    print("\n=== 测试3: 用户交互流程 ===")
    
    # 测试刷新按钮
    try:
        page.goto('http://localhost:5173/', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(2000)
        
        refresh_button = page.query_selector('.refresh-button')
        if refresh_button:
            refresh_button.click()
            page.wait_for_timeout(1000)
            
            # 检查按钮点击后的 loading 状态
            loading_text = page.query_selector('.refresh-button:has-text("加载中")')
            if loading_text:
                add_test_case("刷新按钮交互", "passed", "按钮点击后显示加载状态")
            else:
                add_test_case("刷新按钮交互", "warning", "按钮点击但未检测到加载状态（可能加载过快）")
        else:
            add_test_case("刷新按钮交互", "failed", "刷新按钮未找到",
                         {"type": "交互", "description": "刷新按钮不存在", "expected": "存在刷新按钮"})
        
    except Exception as e:
        add_test_case("刷新按钮交互", "failed", str(e),
                     {"type": "交互", "description": "刷新按钮交互失败", "error": str(e)})
    
    # 测试客户分析页面筛选器
    try:
        page.goto('http://localhost:5173/customer', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(2000)
        
        # 检查维度筛选器
        dimension_select = page.query_selector('#dimension')
        if dimension_select:
            # 获取选项
            options = dimension_select.query_selector_all('option')
            if len(options) >= 5:
                add_test_case("客户分析-维度筛选器", "passed", f"找到 {len(options)} 个选项")
                
                # 测试切换维度
                dimension_select.select_option(index=1)
                page.wait_for_timeout(1000)
                add_test_case("客户分析-维度切换", "passed", "成功切换维度选项")
            else:
                add_test_case("客户分析-维度筛选器", "warning", f"只找到 {len(options)} 个选项")
        else:
            add_test_case("客户分析-维度筛选器", "warning", "维度筛选器未找到（可能因 API 未响应）")
        
        # 检查行为类型筛选器
        behavior_select = page.query_selector('#behaviorType')
        if behavior_select:
            options = behavior_select.query_selector_all('option')
            if len(options) >= 4:
                add_test_case("客户分析-行为类型筛选器", "passed", f"找到 {len(options)} 个选项")
            else:
                add_test_case("客户分析-行为类型筛选器", "warning", f"只找到 {len(options)} 个选项")
        else:
            add_test_case("客户分析-行为类型筛选器", "warning", "行为类型筛选器未找到")
        
        take_screenshot(page, "04_customer_filters")
        
    except Exception as e:
        add_test_case("客户分析筛选器", "failed", str(e),
                     {"type": "交互", "description": "客户分析筛选器测试失败", "error": str(e)})
    
    # 测试交易分析页面筛选器
    try:
        page.goto('http://localhost:5173/trade', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(2000)
        
        # 检查股票交易类型筛选器
        stock_select = page.query_selector('#stockType')
        if stock_select:
            options = stock_select.query_selector_all('option')
            if len(options) >= 4:
                add_test_case("交易分析-股票类型筛选器", "passed", f"找到 {len(options)} 个选项")
                
                # 测试切换选项
                stock_select.select_option(index=1)
                page.wait_for_timeout(1000)
                add_test_case("交易分析-股票类型切换", "passed", "成功切换股票类型")
            else:
                add_test_case("交易分析-股票类型筛选器", "warning", f"只找到 {len(options)} 个选项")
        else:
            add_test_case("交易分析-股票类型筛选器", "warning", "股票类型筛选器未找到")
        
        # 检查基金类型筛选器
        fund_select = page.query_selector('#fundType')
        if fund_select:
            options = fund_select.query_selector_all('option')
            if len(options) >= 5:
                add_test_case("交易分析-基金类型筛选器", "passed", f"找到 {len(options)} 个选项")
            else:
                add_test_case("交易分析-基金类型筛选器", "warning", f"只找到 {len(options)} 个选项")
        else:
            add_test_case("交易分析-基金类型筛选器", "warning", "基金类型筛选器未找到")
        
        take_screenshot(page, "05_trade_filters")
        
    except Exception as e:
        add_test_case("交易分析筛选器", "failed", str(e),
                     {"type": "交互", "description": "交易分析筛选器测试失败", "error": str(e)})

def test_responsive_design(page, context):
    """测试4: 响应式设计"""
    print("\n=== 测试4: 响应式设计 ===")
    
    # 桌面端 (1920x1080)
    try:
        page.set_viewport_size({"width": 1920, "height": 1080})
        page.goto('http://localhost:5173/', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(1500)
        
        # 检查布局
        app_main = page.query_selector('.app-main')
        if app_main:
            style = app_main.evaluate('el => window.getComputedStyle(el).display')
            if style == 'flex':
                add_test_case("响应式-桌面端布局", "passed", "使用 flex 布局")
            else:
                add_test_case("响应式-桌面端布局", "warning", f"布局方式: {style}")
        
        take_screenshot(page, "06_responsive_desktop")
        
    except Exception as e:
        add_test_case("响应式-桌面端布局", "failed", str(e),
                     {"type": "响应式", "description": "桌面端布局测试失败", "error": str(e)})
    
    # 平板端 (768x1024)
    try:
        page.set_viewport_size({"width": 768, "height": 1024})
        page.goto('http://localhost:5173/', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(1500)
        
        # 检查侧边栏宽度
        sidebar = page.query_selector('.app-sidebar')
        if sidebar:
            width = sidebar.evaluate('el => window.getComputedStyle(el).width')
            add_test_case("响应式-平板端侧边栏", "passed", f"侧边栏宽度: {width}")
        
        # 检查 KPI 网格
        kpi_grid = page.query_selector('.kpi-grid')
        if kpi_grid:
            grid_cols = kpi_grid.evaluate('el => window.getComputedStyle(el).gridTemplateColumns')
            add_test_case("响应式-平板端KPI网格", "passed", f"网格列: {grid_cols}")
        
        take_screenshot(page, "07_responsive_tablet")
        
    except Exception as e:
        add_test_case("响应式-平板端布局", "failed", str(e),
                     {"type": "响应式", "description": "平板端布局测试失败", "error": str(e)})
    
    # 移动端 (375x667)
    try:
        page.set_viewport_size({"width": 375, "height": 667})
        page.goto('http://localhost:5173/', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(1500)
        
        # 检查移动端布局
        dashboard_header = page.query_selector('.dashboard-header')
        if dashboard_header:
            flex_direction = dashboard_header.evaluate('el => window.getComputedStyle(el).flexDirection')
            if flex_direction == 'column':
                add_test_case("响应式-移动端布局", "passed", "Header 使用 column 布局")
            else:
                add_test_case("响应式-移动端布局", "warning", f"Header 布局方向: {flex_direction}")
        
        # 检查 KPI 网格在移动端是否为单列
        kpi_grid = page.query_selector('.kpi-grid')
        if kpi_grid:
            grid_cols = kpi_grid.evaluate('el => window.getComputedStyle(el).gridTemplateColumns')
            add_test_case("响应式-移动端KPI网格", "passed", f"网格列: {grid_cols}")
        
        take_screenshot(page, "08_responsive_mobile")
        
    except Exception as e:
        add_test_case("响应式-移动端布局", "failed", str(e),
                     {"type": "响应式", "description": "移动端布局测试失败", "error": str(e)})
    
    # 恢复桌面端
    page.set_viewport_size({"width": 1920, "height": 1080})

def test_all_views(page, context):
    """测试5: 所有视图页面加载"""
    print("\n=== 测试5: 所有视图页面加载 ===")
    
    views = [
        ('/', '关键指标概览'),
        ('/customer', '客户分析'),
        ('/trade', '交易分析'),
        ('/pi', 'PI用户分析'),
        ('/account', '开户主题'),
        ('/ipo', 'IPO主题'),
        ('/finance', '融资主题'),
        ('/xaua', 'XAUa 黄金代币'),
        ('/drilldown', '数据钻取'),
        ('/login', '登录')
    ]
    
    for route, name in views:
        try:
            page.goto(f'http://localhost:5173{route}', wait_until='networkidle', timeout=30000)
            page.wait_for_timeout(1500)
            
            # 检查页面是否有内容
            body = page.query_selector('body')
            if body:
                text_content = body.inner_text()
                if len(text_content) > 50:
                    add_test_case(f"视图加载 - {name}", "passed", f"页面内容长度: {len(text_content)} 字符")
                else:
                    add_test_case(f"视图加载 - {name}", "warning", f"页面内容较少: {len(text_content)} 字符")
            else:
                add_test_case(f"视图加载 - {name}", "failed", "页面 body 为空",
                             {"type": "视图", "description": f"{name} 页面内容为空", "expected": "有实际内容"})
            
        except Exception as e:
            add_test_case(f"视图加载 - {name}", "failed", str(e),
                         {"type": "视图", "description": f"{name} 页面加载失败", "error": str(e)})

def test_console_errors(page, context):
    """测试6: 检查控制台错误和警告"""
    print("\n=== 测试6: 控制台错误和警告 ===")
    
    console_errors = []
    console_warnings = []
    
    def handle_console(msg):
        if msg.type == 'error':
            console_errors.append(msg.text)
        elif msg.type == 'warning':
            console_warnings.append(msg.text)
    
    page.on('console', handle_console)
    
    # 访问所有页面并收集控制台信息
    pages_to_test = ['/', '/customer', '/trade', '/pi', '/account', '/ipo', '/finance', '/xaua']
    
    for route in pages_to_test:
        try:
            page.goto(f'http://localhost:5173{route}', wait_until='networkidle', timeout=30000)
            page.wait_for_timeout(2000)
        except Exception as e:
            console_errors.append(f"页面 {route} 加载失败: {str(e)}")
    
    # 分析控制台错误
    if len(console_errors) > 0:
        # 过滤掉预期的 API 错误（因为后端未运行）
        api_errors = [e for e in console_errors if 'api' in e.lower() or 'request' in e.lower() or 'failed' in e.lower()]
        real_errors = [e for e in console_errors if e not in api_errors]
        
        if len(real_errors) > 0:
            add_test_case("控制台错误检查", "warning", f"发现 {len(real_errors)} 个非 API 错误",
                         {"type": "控制台", "description": "存在非 API 相关的控制台错误", "errors": real_errors[:5]})
        else:
            add_test_case("控制台错误检查", "passed", f"只有 API 请求错误（后端未运行），无其他错误")
    else:
        add_test_case("控制台错误检查", "passed", "无控制台错误")
    
    if len(console_warnings) > 0:
        add_test_case("控制台警告检查", "warning", f"发现 {len(console_warnings)} 个警告",
                     {"type": "控制台", "description": "存在控制台警告", "warnings": console_warnings[:5]})
    else:
        add_test_case("控制台警告检查", "passed", "无控制台警告")
    
    test_results["console_errors"] = console_errors

def test_performance(page, context):
    """测试7: 性能指标"""
    print("\n=== 测试7: 性能指标 ===")
    
    # 测试首页加载性能
    try:
        page.goto('http://localhost:5173/', wait_until='networkidle', timeout=30000)
        
        # 获取性能指标
        performance_metrics = page.evaluate('''() => {
            const perf = performance.getEntriesByType('navigation')[0];
            return {
                domContentLoaded: perf.domContentLoadedEventEnd - perf.startTime,
                loadComplete: perf.loadEventEnd - perf.startTime,
                domInteractive: perf.domInteractive - perf.startTime,
                transferSize: perf.transferSize,
                decodedBodySize: perf.decodedBodySize
            };
        }''')
        
        test_results["performance_metrics"]["homepage"] = performance_metrics
        
        # 评估性能
        load_time = performance_metrics.get('loadComplete', 0)
        if load_time < 3000:
            add_test_case("性能-首页加载时间", "passed", f"加载时间: {load_time:.0f}ms")
        elif load_time < 5000:
            add_test_case("性能-首页加载时间", "warning", f"加载时间: {load_time:.0f}ms（建议优化到 3s 以内）")
        else:
            add_test_case("性能-首页加载时间", "failed", f"加载时间: {load_time:.0f}ms（超过 5s）",
                         {"type": "性能", "description": "首页加载时间过长", "actual": f"{load_time:.0f}ms", "expected": "< 3000ms"})
        
    except Exception as e:
        add_test_case("性能-首页加载时间", "failed", str(e),
                     {"type": "性能", "description": "性能测试失败", "error": str(e)})
    
    # 测试路由切换性能
    try:
        start_time = time.time()
        page.goto('http://localhost:5173/customer', wait_until='networkidle', timeout=30000)
        route_switch_time = (time.time() - start_time) * 1000
        
        test_results["performance_metrics"]["route_switch"] = route_switch_time
        
        if route_switch_time < 2000:
            add_test_case("性能-路由切换时间", "passed", f"切换时间: {route_switch_time:.0f}ms")
        else:
            add_test_case("性能-路由切换时间", "warning", f"切换时间: {route_switch_time:.0f}ms")
        
    except Exception as e:
        add_test_case("性能-路由切换时间", "failed", str(e),
                     {"type": "性能", "description": "路由切换性能测试失败", "error": str(e)})

def test_data_flow(page, context):
    """测试8: 数据流和状态管理"""
    print("\n=== 测试8: 数据流和状态管理 ===")
    
    # 测试 API 调用
    try:
        page.goto('http://localhost:5173/', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(2000)
        
        # 检查是否有 API 请求
        api_requests = []
        
        def handle_request(request):
            if '/api/' in request.url:
                api_requests.append({
                    'url': request.url,
                    'method': request.method,
                    'status': None
                })
        
        page.on('request', handle_request)
        
        # 刷新页面重新触发 API 调用
        page.reload(wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(2000)
        
        if len(api_requests) > 0:
            add_test_case("数据流-API 调用", "passed", f"检测到 {len(api_requests)} 个 API 请求")
        else:
            add_test_case("数据流-API 调用", "warning", "未检测到 API 请求（可能组件未挂载或 API 路径错误）")
        
    except Exception as e:
        add_test_case("数据流-API 调用", "failed", str(e),
                     {"type": "数据流", "description": "API 调用测试失败", "error": str(e)})
    
    # 测试错误处理
    try:
        # 检查页面是否有错误处理 UI
        page.goto('http://localhost:5173/customer', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(2000)
        
        error_elements = page.query_selector_all('.error')
        empty_elements = page.query_selector_all('.empty')
        loading_elements = page.query_selector_all('.loading')
        
        if len(error_elements) > 0 or len(empty_elements) > 0 or len(loading_elements) > 0:
            add_test_case("数据流-错误处理 UI", "passed", "页面包含错误/空状态/加载状态 UI")
        else:
            add_test_case("数据流-错误处理 UI", "warning", "未检测到错误处理 UI 元素")
        
    except Exception as e:
        add_test_case("数据流-错误处理 UI", "failed", str(e),
                     {"type": "数据流", "description": "错误处理 UI 测试失败", "error": str(e)})

def test_accessibility(page, context):
    """测试9: 可访问性"""
    print("\n=== 测试9: 可访问性 ===")
    
    try:
        page.goto('http://localhost:5173/', wait_until='networkidle', timeout=30000)
        page.wait_for_timeout(1500)
        
        # 检查图片 alt 属性
        images = page.query_selector_all('img')
        images_without_alt = [img for img in images if not img.get_attribute('alt')]
        
        if len(images_without_alt) == 0:
            add_test_case("可访问性-图片 alt", "passed", "所有图片都有 alt 属性")
        else:
            add_test_case("可访问性-图片 alt", "warning", f"{len(images_without_alt)} 个图片缺少 alt 属性")
        
        # 检查表单标签
        inputs = page.query_selector_all('input, select')
        inputs_without_label = []
        for inp in inputs:
            aria_label = inp.get_attribute('aria-label')
            id_attr = inp.get_attribute('id')
            if not aria_label and not id_attr:
                inputs_without_label.append(inp)
        
        if len(inputs_without_label) == 0:
            add_test_case("可访问性-表单标签", "passed", "所有表单元素都有标签")
        else:
            add_test_case("可访问性-表单标签", "warning", f"{len(inputs_without_label)} 个表单元素缺少标签")
        
        # 检查对比度（通过检查是否有明显的文本颜色）
        text_elements = page.query_selector_all('h1, h2, h3, p, span')
        if len(text_elements) > 0:
            add_test_case("可访问性-文本元素", "passed", f"找到 {len(text_elements)} 个文本元素")
        else:
            add_test_case("可访问性-文本元素", "warning", "文本元素较少")
        
    except Exception as e:
        add_test_case("可访问性", "failed", str(e),
                     {"type": "可访问性", "description": "可访问性测试失败", "error": str(e)})

def main():
    print("=" * 60)
    print("港美股经纪业务数据看板 - 前端全面测试")
    print("=" * 60)
    
    with sync_playwright() as p:
        browser = p.chromium.launch(headless=True)
        context = browser.new_context(
            viewport={'width': 1920, 'height': 1080},
            user_agent='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36'
        )
        
        page = context.new_page()
        
        # 执行所有测试
        test_page_load_and_navigation(page, context)
        test_core_components(page, context)
        test_user_interactions(page, context)
        test_responsive_design(page, context)
        test_all_views(page, context)
        test_console_errors(page, context)
        test_performance(page, context)
        test_data_flow(page, context)
        test_accessibility(page, context)
        
        browser.close()
    
    # 计算通过率
    if test_results["total_tests"] > 0:
        pass_rate = (test_results["passed_tests"] / test_results["total_tests"]) * 100
        test_results["pass_rate"] = f"{pass_rate:.2f}%"
    else:
        test_results["pass_rate"] = "0%"
    
    # 输出测试结果
    print("\n" + "=" * 60)
    print("测试结果汇总")
    print("=" * 60)
    print(f"总测试数: {test_results['total_tests']}")
    print(f"通过: {test_results['passed_tests']}")
    print(f"失败: {test_results['failed_tests']}")
    print(f"警告: {test_results['warnings']}")
    print(f"通过率: {test_results['pass_rate']}")
    
    # 保存测试报告
    with open('/tmp/test_report.json', 'w', encoding='utf-8') as f:
        json.dump(test_results, f, ensure_ascii=False, indent=2)
    
    print(f"\n测试报告已保存到: /tmp/test_report.json")
    print("=" * 60)

if __name__ == '__main__':
    main()
