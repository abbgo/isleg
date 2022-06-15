// import request from '@/api/apiRequest'

const fetchHeader = async ({ commit }, { url, $nuxt }) => {
  // try {
  //   const { data } = await request(url)
  //   commit('SET_HEADER', data)
  // } catch (e) {
  //   if (e && e.response && e.response.status === 404) {
  //     return $nuxt.error({ statusCode: 404, message: e.message })
  //   }
  // }
}

export default {
  fetchHeader,
}
