<script setup lang="ts">
import { Button } from 'ui-vue'
import IconFacebook from '~icons/fa6-brands/facebook'
import IconGoogle from '~icons/gg/google'
import IconGithub from '~icons/mdi/github'
import IconGitlab from '~icons/ri/gitlab-fill'

const route = useRoute()
const isLoginPage = computed(() => route.name === 'login')
const buttonText = computed(() => (isLoginPage.value ? 'Continue' : 'Sign up'))
</script>

<template>
	<NuxtLayout name="default">
		<main
			class="flex-1 grid place-items-center lg:flex-row gap-8 bg-gray-100 lg:p-8"
		>
			<section
				class="w-full grid items-stretch lg:max-w-7xl place-items-center"
			>
				<div
					class="flex flex-col items-center bg-white p-6 gap-6 lg:p-12 lg:rounded-md lg:max-w-[640px] lg:w-full lg:shadow"
				>
					<header class="w-full">
						<h1
							class="text-center font-bold text-2xl lg:text-3xl"
						>
							Welcome to COPE
							Community
						</h1>
						<p
							class="text-center lg:text-base"
						>
							Copium Community is a
							Community of 1,026,041
							amazing copers
						</p>
					</header>
					<div
						class="w-full flex-1 flex flex-col gap-4"
					>
						<div
							class="flex flex-col gap-2"
						>
							<Button
								color="yellow"
								size="md"
								class="!py-3"
								block
							>
								<template
									#leftIcon
								>
									<IconGoogle />
								</template>
								{{ buttonText }}
								with Google
							</Button>
							<Button
								color="dark"
								size="md"
								class="!py-3"
								block
							>
								<template
									#leftIcon
								>
									<IconGithub />
								</template>
								{{ buttonText }}
								with Github
							</Button>
							<Button
								color="orange"
								size="md"
								class="!py-3"
								block
							>
								<template
									#leftIcon
								>
									<IconGitlab />
								</template>
								{{ buttonText }}
								with Gitlab
							</Button>
							<Button
								size="md"
								class="!py-3"
								block
							>
								<template
									#leftIcon
								>
									<IconFacebook />
								</template>
								{{ buttonText }}
								with Facebook
							</Button>
						</div>
						<div
							class="h-[44px] bg-black text-white"
						>
							Divider
						</div>
						<slot />
						<NuxtLink
							v-if="isLoginPage"
							class="text-center text-sm text-indigo-500"
							to="/forgot-password"
							>I forgot my
							password</NuxtLink
						>
					</div>
				</div>
			</section>
		</main>
	</NuxtLayout>
</template>
