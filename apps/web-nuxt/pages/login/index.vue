<script setup lang="ts">
import { SignInForm } from './components'

definePageMeta({
	layout: 'enter',
	middleware: ['redirect']
})
</script>

<template>
	<SignInForm />
</template>
