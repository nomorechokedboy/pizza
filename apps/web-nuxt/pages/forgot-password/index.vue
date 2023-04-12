<script setup lang="ts">
import TextInput from '@@/components/TextInput.vue'
import { useVuelidate } from '@vuelidate/core'
import { email as emailRule, helpers, required } from '@vuelidate/validators'
import { Button } from 'ui-vue'
import { notify } from '~~/composables/useNotification'

definePageMeta({
	middleware: ['redirect']
})

const { $blogApi } = useNuxtApp()
const formData = reactive({ email: '' })
const emailInput = ref<InstanceType<typeof TextInput> | null>(null)
async function handleSubmit() {
	const isValidForm = await v$.value.$validate()
	if (!isValidForm) {
		return
	}

	try {
		await $blogApi.auth.authForgotPasswordPost({
			email: formData.email
		})
		notify({
			type: 'success',
			k: crypto.randomUUID(),
			content: 'We have sent you an email'
		})
	} catch (e) {
		notifyError(e)
	}
}
function fotgotPasswordFormRules() {
	return {
		email: {
			required: helpers.withMessage(
				'The email field is required',
				required
			),
			email: helpers.withMessage(
				'Invalid email format',
				emailRule
			)
		}
	}
}
const rules = computed(fotgotPasswordFormRules)
const v$ = useVuelidate(rules, formData)

onMounted(() => {
	if (emailInput.value) {
		emailInput.value.focus()
	}
})
</script>

<template>
	<div class="pt-12 pb-32 bg-[rgb(245,245,245)]">
		<div
			class="flex flex-col gap-4 p-8 pt-3 border bg-white lg:max-w-[490px] m-auto"
		>
			<h1 class="font-bold text-[27px]">
				Forgot your password?
			</h1>
			<form class="flex flex-col">
				<TextInput
					:error="v$.email.$error"
					:errors="v$.email.$errors"
					id="email"
					name="email"
					type="email"
					v-model="formData.email"
					ref="emailInput"
					@blur="v$.email.$touch"
				>
					<template #label>Email</template>
				</TextInput>
				<Button
					color="indigo"
					size="md"
					class="!font-medium"
					@click.prevent="handleSubmit"
					block
					>Send me reset password
					instructions</Button
				>
			</form>
		</div>
	</div>
</template>
