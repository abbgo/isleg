import { request } from './generic.api'

export const productAdd = ({ url, data, accessToken }) =>
  request({ url: url, data: data, accessToken: accessToken })
export const productLike = ({ url, accessToken }) =>
  request({ url: url, accessToken: accessToken })
export const getRefreshToken = ({ url, refreshToken }) =>
  request({ url: url, refreshToken: refreshToken })
export const deleteAllProductsFromBasket = ({ url, accessToken }) =>
  request({ url: url, accessToken: accessToken })
export const getMyInformation = ({ url, accessToken }) =>
  request({ url: url, accessToken: accessToken })
export const userLogin = ({ url, data }) => request({ url: url, data: data })
export const sendMail = ({ url, data }) => request({ url: url, data: data })
export const translationContact = ({ url }) =>
  request({ url: url, method: 'GET' })
export const companyPhones = ({ url }) => request({ url: url, method: 'GET' })
