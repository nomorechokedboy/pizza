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

async function handleSubmit(published: boolean) {
	const isValidData = await v$.value.$validate()
	if (!isValidData) {
		return
	}

	try {
		const { data } = await $blogApi.post.postsPost({
			content: formData.content,
			title: formData.title,
			published
		})

		await navigateTo(`/${data.user?.id}/${data.slug}`)
	} catch (e) {
		notifyError(e)
	}
}

const { $blogApi } = useNuxtApp()
const rules = computed(getFormRules)
const formData = reactive({
	title: '',
	content: ''
})
const v$ = useVuelidate(rules, formData)
definePageMeta({ layout: 'new', middleware: ['authn'] })
</script>

<template>
	<main class="max-w-7xl p-5 w-full h-full m-auto md:py-10">
		<form class="flex flex-col gap-2 max-h-full h-full">
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
					>
						<span
							class="text-neutral-800 font-normal group-hover:text-indigo-500"
							>Create as draft</span
						>
					</Button>
					<Button
						@click.prevent="
							handleSubmit(true)
						"
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
			<div class="flex item-center w-full h-full">
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
