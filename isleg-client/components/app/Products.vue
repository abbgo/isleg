<template>
  <div class="products">
    <ProductBox
      v-for="product in products"
      :key="product.id"
      :product="product"
      @removeFromWishlist="removeFromWishlist"
    />
  </div>
</template>

<script>
const ProductBox = () => import('./ProductBox.vue')
export default {
  props: {
    products: {
      type: Array,
      default: () => [],
    },
  },
  components: {
    ProductBox,
  },
  methods: {
    removeFromWishlist(data) {
      this.$emit('remove', data)
    },
  },
}
</script>
