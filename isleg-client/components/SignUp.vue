<template>
  <div :class="['pop-up', { active: isOpenSignUp }]" @click="closeSignUp">
    <div class="pop-up__body" @click.stop>
      <div class="pop-up__wrapper">
        <div class="pop-up__close" @click="closeSignUp">
          <img src="@/assets/img/close.svg" alt="" />
        </div>
        <div class="pop-up_form">
          <div class="form__input">
            <h4>{{ phone }}</h4>
            <div class="form__input-container">
              <img src="@/assets/img/tel.svg" alt="" />
              <input
                type="tel"
                :placeholder="phone"
                v-model="$v.signUp.phone_number.$model"
                @input="enforcePhoneFormat"
              />
            </div>
            <span class="error" v-if="isPhoneNumber">
              {{ $t('register.phoneNumberIsRequired') }}
            </span>
          </div>
          <div class="form__input">
            <h4>{{ password }}</h4>
            <div class="form__input-container">
              <img src="@/assets/img/lock.svg" alt="" />
              <input
                :type="showPass ? 'password' : 'text'"
                :placeholder="password"
                v-model="$v.signUp.password.$model"
              />
              <img
                @click="showPass = !showPass"
                :src="showPass ? '/img/Hide.svg' : '/img/Show.svg'"
                alt=""
              />
            </div>
            <span
              class="error"
              v-if="$v.signUp.password.$error && !$v.signUp.password.required"
            >
              {{ $t('register.passwordIsRequired') }}
            </span>
            <span class="error" v-if="!$v.signUp.password.minLength">
              {{ $t('register.passwordMustHavetletters') }}
            </span>
          </div>
          <div class="form-chek__password">
            <p>{{ forgotPassword }}</p>
          </div>
          <div class="pop-up__btns">
            <button class="left_btn" @click="openRegister">
              {{ userSignUp }}
            </button>
            <button type="button" class="right__btn" @click="logIn">
              {{ signIn }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators'
export default {
  props: {
    isOpenSignUp: {
      type: Boolean,
      default: () => false,
    },
    phone: {
      type: String,
      default: () => '',
    },
    password: {
      type: String,
      default: () => '',
    },
    forgotPassword: {
      type: String,
      default: () => '',
    },
    signIn: {
      type: String,
      default: () => '',
    },
    userSignUp: {
      type: String,
      default: () => '',
    },
  },
  data() {
    return {
      showPass: true,
      isPhoneNumber: false,
      signUp: {
        phone_number: '+9936',
        password: '',
      },
    }
  },
  validations: {
    signUp: {
      phone_number: {
        required,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    enforcePhoneFormat() {
      this.isPhoneNumber = false
      let x = this.signUp.phone_number
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
      this.signUp.phone_number = !x[2]
        ? '+9936'
        : '+9936' +
          (x[3] ? x[3] : '') +
          (x[4] ? x[4] : '') +
          (x[5] ? x[5] : '') +
          (x[6] ? x[6] : '')
    },
    async logIn() {
      this.$v.$touch()
      if (this.signUp.phone_number.length >= 12) {
        try {
          let response = await this.$auth.loginWith('userLogin', {
            data: {
              phone_number: this.signUp.phone_number,
              password: this.signUp.password,
            },
          })
          if (this.$auth.loggedIn) {
            const { access_token, customer_id, refresh_token } = response.data
            this.$cookies.set('access_token', access_token)
            this.$cookies.set('customer_id', customer_id)
            this.$cookies.set('refresh_token', refresh_token)
            this.$router.push({ name: this.$route.name })
            this.closeSignUp()
            this.$toast(this.$t('register.success.logIn'))
          }
        } catch (err) {
          console.log('err', err)
          if (err) {
            if (err?.response?.status == 401) {
              this.$toast(this.$t('register.phoneNumberOrPassValid'))
            } else {
              this.$toast(this.$t('register.error'))
            }
          }
        }
      } else {
        this.isPhoneNumber = true
      }
    },
    openRegister() {
      this.clear()
      this.$emit('openRegisterPopUp')
    },
    closeSignUp() {
      this.clear()
      this.$emit('closeSignUpPopUp')
    },
    clear() {
      this.signUp.phone_number = '+9936'
      this.signUp.password = ''
      this.isPhoneNumber = false
      this.$v.$reset()
    },
  },
}
</script>
