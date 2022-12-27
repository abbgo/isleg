import { request } from './generic.api'

export const getTranslationBasketPage = ({ url }) =>
  request({ url: url, method: 'GET' })
