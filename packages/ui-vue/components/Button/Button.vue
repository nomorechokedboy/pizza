<script lang="ts" setup>
import type { UISize } from '../common'

export interface ButtonProps {
	size?: UISize
	radius?: UISize
	fullWidth?: boolean
}

const {
	radius = 'sm',
	size = 'sm',
	fullWidth = false
} = defineProps<ButtonProps>()
</script>

<template>
	<button
		:class="[
			`size-${size} radius-${radius}`,
			{ 'w-full': fullWidth }
		]"
		class="button"
	>
		<slot></slot>
	</button>
</template>

<style lang="css" scoped>
.radius-xs {
	@apply rounded-sm;
}

.radius-sm {
	@apply rounded;
}

.radius-md {
	@apply rounded-lg;
}

.radius-lg {
	@apply rounded-2xl;
}

.radius-xl {
	@apply rounded-4xl;
}

.size-xs {
	@apply px-3.5 py-1;
}

.size-sm {
	@apply px-4 py-1.5;
}

.size-md {
	@apply px-5 py-2;
}
.size-lg {
	@apply px-6.5 py-3.5;
}
.size-xl {
	@apply px-8 py-5;
}

.button {
	@apply bg-sky-300 active:translate-y-px flex flex-row gap-4 items-center justify-center font-semibold overflow-hidden disabled:bg-slate-200 disabled:text-slate-400 disabled:translate-y-0;
}
</style>
