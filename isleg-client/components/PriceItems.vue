<template>
  <div class="price__items">
    <div class="item__title">
      <h4>Sebet</h4>
      <button @click="$emit('cartTheEmpty')">
        <p>Sebedi boşat</p>
        <img src="@/assets/img/trush.svg" alt="" />
      </button>
    </div>
    <div class="price__texts">
      <span>Haryt mukdary:</span>
      <span>{{ count }} sany</span>
    </div>
    <div class="price__texts">
      <span>Doly bahasy:</span>
      <span>{{ totalPrice }} TMT</span>
    </div>
    <!-- <div class="price__texts">
      <span>Arzanladyş:</span>
      <span>-{{ productsChek.discount }} TMT</span>
    </div> -->
    <div class="price__texts">
      <span>Eltip berme:</span>
      <span>{{ productsChek.delivery }} TMT</span>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    count: {
      type: Number,
      default: () => 0,
    },
    productsChek: {
      type: Object,
      default: () => {},
    },
    totalPrice: {
      type: [String, Number],
      default: () => '',
    },
  },
}
</script>
