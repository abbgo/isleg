import { request } from './generic.api'

export const getCategoryProducts = ({ url }) =>
  request({ url: url, method: 'GET' })
export const getBrendProducts = ({ url }) =>
  request({ url: url, method: 'GET' })
