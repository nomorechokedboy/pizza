<script setup lang="ts">
import { authApi } from '~~/external_modules'
import { Button } from 'ui-vue'

const formData = reactive({
	identifier: '',
	password: ''
})

const handleLogin = async () => {
	try {
		await authApi.authLoginPost({
			password: formData.password,
			identifier: formData.identifier
		})
		await navigateTo('/')
	} catch (e) {
		console.error('Login error', e)
	}
}
</script>

<template>
	<form class="flex flex-col gap-3">
		<TextInput
			id="email"
			name="email"
			type="email"
			v-model="formData.identifier"
		>
			<template #label>Email</template>
		</TextInput>
		<TextInput
			id="password"
			name="password"
			type="password"
			v-model="formData.password"
		>
			<template #label>Password</template>
		</TextInput>
		<div class="flex gap-2">
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
