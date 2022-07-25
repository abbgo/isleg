<template>
  <section class="basket __container">
    <basket-products></basket-products>
    <basket-prices @openPayment="openPayment"></basket-prices>
    <pop-up-payment
      :isPayment="isPayment"
      @close="closePaymentPopup"
    ></pop-up-payment>
  </section>
</template>

<script>
import BasketProducts from '@/components/BasketProducts.vue'
import BasketPrices from '@/components/BasketPrices.vue'
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
