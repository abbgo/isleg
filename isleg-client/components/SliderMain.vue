<template>
  <section class="slider __container">
    <swiper ref="mySwiper" :options="swiperOptions" class="big__slider">
      <swiper-slide v-for="brend in brends" :key="brend.id">
        <img :data-src="`${imgURL}/${brend.image}`" loading="lazy" alt="" />
      </swiper-slide>
    </swiper>
    <div class="swiper-pagination"></div>
    <button class="arrow arrow__left">
      <svg
        width="12"
        height="18"
        viewBox="0 0 12 18"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M10.775 1.28906L2.401 8.64306L10.776 16.2731"
          stroke="#1B3254"
          stroke-width="2"
        />
      </svg>
    </button>
    <button class="arrow arrow__right">
      <svg
        width="12"
        height="18"
        viewBox="0 0 12 18"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M1.22595 1.28906L9.59995 8.64306L1.22595 16.2731"
          stroke="#1B3254"
          stroke-width="2"
        />
      </svg>
    </button>
  </section>
</template>

<script>
import { Swiper, SwiperSlide } from 'vue-awesome-swiper'
import observer from '@/mixins/observer'

export default {
  mixins: [observer],
  components: {
    Swiper,
    SwiperSlide,
  },
  props: {
    imgURL: {
      type: String,
      default: () => '',
    },
    brends: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      swiperOptions: {
        autoplay: {
          delay: 3000,
          disableOnInteraction: false,
        },
        loop: true,
        speed: 1500,
        autoplayStopOnLast: false,
        navigation: {
          nextEl: '.arrow__right',
          prevEl: '.arrow__left',
        },
        pagination: {
          el: '.swiper-pagination',
          clickable: true,
        },
        keyboard: {
          enabled: true,
          onlyInViewport: true,
          pageUpDown: true,
        },
      },
    }
  },
  computed: {
    swiper() {
      return this.$refs?.mySwiper?.$swiper
    },
  },
}
</script>

<style scoped>
.swiper-slide {
  background: url('../assets/img/isloading.svg') center no-repeat;
}
</style>
