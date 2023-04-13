<script setup lang="ts">
export interface ArticleProps {
	src?: string
	owner: {
		name: string
		src: string
		id: number
	}
	publishedAt: string
	title: string
	tags?: string[]
	like: number
	comments: number
	slug: string
	showImage?: boolean
	loading?: boolean
}

const {
	comments,
	publishedAt,
	title,
	// tags,
	like,
	owner,
	src,
	showImage,
	slug,
	loading
} = defineProps<ArticleProps>()
function haveComments() {
	return comments === 0 ? 'Add Comment' : `${comments} comments`
}
function calculateTo() {
	return `/${owner.name}/${slug}`
}

const comment = computed(haveComments)
const to = computed(calculateTo)
const linkLoadingClass = { 'pointer-events-none': loading }
</script>

<template>
	<NuxtLink :to="to" :class="linkLoadingClass">
		<div class="bg-white shadow rounded overflow-hidden">
			<div
				v-if="loading"
				:class="{ skeleton: loading }"
				class="max-w-3xl w-full h-72"
			/>
			<nuxt-img
				v-if="showImage && src && !loading"
				:src="src"
				width="684"
				height="275"
			/>
			<div class="flex flex-col gap-3 p-4">
				<NoClue
					:alt="`${owner.name} avatar`"
					:description="publishedAt"
					:src="owner.src"
					:title="owner.name"
					:loading="loading"
				/>
				<div class="flex flex-col pl-8">
					<h2
						class="mb-1 text-xl font-bold text-neutral-900"
						:class="{ skeleton: loading }"
					>
						<NuxtLink
							:class="
								linkLoadingClass
							"
							:to="to"
						>
							{{ title }}</NuxtLink
						>
					</h2>
					<!-- <Tags :tags="tags" /> -->
					<ArticleFooter
						:user="owner.name"
						:slug="slug"
						:comment="comment"
						:like="like"
						:loading="loading"
					/>
				</div>
			</div>
		</div>
	</NuxtLink>
</template>
