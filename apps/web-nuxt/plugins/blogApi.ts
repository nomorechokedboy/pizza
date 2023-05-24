import axiosGlobal, { AxiosError, CreateAxiosDefaults } from 'axios'
import { NotificationApi } from '~/codegen/noisy-boi'
import {
	AuthApi,
	CommentsApi,
	Configuration,
	MediaApi,
	PostsApi,
	UserApi
} from '~~/codegen/api'

const configs: CreateAxiosDefaults = {
	timeout: 5000,
	headers: {
		'Content-Type': 'application/json'
	}
}

type BlogApi = {
	auth: AuthApi
	comment: CommentsApi
	post: PostsApi
	user: UserApi
	notification: NotificationApi
	media: MediaApi
}

function autoAttachTokenInterceptor(config: any) {
	const customHeaders = {
		Authorization: ''
	}

	const token = useAuthToken()
	const accessToken = token.value

	if (accessToken) customHeaders.Authorization = `Bearer ${accessToken}`

	return {
		...config,
		headers: {
			...customHeaders,
			...config.headers
		}
	}
}

export default defineNuxtPlugin(() => {
	const axios = axiosGlobal.create(configs)
	const nuxtConfig = useRuntimeConfig()

	axios.interceptors.request.use(autoAttachTokenInterceptor)

	const blogApiConfig = new Configuration({
		baseOptions: { baseURL: nuxtConfig.public.apiUrl }
	})
	const notificationApiConfig = new Configuration({
		baseOptions: { baseURL: nuxtConfig.public.notificationUrl }
	})
	const auth = new AuthApi(blogApiConfig, undefined, axios)
	const comment = new CommentsApi(blogApiConfig, undefined, axios)
	const post = new PostsApi(blogApiConfig, undefined, axios)
	const user = new UserApi(blogApiConfig, undefined, axios)
	const media = new MediaApi(blogApiConfig, undefined, axios)
	const notification = new NotificationApi(
		notificationApiConfig,
		undefined,
		axios
	)
	const blogApi: BlogApi = {
		auth,
		comment,
		post,
		user,
		notification,
		media
	}

	function createAxiosResponseInterceptor() {
		const interceptor = axios.interceptors.response.use(
			(response) => response,
			async (error) => {
				/* const id = crypto.randomUUID()
                const type: NotificationMessage['type'] = 'error'
                let description = 'Internet error' */

				if (error.response?.data?.message) {
					// description = error.response?.data?.message
				}

				if (
					error instanceof AxiosError &&
					(error.response?.status !== 401 ||
						!error.response.headers)
				) {
					// openNotification({ description, id, type })
					return Promise.reject(error)
				}

				axios.interceptors.response.eject(interceptor)

				const refreshToken = useRefreshToken()
				return blogApi.auth
					.authRefreshTokenPost({
						refresh_token:
							refreshToken.value
					})
					.then((resp) => {
						onRefreshToken(resp)
						error.response.config.headers[
							'Authorization'
						] = `Bearer ${resp.data.token}`
						return axios(
							error.response.config
						)
					})
					.catch(async (e) => {
						onRefreshTokenError()
						await navigateTo('/login')
						return Promise.reject(e)
					})
					.finally(createAxiosResponseInterceptor)
			}
		)
	}

	createAxiosResponseInterceptor()

	return {
		provide: {
			blogApi
		}
	}
})
