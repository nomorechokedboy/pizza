<script lang="ts" setup>
const { data: userProfile } = useUserProfile()

definePageMeta({ middleware: ['authn'] })
useHead({ title: userProfile.value?.name })
</script>

<template>
	<main
		class="w-full h-screen mt-5 overflow-y-auto flex flex-col items-center"
	>
		<div
			class="w-full flex-grow flex flex-col max-w-5xl overflow-auto"
		>
			<UserProfileHeader />
			<div
				class="flex flex-col gap-4 flex-grow overflow-y-auto h-auto py-5 md:!flex-row"
			>
				<UserSectionContainer
					class="flex-1 md:flex-none md:w-1/3"
					title="Recent comments"
				>
					<UserListComment />
				</UserSectionContainer>
				<UserSectionContainer
					class="flex-1 md:flex-none md:w-2/3"
					title="Your posts"
				>
					<UserListArticle />
				</UserSectionContainer>
			</div>
		</div>
	</main>
</template>
