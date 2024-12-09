import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { VitePWA } from 'vite-plugin-pwa'
import viteCompression from 'vite-plugin-compression'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), 
    viteCompression({
      threshold: 1024,
    }),VitePWA({
    registerType: 'autoUpdate',
    manifest: {
      name: 'Obcsapi',
      short_name: 'Note',
      description: 'Obcsapi',
      theme_color: '#ffffff',
      icons: [
        {
          src: 'pwa-192x192.png',
          sizes: '192x192',
          type: 'image/png'
        }
      ]
    }
  })],
  server: {
    host: '0.0.0.0', // 设置为0.0.0.0以监听所有网络接口
    port: 5173 // 你可以根据需要更改端口
  },
  base: "/web/",
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
})
