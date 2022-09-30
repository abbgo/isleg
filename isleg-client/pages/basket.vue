<template>
  <section class="basket __container">
    <BasketProducts />
    <BasketPrices @openPayment="openPayment" />
    <LazyPopUpPayment :isPayment="isPayment" @close="closePaymentPopup" />
  </section>
</template>

<script>
const BasketProducts = () => import('@/components/BasketProducts.vue')
const BasketPrices = () => import('@/components/BasketPrices.vue')

export default {
  name: 'BasketPage',
  components: { BasketProducts, BasketPrices },
  data() {
    return {
      isPayment: false,
    }
  },
  methods: {
    openPayment() {
      this.isPayment = true
      document.body.classList.add('_lock')
    },
    closePaymentPopup() {
      this.isPayment = false
      document.body.classList.remove('_lock')
    },
  },
}
</script>
