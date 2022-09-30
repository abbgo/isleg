<template>
  <div>
    <LazySliderMain
      v-if="brends && brends.length"
      :imgURL="imgURL"
      :brends="brends"
    />
    <section class="product__categoty __container">
      <ProductColumn
        v-for="productsCategory in productsCategories.filter(
          (item) => item.products !== null
        )"
        :key="productsCategory.id"
        :productsCategory="productsCategory"
      ></ProductColumn>
    </section>
    <LazySliderBrends />
  </div>
</template>

<script>
const ProductColumn = () => import('@/components/app/ProductColumn.vue')
import { mapGetters } from 'vuex'
export default {
  name: 'IndexPage',
  components: { ProductColumn },
  async fetch() {
    await this.$store.dispatch('ui/fetchBrends', {
      url: `${process.env.BASE_API}/${this.$i18n.locale}/brends`,
      $nuxt: this.$nuxt,
    })
    await this.$store.dispatch('products/fetchProductsCategories', {
      url: `${process.env.BASE_API}/${this.$i18n.locale}/homepage-categories`,
      $nuxt: this.$nuxt,
    })
  },
  computed: {
    ...mapGetters('ui', ['imgURL', 'brends']),
    ...mapGetters('products', ['productsCategories']),
  },
}
</script>
