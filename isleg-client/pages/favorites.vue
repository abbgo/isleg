<template>
  <div class="pages__container">
    <section class="menu__filter menu__like">
      <div class="like__product">
        <span>
          <img src="@/assets/img/favorite.svg" alt="" />
        </span>
        <span>Halanlarym</span>
      </div>
    </section>
    <section class="product__categoty __container">
      <div class="category__section">
        <products></products>
      </div>
    </section>
  </div>
</template>

<script>
import Products from '@/components/app/Products.vue'
export default {
  components: {
    Products,
  },
  mounted() {
    console.log(this.$route.name)
  },
}
</script>
