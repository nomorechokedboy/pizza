<script lang="ts" setup>
import type { UISize } from '$common'

export interface PaperProps {
	shadow?: UISize
	radius?: UISize
	withBorder?: boolean
	padding?: UISize
}

const { shadow, radius, withBorder, padding } = defineProps<PaperProps>()
</script>

<template>
	<div
		class="paper"
		:data-padding="padding ?? 'xs'"
		:data-shadow="shadow ?? 'xs'"
		:data-border="withBorder"
		:data-radius="radius ?? 'xs'"
	>
		<slot />
	</div>
</template>

<style lang="css" scoped>
.paper {
	@apply bg-slate-100 w-fit h-fit;
}
.paper[data-border='true'] {
	@apply border-solid border;
}

.paper[data-padding='xs'] {
	@apply px-[10px] py-[10px];
}

.paper[data-padding='sm'] {
	@apply px-[12px] py-[12px];
}

.paper[data-padding='md'] {
	@apply px-[16px] py-[16px];
}

.paper[data-padding='lg'] {
	@apply px-[20px] py-[20px];
}

.paper[data-padding='xl'] {
	@apply px-[24px] py-[24px];
}

.paper[data-shadow='xs'] {
	@apply drop-shadow;
}

.paper[data-shadow='sm'] {
	@apply drop-shadow-sm;
}

.paper[data-shadow='md'] {
	@apply drop-shadow-md;
}

.paper[data-shadow='lg'] {
	@apply drop-shadow-lg;
}

.paper[data-shadow='xl'] {
	@apply drop-shadow-xl;
}
</style>
