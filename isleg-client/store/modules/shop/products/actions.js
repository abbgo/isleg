import axios from 'axios'

const fetchProductsCategories = async ({ commit }, { url, $nuxt }) => {
  try {
    const { data } = await axios.get(url)
    console.log(data)
    commit('SET_PRODUCTS_CATEGORIES', data.homepage_categories)
  } catch (e) {
    if (e && e.response && e.response.status === 404) {
      return $nuxt.error({ statusCode: 404, message: e.message })
    }
  }
}

export default {
  fetchProductsCategories,
}
