import actions from './actions'
import mutations from './mutations'
import getters from './getters'

const state = () => ({
  imgURL: process.env.IMAGE_URL,
  isOpenSignUp: false,
  logoFavicon: null,
  translationHeader: null,
  languages: null,
  categories: null,
  footerDatas: null,
  brends: null,
})

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
