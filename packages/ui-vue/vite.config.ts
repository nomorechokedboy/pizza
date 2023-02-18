import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig({
	build: {
		lib: {
			entry: path.resolve(__dirname, 'index.ts'),
			name: 'PizzaUI',
			fileName: (format) => `pizza-ui.${format}.js`
		},
		rollupOptions: {
			external: ['vue'],
			output: {
				globals: {
					vue: 'Vue'
				}
			}
		}
	},
	plugins: [vue()]
})
