<template>
  <div class="basket__products">
    <product-check
      v-for="basketProduct in basketProducts"
      :key="basketProduct.id"
      :basketProduct="basketProduct"
    ></product-check>
  </div>
</template>

<script>
export default {
  data() {
    return {
      basketProducts: [],
    }
  },
  async mounted() {
    await this.getBasketProducts()
  },
  methods: {
    async getBasketProducts() {
      let cart = JSON.parse(localStorage.getItem('lorem'))
      console.log(cart)
      if (cart) {
        let totalCount = cart.cart.reduce((total, num) => {
          return total + num.quantity
        }, 0)
        this.basketProducts = cart.cart
        this.$store.commit(
          'products/SET_PRODUCT_COUNT',
          totalCount == 0 ? null : totalCount
        )
      } else {
        this.basketProducts = []
      }
    },
  },
}
</script>
