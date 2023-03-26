<script lang="ts" setup>
import type { UIColor, UISize } from '$common'

export interface Item {
	title: string
	href?: string
	icon?: string
}

export interface BreadCrumbProps {
	separator?: string
	color?: UIColor
	size?: UISize
	items: Array<Item>
}

const {
	color,
	size,
	separator,
	items
	// compact
} = defineProps<BreadCrumbProps>()
</script>

<template>
	<div class="breadcrumb" :data-size="size ?? 'sm'">
		<template v-for="(item, index) in items">
			<span>
				{{ item?.icon }}
			</span>
			<a :data-color="color ?? 'blue'" :href="item.href">
				{{ item.title }}
			</a>
			<span
				class="separator"
				v-if="index != items.length - 1"
			>
				{{ separator ?? '/' }}
			</span>
		</template>
	</div>
</template>

<style lang="css" scoped>
.breadcrumb[data-size='xs'] {
	@apply px-[2px] py-[7px] text-xs;
}

.breadcrumb[data-size='sm'] {
	@apply px-[2.5px] py-2 text-sm;
}

.breadcrumb[data-size='md'] {
	@apply px-[3px] py-[9px] text-base;
}

.breadcrumb[data-size='lg'] {
	@apply px-[3.5px] py-[11px] text-lg;
}

.breadcrumb[data-size='xl'] {
	@apply px-[4px] py-4 text-xl;
}

.breadcrumb [data-color='default'] {
	@apply hover:underline text-dark-9;
}

.breadcrumb [data-color='dark'] {
	@apply hover:underline text-dark-8;
}

.breadcrumb [data-color='gray'] {
	@apply hover:underline text-gray-8;
}

.breadcrumb [data-color='red'] {
	@apply hover:underline text-red-8;
}

.breadcrumb [data-color='pink'] {
	@apply hover:underline text-pink-8;
}

.breadcrumb [data-color='grape'] {
	@apply hover:underline text-grape-8;
}

.breadcrumb [data-color='violet'] {
	@apply hover:underline text-violet-8;
}

.breadcrumb [data-color='indigo'] {
	@apply hover:underline text-indigo-8;
}

.breadcrumb [data-color='blue'] {
	@apply hover:underline text-blue-8;
}

.breadcrumb [data-color='cyan'] {
	@apply hover:underline text-cyan-8;
}

.breadcrumb [data-color='teal'] {
	@apply hover:underline text-teal-8;
}

.breadcrumb [data-color='green'] {
	@apply hover:underline text-green-8;
}

.breadcrumb [data-color='lime'] {
	@apply hover:underline text-lime-8;
}

.breadcrumb [data-color='yellow'] {
	@apply hover:underline text-yellow-5;
}

.breadcrumb [data-color='orange'] {
	@apply hover:underline text-orange-8;
}

.separator {
	@apply ml-2 mr-2 text-gray-700 leading-4;
}
</style>
