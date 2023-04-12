<script setup lang="ts">
import { SignUpForm } from './components'

definePageMeta({
	layout: 'enter',
	middleware: ['redirect']
})
</script>

<template>
	<SignUpForm />
</template>
