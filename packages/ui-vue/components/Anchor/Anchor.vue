<script lang="ts" setup>
import type { UIColor, UISize, UIVariant } from '$common'

export interface AnchorProps {
	variant?: UIVariant
	underline: boolean
	color?: UIColor
	size?: UISize
}

const { variant, underline, color, size } = defineProps<AnchorProps>()
</script>

<template>
	<div class="Anchor">
		<a
			class="Anchor"
			:data-size="size ?? 'sm'"
			:class="[
				color ?? 'blue',
				variant ?? 'default',
				{ underline: underline }
			]"
		>
			<slot />
		</a>
	</div>
</template>

<style lang="css" scoped></style>
