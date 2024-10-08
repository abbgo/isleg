const imgURL = (state) => {
  return state.imgURL
}
const productsCategories = (state) => {
  return state.productsCategories
}
const productCount = (state) => {
  return state.productCount
}
const fillEmpty = (state) => {
  return state.fillEmpty
}
const fillColor = (state) => {
  return state.fillColor
}
const quantity = (state) => {
  return state.quantity
}
const isFavorite = (state) => {
  return state.isFavorite
}
const removedFromBasket = (state) => {
  return state.removedFromBasket
}
const likesCount = (state) => {
  return state.likesCount
}
export default {
  imgURL,
  productsCategories,
  productCount,
  fillEmpty,
  fillColor,
  quantity,
  isFavorite,
  removedFromBasket,
  likesCount,
}
