<template>
  <div class="category__section">
    <ProductTitle :title="productsCategory.name" />
    <Products :products="productsCategory.products" />
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
}
</script>
