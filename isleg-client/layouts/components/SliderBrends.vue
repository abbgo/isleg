<template>
  <div class="brends">
    <div class="brends__container __container">
      <div>
        <swiper
          ref="mySwiperSmall"
          :options="swiperSmallOptions"
          class="brends__animation swiper"
        >
          <swiper-slide class="brends__animation-slide swiper-slide">
            <img src="@/assets/img/slider/1.png" alt="" />
          </swiper-slide>
          <swiper-slide class="brends__animation-slide swiper-slide">
            <img src="@/assets/img/slider/1.png" alt="" />
          </swiper-slide>
          <swiper-slide class="brends__animation-slide swiper-slide">
            <img src="@/assets/img/slider/1.png" alt="" />
          </swiper-slide>
          <swiper-slide class="brends__animation-slide swiper-slide">
            <img src="@/assets/img/slider/1.png" alt="" />
          </swiper-slide>
          <swiper-slide class="brends__animation-slide swiper-slide">
            <img src="@/assets/img/slider/1.png" alt="" />
          </swiper-slide>
          <swiper-slide class="brends__animation-slide swiper-slide">
            <img src="@/assets/img/slider/1.png" alt="" />
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </div>
</template>

<script>
import { Swiper, SwiperSlide } from 'vue-awesome-swiper'
export default {
  components: {
    Swiper,
    SwiperSlide,
  },
  data() {
    return {
      swiperSmallOptions: {
        slidesPerView: 3,
        spaceBetween: 25,
        speed: 1000,
        loop: true,
        keyboard: {
          enabled: true,
          onlyInViewport: true,
          pageUpDown: true,
        },
        autoplay: {
          delay: 2000,
          reverseDirection: true,
          disableOnInteraction: false,
          pauseOnMouseEnter: true,
        },
        breakpoints: {
          600: {
            slidesPerView: 4,
          },
          900: {
            slidesPerView: 5,
          },
          1200: {
            slidesPerView: 6,
          },
        },
      },
    }
  },
  computed: {
    swiperSmall() {
      return this.$refs.mySwiperSmall.$swiper
    },
  },
  async mounted() {
    await this.swiperSmall
  },
}
</script>
