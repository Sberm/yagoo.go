import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  // port 8080 
  server: {
    port : 8080,
    proxy : {
      '/search' : {
        target: 'http://127.0.0.1:3030',
        changeOrigin: true,
        secure: false,
      },
      '/albums' : 'http://localhost:3030',
    }
  },
})

