<script lang="ts" setup>
import type { UISize } from '../common'
import Loader, { LoaderProps } from '../Loader/Loader.vue'
import LoadingOverlay from '../LoadingOverlay/LoadingOverlay.vue'
import Overlay from '../Overlay/Overlay.vue'

export interface ButtonProps {
	size?: UISize
	radius?: UISize
	block?: boolean
	loading?: boolean
	loaderPosition?: 'left' | 'right' | 'center'
	loaderProps?: LoaderProps
}

const {
	radius,
	size,
	loaderPosition,
	block,
	loading,
	loaderProps: userLoaderProps
} = defineProps<ButtonProps>()
const loaderProps = userLoaderProps ?? { size: 'sm' }
</script>

<template>
	<button
		:data-block="block"
		:data-size="size ?? 'sm'"
		:data-radius="radius ?? 'sm'"
		:data-loading="loading"
		class="button"
	>
		<Overlay v-if="loading" class="overide-overlay" />
		<Loader
			v-if="loading && loaderPosition === 'left'"
			v-bind="loaderProps"
		/>
		<slot v-else name="leftIcon" />
		<LoadingOverlay
			v-if="loading && loaderPosition === 'center'"
			:visible="loading && loaderPosition === 'center'"
			:loader-props="loaderProps"
		/>
		<slot />
		<Loader
			v-if="loading && loaderPosition === 'right'"
			v-bind="loaderProps"
		/>
		<slot v-else name="rightIcon" />
	</button>
</template>

<style lang="css" scoped>
.button[data-size='xs'] {
	@apply px-3.5 py-1;
}

.button[data-size='sm'] {
	@apply px-4 py-1.5;
}

.button[data-size='md'] {
	@apply px-5 py-2;
}

.button[data-size='lg'] {
	@apply px-6.5 py-3.5;
}

.button[data-size='xl'] {
	@apply px-8 py-5;
}

.button {
	@apply bg-sky-300 active:translate-y-px flex flex-row gap-4 items-center justify-center font-semibold overflow-hidden relative disabled:bg-slate-200 disabled:text-slate-400 disabled:translate-y-0;
}

.button[data-block='true'] {
	@apply w-full;
}

.button[data-loading='true'] {
	@apply pointer-events-none;
}

.overide-overlay {
	@apply bg-black/20;
}
</style>
