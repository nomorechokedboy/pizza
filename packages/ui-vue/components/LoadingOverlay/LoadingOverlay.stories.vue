<script lang="ts" setup>
import LoadingOverlay, { LoadingOverlayProps } from './LoadingOverlay.vue'
import { reactive } from 'vue'

const state = reactive<LoadingOverlayProps>({
	visible: false
})
</script>

<template>
	<Story>
		<Variant title="Usage" auto-props-disabled>
			<template #controls>
				<HstCheckbox
					v-model="state.visible"
					title="Visible"
				/>
			</template>
			<div class="relative w-[600px]">
				<LoadingOverlay :visible="state.visible" />
				<img
					src="https://images.unsplash.com/photo-1618359057154-e21ae64350b6?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=720&q=80"
					alt="Dog photo"
				/>
			</div>
		</Variant>
	</Story>
</template>
