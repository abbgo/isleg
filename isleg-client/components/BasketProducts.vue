<template>
  <div class="basket__products">
    <product-check
      v-for="basketProduct in products"
      :key="basketProduct.id"
      :basketProduct="basketProduct"
      @popUpSureOpen="popUpSure"
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
    popUpSure(item) {
      this.$emit('sure', item)
    },
  },
}
</script>
