<template>
  <div class="main">
    <TheHeader />
    <main><Nuxt /></main>
    <div
      class="toast-container __container"
      style="align-items: center; z-index: 1000000"
    ></div>
    <TheMobileFooter v-if="!mobileShow"></TheMobileFooter>
    <LazyTheFooter v-if="mobileShow" />
  </div>
</template>
<script>
export default {
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
