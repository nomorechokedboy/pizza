<script lang="ts" setup>
import { Size } from '$common'
import Paper, { PaperProps } from './Paper.vue'

const initStateBasicUsage: () => PaperProps = () => ({
	shadow: 'xs',
	radius: 'sm',
	withBorder: true,
	padding: 'xs'
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
				<template v-if="state != undefined">
					{{ state }}
				</template>
				<HstSelect
					v-model="state.shadow"
					title="Shadow"
					:option="Size"
				/>
				<HstSelect
					v-model="state.radius"
					title="Radius"
					:option="Size"
				/>
				<HstSelect
					v-model="state.padding"
					title="Padding"
					:option="Size"
				/>
				<HstCheckbox
					v-model="state.withBorder"
					title="withBorder"
				/>
			</template>

			<template #default="{ state }">
				<Paper
					:radius="state.radius"
					:padding="state.padding"
					:shadow="state.shadow"
					:withBorder="state.withBorder"
				>
					<p>Hello, I'm a vue developer</p>
					<p>Now, I'm a senior in HCMUTE</p>
				</Paper>
			</template>
		</Variant>
	</Story>
</template>
