<template>
  <div class="basket__products">
    <product-check
      v-for="basketProduct in products"
      :key="basketProduct.id"
      :basketProduct="basketProduct"
      @popUpSureOpen="popUpSureOpen"
    ></product-check>
  </div>
</template>

<script>
export default {
  props: {
    products: {
      type: Array,
      default: () => [],
    },
  },
  methods: {
    popUpSureOpen(data) {
      this.$emit('popUpSureOpen', data)
    },
  },
}
</script>
