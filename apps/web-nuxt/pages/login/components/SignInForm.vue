<script setup lang="ts">
import useVuelidate from '@vuelidate/core'
import { email, helpers, required } from '@vuelidate/validators'
import { Button } from 'ui-vue'
import { notify } from '~~/composables/useNotification'

const formData = reactive({
	identifier: '',
	password: ''
})
const { $blogApi } = useNuxtApp()
const token = useAuthToken()
function formDataRules() {
	return {
		identifier: {
			required: helpers.withMessage(
				'The email field is required',
				required
			),
			email: helpers.withMessage(
				'Invalid email format',
				email
			)
		},
		password: {
			required: helpers.withMessage(
				'The password field is required',
				required
			)
		}
	}
}
const rules = computed(formDataRules)
const v$ = useVuelidate(rules, formData)
async function handleLogin() {
	const isValidForm = await v$.value.$validate()
	if (!isValidForm) {
		return
	}

	try {
		const { data: tokenData } = await $blogApi.auth.authLoginPost({
			password: formData.password,
			identifier: formData.identifier
		})

		token.value.refreshToken = tokenData.refresh_token
		token.value.accessToken = tokenData.token
		notify({
			content: 'Login success',
			type: 'success',
			k: crypto.randomUUID()
		})
		await navigateTo('/')
	} catch (e) {
		notifyError(e)
	}
}
</script>

<template>
	<form class="flex flex-col">
		<TextInput
			:error="v$.identifier.$error"
			:errors="v$.identifier.$errors"
			id="email"
			name="email"
			type="email"
			v-model="formData.identifier"
			@blur="v$.identifier.$touch"
		>
			<template #label>Email</template>
		</TextInput>
		<TextInput
			:error="v$.password.$error"
			:errors="v$.password.$errors"
			id="password"
			name="password"
			type="password"
			v-model="formData.password"
			@blur="v$.password.$touch"
		>
			<template #label>Password</template>
		</TextInput>
		<div class="flex gap-2 mb-4">
			<input type="checkbox" />
			<label for="">Remember me</label>
		</div>
		<Button
			color="indigo"
			size="md"
			class="!py-3"
			@click.prevent="handleLogin"
			block
			>Continue</Button
		>
	</form>
</template>
