<template>
  <section
    class="basket __container"
    :class="{ active: mobileBasket }"
    v-if="products.length"
  >
    <BasketProducts :products="products" @sure="openPopUpSure" />
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
import { getTranslationBasketPage } from '@/api/ui.api'
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
      translationBasketPage: null,
    }
  },
  watch: {
    isUserLoggined: async function (val) {
      if (val) {
        await this.fetchProductsFromDataBase()
      }
    },
  },
  computed: {
    ...mapGetters('products', ['productCount']),
    ...mapGetters('ui', ['phone', 'name', 'signIn', 'isUserLoggined']),
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
        return Number(totalSum).toFixed(2) > 0
          ? Number(totalSum).toFixed(2)
          : '0'
      }, 0)
    },
  },
  async fetch() {
    try {
      const { status, translation_basket_page } = (
        await getTranslationBasketPage({
          url: `${this.$i18n.locale}/translation-basket-page`,
        })
      ).data
      if (status) {
        this.translationBasketPage = translation_basket_page
      }
    } catch (error) {
      console.log('error', error)
    }
  },
  async mounted() {
    await this.fetchProductsFromDataBase()
    if (window.innerWidth <= 950) {
      if (document.body.classList.contains('_lock')) {
        document.body.classList.remove('_lock')
      }
    }
  },
  methods: {
    async fetchProductsFromDataBase() {
      let product_ids = []
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (cart?.cart) {
        product_ids = cart.cart
          .filter((product) => product.quantity > 0)
          .map((item) => item.id)
        try {
          const { products, status } = await this.$axios.$post(
            `/${this.$i18n.locale}/likes-or-orders-without-customer`,
            { product_ids: product_ids }
          )
          if (status) {
            let res = null
            if (products) {
              res = product_ids.filter(function (o1) {
                return !products.some(function (o2) {
                  return o1 === o2.id
                })
              })
              cart?.cart?.forEach((elem) => {
                products.forEach((product) => {
                  if (elem.id === product.id) {
                    if (elem.amount > product.amount) {
                      elem.amount = product.amount
                      elem.quantity = product.amount
                    }
                    if (elem.limit_amount > product.limit_amount) {
                      elem.limit_amount = product.limit_amount
                      elem.quantity = product.limit_amount
                    }
                  }
                })
              })
            }
            localStorage.setItem('lorem', JSON.stringify(cart))
            this.products =
              cart.cart.filter((product) => product.quantity > 0) || []
            let totalCount = cart?.cart.reduce((total, num) => {
              return total + num.quantity
            }, 0)
            this.$store.commit(
              'products/SET_PRODUCT_COUNT',
              totalCount == 0 ? null : totalCount
            )
          }
        } catch (err) {
          console.log(err)
        }
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
        if (status) {
          for (let i = 0; i < payment_types.length; i++) {
            this.paymentDatas.types.push({
              id: i + Math.random() * 1,
              name: payment_types[i],
              checked: false,
            })
          }
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
        if (status) {
          for (let i = 0; i < order_times.length; i++) {
            for (let j = 0; j < order_times[i].times.length; j++) {
              order_times[i].times[j]['checked'] = false
            }
            this.paymentDatas.orderTimes.times.push({
              id: i + Math.random() * 2,
              translation_date: order_times[i].translation_date,
              time: order_times[i].times,
            })
          }
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
    openPopUpSure(product) {
      this.isSure = true
      this.productRemoveItem = product
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
      this.isCartEmpty = false
      document.body.classList.remove('_lock')
    },
    async confirm() {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (this.isCartEmpty) {
        this.products = []
        cart.cart = cart.cart.filter((product) => product.is_favorite === true)
        for (let i = 0; i < cart.cart.length; i++) {
          cart.cart[i].quantity = 0
        }
        this.$store.commit('products/SET_PRODUCT_COUNT_WHEN_PAYMENT', null)
        localStorage.setItem('lorem', JSON.stringify(cart))
        if (cart && cart?.auth?.accessToken) {
          try {
            const res = (
              await deleteAllProductsFromBasket({
                url: `${this.$i18n.locale}/remove-cart`,
                accessToken: `Bearer ${cart?.auth?.accessToken}`,
              })
            ).data
          } catch (error) {
            console.log(error)
          }
        }
        this.isCartEmpty = false
      } else {
        if (this.productRemoveItem.quantity > 0) {
          const findProduct = cart?.cart.find(
            (product) => product.id === this.productRemoveItem.id
          )
          findProduct.quantity = 0
        } else {
          cart.cart = cart?.cart.filter(
            (product) => product.id != this.productRemoveItem.id
          )
        }
        this.products = this.products.filter(
          (product) => product.id != this.productRemoveItem.id
        )
        localStorage.setItem('lorem', JSON.stringify(cart))
        this.$store.commit(
          'products/SET_BASKET_PRODUCT_COUNT',
          this.productRemoveItem.quantity
        )
        if (cart && cart?.auth?.accessToken) {
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
                        accessToken: `Bearer ${access_token}`,
                      })
                    ).data
                  } catch (error) {
                    console.log('productAdd1', error)
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
      this.$store.commit('products/SET_PRODUCT_COUNT_WHEN_PAYMENT', null)
      this.isPayment = false
      document.body.classList.remove('_lock')
    },
  },
}
</script>
