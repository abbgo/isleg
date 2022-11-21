<template>
  <div class="basket__prices">
    <price-items
      :count="productCount"
      :productsChek="productsChek"
      :totalPrice="totalPrice"
      @cartTheEmpty="$emit('cartTheEmpty')"
    ></price-items>
    <total-price
      :totalPrice="total"
      @payment="$emit('openPayment')"
    ></total-price>
  </div>
</template>

<script>
export default {
  props: {
    productCount: {
      type: Number,
      default: () => 0,
    },
    productsChek: {
      type: Object,
      default: () => {},
    },
    totalPrice: {
      type: [String, Number],
      default: () => '',
    },
  },
  computed: {
    total() {
      return Number(
        Number(this.totalPrice) + Number(this.productsChek.delivery)
      ).toFixed(2)
    },
  },
}
</script>
