<template>
  <footer class="footer">
    <div class="footer__content">
      <div class="footer__items">
        <a
          href="#"
          @click.prevent="$router.push(localeLocation('/about-us'))"
          class="footer__item"
          >{{ about }}</a
        >
        <a
          href="#"
          @click.prevent="
            $router.push(localeLocation('/delivery-and-payment-order'))
          "
          class="footer__item"
          >{{ payment }}</a
        >
        <a
          href="#"
          @click.prevent="$router.push(localeLocation('/communication'))"
          class="footer__item"
          >{{ contact }}</a
        >
        <a
          href="#"
          @click.prevent="
            $router.push(localeLocation('/terms-of-service-and-privacy-policy'))
          "
          class="footer__item"
          >{{ secure }}</a
        >
      </div>
      <div class="footer__bootom">
        <h4>© {{ new Date().getFullYear() }} {{ word }}.</h4>
      </div>
    </div>
  </footer>
</template>

<script>
import { mapGetters } from 'vuex'
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
  computed: {
    ...mapGetters('ui', ['about', 'payment', 'contact', 'secure', 'word']),
  },
}
</script>
