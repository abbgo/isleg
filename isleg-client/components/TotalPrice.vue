<template>
  <div class="total__price">
    <div class="price__number">
      <span>Jemi:</span>
      <p>{{ totalPrice }} TMT</p>
    </div>
    <div class="price__send-btn" @click="$emit('payment')">
      <button class="btn-send">Sargyt et</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    totalPrice: {
      type: String,
      default: () => '',
    },
  },
}
</script>
