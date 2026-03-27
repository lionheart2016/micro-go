<template>
  <div class="app-container">
    <header class="app-header">
      <h1>港美股经纪业务数据看板</h1>
    </header>
    <div class="app-main">
      <aside class="app-sidebar">
        <el-menu
          :default-active="activeMenu"
          class="el-menu-vertical-demo"
          router
        >
          <el-menu-item index="/">
            <span>关键指标概览</span>
          </el-menu-item>
          <el-menu-item index="/customer">
            <span>客户分析</span>
          </el-menu-item>
          <el-menu-item index="/trade">
            <span>交易分析</span>
          </el-menu-item>
          <el-menu-item index="/pi">
            <span>PI用户分析</span>
          </el-menu-item>
          <el-menu-item index="/account">
            <span>开户主题</span>
          </el-menu-item>
          <el-menu-item index="/ipo">
            <span>IPO主题</span>
          </el-menu-item>
          <el-menu-item index="/finance">
            <span>融资主题</span>
          </el-menu-item>
        </el-menu>
      </aside>
      <main class="app-content">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const activeMenu = computed(() => route.path)
</script>

<style scoped>
.app-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: var(--bg-secondary);
}

.app-header {
  background-color: var(--primary);
  color: var(--text-light);
  padding: 0 24px;
  line-height: 64px;
  box-shadow: var(--shadow-lg);
  position: relative;
  z-index: 100;
}

.app-header h1 {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  letter-spacing: 0.5px;
}

.app-main {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.app-sidebar {
  width: 240px;
  background-color: var(--bg-primary);
  border-right: 1px solid var(--border-light);
  box-shadow: var(--shadow-sm);
  transition: all 0.3s ease;
}

.app-sidebar:hover {
  box-shadow: var(--shadow-md);
}

.app-content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
  background-color: var(--bg-secondary);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

/* Element Plus 菜单样式覆盖 */
:deep(.el-menu-vertical-demo) {
  background-color: var(--bg-primary) !important;
  border-right: none !important;
}

:deep(.el-menu-item) {
  height: 56px !important;
  line-height: 56px !important;
  font-size: 14px !important;
  color: var(--text-secondary) !important;
  margin: 4px 12px !important;
  border-radius: 8px !important;
  transition: all 0.3s ease !important;
}

:deep(.el-menu-item:hover) {
  background-color: var(--bg-tertiary) !important;
  color: var(--primary) !important;
}

:deep(.el-menu-item.is-active) {
  background-color: rgba(37, 99, 235, 0.1) !important;
  color: var(--primary) !important;
  font-weight: 600 !important;
}

:deep(.el-menu-item.is-active::before) {
  content: '';
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 3px;
  height: 20px;
  background-color: var(--primary);
  border-radius: 0 2px 2px 0;
}

@media (max-width: 768px) {
  .app-sidebar {
    width: 200px;
  }
  
  .app-content {
    padding: 16px;
  }
  
  .app-header {
    padding: 0 16px;
  }
  
  .app-header h1 {
    font-size: 18px;
  }
}
</style>