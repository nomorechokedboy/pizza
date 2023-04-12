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
			<slot name="header">
				<Header />
			</slot>
		</header>
		<div
			class="fixed flex flex-col gap-5 top-20 right-0 z-40 max-w-xs lg:mr-5"
		>
			<Alert
				v-for="props in notificationList"
				:key="props.k"
				:="props"
			/>
		</div>
		<slot />
		<footer>
			<slot name="footer">
				<div class="min-h-[233px] bg-neutral-200">
					Footer
				</div>
			</slot>
		</footer>
	</div>
</template>

<style lang="postcss">
.skeleton {
	@apply pointer-events-none relative z-10 inline-block w-full animate-pulse select-none overflow-hidden rounded bg-slate-200 text-transparent;
}
</style>
