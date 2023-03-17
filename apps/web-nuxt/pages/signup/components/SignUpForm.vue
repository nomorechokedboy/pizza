<script setup lang="ts">
import { Button } from 'ui-vue'
import { authApi } from '~~/external_modules'

const formData = reactive({ identifier: '', password: '', confirmPassword: '' })
const handleSignUp = async () => {
	if (formData.password !== formData.confirmPassword) {
		alert('Password does not match')
		return
	}

	try {
		await authApi.authRegisterPost({
			identifier: formData.identifier,
			password: formData.password
		})
		await navigateTo('/login')
	} catch (e) {
		console.log('Sign up error: ', e)
	}
}
</script>

<template>
	<form class="flex flex-col gap-3">
		<div class="flex flex-col gap-2">
			<label for="">Full name</label>
			<input class="border" type="text" />
		</div>
		<div class="flex flex-col gap-3">
			<div class="flex flex-col gap-2">
				<label for="">Email</label>
				<input
					class="border"
					type="text"
					v-model="formData.identifier"
				/>
			</div>
			<div class="flex flex-col gap-2">
				<label for="">User name</label>
				<input class="border" type="text" />
			</div>
		</div>
		<div class="flex flex-col gap-2">
			<label for="">Password</label>
			<input
				class="border"
				type="password"
				v-model="formData.password"
			/>
		</div>
		<div class="flex flex-col gap-2">
			<label for="">Confirm Password</label>
			<input
				class="border"
				type="password"
				v-model="formData.confirmPassword"
			/>
		</div>
		<Button
			size="md"
			class="!py-3"
			@click.prevent="handleSignUp"
			block
			>Continue</Button
		>
	</form>
</template>
