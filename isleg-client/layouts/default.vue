<template>
  <div class="main">
    <the-header></the-header>
    <main><nuxt></nuxt></main>
    <div
      class="toast-container __container"
      style="align-items: center; z-index: 1000000"
    ></div>
    <the-mobile-footer v-if="!mobileShow"></the-mobile-footer>
    <the-footer v-if="mobileShow"></the-footer>
  </div>
</template>
<script>
import TheMobileFooter from '~/components/TheMobileFooter.vue'
export default {
  components: { TheMobileFooter },
  data() {
    return {
      mobileShow: true,
      width: 950,
    }
  },
  methods: {
    windowWidth() {
      if (window.innerWidth <= this.width) {
        this.mobileShow = false
      } else {
        this.mobileShow = true
      }
    },
  },
  mounted() {
    this.windowWidth()
    window.addEventListener('resize', () => this.windowWidth())
  },
  beforeDestroy() {
    window.removeEventListener('resize')
  },
}
</script>
