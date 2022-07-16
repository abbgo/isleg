<template>
  <main>
    <slider-main></slider-main>
    <section class="product__categoty __container">
      <product-column @productPopUpOpen="isProduct = true"></product-column>
      <product-column @productPopUpOpen="isProduct = true"></product-column>
      <product-column @productPopUpOpen="isProduct = true"></product-column>
    </section>
    <slider-brends></slider-brends>
    <pop-up-product
      :images="images"
      :bigSlider="bigSlider"
      @changeImagePath="changeImagePath"
      @currentImagePath="currentImagePath"
      @close="isProduct = false"
    ></pop-up-product>
  </main>
</template>

<script>
import ProductColumn from '~/components/app/ProductColumn.vue'
import SliderBrends from '~/components/SliderBrends.vue'
export default {
  name: 'IndexPage',
  components: { ProductColumn, SliderBrends },
  data() {
    return {
      isProduct: false,
      bigSlider: 'bigSlider.jpg',
      images: [
        { id: 1, src: '1.jpg' },
        { id: 2, src: '2.jpg' },
        { id: 3, src: '3.jpg' },
        { id: 4, src: '1.jpg' },
        { id: 5, src: '2.jpg' },
        { id: 6, src: '3.jpg' },
      ],
    }
  },
  methods: {
    changeImagePath(imagePath) {
      this.bigSlider = imagePath
    },
    currentImagePath() {
      this.bigSlider = 'bigSlider.jpg'
    },
  },
}
</script>
