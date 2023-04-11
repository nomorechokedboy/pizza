<script setup lang="ts"></script>

<template>
	<div class="flex flex-col min-h-screen relative">
		<Header />
		<slot />
		<footer class="min-h-[233px]"></footer>
	</div>
</template>
