<template>
  <main>
    <slider-main></slider-main>
    <section class="product__categoty __container">
      <product-column></product-column>
      <product-column></product-column>
      <product-column></product-column>
    </section>
    <slider-brends></slider-brends>
  </main>
</template>

<script>
import ProductColumn from '~/components/app/ProductColumn.vue'
import SliderBrends from '~/components/SliderBrends.vue'
export default {
  name: 'IndexPage',
  components: { ProductColumn, SliderBrends },
}
</script>
