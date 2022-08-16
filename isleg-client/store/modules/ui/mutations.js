const SET_HEADER = (state, payload) => {
  state.logoFavicon = payload.logo_favicon
  state.translationHeader = payload.translation_header
  state.languages = payload.languages
  state.categories = payload.categories
}
const SET_CLOSE_ISOPENSIGNUP = (state) => {
  state.isOpenSignUp = false
}
const SET_OPEN_ISOPENSIGNUP = (state) => {
  state.isOpenSignUp = true
}
export default {
  SET_HEADER,
  SET_CLOSE_ISOPENSIGNUP,
  SET_OPEN_ISOPENSIGNUP,
}
