<template>
  <div class="category__section">
    <ProductTitle :title="productsCategory.name" />
    <Products :productsCategory="productsCategory" />
  </div>
</template>

<script>
const ProductTitle = () => import('./ProductTitle.vue')
const Products = () => import('./Products.vue')
export default {
  components: { ProductTitle, Products },
  props: {
    productsCategory: {
      type: Object,
      default: () => {},
    },
  },
}
</script>
