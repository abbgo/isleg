import { request } from './generic.api'

export const getPaymentText = ({ url }) => request({ url: url, method: 'GET' })
export const getPaymentTypes = ({ url }) => request({ url: url, method: 'GET' })
export const getPaymentTime = ({ url }) => request({ url: url, method: 'GET' })
export const postPaymentDatas = ({ url, data }) =>
  request({ url: url, data: data })
