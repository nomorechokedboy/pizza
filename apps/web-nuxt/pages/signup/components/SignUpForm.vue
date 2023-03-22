<script setup lang="ts">
import { Button } from 'ui-vue'

const formData = reactive({
	identifier: '',
	password: '',
	confirmPassword: '',
	fullName: '',
	userName: ''
})
const { $blogApi } = useNuxtApp()
const handleSignUp = async () => {
	if (formData.password !== formData.confirmPassword) {
		alert('Password does not match')
		return
	}

	try {
		await $blogApi.auth.authRegisterPost({
			email: formData.identifier,
			password: formData.password,
			fullname: formData.fullName,
			username: formData.userName
		})
		await navigateTo('/login')
	} catch (e) {
		console.log('Sign up error: ', e)
	}
}
</script>

<template>
	<form class="flex flex-col gap-3">
		<TextInput name="fullname" v-model="formData.fullName">
			<template #label>Full name</template>
		</TextInput>
		<div class="flex flex-col gap-3">
			<TextInput name="email" v-model="formData.identifier">
				<template #label>Email</template>
			</TextInput>
			<TextInput name="username" v-model="formData.userName">
				<template #label>User name</template>
			</TextInput>
		</div>
		<TextInput
			name="password"
			type="password"
			v-model="formData.password"
		>
			<template #label>Password</template>
		</TextInput>
		<TextInput
			name="confirmPassword"
			type="password"
			v-model="formData.confirmPassword"
		>
			<template #label>Confirm Password</template>
		</TextInput>
		<Button
			color="indigo"
			size="md"
			class="!py-3"
			@click.prevent="handleSignUp"
			block
			>Sign up</Button
		>
	</form>
</template>
