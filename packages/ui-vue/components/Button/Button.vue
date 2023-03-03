<script lang="ts" setup>
import type { UIColor, UISize, UIVariant } from '../common'
import Loader, { LoaderProps } from '../Loader/Loader.vue'
import LoadingOverlay from '../LoadingOverlay/LoadingOverlay.vue'
import Overlay from '../Overlay/Overlay.vue'

export interface ButtonProps {
	block?: boolean
	color?: UIColor
	compact?: boolean
	size?: UISize
	radius?: UISize
	loading?: boolean
	loaderPosition?: 'left' | 'right' | 'center'
	loaderProps?: LoaderProps
	variant?: UIVariant
	uppercase?: boolean
}

const {
	color,
	radius,
	size,
	loaderPosition,
	block,
	loading,
	loaderProps,
	variant,
	compact,
	uppercase
} = defineProps<ButtonProps>()
</script>

<template>
	<button
		:data-block="block"
		:data-size="size ?? 'sm'"
		:data-radius="radius ?? 'sm'"
		:data-loading="loading"
		class="button"
		:class="[
			color ?? 'blue',
			variant ?? 'filled',
			{ uppercase: uppercase },
			{ compact: compact }
		]"
	>
		<Overlay
			v-if="loading && loaderPosition !== 'center'"
			class="overide-overlay"
			:radius="radius"
		/>
		<Loader
			v-if="
				loading &&
				(loaderPosition === 'left' || !loaderPosition)
			"
			v-bind="loaderProps"
		/>
		<slot v-else name="leftIcon" />
		<LoadingOverlay
			v-if="loading && loaderPosition === 'center'"
			:visible="loading && loaderPosition === 'center'"
			:loader-props="loaderProps"
			class="overide-overlay"
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
	@apply px-3.5 py-[7px] text-xs;
}

.button[data-size='sm'] {
	@apply px-4 py-2 text-sm;
}

.button[data-size='md'] {
	@apply px-5 py-[9px] text-base;
}

.button[data-size='lg'] {
	@apply px-6.5 py-[11px] text-lg;
}

.button[data-size='xl'] {
	@apply px-8 py-4 text-xl;
}

.button {
	@apply active:translate-y-px flex flex-row gap-4 items-center justify-center font-semibold relative disabled:bg-slate-200 disabled:text-slate-400 disabled:translate-y-0;
}

.button[data-block='true'] {
	@apply w-full overflow-hidden;
}

.button[data-loading='true'] {
	@apply pointer-events-none;
}

div.overide-overlay {
	@apply bg-black/20;
}

button.button.compact {
	@apply py-0;
}
</style>
