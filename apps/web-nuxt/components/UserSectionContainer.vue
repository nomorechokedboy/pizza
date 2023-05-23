<script lang="ts" setup>
export interface UserSectionContainerProps {
	title: string
}

defineProps<UserSectionContainerProps>()
</script>

<template>
	<section class="flex flex-col bg-white rounded">
		<h3 class="flex-shrink-0 text-2xl font-bold p-4">
			{{ title }}
		</h3>
		<slot />
	</section>
</template>
