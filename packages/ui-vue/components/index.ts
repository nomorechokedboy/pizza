// import { defineNuxtModule } from '@nuxt/kit'
// import { join } from 'path'

// export default defineNuxtModule({
// 	setup(_, nuxt) {
// 		// here we need to setup our components
// 		nuxt.hook('components:dirs', (dirs) => {
// 			dirs.push({
// 				path: join(__dirname, 'components'),
// 				prefix: 'pizza'
// 			})
// 		}),
// 			nuxt.options.modules.push('@nuxtjs/tailwindcss')
// 	}
// })

export { default as ActionIcon } from './ActionIcon/ActionIcon.vue'
export { default as Button } from './Button/Button.vue'
export { default as CloseButton } from './CloseButton/CloseButton.vue'
export { default as CopyButton } from './CopyButton/CopyButton.vue'
export { default as FileButton } from './FileButton/FileButton.vue'
export { default as Loader } from './Loader/Loader.vue'
export { default as LoadingOverlay } from './LoadingOverlay/LoadingOverlay.vue'
export { default as Overlay } from './Overlay/Overlay.vue'
