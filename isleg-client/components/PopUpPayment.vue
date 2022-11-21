<template>
  <div :class="['pop-up', { active: isPaymentComputed }]">
    <div class="pop-up__product-body" style="width: 900px">
      <div class="pop-up__wrapper">
        <div class="pop-up__close" @click="close">
          <svg
            width="60"
            height="60"
            viewBox="0 0 90 90"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              opacity="0.8"
              d="M41.1636 73.4122C54.9602 73.4122 66.1446 62.2278 66.1446 48.4312C66.1446 34.6346 54.9602 23.4502 41.1636 23.4502C27.367 23.4502 16.1826 34.6346 16.1826 48.4312C16.1826 62.2278 27.367 73.4122 41.1636 73.4122Z"
              fill="#1B3254"
            />
            <path
              d="M29.4334 58.6015L31.0243 60.1925L41.3658 49.851L51.4421 59.9273L53.033 58.3363L42.9568 48.26L53.2982 37.9186L51.7072 36.3276L41.3658 46.669L31.2895 36.5928L29.6985 38.1838L39.7748 48.26L29.4334 58.6015Z"
              fill="white"
            />
          </svg>
        </div>
        <div class="payment__container">
          <div class="payment__text">
            <p>
              {{
                paymentDatas && paymentDatas.text && paymentDatas.text.content
              }}
            </p>
            <br />
          </div>
          <div class="payment__settings">
            <div class="settings__chekbox-container">
              <div class="payment__chekbox">
                <div class="payment__chekbox-title">
                  <h3>
                    {{
                      paymentDatas &&
                      paymentDatas.text &&
                      paymentDatas.text.type_of_payment
                    }}
                  </h3>
                </div>
                <div
                  class="payment__chekbox-input"
                  v-for="item in paymentDatas.types"
                  :key="item.id"
                >
                  <input
                    :checked="item.checked"
                    class="top__input"
                    name="top"
                    :id="item.name"
                    type="radio"
                    @change="paymentChecked(item)"
                  />
                  <label :for="item.name">{{ item.name }}</label>
                </div>
                <span class="error" v-if="isPaymentForm">
                  {{ $t('payment.paymentForm') }}
                </span>
              </div>
              <div class="payment__chekbox">
                <div class="payment__chekbox-title">
                  <h3>
                    {{
                      paymentDatas &&
                      paymentDatas.text &&
                      paymentDatas.text.choose_a_delivery_time
                    }}
                  </h3>
                </div>
                <div class="payment__chekbox-subtitle">
                  <h4>{{ paymentDatas.orderTimes.title }}</h4>
                </div>
                <div
                  class="payment__chekbox-input"
                  v-for="item in paymentDatas.orderTimes.times"
                  :key="item.id"
                >
                  <h4>{{ item.translation_date }}</h4>
                  <template v-for="time in item.time"
                    ><input
                      :checked="time.checked"
                      class="bottom__input"
                      name="bottom"
                      :id="time.time"
                      type="radio"
                      @change="theDeliveryTimeChecked(time)"
                    />
                    <label :for="time.time">{{ time.time }}</label></template
                  >
                </div>
                <span class="error" v-if="isTheDeliveryTime">
                  {{ $t('payment.theDeliveryTime') }}
                </span>
              </div>
            </div>
            <div class="payment__form">
              <div class="payment__form-box">
                <label for="">{{ name }} <span>*</span></label>
                <input
                  type="text"
                  v-model.trim="$v.payment.fullName.$model"
                  :placeholder="name"
                />
                <span
                  class="error"
                  v-if="
                    $v.payment.fullName.$error && !$v.payment.fullName.required
                  "
                >
                  {{ $t('register.nameIsRequired') }}
                </span>
                <span class="error" v-if="!$v.payment.fullName.minLength">
                  {{ $t('register.nameMustHavetletters') }}</span
                >
              </div>
              <div class="payment__form-box">
                <label for="">{{ phone }} <span>*</span></label>
                <input
                  type="tel"
                  v-model="$v.payment.phone_number.$model"
                  @input="enforcePhoneFormat"
                />
                <span class="error" v-if="isPhoneNumber">
                  {{ $t('register.phoneNumberIsRequired') }}
                </span>
              </div>
              <div class="payment__form-box">
                <label for=""
                  >{{
                    paymentDatas &&
                    paymentDatas.text &&
                    paymentDatas.text.your_address
                  }}
                  <span>*</span></label
                >
                <input
                  type="text"
                  v-model.trim="$v.payment.address.$model"
                  :placeholder="
                    paymentDatas &&
                    paymentDatas.text &&
                    paymentDatas.text.your_address
                  "
                />
                <span
                  class="error"
                  v-if="
                    $v.payment.address.$error && !$v.payment.address.required
                  "
                >
                  {{ $t('payment.address') }}
                </span>
              </div>
              <div class="payment__form-box">
                <label for="">{{
                  paymentDatas && paymentDatas.text && paymentDatas.text.mark
                }}</label>
                <input
                  type="text"
                  v-model="payment.note"
                  :placeholder="
                    paymentDatas && paymentDatas.text && paymentDatas.text.mark
                  "
                />
              </div>
              <div class="payment__form-btn">
                <button
                  class="disable__btn"
                  v-if="isAuth"
                  @click="$emit('paymentRegister')"
                >
                  {{ signIn }}
                </button>
                <button v-else></button>

                <button class="order__btn" @click="order">
                  {{
                    paymentDatas &&
                    paymentDatas.text &&
                    paymentDatas.text.to_order
                  }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators'
import { postPaymentDatas } from '@/api/payment.api'
import { getMyProfile } from '@/api/myProfile.api'
import { getRefreshToken } from '@/api/user.api'

export default {
  props: {
    isPayment: {
      type: Boolean,
      default: false,
    },
    name: {
      type: String,
      default: () => '',
    },
    phone: {
      type: String,
      default: () => '',
    },
    signIn: {
      type: String,
      default: () => '',
    },
    totalPrice: {
      type: [String, Number],
      default: () => '',
    },
    paymentDatas: {
      type: Object,
      default: () => {},
    },
    userInformation: {
      type: Object,
      default: () => {},
    },
  },
  data() {
    return {
      isPhoneNumber: false,
      isPaymentForm: false,
      isTheDeliveryTime: false,
      selectedPaymentType: null,
      selectedPaymentTime: null,
      payment: {
        fullName: '',
        phone_number: '+9936',
        address: '',
        note: '',
      },
    }
  },
  validations: {
    payment: {
      fullName: {
        required,
        minLength: minLength(4),
      },
      phone_number: {
        required,
      },
      address: {
        required,
      },
    },
  },
  computed: {
    isPaymentComputed() {
      if (this.isPayment) {
        const cart = JSON.parse(localStorage.getItem('lorem'))
        if (cart && cart?.auth?.accessToken) {
          this.fetchUserInformation(cart)
        }
        return true
      } else {
        return false
      }
    },
    isAuth() {
      const cart = JSON.parse(localStorage.getItem('lorem'))
      console.log('cart?.auth?.accessToken1', cart?.auth?.accessToken)
      if (cart?.auth?.accessToken) {
        console.log('cart?.auth?.accessToken2', cart?.auth?.accessToken)
        return false
      } else {
        return true
      }
    },
  },
  methods: {
    async fetchUserInformation(cart) {
      try {
        const { customer_informations, status } = (
          await getMyProfile({
            url: `${this.$i18n.locale}/my-information`,
            accessToken: `Bearer ${cart.auth.accessToken}`,
          })
        ).data
        console.log(
          'deeeeeeeeeeeeeeeeeeeeeeeeeede',
          customer_informations,
          status
        )
        if (status) {
          this.payment.fullName = customer_informations.full_name
          this.payment.phone_number = customer_informations.phone_number
          this.payment.address = customer_informations.addresses[0]?.address
          console.log(' this.userInformation', this.userInformation)
        }
      } catch (error) {
        console.log('err', error)
        if (error?.response?.status === 403) {
          try {
            const { access_token, refresh_token, status } = (
              await getRefreshToken({
                url: `auth/refresh`,
                refreshToken: `Bearer ${cart.auth.refreshToken}`,
              })
            ).data
            console.log('new', access_token, refresh_token, status)
            if (status) {
              const lorem = await JSON.parse(localStorage.getItem('lorem'))
              if (lorem) {
                lorem.auth = {
                  accessToken: access_token,
                  refreshToken: refresh_token,
                }
                localStorage.setItem('lorem', JSON.stringify(lorem))
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
              try {
                const { customer_informations, status } = (
                  await getMyProfile({
                    url: `${this.$i18n.locale}/my-information`,
                    accessToken: `Bearer ${lorem.auth.accessToken}`,
                  })
                ).data
                console.log(
                  'deeeeeeeeeeeeeeeeeeeeeeeeeede',
                  customer_informations,
                  status
                )
                if (status) {
                  this.payment.fullName = customer_informations.full_name
                  this.payment.phone_number = customer_informations.phone_number
                  this.payment.address =
                    customer_informations.addresses[0]?.address
                  console.log(' this.userInformation', this.userInformation)
                }
              } catch (error) {
                console.log('getMyProfile2', error)
              }
            }
          } catch (error) {
            console.log('ref', error.response.status)
            if (error.response.status === 403) {
              cart.auth.accessToken = null
              cart.auth.refreshToken = null
              localStorage.setItem('lorem', JSON.stringify(cart))
              this.$router.push(this.localeLocation('/'))
            }
          }
        }
      }
    },
    enforcePhoneFormat() {
      this.isPhoneNumber = false
      let x = this.payment.phone_number
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
      if (!x[2]) {
        this.payment.phone_number = '+9936'
      } else {
        this.payment.phone_number =
          '+9936' +
          (x[3] ? x[3] : '') +
          (x[4] ? x[4] : '') +
          (x[5] ? x[5] : '') +
          (x[6] ? x[6] : '')
      }
    },
    paymentChecked(payload) {
      const findItem = this.paymentDatas.types.find(
        (item) => item.checked == true
      )
      if (findItem) {
        findItem.checked = false
      }
      payload.checked = true
      this.selectedPaymentType = payload
      console.log(this.selectedPaymentType)
      this.isPaymentForm = false
    },
    theDeliveryTimeChecked(payload) {
      const findItem = this.paymentDatas.orderTimes.times.find(
        (item) => item.checked == true
      )
      if (findItem) {
        findItem.checked = false
      }
      payload.checked = true
      this.selectedPaymentTime = payload
      console.log(this.selectedPaymentTime)
      this.isTheDeliveryTime = false
    },
    order: async function () {
      this.$v.$touch()
      const paymentForm = this.paymentDatas.types.filter(
        (pay) => pay.checked == true
      )
      const theDeliveryTime = this.paymentDatas.orderTimes.times.filter(
        (pay) => pay.checked == true
      )
      if (paymentForm.length == 0) {
        this.isPaymentForm = true
      }
      if (theDeliveryTime.length == 0) {
        this.isTheDeliveryTime = true
      }
      if (this.$v.$invalid) {
        if (this.payment.phone_number.length < 12) {
          this.isPhoneNumber = true
        } else {
          this.isPhoneNumber = false
        }
      } else {
        if (
          this.payment.phone_number.length >= 12 &&
          paymentForm.length > 0 &&
          theDeliveryTime.length > 0
        ) {
          const cart = await JSON.parse(localStorage.getItem('lorem'))
          let products = []
          if (cart) {
            for (let i = 0; i < cart?.cart?.length; i++) {
              if (cart.cart[i]?.quantity > 0) {
                products.push({
                  product_id: cart.cart[i].id,
                  quantity_of_product: cart.cart[i].quantity,
                })
              }
            }
          }
          console.log(products, {
            full_name: this.payment.fullName,
            phone_number: this.payment.phone_number,
            address: this.payment.address,
            customer_mark: this.payment.note,
            order_time: this.selectedPaymentTime.time,
            payment_type: this.selectedPaymentType.name,
            total_price: this.totalPrice,
            products: products,
          })
          try {
            const { status } = (
              await postPaymentDatas({
                url: `${this.$i18n.locale}/to-order`,
                data: {
                  full_name: this.payment.fullName,
                  phone_number: this.payment.phone_number,
                  address: this.payment.address,
                  customer_mark: this.payment.note,
                  order_time: this.selectedPaymentTime.time,
                  payment_type: this.selectedPaymentType.name,
                  total_price: Number(this.totalPrice),
                  products: products,
                },
              })
            ).data
            console.log(status)
            if (status) {
              cart.cart = cart.cart
                .filter((product) => product.is_favorite === true)
                .filter((item) => {
                  item.quantity = 0
                  return item
                })
              localStorage.setItem('lorem', JSON.stringify(cart))
              this.$emit('paymentSuccesfullySended')
            }
          } catch (error) {
            console.log('payment', error)
            this.$toast(this.$t('register.error'))
          }
        } else {
          if (this.payment.phone_number.length < 12) {
            this.isPhoneNumber = true
          } else {
            this.isPhoneNumber = false
          }
          if (paymentForm.length == 0) {
            this.isPaymentForm = true
          }
          if (theDeliveryTime.length == 0) {
            this.isTheDeliveryTime = true
          }
        }
      }
    },
    close() {
      this.$emit('close')
      this.isPhoneNumber = false
      this.isPaymentForm = false
      this.isTheDeliveryTime = false
      this.$v.$reset()
    },
  },
}
</script>
