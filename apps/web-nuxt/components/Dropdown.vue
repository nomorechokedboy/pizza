<script lang="ts" setup>
export interface DropdownProps {
	user: {
		name: string
		username: string
	}
	open: boolean
}

async function logout() {
	removeToken()
	navigateTo('/login')
}

const { open } = defineProps<DropdownProps>()
</script>

<template>
	<div
		class="dropdown absolute mt-2 px-2 py-2 z-10 text-md bg-white divide-y divide-gray-100 rounded-lg shadow w-[15.625rem] dark:bg-gray-700"
		:data-biding="open"
	>
		<div
			class="pb-2 mb-2 hover:underline hover:rounded-lg px-4 py-3 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
		>
			<a href="#">
				<div>
					<span class="fw-medium block">
						{{ user.name }}
					</span>
					<small class="text-gray-400"
						>@{{ user.username }}
					</small>
				</div>
			</a>
		</div>
		<ul
			class="py-2 text-gray-700 dark:text-gray-200"
			aria-labelledby="dropdownDefaultButton"
		>
			<li>
				<a
					href="#"
					class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
					>Dashboard</a
				>
			</li>
			<li>
				<a
					href="#"
					class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
					>Create Post</a
				>
			</li>
			<li>
				<a
					href="#"
					class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
					>Reading list</a
				>
			</li>
			<li>
				<a
					href="#"
					class="block px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
					>Settings</a
				>
			</li>
		</ul>
		<div class="pt-2 dark:text-gray-200">
			<div
				class="cursor-pointer px-4 py-2 hover:underline hover:rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
				@click="logout"
			>
				Sign out
			</div>
		</div>
	</div>
</template>

<style lang="css" scoped>
.dropdown {
	@apply -right-2;
}

.dropdown[data-biding='true'] {
	@apply block overflow-visible;
}

.dropdown[data-biding='false'] {
	@apply hidden;
}
</style>
