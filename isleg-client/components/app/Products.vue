<template>
  <div class="products">
    <product-box></product-box>
    <product-box></product-box>
    <product-box></product-box>
    <product-box></product-box>
    <product-box></product-box>
    <product-box></product-box>
  </div>
</template>

<script>
import ProductBox from './ProductBox.vue'
export default {
  components: {
    ProductBox,
  },
}
</script>
