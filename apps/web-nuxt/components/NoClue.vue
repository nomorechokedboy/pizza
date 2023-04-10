<script lang="ts" setup>
export interface NoClueProps {
	src: string
	title: string
	description: string
	alt: string
}

const { description, src, title, alt } = defineProps<NoClueProps>()
</script>

<template>
	<div class="flex flex-row gap-2">
		<div>
			<nuxt-img
				:alt="alt"
				class="rounded-full"
				:src="src"
				width="32"
				height="32"
			/>
		</div>
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
