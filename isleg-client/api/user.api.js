import { request } from './generic.api'

export const productAdd = ({ url, data, accessToken }) =>
  request({ url: url, data: data, accessToken: accessToken })
export const productLike = ({ url, accessToken }) =>
  request({ url: url, accessToken: accessToken })
export const getRefreshToken = ({ url, refreshToken }) =>
  request({ url: url, refreshToken: refreshToken })
export const deleteAllProductsFromBasket = ({ url, accessToken }) =>
  request({ url: url, accessToken: accessToken })
