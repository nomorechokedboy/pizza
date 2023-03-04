<script lang="ts">
export default {
	inheritAttrs: false
}
</script>

<script lang="ts" setup>
import Button from '$/Button/Button.vue'
import type { LoaderProps } from '$/Loader/Loader.vue'
import type { UIColor, UISize, UIVariant } from '$common'

export interface ActionIconProps {
	color?: UIColor
	size?: UISize
	radius?: UISize
	loading?: boolean
	loaderProps?: LoaderProps
	variant?: UIVariant | 'transparent'
}

const { loading, loaderProps, radius, size, color, variant } =
	defineProps<ActionIconProps>()
</script>

<template>
	<div :data-size="size ?? 'sm'" class="action-icon-wrapper">
		<Button
			:color="color ?? 'blue'"
			:loading="loading"
			:loader-props="loaderProps"
			:radius="radius"
			:size="size"
			:variant="variant as UIVariant"
			loader-position="center"
			class="overide-button"
			v-bind="$attrs"
			block
		>
			<slot v-if="!loading" />
		</Button>
	</div>
</template>

<style lang="css" scoped>
.action-icon-wrapper[data-size='xs'] {
	@apply w-[18px] h-[18px];
}
.action-icon-wrapper[data-size='sm'] {
	@apply w-[22px] h-[22px];
}
.action-icon-wrapper[data-size='md'] {
	@apply w-[28px] h-[28px];
}
.action-icon-wrapper[data-size='lg'] {
	@apply w-[34px] h-[34px];
}
.action-icon-wrapper[data-size='xl'] {
	@apply w-[44px] h-[44px];
}
button.overide-button[data-size='xs'] {
	@apply p-0;
}
button.overide-button[data-size='sm'] {
	@apply p-0;
}
button.overide-button[data-size='md'] {
	@apply p-0;
}
button.overide-button[data-size='lg'] {
	@apply p-0;
}
button.overide-button[data-size='xl'] {
	@apply p-0;
}
button.overide-button {
	@apply h-full font-medium;
}
</style>
