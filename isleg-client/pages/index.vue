<template>
  <div>
    <client-only>
      <LazySliderMain
        v-if="brends && brends.length"
        :imgURL="imgURL"
        :brends="brends"
      />
    </client-only>
    <section class="product__categoty __container">
      <client-only>
        <ProductColumn
          v-for="productsCategory in productsCategories.filter(
            (item) => item.products !== null
          )"
          :key="productsCategory.id"
          :productsCategory="productsCategory"
        />
      </client-only>
    </section>
    <client-only>
      <LazySliderBrends
        v-if="brends && brends.length"
        :imgURL="imgURL"
        :brends="brends"
      />
    </client-only>
  </div>
</template>

<script>
import ProductColumn from '@/components/app/ProductColumn.vue'
import { mapGetters } from 'vuex'
export default {
  name: 'IndexPage',
  components: { ProductColumn },
  async fetch() {
    await this.$store.dispatch('products/fetchProductsCategories', {
      url: `${process.env.BASE_API}/${this.$i18n.locale}/homepage-categories`,
      $nuxt: this.$nuxt,
    })
    await this.$store.dispatch('ui/fetchBrends', {
      url: `${process.env.BASE_API}/${this.$i18n.locale}/brends`,
      $nuxt: this.$nuxt,
    })
  },
  computed: {
    ...mapGetters('ui', ['imgURL', 'brends']),
    ...mapGetters('products', ['productsCategories']),
  },
  mounted() {
    if (window.innerWidth <= 950) {
      if (window.innerWidth <= 950) {
        if (document.body.classList.contains('_lock')) {
          document.body.classList.remove('_lock')
        }
      }
    }
  },
}
</script>
