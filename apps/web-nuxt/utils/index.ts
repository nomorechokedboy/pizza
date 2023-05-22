import {
	EventStreamContentType,
	fetchEventSource
} from '@microsoft/fetch-event-source'
import { InfiniteData } from '@tanstack/vue-query'
import { AxiosError } from 'axios'
import dayjs from 'dayjs'
import calendar from 'dayjs/plugin/calendar'
import relativeTime from 'dayjs/plugin/relativeTime'
import jwtDecode, { JwtPayload } from 'jwt-decode'
import { CommonBasePaginationResponseEntitiesPostResponse } from '~/codegen/api'
import { ArticleData } from '~/components/ListArticle.vue'
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

export function timeFromNow(
	d?: string,
	withSuffix: boolean | undefined = undefined
) {
	let date = d
	if (!date) {
		date = dayjs().toISOString()
	}

	return dayjs(d).fromNow(withSuffix)
}

export function flattenPostData(
	postData:
		| globalThis.Ref<undefined>
		| globalThis.Ref<
				InfiniteData<CommonBasePaginationResponseEntitiesPostResponse>
		  >
) {
	const appConfig = useRuntimeConfig()
	const data: ArticleData[] = []
	if (!postData.value) {
		return data
	}

	for (let i = 0; i < postData.value.pages.length; i++) {
		const pageData = postData.value.pages[i]
		if (!pageData.items) {
			return data
		}

		for (let j = 0; j < pageData.items.length; j++) {
			const {
				user,
				id,
				publishedAt,
				title,
				slug,
				commentCount,
				image
			} = pageData.items[j]
			const item: ArticleData = {
				commentCount: commentCount || 0,
				id: id || 0,
				image,
				publishedAt,
				slug: slug || '',
				title: title || '',
				user: {
					avatar:
						user?.avatar ||
						`${appConfig.public.dicebearMedia}${user?.fullName}`,
					id: id || 0,
					name: user?.fullName || 'User name'
				}
			}
			data.push(item)
		}
	}

	return data
}

class AuthError extends Error {}
class RetriableError extends Error {}
class FatalError extends Error {}

export function setupNotifyConnection(
	refetch: any,
	notifyController: AbortController | undefined
) {
	const authToken = useAuthToken()
	const config = useRuntimeConfig()

	return fetchEventSource(config.public.notifyUrl, {
		signal: notifyController?.signal,
		headers: {
			Authorization: `Bearer ${authToken.value}`
		},
		async onopen(response) {
			if (
				response.ok &&
				response.headers.get('content-type') ===
					EventStreamContentType
			) {
				return // everything's good
			} else if (response.status === 401) {
				throw new AuthError()
			} else if (response.status >= 500) {
				// client-side errors are usually non-retriable:
				throw new FatalError()
			} else {
				throw new RetriableError()
			}
		},
		onmessage(msg) {
			// if the server emits an error message, throw an exception
			// so it gets handled by the onerror callback below:
			if (msg.event === 'notification') {
				refetch()
			}
		},
		onclose() {
			// if the server closes the connection unexpectedly, retry:
			throw new RetriableError()
		},
		onerror(err) {
			if (
				err instanceof FatalError ||
				err instanceof AuthError
			) {
				throw err // rethrow to stop the operation
			} else {
				// do nothing to automatically retry. You can also
				// return a specific retry interval here.
				return 5000
			}
		}
	})
}

export function isAuthError(e: unknown): boolean {
	if (e instanceof AuthError) {
		return true
	}

	return false
}
