<script lang="ts" setup>
import { reactive } from 'vue'
import CopyButton, { CopyButtonProps } from './CopyButton.vue'
import DemoCopyButton from './DemoCopyButton.vue'

const state = reactive<CopyButtonProps & { disabled?: boolean }>({
	disabled: false,
	value: 'https://mantine.dev'
})
</script>

<template>
	<Story>
		<Variant title="Usage" auto-props-disabled>
			<template #controls>
				<HstText v-model="state.value" title="Value"
			/></template>
			<template #default>
				<div class="max-w-xs w-full">
					<CopyButton
						:value="state.value"
						v-slot="slotProps"
					>
						<DemoCopyButton :="slotProps" />
					</CopyButton>
				</div>
			</template>
		</Variant>
	</Story>
</template>
