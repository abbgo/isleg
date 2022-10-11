<template>
  <div class="price__items">
    <div class="item__title">
      <h4>Sizin sebediniz</h4>
    </div>
    <div class="price__texts">
      <span>Haryt mukdary:</span>
      <span>{{ count }} sany</span>
    </div>
    <div class="price__texts">
      <span>Doly bahasy:</span>
      <span>{{ totalPrice }} manat</span>
    </div>
    <!-- <div class="price__texts">
      <span>Arzanladyş:</span>
      <span>-{{ productsChek.discount }} manat</span>
    </div> -->
    <div class="price__texts">
      <span>Eltip berme:</span>
      <span>{{ productsChek.delivery }} manat</span>
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
      type: Number,
      default: () => 0,
    },
  },
}
</script>
