<template>
  <div class="header__items">
    <span class="item___box dropdown">
      <span class="item__active">
        <img :src="`${imgURL}/${activeLang && activeLang.flag}`" alt="" />
      </span>
      <span class="item__dropdown">
        <img
          v-for="(item, i) in withOutLocaleLang"
          :key="i"
          :src="`${imgURL}/${item && item.flag}`"
          @click="$i18n.setLocale(item.name_short)"
          alt=""
        />
      </span>
    </span>
    <span class="item__border"></span>
    <span class="item___box account" @click.stop="$emit('openSignUp')">
      <svg
        class="sign__up"
        width="30"
        height="22"
        viewBox="0 0 30 22"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M0.99292 21.6281V19.05C0.99292 15.62 7.86737 13.8911 11.3074 13.8911C14.7474 13.8911 21.6257 15.62 21.6257 19.05V21.6281H0.99292ZM6.1521 6.15405C6.15183 4.78641 6.6942 3.47464 7.66089 2.5072C8.62758 1.53976 9.93973 0.995913 11.3074 0.995117C11.9849 0.994986 12.655 1.12823 13.281 1.38745C13.907 1.64667 14.4757 2.02678 14.9548 2.50586C15.4339 2.98494 15.8138 3.55371 16.073 4.17969C16.3322 4.80567 16.4657 5.47653 16.4656 6.15405C16.4657 6.83158 16.3322 7.50256 16.073 8.12854C15.8138 8.75452 15.4339 9.32328 14.9548 9.80237C14.4757 10.2815 13.907 10.6614 13.281 10.9207C12.655 11.1799 11.9849 11.3132 11.3074 11.3131C9.94025 11.3115 8.62889 10.7672 7.66284 9.7998C6.6968 8.83245 6.15476 7.52117 6.15503 6.15405H6.1521Z"
        />
        <path
          d="M11.313 11.3121C10.2928 11.3121 9.29557 11.0096 8.44734 10.4428C7.59911 9.87605 6.938 9.07048 6.5476 8.12798C6.1572 7.18548 6.05506 6.14837 6.25408 5.14782C6.4531 4.14727 6.94436 3.2282 7.66571 2.50684C8.38707 1.78548 9.30614 1.29423 10.3067 1.09521C11.3073 0.896183 12.3444 0.998329 13.2869 1.38873C14.2294 1.77912 15.0349 2.44024 15.6017 3.28846C16.1685 4.13669 16.471 5.13394 16.471 6.1541C16.4711 6.83149 16.3378 7.50228 16.0786 8.12813C15.8194 8.75399 15.4395 9.32266 14.9605 9.80165C14.4815 10.2806 13.9129 10.6606 13.287 10.9197C12.6612 11.1789 11.9904 11.3122 11.313 11.3121ZM22.919 8.7331V4.8641H25.496V8.7331H29.365V11.3121H25.496V15.1811H22.919V11.3121H19.05V8.7321L22.919 8.7331ZM11.313 13.8951C14.756 13.8951 21.63 15.6231 21.63 19.0531V21.6321H0.995972V19.0491C0.995972 15.6191 7.86897 13.8951 11.313 13.8951Z"
        />
      </svg>
      <profile-box
        :isProfileProps="isProfile"
        :information="myInformation"
        :orders="myOrders"
        :logOut="logOut"
        @close="$emit('closeProfilePopUp')"
      ></profile-box>
    </span>
    <span class="item__border"></span>
    <span
      class="item___box like__container"
      @click.stop="$router.push(localeLocation('/wishlist'))"
    >
      <svg
        class="like__svg"
        width="22"
        height="20"
        viewBox="0 0 22 20"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M11.1009 20.0001L9.52095 18.5611C3.90795 13.4711 0.201946 10.1141 0.201946 5.99405C0.194196 5.20475 0.343943 4.42184 0.642425 3.69111C0.940907 2.96038 1.38213 2.29653 1.94027 1.73838C2.49842 1.18024 3.16228 0.739012 3.893 0.44053C4.62373 0.142048 5.40665 -0.00769852 6.19595 5.20535e-05C7.12799 0.00787245 8.04755 0.215225 8.89276 0.60816C9.73796 1.00109 10.4892 1.57049 11.0959 2.27805C11.7027 1.57049 12.4539 1.00109 13.2991 0.60816C14.1443 0.215225 15.0639 0.00787245 15.9959 5.20535e-05C16.7854 -0.00783452 17.5685 0.141845 18.2994 0.440329C19.0303 0.738814 19.6944 1.1801 20.2526 1.73837C20.8109 2.29664 21.2522 2.96066 21.5507 3.69157C21.8492 4.42248 21.9988 5.20559 21.9909 5.99505C21.9909 10.1151 18.2849 13.4721 12.6719 18.5731L11.1009 20.0001Z"
          fill="#8D98A9"
        />
      </svg>
      <client-only
        ><span class="like__count" v-if="likesCount && likesCount > 0">{{
          likesCount > 99 ? '99+' : likesCount
        }}</span></client-only
      >
    </span>
    <span id="oppacity" class="item__border"></span>
    <button
      @mouseenter="mouseEnter"
      @mouseleave="mouseLeave"
      @click="$router.push(localeLocation('/basket'))"
      class="item___box shop"
    >
      <svg
        class="shop__svg"
        width="22"
        height="21"
        viewBox="0 0 22 21"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M7.27123 16.5C6.8572 16.5 6.45248 16.621 6.10823 16.8477C5.76398 17.0744 5.49567 17.3966 5.33723 17.7735C5.17879 18.1505 5.13734 18.5653 5.21811 18.9655C5.29888 19.3657 5.49825 19.7332 5.79101 20.0218C6.08377 20.3103 6.45677 20.5068 6.86284 20.5864C7.2689 20.666 7.6898 20.6251 8.07231 20.469C8.45482 20.3128 8.78175 20.0484 9.01177 19.7091C9.24179 19.3699 9.36456 18.971 9.36456 18.563C9.36496 18.292 9.31109 18.0235 9.20603 17.7731C9.10098 17.5226 8.9468 17.295 8.75233 17.1034C8.55787 16.9117 8.32694 16.7598 8.07279 16.6562C7.81863 16.5527 7.54624 16.4996 7.27123 16.5V16.5ZM0.991211 0V2.063H3.08455L6.85215 9.892L5.43968 12.419C5.26654 12.7209 5.1763 13.0622 5.17789 13.409C5.17949 13.9557 5.40056 14.4795 5.79279 14.866C6.18502 15.2526 6.71653 15.4704 7.27123 15.472H19.8302V13.409H7.71059C7.6761 13.4094 7.64187 13.403 7.60993 13.3902C7.57798 13.3774 7.54896 13.3584 7.52457 13.3343C7.50018 13.3103 7.48091 13.2817 7.4679 13.2502C7.45488 13.2187 7.44839 13.185 7.4488 13.151L7.48026 13.027L8.4219 11.346H16.2189C16.5928 11.3473 16.9602 11.2493 17.2824 11.0625C17.6047 10.8756 17.87 10.6067 18.0504 10.284L21.7978 3.59C21.8829 3.43689 21.9263 3.26459 21.9236 3.09C21.9217 2.81784 21.811 2.55739 21.6156 2.36513C21.4201 2.17287 21.1556 2.06431 20.8794 2.063H5.39706L4.41381 0H0.991211ZM17.7339 16.5C17.3198 16.5 16.9151 16.621 16.5709 16.8477C16.2266 17.0744 15.9583 17.3966 15.7999 17.7735C15.6414 18.1505 15.6 18.5653 15.6807 18.9655C15.7615 19.3657 15.9609 19.7332 16.2536 20.0218C16.5464 20.3103 16.9194 20.5068 17.3255 20.5864C17.7315 20.666 18.1524 20.6251 18.5349 20.469C18.9175 20.3128 19.2444 20.0484 19.4744 19.7091C19.7044 19.3699 19.8272 18.971 19.8272 18.563C19.8276 18.292 19.7737 18.0235 19.6687 17.7731C19.5636 17.5226 19.4094 17.295 19.215 17.1034C19.0205 16.9117 18.7896 16.7598 18.5354 16.6562C18.2813 16.5527 18.0089 16.4996 17.7339 16.5V16.5Z"
          fill="#8D98A9"
        />
      </svg>
      <client-only
        ><span class="shop__count" v-if="productCount && productCount > 0">{{
          productCount > 99 ? '99+' : productCount
        }}</span></client-only
      >
      <span class="shop__span">Sebet</span>
    </button>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  props: {
    isProfile: {
      type: Boolean,
      default: false,
    },
    myInformation: {
      type: String,
      default: () => '',
    },
    myOrders: {
      type: String,
      default: () => '',
    },
    logOut: {
      type: String,
      default: () => '',
    },
    withOutLocaleLang: {
      type: Array,
      default: () => [],
    },
    activeLang: {
      type: Object,
      default: () => {},
    },
    imgURL: {
      type: String,
      default: () => '',
    },
  },
  data() {
    return {
      totalCount: null,
      likes: null,
    }
  },
  computed: {
    ...mapGetters('products', ['productCount', 'likesCount']),
  },
  mounted() {
    const cart = JSON.parse(localStorage.getItem('lorem'))
    if (cart && cart.cart) {
      if (!this.productCount) {
        this.totalCount =
          cart.cart.reduce((total, num) => {
            return total + num?.quantity
          }, 0) || null
        this.$store.commit(
          'products/SET_PRODUCT_COUNT_WHEN_PAYMENT',
          this.totalCount
        )
      }
      if (!this.likesCount) {
        this.likes =
          cart.cart.filter((product) => product.is_favorite === true).length ||
          null
        this.$store.commit('products/SET_PRODUCT_LIKES', this.likes)
      }
    } else {
      this.$store.commit('products/SET_PRODUCT_COUNT_WHEN_PAYMENT', null)
      this.$store.commit('products/SET_PRODUCT_LIKES', null)
    }
  },
  methods: {
    mouseEnter() {
      let item__border = document.getElementById('oppacity')
      item__border.classList.add('active__item-border')
    },
    mouseLeave() {
      let item__border = document.getElementById('oppacity')
      item__border.classList.remove('active__item-border')
    },
  },
}
</script>
