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
      commit('SET_CATEGORY_PRODUCTS', data)
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
    console.log(data)
    if (data?.status) {
      commit('SET_MY_PROFILE', data.customer_informations)
    }
  } catch (e) {
    console.log(e)
    // if (e.response.status == 401) {
    //   try {
    //     const res = await axios.post(`${process.env.BASE_API}/auth/refresh`, {
    //       headers: {
    //         RefreshToken: `Bearer ${refreshToken}`,
    //       },
    //     })
    //     console.log(res);
    //   } catch (e) { console.log('sonky', e.response) }
    // } else {
    //   return $nuxt.error({
    //     statusCode: e?.response?.status,
    //     message: e?.message,
    //   })
    // }
  }
}

export default {
  fetchFooter,
  fetchBrends,
  fetchCategoryProducts,
  fetchMyInformation,
}
