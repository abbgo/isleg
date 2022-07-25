<template>
  <a
    @click.prevent="$router.push(localeLocation('/'))"
    @mouseenter="shadowLogo"
    href="#"
    class="header__logo"
  >
    <img src="/img/logo.svg" alt="" />
  </a>
</template>

<script>
export default {
  methods: {
    shadowLogo() {},
  },
}
</script>
