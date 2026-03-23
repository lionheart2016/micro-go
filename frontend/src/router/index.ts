import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: () => import('../views/dashboard/index.vue'),
    meta: { title: '关键指标概览' }
  },
  {
    path: '/customer',
    name: 'Customer',
    component: () => import('../views/customer/index.vue'),
    meta: { title: '客户分析' }
  },
  {
    path: '/trade',
    name: 'Trade',
    component: () => import('../views/trade/index.vue'),
    meta: { title: '交易分析' }
  },
  {
    path: '/pi',
    name: 'PI',
    component: () => import('../views/pi/index.vue'),
    meta: { title: 'PI用户分析' }
  },
  {
    path: '/account',
    name: 'Account',
    component: () => import('../views/account/index.vue'),
    meta: { title: '开户主题' }
  },
  {
    path: '/ipo',
    name: 'IPO',
    component: () => import('../views/ipo/index.vue'),
    meta: { title: 'IPO主题' }
  },
  {
    path: '/finance',
    name: 'Finance',
    component: () => import('../views/finance/index.vue'),
    meta: { title: '融资主题' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router