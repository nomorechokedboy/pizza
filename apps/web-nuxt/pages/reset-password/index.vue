<script setup lang="ts">
import { Button } from 'ui-vue'
import TextInput from '~~/components/TextInput.vue'

definePageMeta({
	middleware: ['authn', 'reset-password-guard']
})

const { $blogApi } = useNuxtApp()
const route = useRoute()
const token = route.query.token?.toString()
const formData = reactive({ password: '', confirmPassword: '' })
const passwordInput = ref<InstanceType<typeof TextInput> | null>(null)
const handleSubmit = async () => {
	try {
		await $blogApi.auth.authResetPasswordPut({
			password: formData.password,
			token
		})
		notify({
			content: 'Reset password success!',
			k: crypto.randomUUID(),
			type: 'success'
		})
	} catch (e) {
		notifyError(e)
	}
}
</script>

<template>
	<div class="pt-12 pb-32 bg-[rgb(245,245,245)]">
		<div
			class="flex flex-col gap-4 p-8 pt-3 border bg-white lg:max-w-[490px] m-auto"
		>
			<h1 class="font-bold text-[27px]">
				Reset your password
			</h1>
			<form class="flex flex-col gap-3">
				<TextInput
					id="password"
					name="password"
					type="password"
					v-model="formData.password"
					ref="passwordInput"
				>
					<template #label>Password</template>
				</TextInput>
				<TextInput
					name="confirmPassword"
					type="password"
					v-model="formData.confirmPassword"
				>
					<template #label
						>Confirm Password</template
					>
				</TextInput>
				<Button
					color="indigo"
					size="md"
					class="!font-medium"
					@click.prevent="handleSubmit"
					block
				>
					Reset
				</Button>
			</form>
		</div>
	</div>
</template>
