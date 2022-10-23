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
            <button
              type="button"
              :disabled="disabled"
              class="right__btn"
              @click="logIn"
            >
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
      disabled: false,
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
        this.disabled = true
        try {
          let response = await this.$auth.loginWith('userLogin', {
            data: {
              phone_number: this.signUp.phone_number,
              password: this.signUp.password,
            },
          })
          console.log(response)
          console.log(this.$auth.loggedIn, this.$store.$auth)
          if (response.status === 200) {
            const { access_token, customer_id, refresh_token } = response.data
            console.log(access_token, customer_id, refresh_token)
            this.$cookies.set('access_token', access_token)
            this.$cookies.set('customer_id', customer_id)
            this.$cookies.set('refresh_token', refresh_token)
            await this.postCarts()
            await this.postFishlists()
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
        } finally {
          this.disabled = false
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
    async postCarts() {
      const customerId = await this.$cookies.get('customer_id')
      const accessToken = await this.$cookies.get('access_token')
      console.log('accessToken', accessToken)
      let products = []
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (cart) {
        for (let i = 0; i < cart.cart.length; i++) {
          if (cart.cart[i].quantity > 0) {
            products.push({
              product_id: cart.cart[i].id,
              quantity_of_product: cart.cart[i].quantity,
            })
          }
        }
      }
      console.log(products)
      try {
        const res = await this.$axios.$post(
          `/${this.$i18n.locale}/add-cart`,
          {
            customer_id: customerId,
            products: products,
          },
          {
            headers: {
              Authorization: accessToken,
            },
          }
        )
        console.log(res)
      } catch (e) {
        console.log(e)
      }
    },
    async postFishlists() {
      const formData = new FormData()
      const customerId = await this.$cookies.get('customer_id')
      const accessToken = await this.$cookies.get('access_token')
      console.log('accessToken', accessToken)
      let wishlists = []
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (cart) {
        wishlists = cart.cart
          .filter((product) => product.is_favorite === true)
          .map((item) => item.id)
      }
      formData.append('customer_id', customerId)
      for (let i = 0; i < wishlists.length; i++) {
        formData.append('product_ids', wishlists[i])
      }
      try {
        const res = await this.$axios.$post(
          `/${this.$i18n.locale}/like`,
          formData,
          {
            headers: {
              Authorization: accessToken,
            },
          }
        )
        console.log(res)
      } catch (e) {
        console.log(e)
      }
    },
  },
}
</script>
