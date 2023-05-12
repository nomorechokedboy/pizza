<script lang="ts" setup>
export interface DropdownProps {
	open: boolean
}

const { open } = defineProps<DropdownProps>()
</script>

<template>
	<div
		class="absolute right-0 w-max p-4 z-10 text-md bg-white divide-y divide-gray-100 rounded-lg shadow"
		v-if="open"
	>
		<slot />
	</div>
</template>
