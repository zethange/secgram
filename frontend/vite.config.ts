import { join } from 'node:path';
import { defineConfig } from 'vite';
import solid from 'vite-plugin-solid';

export default defineConfig({
  plugins: [solid()],
  server: {
    host: '0.0.0.0'
  },
  resolve: {
    alias: {
      "@": join(__dirname, "src")
    }
  },
})
