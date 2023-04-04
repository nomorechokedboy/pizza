import { AxiosError } from 'axios'
import { AlertProps } from '~~/components/Alert.vue'
import { notify } from '~~/composables/useNotification'

export function notifyError(e: unknown) {
	if (e instanceof AxiosError) {
		const notification: AlertProps = {
			content: '',
			k: crypto.randomUUID(),
			type: 'error'
		}

		if (e.response?.data) {
			notification.content = checkErrorType(e.response.data)
		} else if (e.message) {
			notification.content = e.message
		} else {
			notification.content = 'Unknown error'
		}

		notify(notification)
	}
}

function checkErrorType(e: unknown): string {
	if (typeof e === 'string') {
		return e
	}

	return JSON.stringify(e)
}
