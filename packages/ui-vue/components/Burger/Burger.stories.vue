<script lang="ts" setup>
import { openColor, Size, Variant as VariantOptions } from '$common'
import { reactive } from 'vue'
import Burger, { BurgerProps } from './Burger.vue'

const state = reactive<BurgerProps & { disabled: boolean }>({
	color: 'blue',
	disabled: false,
	size: 'sm',
	opened: false,
	variant: 'default',
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
			<HstCheckbox
				v-model="state.disabled"
				title="Disabled"
			/>
			<HstSelect
				v-model="state.iconSize"
				title="iconSize"
				:options="Size"
			/>
		</template>

		<Variant title="Usage" auto-props-disabled>
			<Burger
				:opened="state.opened"
				:color="state.color"
				:size="state.size"
				:disabled="state.disabled"
				:variant="state.variant"
				:icon-size="state.iconSize"
			/>
		</Variant>
	</Story>
</template>
