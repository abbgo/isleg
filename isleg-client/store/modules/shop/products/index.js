import actions from './actions'
import mutations from './mutations'
import getters from './getters'

const state = () => ({
  imgURL: `${process.env.IMAGE_URL}`,
  productsCategories: [],
  productCount: null,
  likesCount: null,
  fillEmpty: null,
  fillColor: '#FD5E29',
  quantity: 0,
  isFavorite: false,
  removedFromBasket: false,
})

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
