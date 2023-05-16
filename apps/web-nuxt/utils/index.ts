import { AxiosError } from 'axios'
import dayjs from 'dayjs'
import calendar from 'dayjs/plugin/calendar'
import relativeTime from 'dayjs/plugin/relativeTime'
import jwtDecode, { JwtPayload } from 'jwt-decode'
import { AlertProps } from '~~/components/Alert.vue'
import { notify } from '~~/composables/useNotification'

export * from './astToVue'
export * from './rehypeFilter'
export * from './storyFixture'
export * from './types'
export * from './uriTransformer'

dayjs.extend(calendar)
dayjs.extend(relativeTime)

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

export function convertDate(d?: string, format = 'MMM D'): string {
	let date = d
	if (!date) {
		date = dayjs().toISOString()
	}

	return dayjs(date).format(format)
}

export function isTokenExpired(token: string): boolean {
	const { exp } = jwtDecode<JwtPayload>(token)
	if (!exp) return false

	return !(Date.now() >= exp * 1000)
}

export function getCalendarTime(d?: string) {
	let date = d
	if (!date) {
		date = dayjs().toISOString()
	}
	return dayjs(date).calendar()
}

export function timeFromNow(d: string) {
	return dayjs(d).fromNow()
}
