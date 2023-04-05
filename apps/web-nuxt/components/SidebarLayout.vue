<script setup lang="ts">
export interface SidebarLayoutProps {
	reverse?: boolean
}

const { reverse } = defineProps<SidebarLayoutProps>()
</script>

<template>
	<div
		class="w-full m-auto h-full flex gap-4"
		:class="{ 'flex-row-reverse': reverse }"
	>
		<slot name="sidebar" />
		<slot />
	</div>
</template>
