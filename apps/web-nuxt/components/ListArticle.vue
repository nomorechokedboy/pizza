<script lang="ts" setup>
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'

dayjs.extend(relativeTime)

export interface ArticleData {
	user: {
		id: number
		name: string
		avatar: string
	}
	id: number
	publishedAt: string | undefined
	title: string
	slug: string
	commentCount: number
	image: string | undefined
}

export interface ListArticleProps {
	loading?: boolean
	data: ArticleData[]
}

defineProps<ListArticleProps>()
</script>

<template>
	<div class="flex flex-col gap-2">
		<Article
			v-for="{
				user,
				id,
				publishedAt,
				title,
				slug,
				commentCount,
				image
			} in data"
			:owner="{
				src: user.avatar,
				name: user.name,
				id: id ?? 0
			}"
			:slug="slug || ''"
			:tags="['tags']"
			:publishedAt="`${convertDate(publishedAt)} (${dayjs(
				publishedAt
			).fromNow()})`"
			:title="title || ''"
			:src="image"
			:comments="commentCount || 0"
			:like="0"
			:key="id"
		/>
		<template v-if="loading">
			<Article
				v-for="n in 3"
				:owner="{
					id: n,
					name: '',
					src: ''
				}"
				slug="test"
				publishedAt=""
				title=""
				:comments="0"
				:like="0"
				src=""
				:key="n"
				showImage
				loading
			/>
		</template>
	</div>
</template>
