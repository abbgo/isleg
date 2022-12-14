const SET_PRODUCTS_CATEGORIES = (state, payload) => {
  state.productsCategories = payload
  for (let i = 0; i < state.productsCategories.length; i++) {
    if (state.productsCategories[i].products) {
      for (let j = 0; j < state.productsCategories[i].products.length; j++) {
        state.productsCategories[i].products[j]['quantity'] = Number(0)
        state.productsCategories[i].products[j]['is_favorite'] = false
      }
    }
  }
}
const SET_PRODUCT_TOTAL_INCREMENT = (state, { data, quantity }) => {
  data.quantity = quantity
  state.productCount += 1
}
const SET_PRODUCT_TOTAL_DECREMENT = (state, { data, quantity }) => {
  data.quantity = quantity
  state.productCount -= 1
}
const SET_PRODUCT_COUNT = (state, count) => {
  state.productCount = count
}
const SET_PRODUCT_FAVORITE = (state, { data, isFavorite }) => {
  data.is_favorite = isFavorite
}
const SET_BASKET_PRODUCT_COUNT = (state, payload) => {
  state.productCount -= payload
}
const SET_REMOVED_FROM_BASKET = (state, payload) => {
  state.removedFromBasket = payload
}
const SET_PRODUCT_COUNT_WHEN_PAYMENT = (state, payload) => {
  state.productCount = payload
}
const SET_LIKES_COUNT_INCREMENT = (state) => {
  state.likesCount += 1
}
const SET_LIKES_COUNT_DECREMENT = (state) => {
  state.likesCount -= 1
}
const SET_PRODUCT_LIKES = (state, payload) => {
  state.likesCount = payload
}
export default {
  SET_PRODUCTS_CATEGORIES,
  SET_PRODUCT_TOTAL_INCREMENT,
  SET_PRODUCT_TOTAL_DECREMENT,
  SET_PRODUCT_COUNT,
  SET_PRODUCT_FAVORITE,
  SET_BASKET_PRODUCT_COUNT,
  SET_REMOVED_FROM_BASKET,
  SET_PRODUCT_COUNT_WHEN_PAYMENT,
  SET_LIKES_COUNT_INCREMENT,
  SET_LIKES_COUNT_DECREMENT,
  SET_PRODUCT_LIKES,
}
