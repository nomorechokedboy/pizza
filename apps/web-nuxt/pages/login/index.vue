<script setup lang="ts">
import { Button } from 'ui-vue'
import { SignInForm } from './components/'

definePageMeta({
	layout: 'enter'
})
</script>

<template>
	<div class="w-full flex-1 flex flex-col gap-4">
		<div class="flex flex-col gap-2">
			<Button size="md" class="!py-3" block
				>Continue with Google</Button
			>
			<Button size="md" class="!py-3" block
				>Continue with Github</Button
			>
			<Button size="md" class="!py-3" block
				>Continue with Gitlab</Button
			>
			<Button size="md" class="!py-3" block
				>Continue with Facebook</Button
			>
		</div>
		<div class="h-[44px] bg-black text-white">Divider</div>
		<SignInForm />
		<a class="text-center text-sm" href="">I forgot my password</a>
	</div>
</template>
