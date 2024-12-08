<template>
  <div class="greetings">
    <h3 class="welcome">{{ msg }}</h3>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  msg: string
}>()
</script>

<style scoped>
.welcome {
  font-size: 2rem; /* 增加字体大小 */
  color: #fff; /* 字体颜色 */
  background: linear-gradient(90deg, #ff7e5f, #feb47b); /* 渐变背景 */
  -webkit-background-clip: text; /* 文字背景裁剪 */
  background-clip: text; /* 文字背景裁剪 */
  text-fill-color: transparent; /* 文字填充颜色透明 */
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3); /* 添加阴影效果 */
  animation: fadeIn 1s ease-in-out; /* 添加动画效果 */
}

/* 渐变动画 */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-20px); /* 初始位置向上偏移 */
  }
  to {
    opacity: 1;
    transform: translateY(0); /* 最终位置 */
  }
}

.greetings {
  text-align: center;
  padding: 20px; /* 增加内边距 */
}
</style>