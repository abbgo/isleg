<template>
  <section class="basket __container" v-if="products.length">
    <BasketProducts :products="products" @popUpSureOpen="openPopUpSure" />
    <BasketPrices
      :productCount="productCount"
      :totalPrice="totalPrice"
      :productsChek="productsChek"
      @openPayment="openPayment"
    />
    <LazyPopUpPayment :isPayment="isPayment" @close="closePaymentPopup" />
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
    }
  },
  async mounted() {
    await this.getBasketProducts()
  },
  computed: {
    ...mapGetters('products', ['productCount']),
    totalPrice() {
      return this.products?.reduce((total, num) => {
        return total + num.price * num.quantity
      }, 0)
    },
  },
  methods: {
    async getBasketProducts() {
      let cart = JSON.parse(localStorage.getItem('lorem'))
      if (cart) {
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
    openPayment() {
      this.isPayment = true
      document.body.classList.add('_lock')
    },
    closePaymentPopup() {
      this.isPayment = false
      document.body.classList.remove('_lock')
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
      cart.cart = await cart.cart.filter(
        (product) => product.id !== this.productRemoveItem.id
      )
      this.products = cart.cart
      localStorage.setItem('lorem', JSON.stringify(cart))
      await this.$store.commit(
        'products/SET_BASKET_PRODUCT_COUNT',
        this.productRemoveItem.quantity
      )
      this.closePopUpSure()
    },
  },
}
</script>
