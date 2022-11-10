import { request } from './generic.api'

export const getMyProfile = ({ url, accessToken }) =>
  request({ url: url, accessToken: accessToken, method: 'GET' })
