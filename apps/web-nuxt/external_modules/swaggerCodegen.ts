import axios from './axios'
import { AuthApi, UserApi } from '~~/codegen/api'

export const authApi = new AuthApi(undefined, undefined, axios)
export const userApi = new UserApi(undefined, undefined, axios)
