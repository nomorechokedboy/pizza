<script setup lang="ts">
import useVuelidate from '@vuelidate/core'
import {
	email,
	helpers,
	maxLength,
	minLength,
	required,
	sameAs
} from '@vuelidate/validators'
import { Button } from 'ui-vue'

const formData = reactive({
	identifier: '',
	password: '',
	confirmPassword: '',
	fullName: '',
	userName: ''
})
const { $blogApi } = useNuxtApp()
function formDataRules() {
	return {
		identifier: {
			required: helpers.withMessage(
				'The email field is required',
				required
			),
			email: helpers.withMessage(
				'Invalid email format',
				email
			)
		},
		userName: {
			required: helpers.withMessage(
				'The user name field is required',
				required
			),
			minLength: minLength(3),
			maxLength: maxLength(30)
		},
		fullName: {
			required: helpers.withMessage(
				'The full name field is required',
				required
			),
			minLength: minLength(3),
			maxLength: maxLength(30)
		},
		password: {
			required: helpers.withMessage(
				'The password field is required',
				required
			),
			minLength: minLength(8)
		},
		confirmPassword: {
			required: helpers.withMessage(
				'The password confirmation field is required',
				required
			),
			sameAs: helpers.withMessage(
				"Passwords don't match",
				sameAs(formData.password)
			)
		}
	}
}
const rules = computed(formDataRules)
const v$ = useVuelidate(rules, formData)
async function handleSignUp() {
	const isFormValid = await v$.value.$validate()
	if (!isFormValid) {
		return
	}

	try {
		await $blogApi.auth.authRegisterPost({
			email: formData.identifier,
			password: formData.password,
			fullname: formData.fullName,
			username: formData.userName
		})
		notify({
			k: crypto.randomUUID(),
			content: 'Signup Success',
			type: 'success'
		})
		await navigateTo('/login')
	} catch (e) {
		notifyError(e)
	}
}
const opened = ref(true)
function handleShowDropdown() {
	opened.value = true
}
/* function handleCloseDropdown() {
    opened.value = false
} */
/* const requirements = [
    { re: /[0-9]/, label: 'Includes number' },
    { re: /[a-z]/, label: 'Includes lowercase letter' },
    { re: /[A-Z]/, label: 'Includes uppercase letter' },
    { re: /[$&+,:;=?@#|'<>.^*()%!-]/, label: 'Includes special symbol' },
]; */
</script>

<template>
	<form class="flex flex-col">
		<TextInput
			:errors="v$.fullName.$errors"
			:error="v$.fullName.$error"
			name="fullname"
			v-model="formData.fullName"
			@blur="v$.fullName.$touch"
		>
			<template #label>Full name</template>
		</TextInput>
		<div class="">
			<TextInput
				:error="v$.identifier.$error"
				:errors="v$.identifier.$errors"
				name="email"
				v-model="formData.identifier"
				@blur="v$.identifier.$touch"
			>
				<template #label>Email</template>
			</TextInput>
			<TextInput
				:errors="v$.userName.$errors"
				:error="v$.userName.$error"
				name="username"
				v-model="formData.userName"
				@blur="v$.userName.$touch"
			>
				<template #label>User name</template>
			</TextInput>
		</div>
		<div class="relative">
			<TextInput
				class="peer"
				:error="v$.password.$error"
				:errors="v$.password.$errors"
				name="password"
				type="password"
				v-model="formData.password"
				@focus="handleShowDropdown"
				@blur="v$.password.$touch"
			>
				<template #label>Password</template>
			</TextInput>
			<!-- <Transition>
                <div v-if="opened" id="dropdown"
                    class="z-10 absolute left-0 w-full bg-white divide-y divide-gray-100 rounded-lg shadow dark:bg-gray-700">
                    <ul class="py-2 text-sm text-gray-700 dark:text-gray-200" aria-labelledby="dropdownDefaultButton">
                        <li>
                            <div class="py-2 px-4">
                                <div class="w-full bg-gray-200 rounded-full h-2 dark:bg-gray-700">
                                    <div class="bg-blue-600 h-2.5 rounded-full" style="width: 45%"></div>
                                </div>
                            </div>
                        </li>
                        <li>
                            <a href="#"
                                class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Dashboard</a>
                        </li>
                        <li>
                            <a href="#"
                                class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Settings</a>
                        </li>
                        <li>
                            <a href="#"
                                class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Earnings</a>
                        </li>
                        <li>
                            <a href="#"
                                class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Sign
                                out</a>
                        </li>
                    </ul>
                </div>
            </Transition> -->
		</div>
		<TextInput
			:errors="v$.confirmPassword.$errors"
			:error="v$.confirmPassword.$error"
			name="confirmPassword"
			type="password"
			v-model="formData.confirmPassword"
			@blur="v$.confirmPassword.$touch"
		>
			<template #label>Confirm Password</template>
		</TextInput>
		<Button
			color="indigo"
			size="md"
			class="!py-3"
			@click.prevent="handleSignUp"
			block
			>Sign up</Button
		>
	</form>
</template>

<style scoped>
.v-enter-active,
.v-leave-active {
	transition: opacity 0.2s ease-in-out;
}

.v-enter-from,
.v-leave-to {
	opacity: 0;
}
</style>
