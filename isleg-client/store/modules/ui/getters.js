const imgURL = (state) => {
  return state.imgURL
}
const logo = (state) => {
  return state.logoFavicon?.logo
}
const research = (state) => {
  return state.translationHeader?.research
}
const phone = (state) => {
  return state.translationHeader?.phone
}
const password = (state) => {
  return state.translationHeader?.password
}
const forgotPassword = (state) => {
  return state.translationHeader?.forgot_password
}
const signIn = (state) => {
  return state.translationHeader?.sign_in
}
const name = (state) => {
  return state.translationHeader?.name
}
const passwordVerification = (state) => {
  return state.translationHeader?.password_verification
}
const verifySecure = (state) => {
  return state.translationHeader?.verify_secure
}
const myInformation = (state) => {
  return state.translationHeader?.my_information
}
const myFavorites = (state) => {
  return state.translationHeader?.my_favorites
}
const myOrders = (state) => {
  return state.translationHeader?.my_orders
}
const logOut = (state) => {
  return state.translationHeader?.log_out
}
const userSignUp = (state) => {
  return state.translationHeader?.sign_up
}
const basket = (state) => {
  return state.translationHeader?.basket
}
const languages = (state) => {
  return state.languages
}
const categories = (state) => {
  return state.categories
}
const about = (state) => {
  return state.footerDatas?.about
}
const payment = (state) => {
  return state.footerDatas?.payment
}
const contact = (state) => {
  return state.footerDatas?.contact
}
const secure = (state) => {
  return state.footerDatas?.secure
}
const word = (state) => {
  return state.footerDatas?.word
}
const brends = (state) => {
  return state.brends
}
const isOpenSignUp = (state) => {
  return state.isOpenSignUp
}
const categoryProductsName = (state) => {
  return (
    state.categoryProducts &&
    state.categoryProducts.category &&
    state.categoryProducts.category.name
  )
}
const categoryProducts = (state) => {
  return (
    state.categoryProducts &&
    state.categoryProducts.category &&
    state.categoryProducts.category.products
  )
}
const myProfile = (state) => {
  return state.myProfile
}
const isUserLoggined = (state) => {
  return state.isUserLoggined
}
const isAuthenticated = (state) => {
  return state.authenticated
}
export default {
  isAuthenticated,
  isUserLoggined,
  imgURL,
  logo,
  research,
  phone,
  password,
  forgotPassword,
  signIn,
  userSignUp,
  name,
  passwordVerification,
  verifySecure,
  myInformation,
  myFavorites,
  myOrders,
  logOut,
  basket,
  languages,
  categories,
  about,
  payment,
  contact,
  secure,
  word,
  brends,
  isOpenSignUp,
  categoryProductsName,
  categoryProducts,
  myProfile,
}
