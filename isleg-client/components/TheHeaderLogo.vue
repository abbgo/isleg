<template>
  <a
    @click.prevent="$router.push(localeLocation('/'))"
    href="#"
    class="header__logo"
  >
    <img src="/img/logo.svg" alt="" />
  </a>
</template>
