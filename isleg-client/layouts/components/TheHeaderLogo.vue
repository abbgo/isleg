<template>
  <a
    @click.prevent="$router.push(localeLocation('/'))"
    href="#"
    class="header__logo"
  >
    <img src="@/assets/img/logo.svg" alt="" />
  </a>
</template>

<script>
export default {}
</script>
