<template>
  <section class="basket __container" v-if="products.length">
    <BasketProducts :products="products" @popUpSureOpen="openPopUpSure" />
    <BasketPrices
      :productCount="productCount"
      :totalPrice="totalPrice"
      :productsChek="productsChek"
      @openPayment="openPayment"
    />
    <LazyPopUpPayment
      :isPayment="isPayment"
      :paymentDatas="paymentDatas"
      :phone="phone"
      :name="name"
      :signIn="signIn"
      :totalPrice="totalPrice"
      @close="closePaymentPopup"
      @paymentSuccesfullySended="paymentSuccesfullySended"
      @paymentRegister="paymentRegister"
    />
    <LazyPopUpSure
      :isSure="isSure"
      @close="closePopUpSure"
      @confirm="confirm"
    />
  </section>
</template>

<script>
import BasketProducts from '@/components/BasketProducts.vue'
import BasketPrices from '@/components/BasketPrices.vue'
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
      isPayment: false,
      isSure: false,
      productRemoveItem: null,
      products: [],
      productsChek: {
        discount: 15,
        delivery: 10,
      },
      paymentDatas: {
        text: null,
        types: [],
        orderTimes: {
          title: null,
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
        return total + num.price * num.quantity
      }, 0)
    },
  },
  methods: {
    async getBasketProducts() {
      let cart = JSON.parse(localStorage.getItem('lorem'))
      if (cart?.cart) {
        let totalCount = cart.cart.reduce((total, num) => {
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
        console.log(order_times)
        if (status) {
          this.paymentDatas.orderTimes.title = order_times[0].translation_date
          for (let i = 0; i < order_times[0].times.length; i++) {
            this.paymentDatas.orderTimes.times.push({
              id: i + Math.random() * 2,
              time: order_times[0].times[i].time,
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
    closePopUpSure() {
      this.productRemoveItem = null
      this.isSure = false
      document.body.classList.remove('_lock')
    },
    async confirm() {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      cart.cart = await cart?.cart.filter(
        (product) => product.id != this.productRemoveItem.id
      )
      this.products = cart.cart
      localStorage.setItem('lorem', JSON.stringify(cart))
      await this.$store.commit(
        'products/SET_BASKET_PRODUCT_COUNT',
        this.productRemoveItem.quantity
      )
      this.closePopUpSure()
    },
    paymentRegister() {
      this.isPayment = false
      this.$store.commit('ui/SET_OPEN_ISOPENSIGNUP')
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
