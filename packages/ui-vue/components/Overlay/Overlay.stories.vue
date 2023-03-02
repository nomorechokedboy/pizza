<script lang="ts" setup>
import { reactive } from 'vue'
import Overlay, { OverlayProps } from './Overlay.vue'
import Button from '../Button/Button.vue'
import { Size } from '../common'

const state = reactive<OverlayProps & { visible: boolean }>({
	fixed: false,
	visible: true,
	center: false,
	radius: 'sm'
})
const toggleOverlay = () => {
	state.visible = false
}
</script>

<template>
	<Story>
		<Variant title="Usage" auto-props-disabled>
			<template #controls>
				<HstCheckbox
					v-model="state.center"
					title="Center"
				/>
				<HstCheckbox
					v-model="state.fixed"
					title="Fixed"
				/>
				<HstSelect
					v-model="state.radius"
					title="Radius"
					:options="Size"
				/>
			</template>
			<div
				class="w-[600px] flex flex-col gap-5 rounded-lg overflow-hidden"
			>
				<div class="w-[600px] relative">
					<img
						src="https://images.unsplash.com/photo-1618359057154-e21ae64350b6?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=720&q=80"
						alt="Dog photo"
						@click="state.visible = true"
					/>
					<Overlay
						v-if="state.visible"
						:center="state.center"
						:fixed="state.fixed"
						:radius="state.radius"
					>
						<Button @click="toggleOverlay"
							>Toggle Overlay</Button
						>
					</Overlay>
				</div>
			</div>
		</Variant>
	</Story>
</template>
