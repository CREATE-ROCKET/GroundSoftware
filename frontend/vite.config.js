import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [vue()],
  optimizeDeps: {
    include: ['three', 'chart.js/auto'],
  },
  build: {
    rollupOptions: {
      external: ['chart.js/auto'],
    },
  },
});
