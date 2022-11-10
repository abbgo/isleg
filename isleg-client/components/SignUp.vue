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
import { productAdd } from '@/api/user.api'

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
        try {
          let response = await this.$auth.loginWith('userLogin', {
            data: {
              phone_number: this.signUp.phone_number,
              password: this.signUp.password,
            },
          })
          console.log(response)
          console.log(this.$auth.loggedIn)
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
            console.log('this.$route.name', this.$route.name)
            this.$router.push({ name: this.$route.name })
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
      let products = []
      let array = []
      let wishlists = []
      let newWishlists = []
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
        wishlists = cart.cart.filter((product) => product.is_favorite === true)
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
          if (res.products.length > 0) {
            array = res.products.filter(
              (item) =>
                (item.quantity = item.quantity_of_product
                  ? item.quantity_of_product
                  : 0)
            )
            for (let i = 0; i < array.length; i++) {
              array[i]['is_favorite'] = false
            }
          }
          console.log('array', array)
          console.log('wishlists>>', wishlists)
          if (wishlists.length > 0) {
            for (let i = 0; i < array.length; i++) {
              for (let j = 0; j < wishlists.length; j++) {
                if (array[i].id == wishlists[j].id) {
                  console.log('(array[i] favorite', array[i])
                  array[i]['is_favorite'] = true
                }
              }
            }
          }
          console.log(array)
          this.productsWhenUserSignUp = array
        }
      } catch (error) {
        console.log(error)
      }
    },
    async postWishlists() {
      let wishlists = []
      let array = []
      let withoutWishlists = []
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (cart?.cart) {
        wishlists = cart.cart
          .filter((product) => product.is_favorite === true)
          .map((item) => item.id)
      }
      console.log('wishlists', wishlists)
      withoutWishlists =
        this.productsWhenUserSignUp.filter(
          (product) => product.is_favorite === true && product.quantity > 0
        ) || []
      console.log('withoutWishlists', withoutWishlists)
      try {
        const res = await this.$axios.$post(
          `/${this.$i18n.locale}/like?status=${true}`,
          { product_ids: wishlists },
          {
            headers: {
              Authorization: cart.auth.accessToken,
            },
          }
        )
        console.log('postWishlists', res)
        if (res.status) {
          if (res.products.length > 0) {
            for (let i = 0; i < res.products.length; i++) {
              res.products[i]['quantity'] = 0
              res.products[i]['is_favorite'] = true
              for (let j = 0; j < withoutWishlists.length; j++) {
                if (res.products[i].id !== withoutWishlists[j].id) {
                  console.log('withoutWishlistsAfter', res.products[i])
                  array.push(res.products[i])
                } else {
                  array.push(res.products[i])
                }
              }
            }
          }
          console.log('array', array)
          if (cart && cart.cart) {
            // console.log('postWishlistsssssss', [...withoutWishlists, ...array])
            cart.cart = [...withoutWishlists, ...array]
            console.log(' cart.cart', cart.cart)
            localStorage.setItem('lorem', JSON.stringify(cart))
          } else {
            cart['cart'] = [...withoutWishlists, ...array]
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
