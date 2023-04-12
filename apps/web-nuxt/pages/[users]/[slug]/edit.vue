<script lang="ts" setup>
import useVuelidate from '@vuelidate/core'
import { helpers, required } from '@vuelidate/validators'
import { Button } from 'ui-vue'

function getFormRules() {
	return {
		title: {
			required: helpers.withMessage(
				'Title is required',
				required
			)
		}
	}
}

async function handleLoading(published: boolean) {
	if (published) {
		loading.value = true
	} else {
		draftLoading.value = true
	}
}

async function handleSubmit(published: boolean) {
	const isValidData = await v$.value.$validate()
	if (!isValidData) {
		return
	}

	if (!postDetails.value?.id) {
		return
	}

	handleLoading(published)
	try {
		const { data } = await $blogApi.post.postsIdPut(
			postDetails.value?.id,
			{
				content: formData.content,
				title: formData.title,
				published
			}
		)

		await navigateTo(`/${data.user?.id}/${data.slug}`)
	} catch (e) {
		notifyError(e)
	} finally {
		loading.value = false
		draftLoading.value = false
	}
}

const route = useRoute()
let slug: string = ''
if (typeof route.params.slug === 'object') {
	slug = route.params.slug.pop() || ''
} else {
	slug = route.params.slug
}

const { data: postDetails } = usePostDetails(slug)
const { $blogApi } = useNuxtApp()
const loading = ref(false)
const draftLoading = ref(false)
const rules = computed(getFormRules)
const formData = reactive({
	title: postDetails.value?.title || '',
	content: postDetails.value?.content || ''
})
const v$ = useVuelidate(rules, formData)
definePageMeta({ layout: 'new', middleware: ['authn'] })
</script>

<template>
	<main
		class="flex-1 flex flex-col max-w-7xl p-5 w-full mx-auto md:py-10"
	>
		<form class="flex-1 flex flex-col gap-2 max-h-full">
			<header class="flex flex-col gap-3">
				<div
					class="flex items-center justify-end gap-3"
				>
					<Button
						color="indigo"
						class="group"
						variant="subtle"
						@click.prevent="
							handleSubmit(false)
						"
						:disabled="loading"
						:loading="draftLoading"
					>
						<span
							class="text-neutral-800 font-normal group-hover:text-indigo-500 group-disabled:text-gray-300"
							>Create as draft</span
						>
					</Button>
					<Button
						@click.prevent="
							handleSubmit(true)
						"
						:disabled="draftLoading"
						:loading="loading"
						>Publish</Button
					>
				</div>
				<TextInput
					:error="v$.$error"
					:errors="v$.$errors"
					v-model="formData.title"
					class="w-full rounded focus:outline-none text-3xl font-bold md:w-1/2"
					placeholder="Post title here..."
				/>
			</header>
			<div class="flex flex-1 item-center w-full h-full">
				<RichTextEditor v-model="formData.content" />
				<div
					class="hidden flex-1 bg-white max-w-1/2 border-l pt-10 overflow-auto preview rounded-r md:block"
				>
					<VueMarkdown
						class="markdown"
						:key="formData.content"
						:source="formData.content"
					/>
				</div>
			</div>
		</form>
	</main>
</template>
