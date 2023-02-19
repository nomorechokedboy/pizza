import { defineConfig } from 'histoire'
import { HstVue } from '@histoire/plugin-vue'

export default defineConfig({
	plugins: [HstVue()],
	setupFile: 'histoireSetup.ts',
	storyMatch: ['**/*.stories.vue']
})
