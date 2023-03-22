<script setup lang="ts">
import { SignInForm } from './components'

definePageMeta({
	layout: 'enter',
	middleware: ['authn']
})
</script>

<template>
	<SignInForm />
</template>
