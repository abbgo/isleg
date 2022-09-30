import actions from './actions'
import mutations from './mutations'
import getters from './getters'

const state = () => ({
  imgURL: `${process.env.IMAGE_URL}`,
  productTotal: 1,
})

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
