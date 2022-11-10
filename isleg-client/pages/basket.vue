<template>
  <section
    class="basket __container"
    :class="{ active: mobileBasket }"
    v-if="products.length"
  >
    <BasketProducts :products="products" @popUpSureOpen="openPopUpSure" />
    <BasketPrices
      :productCount="productCount"
      :totalPrice="totalPrice"
      :productsChek="productsChek"
      @openPayment="openPayment"
      @cartTheEmpty="cartEmpty"
    />
    <LazyPopUpPayment
      :isPayment="isPayment"
      :paymentDatas="paymentDatas"
      :phone="phone"
      :name="name"
      :signIn="signIn"
      :totalPrice="totalPrice"
      :userInformation="userInformation"
      @close="closePaymentPopup"
      @paymentSuccesfullySended="paymentSuccesfullySended"
      @paymentRegister="paymentRegister"
    />
    <LazyPopUpSure
      :isSure="isSure"
      @close="closePopUpSure"
      @confirm="confirm"
    />
    <div class="mobile__basket">
      <div class="mobile__basket-content">
        <div class="mobile__content-information">
          <div class="information__price">
            <div class="information__price--row">
              <span>Bahasy</span>
              <span>198 ТМТ</span>
            </div>
            <div class="information__price--row">
              <span>Eltip berme</span>
              <span>15 ТМТ</span>
            </div>
          </div>
          <div class="information__total">
            <div class="information__total--row">
              <span>Jemi:</span>
              <span>198 ТМТ</span>
            </div>
          </div>
        </div>
        <div class="mobile__content-button"></div>
      </div>
    </div>
    <div class="mobile__inside-button">
      <button @click="mobileBasket = !mobileBasket">
        <span></span>
        <span>Sargyt et</span>
        <img src="@/assets/img/mobile__arrow.svg" alt="" />
      </button>
    </div>
  </section>
</template>

<script>
import BasketProducts from '@/components/BasketProducts.vue'
import BasketPrices from '@/components/BasketPrices.vue'
import { getMyProfile } from '@/api/myProfile.api'

import {
  productAdd,
  getRefreshToken,
  deleteAllProductsFromBasket,
} from '@/api/user.api'

import {
  getPaymentText,
  getPaymentTypes,
  getPaymentTime,
} from '@/api/payment.api'
import { mapGetters } from 'vuex'
export default {
  name: 'BasketPage',
  components: { BasketProducts, BasketPrices },
  data() {
    return {
      mobileBasket: true,
      isPayment: false,
      isSure: false,
      productRemoveItem: null,
      isCartEmpty: false,
      products: [],
      userInformation: {
        fullName: null,
        address: null,
        phoneNumber: null,
      },
      productsChek: {
        discount: null,
        delivery: null,
      },
      paymentDatas: {
        text: null,
        types: [],
        orderTimes: {
          times: [],
        },
      },
    }
  },
  async mounted() {
    await this.getBasketProducts()
  },
  computed: {
    ...mapGetters('products', ['productCount']),
    ...mapGetters('ui', ['phone', 'name', 'signIn']),
    totalPrice() {
      return this.products?.reduce((total, num) => {
        let sum = Number(Number(num.price) * Number(num.quantity))
        let totalSum = Number(total) + sum
        if (totalSum > 150) {
          this.productsChek.delivery = Number(0).toFixed(2)
        } else if (totalSum <= 150) {
          this.productsChek.delivery = Number(10).toFixed(2)
        } else {
          this.productsChek.delivery = Number(15).toFixed(2)
        }
        console.log('totalSum', totalSum)
        return Number(totalSum).toFixed(2)
      }, 0)
    },
  },
  methods: {
    async getBasketProducts() {
      let cart = JSON.parse(localStorage.getItem('lorem'))
      if (cart?.cart) {
        let totalCount = cart?.cart.reduce((total, num) => {
          return total + num.quantity
        }, 0)
        this.products = cart.cart.filter((product) => product.quantity > 0)
        this.$store.commit(
          'products/SET_PRODUCT_COUNT',
          totalCount == 0 ? null : totalCount
        )
      } else {
        this.products = []
      }
    },
    async openPayment() {
      await Promise.all([
        this.fetchPaymnetText(),
        this.fetchPaymnetTypes(),
        this.fetchPaymnetTime(),
      ])
      this.isPayment = true
      document.body.classList.add('_lock')
    },
    async fetchPaymnetText() {
      try {
        const { translation_order_page, status } = (
          await getPaymentText({
            url: `${this.$i18n.locale}/translation-order-page`,
          })
        ).data
        console.log(translation_order_page)
        if (status) {
          this.paymentDatas.text = translation_order_page
        }
      } catch (error) {
        console.log('payment', error)
      }
    },
    async fetchPaymnetTypes() {
      try {
        const { payment_types, status } = (
          await getPaymentTypes({
            url: `${this.$i18n.locale}/payment-types`,
          })
        ).data
        console.log(payment_types)
        if (status) {
          for (let i = 0; i < payment_types.length; i++) {
            this.paymentDatas.types.push({
              id: i + Math.random() * 1,
              name: payment_types[i],
              checked: false,
            })
          }
          console.log(this.paymentDatas.types)
        }
      } catch (error) {
        console.log('payment', error)
      }
    },
    async fetchPaymnetTime() {
      try {
        const { order_times, status } = (
          await getPaymentTime({
            url: `${this.$i18n.locale}/order-time`,
          })
        ).data
        console.log('order_times', order_times)
        if (status) {
          for (let i = 0; i < order_times.length; i++) {
            this.paymentDatas.orderTimes.times.push({
              id: i + Math.random() * 2,
              translation_date: order_times[i].translation_date,
              time: order_times[i].times[i].time,
              checked: false,
            })
          }
          console.log(this.paymentDatas.orderTimes)
        }
      } catch (error) {
        console.log('payment', error)
      }
    },
    closePaymentPopup() {
      this.isPayment = false
      document.body.classList.remove('_lock')
      this.paymentDatas.text = null
      this.paymentDatas.types = []
      this.paymentDatas.orderTimes.title = null
      this.paymentDatas.orderTimes.times = []
    },
    openPopUpSure(data) {
      this.isSure = true
      this.productRemoveItem = data
      document.body.classList.add('_lock')
    },
    cartEmpty() {
      this.isCartEmpty = true
      this.isSure = true
      document.body.classList.add('_lock')
    },
    closePopUpSure() {
      this.productRemoveItem = null
      this.isSure = false
      document.body.classList.remove('_lock')
    },
    async confirm() {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (this.isCartEmpty) {
        this.products = []
        cart.cart = []
        localStorage.setItem('lorem', JSON.stringify(cart))
        this.$store.commit('products/SET_PRODUCT_COUNT_WHEN_PAYMENT')
        if (cart && cart?.auth?.accessToken && this.$auth.loggedIn) {
          try {
            const res = (
              await deleteAllProductsFromBasket({
                url: `${this.$i18n.locale}/remove-cart`,
                accessToken: `Bearer ${cart?.auth?.accessToken}`,
              })
            ).data
            console.log(res)
          } catch (error) {
            console.log(error)
          } finally {
            this.isCartEmpty = false
          }
        }
      } else {
        cart.cart = cart?.cart.filter(
          (product) => product.id != this.productRemoveItem.id
        )
        this.products = cart.cart
        localStorage.setItem('lorem', JSON.stringify(cart))
        this.$store.commit('products/SET_BASKET_PRODUCT_COUNT', 1)
        if (cart && cart?.auth?.accessToken && this.$auth.loggedIn) {
          try {
            const res = (
              await productAdd({
                url: `${this.$i18n.locale}/add-cart`,
                data: [
                  {
                    product_id: this.productRemoveItem.id,
                    quantity_of_product: 0,
                  },
                ],
                accessToken: `Bearer ${cart?.auth?.accessToken}`,
              })
            ).data
            console.log('productAdd', res)
          } catch (error) {
            console.log(error)
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
                    const response = (
                      await productAdd({
                        url: `${this.$i18n.locale}/add-cart`,
                        data: [
                          {
                            product_id: this.productRemoveItem.id,
                            quantity_of_product: 0,
                          },
                        ],
                        accessToken: `Bearer ${lorem?.auth?.accessToken}`,
                      })
                    ).data
                    console.log('productAdd1', response)
                  } catch (error) {
                    console.log('productAdd1', error)
                  }
                }
              } catch (error) {
                console.log('ref', error.response.status)
                if (error.response.status === 403) {
                  this.$auth.logout()
                  cart.auth.accessToken = null
                  cart.auth.refreshToken = null
                  localStorage.setItem('lorem', JSON.stringify(cart))
                  console.log(this.$route.name)
                  this.$router.push({ name: this.$route.name })
                }
              }
            }
          }
        }
      }
      this.closePopUpSure()
    },
    paymentRegister() {
      this.isPayment = false
      this.$store.commit('ui/SET_OPEN_ISOPENSIGNUP')
      this.paymentDatas.text = null
      this.paymentDatas.types = []
      this.paymentDatas.orderTimes.title = null
      this.paymentDatas.orderTimes.times = []
    },
    paymentSuccesfullySended() {
      this.products = []
      this.$store.commit('products/SET_PRODUCT_COUNT_WHEN_PAYMENT')
      this.isPayment = false
      document.body.classList.remove('_lock')
    },
  },
}
</script>
