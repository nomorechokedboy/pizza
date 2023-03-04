<script lang="ts" setup>
import { reactive } from 'vue'
import { openColor, Size, Variant as VariantOptions } from '../common'
import CloseButton, { CloseButtonProps } from './CloseButton.vue'

const state = reactive<CloseButtonProps & { disabled: boolean }>({
	color: 'blue',
	disabled: false,
	loading: false,
	size: 'sm',
	radius: 'sm',
	variant: 'filled',
	iconSize: 'sm'
})
</script>

<template>
	<Story>
		<template #controls>
			<HstSelect
				v-model="state.color"
				title="Color"
				:options="openColor"
			/>
			<HstSelect
				v-model="state.variant"
				title="Variant"
				:options="{
					...VariantOptions,
					transparent: 'transparent'
				}"
			/>
			<HstSelect
				v-model="state.size"
				title="Size"
				:options="Size"
			/>
			<HstSelect
				v-model="state.iconSize"
				title="Icon size"
				:options="Size"
			/>
			<HstSelect
				v-model="state.radius"
				title="Radius"
				:options="Size"
			/>
			<HstCheckbox
				v-model="state.disabled"
				title="Disabled"
			/>
			<HstCheckbox v-model="state.loading" title="Loading" />
		</template>
		<Variant title="Usage" auto-props-disabled>
			<CloseButton
				:color="state.color"
				:loading="state.loading"
				:radius="state.radius"
				:size="state.size"
				:disabled="state.disabled"
				:variant="state.variant"
				:icon-size="state.iconSize"
			/>
		</Variant>
	</Story>
</template>
