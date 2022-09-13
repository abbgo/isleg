<template>
  <div class="category__section">
    <product-title :title="productsCategory.name"></product-title>
    <products
      :productsCategory="productsCategory"
      @productPopUpOpen="productPopUpOpen"
    ></products>
  </div>
</template>

<script>
import ProductTitle from './ProductTitle.vue'
import Products from './Products.vue'
export default {
  components: { ProductTitle, Products },
  props: {
    productsCategory: {
      type: Object,
      default: () => {},
    },
  },
  methods: {
    productPopUpOpen(product) {
      this.$emit('openPopUp', product)
    },
  },
}
</script>
