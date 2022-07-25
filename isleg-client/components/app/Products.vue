<template>
  <div class="products">
    <product-box @productPopUp="$emit('productPopUpOpen')"></product-box>
    <product-box @productPopUp="$emit('productPopUpOpen')"></product-box>
    <product-box @productPopUp="$emit('productPopUpOpen')"></product-box>
    <product-box @productPopUp="$emit('productPopUpOpen')"></product-box>
    <product-box @productPopUp="$emit('productPopUpOpen')"></product-box>
    <product-box @productPopUp="$emit('productPopUpOpen')"></product-box>
  </div>
</template>

<script>
import ProductBox from './ProductBox.vue'
export default {
  components: {
    ProductBox,
  },
  methods: {
    openPop() {
      console.log('eee')
    },
  },
}
</script>
