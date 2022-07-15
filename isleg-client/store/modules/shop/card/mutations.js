const SET_HEADER = (state, payload) => {}
const SET_PRODUCT_TOTAL_INCREMENT = (state) => {
  state.productTotal = state.productTotal + 1
}
const SET_PRODUCT_TOTAL_DICREMENT = (state) => {
  state.productTotal = state.productTotal - 1
}
export default {
  SET_HEADER,
  SET_PRODUCT_TOTAL_INCREMENT,
  SET_PRODUCT_TOTAL_DICREMENT,
}
