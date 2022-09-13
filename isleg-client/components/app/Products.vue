<template>
  <div class="products">
    <product-box
      v-for="product in productsCategory.products"
      :key="product.id"
      :product="product"
      @productPopUp="productPopUp"
    ></product-box>
  </div>
</template>

<script>
import ProductBox from './ProductBox.vue'
export default {
  components: {
    ProductBox,
  },
  props: {
    productsCategory: {
      type: Object,
      default: () => {},
    },
  },
  methods: {
    openPop() {
      console.log('eee')
    },
    productPopUp(product) {
      this.$emit('productPopUpOpen', product)
    },
  },
}
</script>
