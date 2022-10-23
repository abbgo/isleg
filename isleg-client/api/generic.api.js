import axios from 'axios'

const baseURL = process.env.BASE_API
const fileURL = process.env.IMAGE_URL
export const request = async ({
  url,
  method = 'POST',
  headers = {},
  params = {},
  data = {},
  accessToken = null,
  refreshToken = null,
  onUploadProgress = {},
  file = false,
}) => {
  if (file) {
    const formData = new FormData()
    headers['Accept'] = 'application/json'
    headers['Content-Type'] = 'multipart/form-data'
    if (data?.files?.length) {
      for (let i = 0; i < data.files.length; i++) {
        formData.append('files', data.files[i])
      }
    } else {
      for (let [key, value] of Object.entries(data)) {
        formData.append(key, value)
      }
    }
    data = formData
  }
  if (accessToken) {
    headers['Authorization'] = accessToken
  }
  if (refreshToken) {
    headers['RefreshToken'] = refreshToken
  }
  return axios({
    url: `${file ? fileURL : baseURL}/${url}`,
    method,
    headers,
    ...onUploadProgress,
    params,
    data,
  })
}
