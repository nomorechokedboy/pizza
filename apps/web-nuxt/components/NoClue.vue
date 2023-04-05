<script lang="ts" setup>
export interface NoClueProps {
	src: string
	title: string
	description: string
}

const { description, src, title } = defineProps<NoClueProps>()
</script>

<template>
	<div class="flex flex-row gap-2">
		<nuxt-img class="rounded-full" :src="src" width="32" />
		<div class="flex flex-col">
			<p class="text-sm font-medium text-neutral-700">
				{{ title }}
			</p>
			<p class="text-xs text-neutral-600">
				{{ description }}
			</p>
		</div>
	</div>
</template>
