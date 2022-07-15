import actions from './actions'
import mutations from './mutations'
import getters from './getters'

const state = () => ({
  imgURL: `${process.env.BASE_API}`,
  productTotal: 1,
})

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
