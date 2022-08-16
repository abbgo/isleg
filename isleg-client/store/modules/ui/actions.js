import axios from 'axios'

const fetchHeader = async ({ commit }, { url, $nuxt }) => {
  console.log(`${process.env.BASE_API}/${url}`)
  try {
    const { data } = await axios.get(`${process.env.BASE_API}/${url}`)
    console.log('headerData', data.header_data)
    if (data?.status) {
      commit('SET_HEADER', data?.header_data)
    }
  } catch (e) {
    console.log(e.response)
    if (e) {
      return $nuxt.error({
        statusCode: e?.response?.status,
        message: e?.message,
      })
    }
  }
}

export default {
  fetchHeader,
}
