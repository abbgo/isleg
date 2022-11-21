import { request } from './generic.api'

export const getCategoryProducts = ({ url }) =>
  request({ url: url, method: 'GET' })
