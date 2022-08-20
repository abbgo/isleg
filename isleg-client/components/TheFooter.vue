<template>
  <footer class="footer">
    <div class="footer__content">
      <div class="footer__items">
        <a
          href="#"
          @click.prevent="$router.push(localeLocation('/about-us'))"
          class="footer__item"
          >Biz barada</a
        >
        <a
          href="#"
          @click.prevent="
            $router.push(localeLocation('/delivery-and-payment-order'))
          "
          class="footer__item"
          >Eltip bermek we töleg tertibi</a
        >
        <a
          href="#"
          @click.prevent="$router.push(localeLocation('/communication'))"
          class="footer__item"
          >Aragatnaşyk</a
        >
        <a
          href="#"
          @click.prevent="
            $router.push(localeLocation('/terms-of-service-and-privacy-policy'))
          "
          class="footer__item"
          >Ulanyş düzgünleri we gizlinlik şertnamasy</a
        >
      </div>
      <div class="footer__bootom">
        <h4>© 2022-2023 isleg.com. Ähli hukuklary goraglydyr.</h4>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  watch: {
    '$i18n.locale': async function () {
      await this.$store.dispatch('ui/fetchFooter', {
        url: `${process.env.BASE_API}/${this.$i18n.locale}/footer`,
        $nuxt: this.$nuxt,
      })
    },
  },
  async fetch() {
    await this.$store.dispatch('ui/fetchFooter', {
      url: `${process.env.BASE_API}/${this.$i18n.locale}/footer`,
      $nuxt: this.$nuxt,
    })
  },
}
</script>
