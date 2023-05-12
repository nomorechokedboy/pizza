<script lang="ts" setup>
export interface NoClueProps {
	src: string
	title: string
	description: string
	alt: string
	loading?: boolean
}

const { description, src, title, alt, loading } = defineProps<NoClueProps>()
const loadingClass = { skeleton: loading }
</script>

<template>
	<div class="flex flex-row gap-2">
		<div :class="{ 'h-8': loading }">
			<div class="skeleton" v-if="loading">
				<div class="w-8 h-8" />
			</div>
			<Avatar
				v-else
				:alt="alt"
				:src="src"
				width="32"
				height="32"
			/>
		</div>
		<div class="flex flex-col w-full" :class="loadingClass">
			<template v-if="!loading">
				<p class="text-sm font-medium text-neutral-700">
					{{ title }}
				</p>
				<p
					class="text-xs text-neutral-600"
					:class="loadingClass"
				>
					{{ description }}
				</p>
			</template>
			<p v-else>
				<br />
			</p>
		</div>
	</div>
</template>
