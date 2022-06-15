import actions from './actions'
import mutations from './mutations'
import getters from './getters'

const state = () => ({
  imgURL: `${process.env.BASE_API}`,
})

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
