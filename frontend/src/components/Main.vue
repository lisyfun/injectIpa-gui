<script setup>
import {OpenFileDlgAndGenerateFile} from "../../wailsjs/go/main/App.js";
import { ref } from 'vue';

const isHovering = ref(false);
const isProcessing = ref(false);
const showSuccess = ref(false);

// 上传并生成文件
const uploadFile = async () => {
  if (isProcessing.value) return;
  
  try {
    isProcessing.value = true;
    await OpenFileDlgAndGenerateFile();
    showSuccess.value = true;
    setTimeout(() => {
      showSuccess.value = false;
    }, 2000);
  } catch (err) {
    alert('注入失败：' + err);
  } finally {
    isProcessing.value = false;
  }
};
</script>

<template>
  <div class="upload-container">
    <div class="upload-card" :class="{ 'success': showSuccess }">
      <div 
        class="upload-area"
        :class="{ 
          'hover': isHovering, 
          'processing': isProcessing,
          'success': showSuccess 
        }"
        @click="uploadFile"
        @mouseenter="isHovering = true"
        @mouseleave="isHovering = false"
      >
        <div class="upload-content">
          <div class="upload-icon" :class="{ 'spin': isProcessing }">
            <svg v-if="!isProcessing && !showSuccess" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
              <polyline points="17 8 12 3 7 8"/>
              <line x1="12" y1="3" x2="12" y2="15"/>
            </svg>
            <svg v-else-if="isProcessing" class="processing-icon" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 12a9 9 0 1 1-6.219-8.56"/>
            </svg>
            <svg v-else class="success-icon" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/>
              <polyline points="22 4 12 14.01 9 11.01"/>
            </svg>
          </div>
          <div class="upload-text">
            <h3>{{ 
              isProcessing ? '处理中...' : 
              showSuccess ? '注入成功！' :
              '点击选择文件'
            }}</h3>
            <p>{{ 
              isProcessing ? '请稍候...' :
              showSuccess ? '文件已成功注入' :
              '选择 IPA 和 Dylib 文件'
            }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.upload-container {
  width: 100%;
  height: 100%;
  padding: 0;
  background: var(--background-color);
}

.upload-card {
  background: var(--card-background);
  height: 100%;
  width: 100%;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
  border-radius: 0;
  padding: 0;
}

.upload-card::before {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(
    135deg,
    color-mix(in srgb, var(--primary-color) 12%, transparent),
    transparent
  );
  opacity: 0;
  transition: opacity 0.3s ease;
}

.upload-card.success::before {
  opacity: 1;
}

.upload-area {
  height: 100%;
  padding: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(
    135deg,
    color-mix(in srgb, var(--primary-color) 3%, transparent),
    transparent 70%
  );
}

.upload-area:hover, 
.upload-area.hover {
  background: linear-gradient(
    135deg,
    color-mix(in srgb, var(--primary-color) 5%, transparent),
    transparent 70%
  );
}

.upload-area.processing {
  cursor: wait;
  background: linear-gradient(
    135deg,
    color-mix(in srgb, var(--primary-light) 5%, transparent),
    transparent 70%
  );
}

.upload-area.success {
  background: linear-gradient(
    135deg,
    color-mix(in srgb, #22c55e 5%, transparent),
    transparent 70%
  );
}

.upload-content {
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.75rem;
  z-index: 1;
  transform: translateY(0);
  transition: transform 0.3s ease;
}

.upload-area:hover .upload-content {
  transform: translateY(-2px);
}

.upload-icon {
  color: var(--primary-color);
  transition: all 0.3s ease;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: color-mix(in srgb, var(--primary-color) 8%, transparent);
  padding: 0.75rem;
  position: relative;
  overflow: hidden;
}

.upload-icon::after {
  content: '';
  position: absolute;
  inset: 0;
  background: radial-gradient(
    circle at center,
    color-mix(in srgb, var(--primary-color) 15%, transparent),
    transparent 70%
  );
  opacity: 0;
  transition: opacity 0.3s ease;
}

.upload-area:hover .upload-icon::after {
  opacity: 1;
}

.upload-area.success .upload-icon {
  color: #22c55e;
  background: color-mix(in srgb, #22c55e 15%, transparent);
}

.upload-text h3 {
  color: var(--text-primary);
  font-size: 0.9rem;
  font-weight: 500;
  margin-bottom: 0.25rem;
  letter-spacing: 0.01em;
}

.upload-text p {
  color: var(--text-secondary);
  font-size: 0.75rem;
  line-height: 1.4;
  opacity: 0.9;
  letter-spacing: 0.01em;
}

@keyframes spin {
  100% { transform: rotate(360deg); }
}

@keyframes success-bounce {
  0% { transform: scale(0.8); opacity: 0; }
  50% { transform: scale(1.1); opacity: 0.8; }
  100% { transform: scale(1); opacity: 1; }
}

.success-icon {
  animation: success-bounce 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

.processing-icon {
  animation: spin 1.2s cubic-bezier(0.4, 0, 0.2, 1) infinite;
}
</style>