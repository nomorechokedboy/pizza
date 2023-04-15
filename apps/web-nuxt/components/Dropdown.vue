<script lang="ts" setup>
export interface DropdownProps {
	open: boolean
}

const { open } = defineProps<DropdownProps>()
</script>

<template>
	<div
		class="dropdown absolute mt-2 px-2 py-2 z-10 text-md bg-white divide-y divide-gray-100 rounded-lg shadow w-[15.625rem] dark:bg-gray-700"
		:data-biding="open"
	>
		<slot />
	</div>
</template>

<style lang="postcss" scoped>
.dropdown {
	@apply -right-2;
}

.dropdown[data-biding='true'] {
	@apply block overflow-visible;
}

.dropdown[data-biding='false'] {
	@apply hidden;
}
</style>
