<template>
  <a
    @click.prevent="$router.push(localeLocation('/'))"
    href="#"
    class="header__logo"
  >
    <img src="@/assets/img/logo.svg" alt="isleg" />
  </a>
</template>

<script>
export default {
  props: {
    imgURL: {
      type: String,
      default: () => '',
    },
    logo: {
      type: String,
      default: () => '',
    },
  },
}
</script>
