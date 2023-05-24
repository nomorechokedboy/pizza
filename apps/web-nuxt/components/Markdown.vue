<script lang="ts" setup>
import raw from 'rehype-raw'
import gfm from 'remark-gfm'

export interface MarkdownProps {
	source: string
}

defineProps<MarkdownProps>()
</script>

<template>
	<VueMarkdown
		:key="source"
		:source="source"
		:rehype-plugins="[raw]"
		:remark-plugins="[gfm]"
	/>
</template>

<style scoped>
.markdown > * {
	max-width: 800px;
	overflow-x: auto;
}
.markdown > h1,
.markdown > h2,
.markdown > h3,
.markdown > h4,
.markdown > h5,
.markdown > h6 {
	font-weight: bold;
}

.markdown > h1 {
	@apply text-4xl;
}
.markdown > h2 {
	@apply text-3xl;
}
.markdown > h3 {
	@apply text-2xl;
}
.markdown > h4 {
	@apply text-xl;
}
.markdown > h5 {
	@apply text-lg;
}
.markdown > h6 {
	@apply text-base;
}

.markdown > code {
	@apply bg-black text-white;
}
</style>
