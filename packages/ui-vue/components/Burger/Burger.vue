<script lang="ts" setup>
import type { UIColor, UISize, UIVariant } from '$common'
import ActionIcon from '$/ActionIcon/ActionIcon.vue'
import XMarkIcon from '$/CloseButton/XMarkIcon.vue'

export interface BurgerProps {
	variant?: UIVariant
	opened: boolean
	color?: UIColor
	size?: UISize
	iconSize?: UISize
}

const { variant, color, size, opened } = defineProps<BurgerProps>()

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
		:size="size ?? 'sm'"
		:variant="variant ?? 'subtle'"
		@click="opened = !opened"
	>
		<Transition name="fade">
			<svg
				v-if="!opened"
				:width="sizeMap[iconSize ?? 'sm']"
				:height="sizeMap[iconSize ?? 'sm']"
				viewBox="0 0 24.00 24.00"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				stroke="currentColor"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
				class="tabler-icon tabler-icon-adjustments"
				transform="matrix(1, 0, 0, 1, 0, 0)"
			>
				<g id="SVGRepo_bgCarrier" stroke-width="0"></g>
				<g
					id="SVGRepo_tracerCarrier"
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke="#ccc"
					stroke-width="0.288"
				></g>
				<g id="SVGRepo_iconCarrier">
					<path
						d="M4 18L20 18"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
					></path>
					<path
						d="M4 12L20 12"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
					></path>
					<path
						d="M4 6L20 6"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
					></path>
				</g>
			</svg>

			<XMarkIcon
				v-else
				:width="sizeMap[iconSize ?? 'sm']"
				:height="sizeMap[iconSize ?? 'sm']"
			/>
		</Transition>
	</ActionIcon>
</template>

<style lang="css" scoped>
.fade-leave-active {
	transition: all 0s ease-in-out;
}

.fade-enter-active {
	transition: all 0.3s ease-in-out;
}
</style>
