<script setup lang="ts">
export interface ArticleProps {
	src?: string
	owner: {
		name: string
		src: string
	}
	publishedAt: string
	title: string
	tags: string[]
	like: number
	comments: number
	slug: string
	showImage?: boolean
}

const {
	comments,
	publishedAt,
	title,
	tags,
	like,
	owner,
	src,
	showImage,
	slug
} = defineProps<ArticleProps>()
function haveComments() {
	return comments === 0 ? 'Add Comment' : `${comments} comments`
}
function calculateTo() {
	return `/${owner.name}/${slug}`
}
const comment = computed(haveComments)
const to = computed(calculateTo)
</script>

<template>
	<div class="bg-white shadow">
		<nuxt-img v-if="showImage" :src="src" />
		<div class="flex flex-col gap-3 p-4">
			<NoClue
				:description="publishedAt"
				:src="owner.src"
				:title="owner.name"
			/>
			<div class="flex flex-col">
				<h2
					class="mb-1 text-xl font-bold text-neutral-900"
				>
					<NuxtLink :to="to">
						{{ title }}</NuxtLink
					>
				</h2>
				<Tags :tags="tags" />
				<ArticleFooter
					:comment="comment"
					:like="like"
				/>
			</div>
		</div>
	</div>
</template>
