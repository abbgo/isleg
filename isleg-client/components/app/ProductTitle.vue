<template>
  <div class="product__title">
    <h3>{{ title }}</h3>
    <svg
      width="11"
      height="16"
      viewBox="0 0 11 16"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M10.5 8L0.5 16L0.5 0L10.5 8Z" fill="#FD5E29" />
    </svg>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      default: () => '',
    },
  },
}
</script>
