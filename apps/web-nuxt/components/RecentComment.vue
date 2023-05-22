<script lang="ts" setup>
export interface RecentCommentProps {
	postTitle: string
	commentContent: string
	createdAt: string
}

defineProps<RecentCommentProps>()
</script>

<template>
	<div class="w-full flex flex-col p-4 gap-2 border-b" style="--lines: 1">
		<h3 class="truncate font-bold text-neutral-800 text-base">
			{{ postTitle }}
		</h3>
		<div class="flex items-center gap-2 text-sm">
			<p class="truncate flex-1 text-neutral-700">
				{{ commentContent }}
			</p>
			<time class="truncate ml-auto text-neutral-500">{{
				createdAt
			}}</time>
		</div>
	</div>
</template>
