import { request } from './generic.api'

export const productAdd = ({ url, data, accessToken }) =>
  request({ url: url, data: data, accessToken: accessToken })
export const productRemove = ({ url, data, accessToken }) =>
  request({ url: url, data: data, accessToken: accessToken })
