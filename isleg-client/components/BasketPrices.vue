<template>
  <div class="basket__prices">
    <price-items></price-items>
    <total-price @payment="$emit('openPayment')"></total-price>
  </div>
</template>

<script>
export default {}
</script>
