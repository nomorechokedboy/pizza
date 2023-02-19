<script lang="ts" setup>
import type { UISize } from '../common'

export interface ButtonProps {
	size?: UISize
	radius?: UISize
	block?: boolean
	loading?: boolean
}

const {
	radius = 'sm',
	size = 'sm',
	block = false,
	loading = false
} = defineProps<ButtonProps>()
</script>

<template>
	<button
		:data-block="block"
		:data-size="size"
		:data-radius="radius"
		:data-loading="loading"
		class="button"
	>
		<slot name="leftIcon" />
		<slot />
		<slot name="rightIcon" />
	</button>
</template>

<style lang="css" scoped>
.button[data-radius='xs'] {
	@apply rounded-sm;
}

.button[data-radius='sm'] {
	@apply rounded;
}

.button[data-radius='md'] {
	@apply rounded-lg;
}

.button[data-radius='lg'] {
	@apply rounded-2xl;
}

.button[data-radius='xl'] {
	@apply rounded-4xl;
}

.button[data-size='xs'] {
	@apply px-3.5 py-1;
}

.button[data-size='sm'] {
	@apply px-4 py-1.5;
}

.button[data-size='md'] {
	@apply px-5 py-2;
}
.button[data-size='lg'] {
	@apply px-6.5 py-3.5;
}
.button[data-size='xl'] {
	@apply px-8 py-5;
}

.button {
	@apply bg-sky-300 active:translate-y-px flex flex-row gap-4 items-center justify-center font-semibold overflow-hidden disabled:bg-slate-200 disabled:text-slate-400 disabled:translate-y-0;
}

.button[data-block='true'] {
	@apply w-full;
}
</style>
