<script lang="ts" setup>
import Loader, { LoaderProps } from '$/Loader/Loader.vue'
import Overlay, { OverlayProps } from '$/Overlay/Overlay.vue'

export interface LoadingOverlayProps {
	visible: boolean
	loaderProps?: LoaderProps
	overlayProps?: OverlayProps
}

const {
	visible,
	loaderProps: userLoaderProps,
	overlayProps: userOverlayProps
} = defineProps<LoadingOverlayProps>()
const loaderProps = userLoaderProps ?? { size: 'md' }
const overlayProps = userOverlayProps ?? { center: true }
</script>

<template>
	<Transition name="fade">
		<Overlay v-if="visible" v-bind="overlayProps">
			<Loader v-bind="loaderProps" />
		</Overlay>
	</Transition>
</template>

<style lang="css" scoped>
.fade-enter-active,
.fade-leave-active {
	transition: opacity 0.1s ease;
}

.fade-enter-from,
.fade-leave-to {
	opacity: 0;
}
</style>
