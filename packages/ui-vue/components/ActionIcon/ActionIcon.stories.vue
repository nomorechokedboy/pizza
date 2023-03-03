<script lang="ts" setup>
import { openColor, Size, Variant as VariantOptions } from '../common'
import ActionIcon, { ActionIconProps } from './ActionIcon.vue'

const initStateBasicUsage: () => ActionIconProps & {
	disabled: boolean
} = () => ({
	color: 'blue',
	disabled: false,
	loading: false,
	size: 'sm',
	radius: 'sm',
	variant: 'filled'
})
</script>

<template>
	<Story>
		<Variant
			title="Usage"
			:init-state="initStateBasicUsage"
			auto-props-disabled
		>
			<template #controls="{ state }">
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
					v-model="state.radius"
					title="Radius"
					:options="Size"
				/>
				<HstCheckbox
					v-model="state.disabled"
					title="Disabled"
				/>
				<HstCheckbox
					v-model="state.loading"
					title="Loading"
				/>
			</template>
			<template #default="{ state }">
				<ActionIcon
					:color="state.color"
					:disabled="state.disabled"
					:size="state.size"
					:radius="state.radius"
					:block="state.block"
					:loading="state.loading"
					:variant="state.variant"
					type="submit"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="75%"
						height="75%"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"
						class="tabler-icon tabler-icon-adjustments"
					>
						<path
							d="M6 10m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0"
						></path>
						<path d="M6 4l0 4"></path>
						<path d="M6 12l0 8"></path>
						<path
							d="M12 16m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0"
						></path>
						<path d="M12 4l0 10"></path>
						<path d="M12 18l0 2"></path>
						<path
							d="M18 7m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0"
						></path>
						<path d="M18 4l0 1"></path>
						<path d="M18 9l0 11"></path>
					</svg>
				</ActionIcon>
			</template>
		</Variant>
	</Story>
</template>
