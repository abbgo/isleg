<template>
  <header class="header __container">
    <div class="header__container">
      <the-header-logo></the-header-logo>
      <div class="serach">
        <span class="search__icon">
          <svg
            width="22"
            height="21"
            viewBox="0 0 22 21"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M19.7734 20.712L14.6728 15.623C13.2047 16.6821 11.4401 17.2514 9.62982 17.25C7.34162 17.2476 5.1479 16.3381 3.52924 14.7207C1.91058 13.1034 0.99929 10.9103 0.995056 8.62207C0.999557 6.33441 1.91147 4.14189 3.53021 2.52539C5.14896 0.908894 7.34216 0.000342889 9.62982 -0.000976562C11.9177 7.75772e-05 14.1123 0.908528 15.7314 2.52502C17.3505 4.14152 18.2618 6.33415 18.2665 8.62207C18.2679 10.5813 17.5997 12.4821 16.372 14.009L21.4228 19.056C21.6214 19.2756 21.7255 19.565 21.7109 19.8607C21.6963 20.1564 21.564 20.4342 21.3447 20.6331C21.1171 20.8638 20.8084 20.9956 20.4843 21C20.3527 21.0012 20.2216 20.9763 20.0995 20.9269C19.9775 20.8775 19.8671 20.8045 19.7734 20.712ZM3.37299 8.59998C3.37616 10.2666 4.04028 11.8639 5.21967 13.0414C6.39906 14.2189 7.99742 14.8805 9.664 14.881C11.3288 14.8783 12.9245 14.2157 14.1015 13.0383C15.2785 11.8609 15.9412 10.2648 15.9433 8.59998C15.9409 6.93569 15.2786 5.34037 14.1015 4.16382C12.9244 2.98727 11.3283 2.32556 9.664 2.32397C7.99865 2.3245 6.40121 2.98556 5.2226 4.16211C4.04398 5.33866 3.38033 6.93463 3.37689 8.59998H3.37299ZM7.07318 6.97705C6.98247 6.81369 6.93451 6.62985 6.93451 6.44299C6.93451 6.25613 6.98247 6.07242 7.07318 5.90906C7.44372 5.53671 7.8846 5.24157 8.37006 5.04089C8.85551 4.84022 9.37601 4.73795 9.90131 4.73999C10.9733 4.73531 12.0028 5.15573 12.7656 5.90906C12.8622 6.0704 12.914 6.2549 12.914 6.44299C12.914 6.63108 12.8622 6.8157 12.7656 6.97705C12.6621 7.10898 12.5246 7.20937 12.3671 7.26697C12.2097 7.32457 12.0391 7.33701 11.8749 7.30298C11.6316 7.01683 11.3293 6.78712 10.9882 6.62964C10.6472 6.47216 10.276 6.39069 9.90033 6.39099C9.52513 6.39165 9.15399 6.47356 8.81342 6.63098C8.47284 6.78841 8.17028 7.01761 7.9267 7.30298C7.90386 7.32512 7.87669 7.34205 7.84662 7.35242C7.81655 7.36279 7.7845 7.3664 7.75287 7.36304C7.61839 7.35482 7.48806 7.31577 7.37103 7.24902C7.254 7.18227 7.1536 7.08959 7.07806 6.97803L7.07318 6.97705Z"
              fill="#FD5E29"
            />
          </svg>
        </span>
        <input
          @focus="focused"
          @blur="focusRemove"
          class="input__border"
          type="text"
          placeholder="Gözleg"
        />
      </div>
      <the-header-items
        @openSignUp="openPopUp"
        :isProfile="isProfile"
      ></the-header-items>
    </div>
    <the-header-nav></the-header-nav>
    <sign-up
      v-if="isOpenSignUp"
      :signUp="signUp"
      @closeSignUp="closeSignUp"
      @openRegister="openRegister"
      @closeSignUpPopUp="closeSignUp"
    ></sign-up>
    <register
      v-if="isOpenRegister"
      :register="register"
      @closeRegister="closeRegister"
      @openSignUpPopUp="openSignUpPopUp"
      @closeRegisterPopUp="closeRegister"
    ></register>
  </header>
</template>

<script>
import TheHeaderLogo from './TheHeaderLogo.vue'
import { mapGetters } from 'vuex'
export default {
  components: { TheHeaderLogo },
  data() {
    return {
      isOpenRegister: false,
      signUp: {
        phone_number: '+993 6',
        newPhoneNumbers: [],
        password: '',
      },
      register: {
        name: '',
        phone_number: '+993 6',
        newPhoneNumbers: [],
        password: '',
      },
      isProfile: false,
    }
  },
  watch: {
    '$i18n.locale': async function () {
      await this.fetchHeader({
        url: `${this.$i18n.locale}/header/${this.calculateHeaderApiPath}`,
        $nuxt: this.$nuxt,
      })
    },
  },
  async fetch() {
    await this.$store.dispatch('ui/fetchHeader', {
      url: `${this.$i18n.locale}/header/${this.calculateHeaderApiPath}`,
      $nuxt: this.$nuxt,
    })
  },
  computed: {
    ...mapGetters('ui', ['isOpenSignUp']),
  },
  methods: {
    focused() {
      let serach = document.querySelector('.serach')
      serach.classList.add('focus__search')
    },
    focusRemove() {
      let serach = document.querySelector('.serach')
      serach.classList.remove('focus__search')
    },
    closeSignUp() {
      this.signUp.phone_number = '+993'
      this.signUp.password = ''
      this.$store.commit('ui/SET_CLOSE_ISOPENSIGNUP')
    },
    closeRegister() {
      this.register.name = ''
      this.register.phone_number = '+993'
      this.register.password = ''
      this.isOpenRegister = false
    },
    openRegister() {
      this.closeSignUp()
      this.isOpenRegister = true
    },
    openSignUpPopUp() {
      this.closeRegister()
      this.$store.commit('ui/SET_OPEN_ISOPENSIGNUP')
    },
    openPopUp() {
      let token = this.$cookies.get('token')
      console.log(token)
      token
        ? (this.isProfile = !this.isProfile)
        : this.$store.commit('ui/SET_OPEN_ISOPENSIGNUP')
    },
  },
}
</script>
