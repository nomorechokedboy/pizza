<script setup lang="ts">
import { SignUpForm } from './components'

definePageMeta({
	layout: 'enter',
	middleware: ['authn']
})
</script>

<template>
	<SignUpForm />
</template>
