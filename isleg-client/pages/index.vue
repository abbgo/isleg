<template>
  <div>
    <slider-main :imgURL="imgURL" :brends="brends"></slider-main>
    <section class="product__categoty __container">
      <product-column @productPopUpOpen="productPopUpOpen"></product-column>
      <product-column @productPopUpOpen="productPopUpOpen"></product-column>
      <product-column @productPopUpOpen="productPopUpOpen"></product-column>
    </section>
    <slider-brends></slider-brends>
    <pop-up-product
      :isProduct="isProduct"
      :images="images"
      :bigSlider="bigSlider"
      @changeImagePath="changeImagePath"
      @currentImagePath="currentImagePath"
      @close="closeProductPopUp"
    ></pop-up-product>
  </div>
</template>

<script>
import ProductColumn from '~/components/app/ProductColumn.vue'
import SliderBrends from '~/components/SliderBrends.vue'
import { mapGetters } from 'vuex'
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
  async fetch() {
    await this.$store.dispatch('ui/fetchBrends', {
      url: `${process.env.BASE_API}/${this.$i18n.locale}/brends`,
      $nuxt: this.$nuxt,
    })
  },
  computed: {
    ...mapGetters('ui', ['imgURL', 'brends']),
  },
  methods: {
    productPopUpOpen() {
      this.isProduct = true
      document.body.classList.add('_lock')
    },
    changeImagePath(imagePath) {
      this.bigSlider = imagePath
    },
    currentImagePath() {
      this.bigSlider = 'bigSlider.jpg'
    },
    closeProductPopUp() {
      this.isProduct = false
      document.body.classList.remove('_lock')
    },
  },
}
</script>
