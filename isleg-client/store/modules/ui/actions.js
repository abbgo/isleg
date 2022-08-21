import axios from 'axios'

const fetchHeader = async ({ commit }, { url, $nuxt }) => {
  try {
    const { data } = await axios.get(url)
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
const fetchFooter = async ({ commit }, { url, $nuxt }) => {
  try {
    const { data } = await axios.get(url)
    if (data?.status) {
      commit('SET_FOOTER', data?.footer_data)
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
const fetchBrends = async ({ commit }, { url, $nuxt }) => {
  try {
    const { data } = await axios.get(url)
    console.log('brends', data.brends)
    if (data?.status) {
      commit('SET_BRENDS', data.brends)
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
  fetchFooter,
  fetchBrends,
}
