import axios from 'axios'

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
const fetchCategoryProducts = async ({ commit }, { url, $nuxt }) => {
  try {
    const { data } = await axios.get(url)
    console.log('data', data)
    if (data?.status) {
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
const fetchMyInformation = async (
  ctx,
  { url, accessToken, refreshToken, $nuxt }
) => {
  try {
    const { data } = await axios.get(url, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    })
    if (data?.status) {
      commit('SET_MY_PROFILE', data.customer_informations)
    }
  } catch (e) {
    console.log(e)
  }
}
const initAuth = ({ commit }) => {
  if (localStorage.getItem('lorem')) {
    commit('SET_AUTHENTICATION', true)
  } else {
    commit('SET_AUTHENTICATION', false)
  }
}
export default {
  fetchFooter,
  fetchBrends,
  fetchCategoryProducts,
  fetchMyInformation,
  initAuth,
}
