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
  background-color: var(--bg-primary);
  border-radius: 12px;
  box-shadow: var(--shadow-md);
  margin-bottom: 24px;
  overflow: hidden;
  border: 1px solid var(--border-light);
  transition: all 0.3s ease;
}

.base-card:hover {
  box-shadow: var(--shadow-lg);
  border-color: var(--primary-light);
}

.card-header {
  padding: 20px 24px;
  border-bottom: 1px solid var(--border-light);
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: var(--bg-primary);
}

.card-header h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  letter-spacing: 0.5px;
}

.card-body {
  padding: 24px;
  background-color: var(--bg-primary);
}

.card-footer {
  padding: 20px 24px;
  border-top: 1px solid var(--border-light);
  text-align: right;
  background-color: var(--bg-primary);
}

@media (max-width: 768px) {
  .base-card {
    margin-bottom: 20px;
  }
  
  .card-header {
    padding: 16px 20px;
  }
  
  .card-body {
    padding: 20px;
  }
  
  .card-footer {
    padding: 16px 20px;
  }
  
  .card-header h3 {
    font-size: 15px;
  }
}
</style>