<template>
  <section class="basket __container">
    <basket-products></basket-products>
    <basket-prices @openPayment="isPayment = true"></basket-prices>
    <pop-up-payment
      v-if="isPayment"
      @close="isPayment = false"
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
}
</script>
