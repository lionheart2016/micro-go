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
}

.app-header {
  background-color: #1a365d;
  color: white;
  padding: 0 20px;
  line-height: 60px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.app-header h1 {
  margin: 0;
  font-size: 20px;
}

.app-main {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.app-sidebar {
  width: 200px;
  background-color: #f5f7fa;
  border-right: 1px solid #e4e7ed;
}

.app-content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  background-color: #f0f2f5;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>