<script setup lang="ts">
import { Button } from 'ui-vue'
</script>

<template>
	<header
		class="grid place-items-center p-2.5 bg-white shadow-md sticky top-0 z-20 lg:px-8"
	>
		<div
			class="max-w-7xl w-full flex flex-row items-center justify-end px-2"
		>
			<div class="flex flex-row gap-2">
				<NuxtLink to="/login">
					<Button
						variant="subtle"
						color="indigo"
						class="text-black group"
					>
						<span
							class="text-neutral-700 font-normal text-base hidden group-hover:underline group-hover:text-indigo-500 md:block"
							>Log in</span
						>
					</Button>
				</NuxtLink>
				<NuxtLink to="/signup">
					<Button color="indigo"
						>Create account</Button
					>
				</NuxtLink>
			</div>
		</div>
	</header>
</template>
