<script lang="ts" setup>
import { reactive } from 'vue'
import { Size } from '../common'
import Loader, { LoaderProps } from './Loader.vue'

const state = reactive<LoaderProps>({
	size: 'sm'
})
</script>

<template>
	<Story>
		<template #controls>
			<HstSelect
				v-model="state.size"
				title="Size"
				:options="Size"
			/>
		</template>
		<Variant auto-props-disabled>
			<Loader :size="state.size" />
		</Variant>
	</Story>
</template>
