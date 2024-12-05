<script setup>
import Main from "./components/Main.vue";
</script>

<template>
  <div class="app-container">
    <Main />
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

:root {
  --primary-color: #4f46e5;
  --primary-light: #818cf8;
  --primary-dark: #4338ca;
  --hover-color: #6366f1;
  --background-color: rgb(28, 28, 28);
  --card-background: rgb(32, 32, 32);
  --text-primary: #ffffff;
  --text-secondary: #999999;
  --border-color: #333333;
  --shadow-sm: 0 1px 2px 0 rgb(0 0 0 / 0.3);
  --shadow: 0 4px 6px -1px rgb(0 0 0 / 0.3);
}

body {
  background: var(--background-color);
  color: var(--text-primary);
  min-height: 100vh;
  transition: all 0.3s ease;
  overflow: hidden;
}

.app-container {
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
  background: var(--background-color);
}

@keyframes gradient {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}
</style>
