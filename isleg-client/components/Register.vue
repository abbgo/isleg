<template>
  <div :class="['pop-up', { active: isOpenRegister }]" @click="closeRegister">
    <div class="pop-up__body" @click.stop>
      <div class="pop-up__wrapper">
        <div class="pop-up__close">
          <img src="@/assets/img/close.svg" alt="" @click="closeRegister" />
        </div>
        <div class="pop-up_form">
          <div class="form__input">
            <h4>{{ name }}</h4>
            <div class="form__input-container">
              <img src="@/assets/img/account.svg" alt="" />
              <input
                type="text"
                :placeholder="name"
                v-model.trim="$v.register.name.$model"
              />
            </div>
            <span
              class="error"
              v-if="$v.register.name.$error && !$v.register.name.required"
            >
              {{ $t('register.nameIsRequired') }}
            </span>
            <span class="error" v-if="!$v.register.name.minLength">
              {{ $t('register.nameMustHavetletters') }}</span
            >
          </div>
          <div class="form__input">
            <h4>Email</h4>
            <div class="form__input-container">
              <img src="@/assets/img/lock.svg" alt="" />
              <input
                type="text"
                placeholder="Email"
                v-model.trim="$v.register.email.$model"
              />
            </div>
            <span
              class="error"
              v-if="$v.register.email.$error && !$v.register.email.required"
            >
              {{ $t('register.emailIsRequired') }}
            </span>
            <span class="error" v-if="inValidEmail">
              {{ $t('register.invalidEmail') }}
            </span>
          </div>
          <div class="form__input">
            <h4>{{ phone }}</h4>
            <div class="form__input-container">
              <img src="@/assets/img/tel.svg" alt="" />
              <input
                type="tel"
                :placeholder="phone"
                v-model="$v.register.phone_number.$model"
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
                v-model="$v.register.password.$model"
              />
              <img
                @click="showPass = !showPass"
                :src="showPass ? '/img/Hide.svg' : '/img/Show.svg'"
                alt=""
              />
            </div>
            <span
              class="error"
              v-if="
                $v.register.password.$error && !$v.register.password.required
              "
            >
              {{ $t('register.passwordIsRequired') }}
            </span>
            <span class="error" v-if="!$v.register.password.minLength">
              {{ $t('register.passwordMustHavetletters') }}
            </span>
          </div>
          <div class="form__input">
            <h4>{{ passwordVerification }}</h4>
            <div class="form__input-container">
              <img src="@/assets/img/lock.svg" alt="" />
              <input
                :type="showConfirmPass ? 'password' : 'text'"
                :placeholder="passwordVerification"
                v-model="$v.register.repeatPassword.$model"
              />
              <img
                @click="showConfirmPass = !showConfirmPass"
                :src="showConfirmPass ? '/img/Hide.svg' : '/img/Show.svg'"
                alt=""
              />
            </div>
            <span
              class="error"
              v-if="
                $v.register.repeatPassword.$error &&
                !$v.register.repeatPassword.required
              "
            >
              {{ $t('register.confirmPasswordIsRequired') }}
            </span>
            <span
              class="error"
              v-if="
                confirm &&
                $v.register.password.$model !== '' &&
                $v.register.repeatPassword.$model !== ''
              "
            >
              {{ $t('register.confirmPasswordMustHavetletters') }}
            </span>
          </div>
          <div class="confirm__chekbox">
            <input
              class="confirm__chekbox-input"
              id="confirm"
              type="checkbox"
              v-model="$v.register.checked.$model"
              @change="termsOfService($event, $v.register.checked.$model)"
            />
            <label for="confirm" class="confirm__chekbox-label">{{
              verifySecure
            }}</label>
          </div>
          <span class="error" v-if="isTearmsOfServices">
            {{ $t('register.checkedIsRequired') }}
          </span>
          <div class="pop-up__btns">
            <button :disabled="disabled" class="right__btn" @click="signUp">
              {{ userSignUp }}
            </button>
            <button type="button" @click="openSignUp" class="left_btn">
              {{ signIn }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { required, sameAs, minLength } from 'vuelidate/lib/validators'
import { productAdd } from '@/api/user.api'
export default {
  props: {
    isOpenRegister: {
      type: Boolean,
      default: () => false,
    },
    name: {
      type: String,
      default: () => '',
    },
    passwordVerification: {
      type: String,
      default: () => '',
    },
    verifySecure: {
      type: String,
      default: () => '',
    },
    phone: {
      type: String,
      default: () => '',
    },
    password: {
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
      show: true,
      showPass: true,
      inValidEmail: false,
      showConfirmPass: true,
      confirm: false,
      isPhoneNumber: false,
      isTearmsOfServices: false,
      register: {
        name: '',
        phone_number: '+9936',
        email: '',
        password: '',
        repeatPassword: '',
        checked: false,
      },
    }
  },
  validations: {
    register: {
      name: {
        required,
        minLength: minLength(2),
      },
      phone_number: {
        required,
      },
      email: {
        required,
      },
      password: {
        required,
        minLength: minLength(5),
      },
      repeatPassword: {
        required,
        sameAsPassword: sameAs('password'),
      },
      checked: {
        required,
      },
    },
  },
  watch: {
    '$v.register.repeatPassword.$model': function (value) {
      if (value === this.$v.register.password.$model) {
        this.confirm = false
      } else {
        this.confirm = true
      }
    },
    '$v.register.password.$model': function (newVal) {
      if (
        this.$v.register.repeatPassword.sameAsPassword == false &&
        newVal !== this.$v.register.repeatPassword.$model
      ) {
        this.confirm = true
      } else {
        this.confirm = false
      }
    },
    '$v.register.email.$model': function (val) {
      if (val === '') {
        if (this.inValidEmail) {
          this.inValidEmail = false
        }
      }
      if (/^[a-z0-9._-]{2,}@[a-z0-9]{2,}\.[a-z]{2,}$/i.test(val)) {
        if (this.inValidEmail) {
          this.inValidEmail = false
        }
      }
    },
  },
  computed: {
    checkValidate() {
      if (
        /^[a-z0-9._-]{2,}@[a-z0-9]{2,}\.[a-z]{2,}$/i.test(
          this.$v.register.email.$model
        )
      ) {
        return true
      } else {
        return false
      }
    },
  },
  methods: {
    enforcePhoneFormat() {
      this.isPhoneNumber = false
      let x = this.register.phone_number
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
      if (!x[2]) {
        this.register.phone_number = '+9936'
      } else {
        this.register.phone_number =
          '+9936' +
          (x[3] ? x[3] : '') +
          (x[4] ? x[4] : '') +
          (x[5] ? x[5] : '') +
          (x[6] ? x[6] : '')
      }
    },
    termsOfService(e, val) {
      val = !val
      this.isTearmsOfServices = false
      if (!this.register.checked) {
        this.isTearmsOfServices = true
      } else {
        this.isTearmsOfServices = false
      }
    },
    signUp: async function () {
      this.$v.$touch()
      if (!this.register.checked) {
        this.isTearmsOfServices = true
      } else {
        this.isTearmsOfServices = false
      }
      if (this.register.phone_number.length < 12) {
        this.isPhoneNumber = true
      } else {
        this.isPhoneNumber = false
      }
      if (this.$v.$invalid) {
        if (
          this.$v.register.email.$model !== '' &&
          this.checkValidate === false
        ) {
          this.inValidEmail = true
        }

        if (
          this.$v.register.password.$model !==
          this.$v.register.repeatPassword.$model
        ) {
          if (this.confirm === false) {
            this.confirm = true
          }
        }
      } else {
        if (
          this.checkValidate == true &&
          this.register.phone_number.length >= 12 &&
          this.$v.register.password.$model ==
            this.$v.register.repeatPassword.$model
        ) {
          this.disabled = true
          const formData = new FormData()
          formData.append('full_name', this.register.name)
          formData.append('phone_number', this.register.phone_number)
          formData.append('password', this.register.password)
          formData.append('email', this.register.email)
          try {
            let response = await this.$auth.loginWith('userRegister', {
              data: {
                full_name: this.register.name,
                phone_number: this.register.phone_number,
                password: this.register.password,
                email: this.register.email,
              },
            })
            console.log(this.$auth.loggedIn, this.$store.$auth)
            console.log(response)
            if (response.status === 200) {
              const { access_token, refresh_token } = response.data
              const cart = await JSON.parse(localStorage.getItem('lorem'))
              if (cart) {
                cart.auth = {
                  accessToken: access_token,
                  fullName: this.register.name,
                  phoneNumber: this.register.phone_number,
                  email: this.register.email,
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
              this.closeRegister()
              await this.postCarts()
              await this.postWishlists()
              console.log('this.$route.name', this.$route.name)
            }
          } catch (err) {
            console.log(err)
            if (err.response.status == 400) {
              this.$toast(this.$t('register.customerExists'))
            } else {
              this.$toast(this.$t('register.error'))
            }
          } finally {
            this.disabled = false
          }
        } else {
          if (this.checkValidate == false) {
            this.inValidEmail = true
          }
          if (this.register.phone_number.length < 12) {
            this.isPhoneNumber = true
          }
        }
      }
    },
    openSignUp() {
      this.clear()
      this.$emit('openSignUpPopUp')
    },
    closeRegister() {
      this.clear()
      this.$emit('closeRegisterPopUp')
    },
    clear() {
      this.register.name = ''
      this.register.phone_number = '+9936'
      this.register.email = ''
      this.register.password = ''
      this.register.repeatPassword = ''
      this.register.checked = false
      this.show = true
      this.showPass = true
      this.inValidEmail = false
      this.showConfirmPass = true
      this.confirm = false
      this.isPhoneNumber = false
      this.isTearmsOfServices = false
      this.$v.$reset()
    },
    async postCarts() {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      let products = []
      if (cart) {
        for (let i = 0; i < cart.cart?.length; i++) {
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
        const res = (
          await productAdd({
            url: `${this.$i18n.locale}/add-cart`,
            data: products,
            accessToken: `Bearer ${cart?.auth?.accessToken}`,
          })
        ).data
        console.log('productAdd', res)
      } catch (error) {
        console.log(error)
      }
    },
    async postWishlists() {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      let wishlists = []
      if (cart) {
        wishlists = cart.cart
          ?.filter((product) => product.is_favorite === true)
          .map((item) => item.id)
      }
      try {
        const { status } = await this.$axios.$post(
          `/${this.$i18n.locale}/like?status=${true}`,
          { product_ids: wishlists },
          {
            headers: {
              Authorization: cart.auth.accessToken,
            },
          }
        )
        console.log(status)
        // if (status) {
        // }
      } catch (e) {
        console.log(e)
      }
    },
  },
}
</script>
