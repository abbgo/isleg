<template>
  <div class="__container">
    <section class="menu__filter menu__like">
      <div class="like__product">
        <span>
          <img src="@/assets/img/favorite.svg" alt="" />
        </span>
        <span>Halanlarym</span>
      </div>
    </section>
    <section class="product__categoty">
      <div class="category__section">
        <Products
          v-if="productsWishList.length"
          :products="productsWishList"
          @remove="remove"
        />
        <client-only v-else><p>Halan harytlaryňyzyň sanawy boş</p></client-only>
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
  data() {
    return {
      productsWishList: [],
    }
  },
  mounted() {
    const cart = JSON.parse(localStorage.getItem('lorem'))
    if (cart && cart.cart) {
      this.productsWishList = cart.cart.filter(
        (product) => product.is_favorite === true
      )
    } else {
      this.productsWishList = []
    }
    if (window.innerWidth <= 950) {
      if (document.body.classList.contains('_lock')) {
        document.body.classList.remove('_lock')
      }
    }
  },
  methods: {
    remove(data) {
      this.productsWishList = this.productsWishList.filter(
        (product) => product.id !== data.id
      )
    },
  },
}
</script>

<style scoped>
[v-cloak] {
  display: none;
}
</style>
