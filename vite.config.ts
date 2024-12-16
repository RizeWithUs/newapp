import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@import "./src/styles/theme.scss";`
      }
    }
  },
  server: {
    watch: {
      usePolling: true
    },
    host: true,
    open: true
  }
})