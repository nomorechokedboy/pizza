<script setup lang="ts">
import { Button } from 'ui-vue'
import TextInput from '@@/components/TextInput.vue'

const { $blogApi } = useNuxtApp()
const email = ref('')
const emailInput = ref<InstanceType<typeof TextInput> | null>(null)
const handleSubmit = async () => {
	try {
		await $blogApi.auth.authForgotPasswordPost({
			email: email.value
		})
		console.log('Success')
	} catch (e) {
		console.error(e)
	}
}

onMounted(() => {
	if (emailInput.value) {
		emailInput.value.focus()
	}
})
</script>

<template>
	<div class="pt-12 pb-32 bg-[rgb(245,245,245)]">
		<div
			class="flex flex-col gap-4 p-8 pt-3 border bg-white lg:max-w-[490px] m-auto"
		>
			<h1 class="font-bold text-[27px]">
				Forgot your password?
			</h1>
			<form class="flex flex-col gap-3">
				<TextInput
					id="email"
					name="email"
					type="email"
					v-model="email"
					ref="emailInput"
				>
					<template #label>Email</template>
				</TextInput>
				<Button
					color="indigo"
					size="md"
					class="!font-medium"
					@click.prevent="handleSubmit"
					block
					>Send me reset password
					instructions</Button
				>
			</form>
		</div>
	</div>
</template>
