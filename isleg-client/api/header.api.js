import { request } from './generic.api'

export const getHeaderDatas = ({ url }) => request({ url: url, method: 'GET' })
