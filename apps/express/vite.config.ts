import { defineConfig } from 'vite'
import { VitePluginNode } from 'vite-plugin-node'
import TypeCheck from 'vite-plugin-checker'
import { resolve } from 'path'

export default defineConfig({
  server: {
    port: parseInt(process.env.PIZZA_SERVER_PORT),
    host: true
  },
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src')
    }
  },
  plugins: [
    ...VitePluginNode({
      adapter: 'express',
      appPath: './src/index.ts',
      exportName: 'pizzaApi'
    }),
    TypeCheck({
      enableBuild: false,
      overlay: true,
      terminal: true
    })
  ],
  optimizeDeps: {
    exclude: ['mock-aws-s3', 'aws-sdk', 'nock']
  }
})
