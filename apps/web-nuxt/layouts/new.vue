<script setup lang="ts">
useHead({
	bodyAttrs: {
		class: 'bg-neutral-100'
	}
})
</script>

<template>
	<div class="flex flex-col h-screen">
		<header>
			<Header />
		</header>
		<slot />
	</div>
</template>
