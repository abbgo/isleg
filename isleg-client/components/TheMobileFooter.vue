<template>
  <footer class="footer">
    <div class="mobile__content">
      <ul>
        <li
          @click="$router.push(localeLocation(`/`))"
          :class="{ active: $route.name === `index___${$i18n.locale}` }"
        >
          <h4>Baş sahypa</h4>
          <svg
            width="20"
            height="21"
            viewBox="0 0 20 21"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M6.85202 19.773V16.715C6.85361 16.3388 7.00449 15.9786 7.27152 15.7137C7.53854 15.4487 7.89985 15.3006 8.27602 15.302H11.15C11.5272 15.3017 11.8891 15.4505 12.157 15.716C12.2891 15.8468 12.394 16.0026 12.4656 16.1742C12.5371 16.3459 12.574 16.53 12.574 16.716V19.774C12.5729 19.9349 12.6038 20.0944 12.6649 20.2432C12.726 20.3921 12.8161 20.5273 12.93 20.641C13.0444 20.7552 13.1803 20.8458 13.3297 20.9076C13.4791 20.9694 13.6393 21.0011 13.801 21.001H15.762C16.6765 21.0036 17.5548 20.644 18.205 20.001C18.5256 19.6841 18.7803 19.3068 18.9541 18.8909C19.128 18.475 19.2177 18.0287 19.218 17.578V8.86596C19.2168 8.50259 19.1355 8.14395 18.98 7.81556C18.8244 7.48717 18.5984 7.19709 18.318 6.96596L11.652 1.67496C11.088 1.22504 10.3846 0.985953 9.66323 0.998914C8.94183 1.01187 8.24755 1.27607 7.70002 1.74596L1.18502 6.96396C0.892384 7.18854 0.653673 7.47571 0.486359 7.80445C0.319044 8.1332 0.22736 8.4952 0.218018 8.86396V17.564C0.221459 18.4772 0.587475 19.3517 1.23557 19.9951C1.88367 20.6385 2.76078 20.9981 3.67402 20.995H5.59002C5.91512 20.9963 6.22756 20.869 6.45913 20.6408C6.69069 20.4126 6.82259 20.1021 6.82602 19.777L6.85202 19.773Z"
              fill="#8D98A9"
            />
          </svg>

          <img
            class="footer__active"
            src="@/assets/img/mobile-menu/active.svg"
            alt=""
          />
        </li>
        <li
          @click="$router.push(localeLocation(`/wishlist`))"
          :class="{ active: $route.name === `wishlist___${$i18n.locale}` }"
        >
          <h4>
            Halanlarym
            <span
              class="footer__circle-like"
              v-if="likesCount && likesCount > 0"
              >{{ likesCount > 99 ? '99+' : likesCount }}</span
            >
          </h4>
          <svg
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

          <img
            class="footer__active"
            src="@/assets/img/mobile-menu/active.svg"
            alt=""
          />
        </li>
        <li
          @click="$router.push(localeLocation(`/basket`))"
          :class="{ active: $route.name === `basket___${$i18n.locale}` }"
        >
          <h4>
            Sebet
            <span
              class="footer__circle"
              v-if="productCount && productCount > 0"
              >{{ productCount > 99 ? '99+' : productCount }}</span
            >
          </h4>

          <svg
            width="21"
            height="20"
            viewBox="0 0 21 20"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M7.98505 16.138C7.60309 16.1378 7.22966 16.2509 6.91197 16.4629C6.59429 16.675 6.34663 16.9765 6.20033 17.3293C6.05402 17.6822 6.01564 18.0705 6.09003 18.4451C6.16443 18.8197 6.34826 19.1639 6.61827 19.4341C6.88829 19.7042 7.23236 19.8882 7.60696 19.9628C7.98157 20.0374 8.36988 19.9992 8.72278 19.8531C9.07568 19.707 9.37732 19.4595 9.58955 19.1419C9.80177 18.8243 9.91505 18.451 9.91505 18.069C9.91531 17.8154 9.86558 17.5643 9.76869 17.33C9.67181 17.0956 9.52967 16.8827 9.35041 16.7033C9.17116 16.524 8.9583 16.3817 8.72402 16.2847C8.48973 16.1877 8.23862 16.1379 7.98505 16.138ZM0.102051 0V1.419H3.28805L5.82705 12.469C5.82705 12.469 6.02705 13.278 6.12705 13.497C6.22664 13.7895 6.40298 14.05 6.63765 14.251C6.87232 14.4521 7.1567 14.5864 7.46105 14.64C8.01405 14.627 16.2751 14.64 16.2751 14.64C16.6203 14.6414 16.9596 14.5498 17.2572 14.3747C17.5547 14.1996 17.7996 13.9475 17.9661 13.645L20.0711 6.168C20.148 6.0259 20.1875 5.86658 20.1861 5.705C20.1855 5.44914 20.0837 5.20389 19.9029 5.02287C19.7221 4.84185 19.4769 4.73979 19.2211 4.739H4.70205L7.46605 13.278H16.2801V12.585H7.91505L4.70205 0H0.102051ZM16.0381 16.138C15.6561 16.1378 15.2827 16.2509 14.965 16.4629C14.6473 16.675 14.3996 16.9765 14.2533 17.3293C14.107 17.6822 14.0686 18.0705 14.143 18.4451C14.2174 18.8197 14.4013 19.1639 14.6713 19.4341C14.9413 19.7042 15.2854 19.8882 15.66 19.9628C16.0346 20.0374 16.4229 19.9992 16.7758 19.8531C17.1287 19.707 17.4303 19.4595 17.6425 19.1419C17.8548 18.8243 17.9681 18.451 17.9681 18.069C17.9683 17.8154 17.9186 17.5643 17.8217 17.33C17.7248 17.0956 17.5827 16.8827 17.4034 16.7033C17.2242 16.524 17.0113 16.3817 16.777 16.2847C16.5427 16.1877 16.2916 16.1379 16.0381 16.138Z"
              fill="#8D98A9"
            />
          </svg>

          <img
            class="footer__active"
            src="@/assets/img/mobile-menu/active.svg"
            alt=""
          />
        </li>
        <li>
          <h4>Profil</h4>
          <svg
            width="21"
            height="20"
            viewBox="0 0 21 20"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M0.0939941 20.001V17.501C0.0939941 14.174 6.75399 12.501 10.094 12.501C13.434 12.501 20.094 14.175 20.094 17.501V20.001H0.0939941ZM5.09399 5.00095C5.09399 4.01204 5.38638 3.04537 5.93579 2.22312C6.4852 1.40087 7.2663 0.759761 8.17993 0.381322C9.09356 0.00288424 10.0987 -0.0957834 11.0686 0.0971427C12.0385 0.290069 12.9299 0.766044 13.6292 1.46531C14.3284 2.16457 14.8054 3.05546 14.9983 4.02537C15.1912 4.99527 15.0916 6.0009 14.7131 6.91453C14.3347 7.82816 13.6936 8.60877 12.8713 9.15818C12.0491 9.70759 11.0829 10.001 10.094 10.001C8.76791 10.001 7.49652 9.47428 6.55884 8.5366C5.62116 7.59891 5.09399 6.32703 5.09399 5.00095Z"
              fill="#8D98A9"
            />
          </svg>

          <img
            class="footer__active"
            src="@/assets/img/mobile-menu/active.svg"
            alt=""
          />
        </li>
      </ul>
    </div>
  </footer>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  data() {
    return {
      total: null,
      likes: null,
    }
  },
  watch: {
    $route: async function () {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      if (cart && cart?.cart) {
        if (!this.productCount) {
          this.total = cart.cart?.reduce((total, num) => {
            return total + num.quantity
          }, 0)
          this.$store.commit(
            'products/SET_PRODUCT_COUNT_WHEN_PAYMENT',
            this.total
          )
        }
        if (!this.likesCount) {
          this.likes =
            cart.cart?.filter((product) => product.is_favorite === true)
              ?.length || null
          this.$store.commit('products/SET_PRODUCT_LIKES', this.likes)
        }
      } else {
        this.$store.commit('products/SET_PRODUCT_COUNT_WHEN_PAYMENT', null)
        this.$store.commit('products/SET_PRODUCT_LIKES', null)
      }
    },
    productCount: function (val) {
      if (!val) {
        this.$store.commit('products/SET_PRODUCT_COUNT_WHEN_PAYMENT', null)
        this.total = null
      }
    },
  },
  computed: {
    ...mapGetters('products', ['productCount', 'likesCount']),
  },
  mounted() {
    const cart = JSON.parse(localStorage.getItem('lorem'))
    if (cart && cart?.cart) {
      if (!this.productCount) {
        this.total =
          cart.cart?.reduce((total, num) => {
            return total + num.quantity
          }, 0) || null
        this.$store.commit(
          'products/SET_PRODUCT_COUNT_WHEN_PAYMENT',
          this.total
        )
      }
      if (!this.likesCount) {
        this.likes =
          cart.cart?.filter((product) => product.is_favorite === true)
            ?.length || null
        this.$store.commit('products/SET_PRODUCT_LIKES', this.likes)
      }
    } else {
      this.$store.commit('products/SET_PRODUCT_COUNT_WHEN_PAYMENT', null)
      this.$store.commit('products/SET_PRODUCT_LIKES', null)
    }
  },
}
</script>

<style>
.mobile__content {
  position: fixed;
  bottom: 0;
  height: 64px;
  width: 100%;
  background: #fff;
  z-index: 20;
}

.mobile__content ul {
  height: 100%;
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.mobile__content ul li {
  display: flex;
  flex-direction: column-reverse;
  justify-content: center;
  align-items: center;
  cursor: pointer;
}
.mobile__content ul li h4 {
  font-family: TTNormsPro;
  padding-top: 6px;
  font-size: 12px;
  color: #707070;
  position: relative;
}
.footer__active {
  position: absolute;
  top: 0;
  display: none;
}
.mobile__content ul li {
}
.mobile__content ul li.active .footer__active {
  display: block;
}
.mobile__content ul li.active h4 {
  color: #fd5e29;
}
.mobile__content ul li.active svg path {
  fill: #fd5e29;
}

.mobile__content ul li.active .footer__circle {
  background-color: #fd5e29;
}
.mobile__content ul li.active .footer__circle-like {
  background-color: #fd5e29;
}
.footer__circle {
  width: 19px;
  height: 19px;
  background: #798393;
  border-radius: 50%;
  position: absolute;
  top: -26px;
  left: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  color: #fff;
}

.footer__circle-like {
  width: 19px;
  height: 19px;
  background: #798393;
  border-radius: 50%;
  position: absolute;
  top: -26px;
  left: 35px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  color: #fff;
}
</style>
