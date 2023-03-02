<script lang="ts" setup>
import type { UISize } from '../common'

export interface LoaderProps {
	size?: UISize
}

const { size } = defineProps<LoaderProps>()
</script>

<template>
	<div class="loader" :data-size="size ?? 'sm'" />
</template>

<style scoped lang="css">
.loader {
	@apply border-4 border-t-blue-400 border-slate-200 rounded-full animate-spin;
}
.loader[data-size='xs'] {
	@apply w-3 h-3;
}
.loader[data-size='sm'] {
	@apply w-4 h-4;
}
.loader[data-size='md'] {
	@apply w-6 h-6;
}
.loader[data-size='lg'] {
	@apply w-8 h-8;
}
.loader[data-size='xl'] {
	@apply w-10 h-10;
}
</style>
