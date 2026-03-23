<template>
  <div class="base-card">
    <div v-if="title" class="card-header">
      <h3>{{ title }}</h3>
      <slot name="header"></slot>
    </div>
    <div class="card-body">
      <slot></slot>
    </div>
    <div v-if="showFooter" class="card-footer">
      <slot name="footer"></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

const props = defineProps({
  title: {
    type: String,
    default: ''
  },
  showFooter: {
    type: Boolean,
    default: false
  }
})
</script>

<style scoped>
.base-card {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
  overflow: hidden;
}

.card-header {
  padding: 16px;
  border-bottom: 1px solid #f0f0f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 500;
  color: #303133;
}

.card-body {
  padding: 16px;
}

.card-footer {
  padding: 16px;
  border-top: 1px solid #f0f0f0;
  text-align: right;
}
</style>