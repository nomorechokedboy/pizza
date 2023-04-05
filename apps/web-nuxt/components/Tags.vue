<script lang="ts" setup>
export interface TagsProps {
	tags: string[]
}

const { tags } = defineProps<TagsProps>()
</script>

<template>
	<div class="mb-2 flex gap-px">
		<span
			v-for="tag in tags"
			class="py-1 px-2 text-sm text-neutral-700"
			>#{{ tag }}</span
		>
	</div>
</template>
