<template>
  <div class="products">
    <ProductBox
      v-for="product in productsCategory.products"
      :key="product.id"
      :product="product"
      @productPopUp="productPopUp"
    />
  </div>
</template>

<script>
const ProductBox = () => import('./ProductBox.vue')
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
