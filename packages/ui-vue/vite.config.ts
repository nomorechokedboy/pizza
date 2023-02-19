import { defineConfig } from 'vitest/config'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig({
	build: {
		lib: {
			entry: path.resolve(__dirname, '/components/index.ts'),
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
	plugins: [vue()],
	test: {
		globals: true,
		environment: 'jsdom',
		includeSource: ['components/**/*.{ts,vue}'],
		setupFiles: ['./setupTest.ts'],
		coverage: {
			exclude: ['./setupTest.ts']
		},
		passWithNoTests: true,
		deps: {
			inline: ['vitest-canvas-mock']
		}
	}
})
