import axiosGlobal, { CreateAxiosDefaults } from 'axios'
import { baseURL } from '~~/constants'

const cfgs: CreateAxiosDefaults = {
	baseURL,
	timeout: 5000,
	headers: {
		'Content-Type': 'application/json'
	},
	withCredentials: true
}

const axios = axiosGlobal.create(cfgs)

/* const createAxiosResponseInterceptor = () => {
    const interceptor = axios.interceptors.response.use(
        (response) => response,
        async (error) => {
            const id = crypto.randomUUID()
            const type: NotificationMessage['type'] = 'error'
            let description = 'Internet error'

            if (error.response?.data?.message) {
                description = error.response?.data?.message
            }

            if (
                error instanceof AxiosError &&
                (error.response?.status !== 401 || !error.response.headers)
            ) {
                openNotification({ description, id, type })
                return Promise.reject(error)
            }

            axios.interceptors.response.eject(interceptor)

            try {
                const { data } = await axios({
                    baseURL: '/api/auth',
                    method: 'POST'
                })
                if (!data) {
                    throw Error('Unknown error from refreshtoken')
                }

                error.config.headers = JSON.parse(
                    JSON.stringify(error.config.headers)
                )
                error.response.config.headers[
                    'Authorization'
                ] = `Bearer ${data.token}`

                return axios(error.response.config)
            } catch (e) {
                openNotification({
                    description: 'Session timeout, please login!',
                    id,
                    type
                })
                return Promise.reject('')
            } finally {
                createAxiosResponseInterceptor()
            }
        }
    )
}

createAxiosResponseInterceptor() */

export default axios
