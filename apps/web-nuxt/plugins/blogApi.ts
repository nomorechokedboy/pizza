import axiosGlobal, { AxiosError, CreateAxiosDefaults } from 'axios'
import { AuthApi, CommentsApi, PostsApi, UserApi } from '~~/codegen/api'
import { baseURL } from '~~/constants'

const configs: CreateAxiosDefaults = {
	baseURL,
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
}

export default defineNuxtPlugin(() => {
	const axios = axiosGlobal.create(configs)
	axios.interceptors.request.use((config) => {
		const customHeaders = {
			Authorization: ''
		}

		const token = useAuthToken()
		const accessToken = token.value.accessToken

		if (accessToken)
			customHeaders.Authorization = `Bearer ${accessToken}`

		return {
			...config,
			headers: {
				...customHeaders,
				...config.headers
			}
		}
	})

	const auth = new AuthApi(undefined, undefined, axios)
	const comment = new CommentsApi(undefined, undefined, axios)
	const post = new PostsApi(undefined, undefined, axios)
	const user = new UserApi(undefined, undefined, axios)
	const blogApi: BlogApi = { auth, comment, post, user }

	const createAxiosResponseInterceptor = () => {
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

				const token = useAuthToken()
				return auth
					.authRefreshTokenPost({
						refresh_token:
							token.value.refreshToken
					})
					.then((resp) => {
						token.value.refreshToken =
							resp.data.refresh_token
						token.value.accessToken =
							resp.data.token
						error.response.config.headers[
							'Authorization'
						] = `Bearer ${resp.data.token}`
						return axios(
							error.response.config
						)
					})
					.catch(async (e) => {
						token.value.accessToken =
							undefined
						token.value.refreshToken =
							undefined
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
