<template>
  <div class="brends">
    <div class="brends__container __container">
      <div>
        <swiper
          ref="mySwiperSmall"
          :options="swiperSmallOptions"
          class="brends__animation swiper"
        >
          <swiper-slide
            v-for="brend in brends"
            :key="brend.id"
            class="brends__animation-slide swiper-slide"
          >
            <img
              :data-src="`${imgURL}/${brend.image}`"
              loading="lazy"
              alt=""
              @click="$router.push(localeLocation(`/brend/${brend.id}`))"
            />
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </div>
</template>

<script>
import { Swiper, SwiperSlide } from 'vue-awesome-swiper'
import observer from '@/mixins/observer'

export default {
  mixins: [observer],
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

<style scoped>
.swiper-slide {
  background: url('../assets/img/isloading.svg') center no-repeat;
}
</style>
