<script lang="ts" setup>
import { reactive } from 'vue'
import { Size } from '../common'
import Button, { ButtonProps } from './Button.vue'

const state = reactive<ButtonProps & { content: string; disabled: boolean }>({
	disabled: false,
	block: false,
	loading: false,
	content: 'Hello world',
	size: 'sm',
	radius: 'sm',
	loaderPosition: 'left'
})
</script>

<template>
	<Story>
		<Variant title="Usage" auto-props-disabled>
			<template #controls>
				<HstText
					v-model="state.content"
					title="Content"
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
				<HstCheckbox
					v-model="state.block"
					title="Block"
				/>
			</template>
			<div class="max-w-xs w-full">
				<Button
					:disabled="state.disabled"
					:size="state.size"
					:radius="state.radius"
					:block="state.block"
					:loading="state.loading"
				>
					{{ state.content }}
				</Button>
			</div>
		</Variant>
		<Variant title="Loading state" auto-props-disabled>
			<template #controls>
				<HstCheckbox
					v-model="state.loading"
					title="Loading"
				/>
				<HstSelect
					v-model="state.loaderPosition"
					title="Loader position"
					:options="{
						left: 'left',
						right: 'right',
						center: 'center'
					}"
				/>
			</template>
			<Button
				:loading="state.loading"
				:size="state.size"
				:radius="state.radius"
				:loader-position="state.loaderPosition"
				>Loading...</Button
			>
		</Variant>
	</Story>
</template>
