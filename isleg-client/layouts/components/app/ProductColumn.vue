<template>
  <div class="category__section">
    <product-title></product-title>
    <products></products>
  </div>
</template>

<script>
import ProductTitle from './ProductTitle.vue'
import Products from './Products.vue'
export default {
  components: { ProductTitle, Products },
}
</script>
