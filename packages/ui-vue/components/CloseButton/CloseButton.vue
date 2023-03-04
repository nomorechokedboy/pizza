<script lang="ts" setup>
import ActionIcon from '$/ActionIcon/ActionIcon.vue'
import { LoaderProps } from '$/Loader/Loader.vue'
import type { UIColor, UISize, UIVariant } from '$common'
import XMarkIcon from './XMarkIcon.vue'

export interface CloseButtonProps {
	color?: UIColor
	size?: UISize
	radius?: UISize
	loading?: boolean
	loaderProps?: LoaderProps
	variant?: UIVariant | 'transparent'
	iconSize?: UISize
}
const { color, variant, loading, loaderProps, radius, size, iconSize } =
	defineProps<CloseButtonProps>()
const sizeMap: Record<UISize, string> = {
	xs: '0.75rem',
	sm: '1rem',
	md: '1.25rem',
	lg: '1.75rem',
	xl: '2.125rem'
}
</script>

<template>
	<ActionIcon
		:color="color ?? 'dark'"
		:variant="variant ?? 'subtle'"
		:loader-props="loaderProps"
		:loading="loading"
		:radius="radius"
		:size="size"
	>
		<slot>
			<XMarkIcon
				:width="sizeMap[iconSize ?? 'sm']"
				:height="sizeMap[iconSize ?? 'sm']"
			/>
		</slot>
	</ActionIcon>
</template>
