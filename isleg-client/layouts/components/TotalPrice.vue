<template>
  <div class="total__price">
    <div class="price__number">
      <span>Jemi:</span>
      <p>112 manat</p>
    </div>
    <div class="price__send-btn" @click="$emit('payment')">
      <button class="btn-send">Sargyt et</button>
    </div>
  </div>
</template>

<script>
export default {}
</script>
