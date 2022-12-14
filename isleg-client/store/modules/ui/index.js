import actions from './actions'
import mutations from './mutations'
import getters from './getters'

const state = () => ({
  isUserLoggined: false,
  imgURL: process.env.IMAGE_URL,
  isOpenSignUp: false,
  logoFavicon: null,
  translationHeader: null,
  languages: null,
  categories: null,
  footerDatas: null,
  brends: null,
  categoryProducts: null,
  myProfile: null,
})

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
