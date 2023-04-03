<script setup lang="ts">
const notificationList = useNotification()

useHead({
	bodyAttrs: {
		class: 'bg-neutral-100'
	}
})
</script>

<template>
	<div class="flex flex-col min-h-screen relative">
		<header class="sticky top-0 z-20">
			<Header />
		</header>
		<div
			class="fixed flex flex-col gap-5 top-20 right-0 max-w-xs lg:mr-5"
		>
			<Alert
				v-for="props in notificationList"
				:key="props.k"
				:="props"
			/>
		</div>
		<slot />
		<footer class="min-h-[233px] bg-neutral-200">Footer</footer>
	</div>
</template>
