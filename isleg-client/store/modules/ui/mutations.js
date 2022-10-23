const SET_HEADER = (state, payload) => {
  state.logoFavicon = payload.logo_favicon
  state.translationHeader = payload.translation_header
  state.languages = payload.languages
  state.categories = payload.categories
}
const SET_FOOTER = (state, payload) => {
  state.footerDatas = payload
}
const SET_BRENDS = (state, payload) => {
  state.brends = payload
}
const SET_CLOSE_ISOPENSIGNUP = (state) => {
  state.isOpenSignUp = false
}
const SET_OPEN_ISOPENSIGNUP = (state) => {
  state.isOpenSignUp = true
}
const SET_CATEGORY_PRODUCTS = (state, payload) => {
  state.categoryProducts = payload
}
const SET_MY_PROFILE = (state, payload) => {
  state.myProfile = payload
}
export default {
  SET_HEADER,
  SET_FOOTER,
  SET_BRENDS,
  SET_CLOSE_ISOPENSIGNUP,
  SET_OPEN_ISOPENSIGNUP,
  SET_CATEGORY_PRODUCTS,
  SET_MY_PROFILE
}
