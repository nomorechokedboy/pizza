<script lang="ts" setup>
import { UISize } from '../common'

export interface OverlayProps {
	center?: boolean
	fixed?: boolean
	radius?: UISize
}

const { center, fixed, radius } = defineProps<OverlayProps>()
</script>

<template>
	<div
		tabindex="-1"
		aria-hidden="true"
		:data-fixed="fixed"
		:data-center="center"
		:data-radius="radius ?? 'sm'"
		class="overlay"
	>
		<slot />
	</div>
</template>

<style lang="css" scoped>
.overlay {
	@apply overflow-y-auto overflow-x-hidden absolute top-0 right-0 left-0 bottom-0 z-40 w-full md:inset-0 bg-black/80;
}
.overlay[data-fixed='true'] {
	@apply fixed;
}
.overlay[data-center='true'] {
	@apply flex flex-col items-center justify-center;
}
</style>
