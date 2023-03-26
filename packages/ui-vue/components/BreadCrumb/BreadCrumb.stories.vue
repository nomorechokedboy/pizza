<script lang="ts" setup>
import { openColor, Size } from '$common'
import BreadCrumb, { BreadCrumbProps } from './BreadCrumb.vue'

const initStateBasicUsage: () => BreadCrumbProps = () => ({
	size: 'sm',
	color: 'blue',
	variant: 'filled',
	items: [
		{ title: 'Home', href: 'https://facebook.com' },
		{ title: 'Product', href: 'https://youtube.com' },
		{ title: 'Sale', href: 'https://instagram.com' }
	]
})

const items = [
	{ title: 'Home', href: 'https://facebook.com' },
	{ title: 'Product', href: 'https://youtube.com' },
	{ title: 'Sale', href: 'https://instagram.com' }
].map((item) => ({ title: item.title, href: item.href }))
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
					<p>{{ state }}</p>
				</template>
				<HstSelect
					v-model="state.size"
					title="size"
					:options="Size"
				/>
				<HstSelect
					v-model="state.color"
					title="Color"
					:options="openColor"
				/>
				<HstText
					v-model="state.separator"
					title="Separator"
				/>
			</template>
			<template #default="{ state }">
				<div class="max-w-xs w-full">
					<BreadCrumb
						:items="items"
						:separator="state.separator"
						:color="state.color"
						:size="state.size"
					/>
				</div>
			</template>
		</Variant>
	</Story>
</template>
