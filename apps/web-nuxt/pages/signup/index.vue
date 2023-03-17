<script setup lang="ts">
import { Button } from 'ui-vue'
import { SignUpForm } from './components'

definePageMeta({
	layout: 'enter'
})
</script>

<template>
	<div class="w-full flex-1 flex flex-col gap-4">
		<div class="flex flex-col gap-2">
			<Button size="md" class="!py-3" block
				>Sign up with Google</Button
			>
			<Button size="md" class="!py-3" block
				>Sign up with Github</Button
			>
			<Button size="md" class="!py-3" block
				>Sign up with Gitlab</Button
			>
			<Button size="md" class="!py-3" block
				>Sign up with Facebook</Button
			>
		</div>
		<div class="h-[44px] bg-black text-white">Divider</div>
		<SignUpForm />
	</div>
</template>
