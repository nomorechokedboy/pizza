<script lang="ts" setup>
import { reactive } from 'vue'
import Button, { ButtonProps } from './Button.vue'
import { Size } from '../common'

const state = reactive<ButtonProps & { content: string; disabled: boolean }>({
	disabled: false,
	fullWidth: false,
	content: 'Hello world',
	size: 'sm',
	radius: 'sm'
})
</script>

<template>
	<Story>
		<template #controls>
			<HstText v-model="state.content" title="Content" />
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
				v-model="state.fullWidth"
				title="Full witdth"
			/>
		</template>
		<Variant auto-props-disabled>
			<div class="max-w-xs w-full">
				<Button
					:disabled="state.disabled"
					:size="state.size"
					:radius="state.radius"
					:full-width="state.fullWidth"
				>
					{{ state.content }}
				</Button>
			</div>
		</Variant>
	</Story>
</template>
