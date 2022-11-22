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
import { productAdd, userLogin } from '@/api/user.api'

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
      productsWhenUserSignUp: [],
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
        const userData = {
          phone_number: this.signUp.phone_number,
          password: this.signUp.password,
        }
        try {
          let response = await userLogin({
            url: 'auth/login',
            data: userData,
          })
          console.log(response)
          if (response.status === 200) {
            const { access_token, refresh_token } = response.data
            console.log(access_token, refresh_token)
            const cart = await JSON.parse(localStorage.getItem('lorem'))
            if (cart) {
              cart.auth = {
                accessToken: access_token,
                refreshToken: refresh_token,
              }
              localStorage.setItem('lorem', JSON.stringify(cart))
            } else {
              localStorage.setItem(
                'lorem',
                JSON.stringify({
                  auth: {
                    accessToken: access_token,
                    refreshToken: refresh_token,
                  },
                })
              )
            }
            this.closeSignUp()
            this.$toast(this.$t('register.success.logIn'))
            await this.postCarts()
            await this.postWishlists()
            this.$store.commit('ui/SET_USER_LOGGINED', true)
          }
        } catch (err) {
          console.log('err', err)
          if (err) {
            if (err?.response?.status == 401) {
              this.$toast(this.$t('register.phoneNumberOrPassValid'))
            } else if (err?.response?.status === 400) {
              this.$toast(this.$t('register.userDoesNotExist'))
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
      let products = []
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (cart?.cart) {
        for (let i = 0; i < cart.cart.length; i++) {
          if (cart.cart[i].quantity > 0) {
            products.push({
              product_id: cart.cart[i].id,
              quantity_of_product: cart.cart[i].quantity,
            })
          }
        }
      }
      console.log('wishlistsPPPPPPPP', products)
      try {
        const res = (
          await productAdd({
            url: `${this.$i18n.locale}/add-cart`,
            data: products,
            accessToken: `Bearer ${cart?.auth?.accessToken}`,
          })
        ).data
        console.log('productAdd', res)
        if (res.status) {
          if (res.products) {
            if (res.products.length > 0) {
              res.products = res.products.filter(
                (item) =>
                  (item.quantity = item.quantity_of_product
                    ? item.quantity_of_product
                    : 0)
              )
              for (let i = 0; i < res.products.length; i++) {
                res.products[i]['is_favorite'] = false
              }
              this.productsWhenUserSignUp = res.products
            }
          } else {
            this.productsWhenUserSignUp = []
          }
          console.log('array', res.products)
          //  if (wishlists.length > 0) {
          //    for (let i = 0; i < res.products.length; i++) {
          //      for (let j = 0; j < wishlists.length; j++) {
          //        if (res.products[i].id == wishlists[j].id) {
          //          console.log('(array[i] favorite', res.products[i])
          //          res.products[i].is_favorite = true
          //        }
          //      }
          //    }
          //  }
        }
      } catch (error) {
        console.log(error)
      }
    },
    async postWishlists() {
      let wishlists = []
      let array = []
      // let withoutWishlists = []
      // let withWishlists = []
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (cart?.cart) {
        wishlists = cart.cart
          .filter((product) => product.is_favorite === true)
          .map((item) => item.id)
      }
      console.log('wishlists', wishlists)
      // withoutWishlists =
      //   this.productsWhenUserSignUp.filter(
      //     (product) => product.is_favorite === false
      //   ) || []
      // withWishlists =
      //   this.productsWhenUserSignUp.filter(
      //     (product) => product.is_favorite === true
      //   ) || []
      // console.log('withoutWishlists', withoutWishlists)
      // console.log('withWishlists', withWishlists)
      try {
        const res = await this.$axios.$post(
          `/${this.$i18n.locale}/like?status=${true}`,
          { product_ids: wishlists },
          {
            headers: {
              Authorization: `Bearer ${cart?.auth?.accessToken}`,
            },
          }
        )
        console.log('postWishlists', res)
        if (res.status) {
          if (res.products) {
            if (res.products.length > 0) {
              for (let i = 0; i < res.products.length; i++) {
                res.products[i]['quantity'] = 0
                res.products[i]['is_favorite'] = true
              }
              if (this.productsWhenUserSignUp.length > 0) {
                for (let i = 0; i < res.products.length; i++) {
                  for (let j = 0; j < this.productsWhenUserSignUp.length; j++) {
                    if (
                      res.products[i].id == this.productsWhenUserSignUp[j].id
                    ) {
                      console.log(
                        'this.productsWhenUserSignUp[j]',
                        this.productsWhenUserSignUp[j]['is_favorite']
                      )
                      this.productsWhenUserSignUp[j]['is_favorite'] = true
                    }
                  }
                }
                let result = res.products.filter(
                  (o1) =>
                    !this.productsWhenUserSignUp.some((o2) => o1.id === o2.id)
                )
                console.log('result', result)
                array = result
              } else {
                array = res.products
              }
            } else {
              array = res.products
            }
          }
          console.log('array', array)
          if (cart && cart.cart) {
            // console.log('postWishlistsssssss', [...withoutWishlists, ...array])
            cart.cart = [...this.productsWhenUserSignUp, ...array]
            console.log(' cart.cart', cart.cart)
            localStorage.setItem('lorem', JSON.stringify(cart))
          } else {
            cart['cart'] = [...this.productsWhenUserSignUp, ...array]
            localStorage.setItem('lorem', JSON.stringify(cart))
          }
        }
      } catch (e) {
        console.log(e)
      }
    },
  },
}
</script>
