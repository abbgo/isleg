<template>
  <header class="header">
    <div class="header__wrapper __container">
      <div class="header__container">
        <the-header-logo :imgURL="imgURL" :logo="logo"></the-header-logo>

        <the-header-items
          @openSignUp="openPopUp"
          :isProfile="isProfile"
          :myInformation="myInformation"
          :myOrders="myOrders"
          :logOut="logOut"
          :activeLang="activeLang"
          :withOutLocaleLang="withOutLocaleLang"
          :imgURL="imgURL"
          @closeProfilePopUp="isProfile = false"
        ></the-header-items>
      </div>
      <the-header-nav
        :categories="categories"
        :imgURL="imgURL"
      ></the-header-nav>
    </div>

    <sign-up
      :isOpenSignUp="isOpenSignUp"
      :phone="phone"
      :password="password"
      :signIn="signIn"
      :userSignUp="userSignUp"
      :forgotPassword="forgotPassword"
      @closeSignUp="closeSignUp"
      @openRegisterPopUp="openRegisterPopUp"
      @closeSignUpPopUp="closeSignUp"
    ></sign-up>
    <register
      :isOpenRegister="isOpenRegister"
      :name="name"
      :passwordVerification="passwordVerification"
      :verifySecure="verifySecure"
      :phone="phone"
      :password="password"
      :signIn="signIn"
      :userSignUp="userSignUp"
      @closeRegister="closeRegister"
      @openSignUpPopUp="openSignUpPopUp"
      @closeRegisterPopUp="closeRegister"
      @registerPost="registerPost"
    ></register>
  </header>
</template>

<script>
import { getHeaderDatas } from '@/api/header.api'
import TheHeaderLogo from './TheHeaderLogo.vue'
import { mapGetters } from 'vuex'
export default {
  components: { TheHeaderLogo },
  data() {
    return {
      isOpenRegister: false,
      isProfile: false,
    }
  },
  watch: {
    $route: async function () {
      await this.fetchHeaderDatas()
    },
  },
  async fetch() {
    await this.fetchHeaderDatas()
  },
  computed: {
    ...mapGetters('ui', [
      'imgURL',
      'isOpenSignUp',
      'logo',
      'research',
      'phone',
      'password',
      'forgotPassword',
      'signIn',
      'userSignUp',
      'name',
      'passwordVerification',
      'verifySecure',
      'myInformation',
      'myFavorites',
      'myOrders',
      'logOut',
      'languages',
      'categories',
    ]),
    withOutLocaleLang() {
      return this.languages.filter(
        (lang) => lang.name_short != this.$i18n.locale
      )
    },
    activeLang() {
      const find = this.languages.find(
        (lang) => lang.name_short == this.$i18n.locale
      )
      return find
    },
  },
  mounted() {
    document.addEventListener('click', (event) => {
      const account = document.querySelector('.account')
      const profileBox = document.querySelector('.profile__box')
      const isAccount = account.contains(event.target)
      const isProfileBox = profileBox.contains(event.target)
      if (!isAccount && !isProfileBox) {
        this.isProfile = false
      }
    })
  },
  methods: {
    async fetchHeaderDatas() {
      try {
        const { header_data, status } = (
          await getHeaderDatas({
            url: `${this.$i18n.locale}/header`,
          })
        ).data
        if (status) {
          this.$store.commit('ui/SET_HEADER', header_data)
        }
      } catch (error) {
        console.log('header', error)
        if (error) {
          return this.$nuxt.error({
            statusCode: error?.response?.status,
            message: error?.message,
          })
        }
      }
    },
    focused() {
      let serach = document.querySelector('.serach')
      serach.classList.add('focus__search')
    },
    focusRemove() {
      let serach = document.querySelector('.serach')
      serach.classList.remove('focus__search')
    },
    closeSignUp() {
      this.$store.commit('ui/SET_CLOSE_ISOPENSIGNUP')
      document.body.classList.remove('_lock')
    },
    closeRegister() {
      this.isOpenRegister = false
      document.body.classList.remove('_lock')
    },
    openRegisterPopUp() {
      this.closeSignUp()
      this.isOpenRegister = true
      document.body.classList.add('_lock')
    },
    openSignUpPopUp() {
      this.closeRegister()
      this.$store.commit('ui/SET_OPEN_ISOPENSIGNUP')
      document.body.classList.add('_lock')
    },
    async openPopUp() {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (cart && cart?.auth?.accessToken && this.$auth.loggedIn) {
        this.isProfile = !this.isProfile
      } else {
        document.body.classList.add('_lock')
        this.$store.commit('ui/SET_OPEN_ISOPENSIGNUP')
      }
    },
    async registerPost() {
      const formData = new FormData()
      formData.append('full_name', this.register.name)
      formData.append('password', this.register.password)
      formData.append('phone_number', this.register.phone_number)
      formData.append('gender', '1')
      formData.append('birthday', '1998-06-23')
      formData.append('adress', ['wekfdnwejk', 'wejdbkwejb'])
      try {
        const res = await this.$axios.post('/tm/register', formData)
      } catch (e) {
        console.log(e.response)
      }
    },
  },
}
</script>
