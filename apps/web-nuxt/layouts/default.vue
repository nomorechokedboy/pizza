<script setup lang="ts"></script>

<template>
	<div class="flex flex-col min-h-screen relative">
		<header
			class="flex flex-row items-center justify-between p-7 bg-white shadow-md sticky top-0 z-20 lg:px-8"
		></header>
		<slot />
		<footer class="min-h-[233px]"></footer>
	</div>
</template>
