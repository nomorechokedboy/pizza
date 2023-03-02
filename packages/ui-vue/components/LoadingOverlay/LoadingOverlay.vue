<script lang="ts" setup>
import Overlay, { OverlayProps } from '../Overlay/Overlay.vue'
import Loader, { LoaderProps } from '../Loader/Loader.vue'

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
	<Overlay v-if="visible" v-bind="overlayProps">
		<Loader v-bind="loaderProps" />
	</Overlay>
</template>
