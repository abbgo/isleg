<template>
  <div :class="['pop-up', { active: isProduct }]" @click="$emit('close')">
    <div class="pop-up__product-body" @click.stop>
      <div class="pop-up__wrapper">
        <div class="pop-up__close" @click="$emit('close')">
          <svg
            width="60"
            height="60"
            viewBox="0 0 90 90"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              opacity="0.8"
              d="M41.1636 73.4122C54.9602 73.4122 66.1446 62.2278 66.1446 48.4312C66.1446 34.6346 54.9602 23.4502 41.1636 23.4502C27.367 23.4502 16.1826 34.6346 16.1826 48.4312C16.1826 62.2278 27.367 73.4122 41.1636 73.4122Z"
              fill="#1B3254"
            />
            <path
              d="M29.4334 58.6015L31.0243 60.1925L41.3658 49.851L51.4421 59.9273L53.033 58.3363L42.9568 48.26L53.2982 37.9186L51.7072 36.3276L41.3658 46.669L31.2895 36.5928L29.6985 38.1838L39.7748 48.26L29.4334 58.6015Z"
              fill="white"
            />
          </svg>
        </div>
        <div class="pop-up__container">
          <div class="product__slider" @mouseleave="$emit('currentImagePath')">
            <div class="product__slider-big">
              <figure
                class="zoom"
                :style="`backgroundImage: url(/img/popup/${bigSlider})`"
                @mousemove="imageZoom"
              >
                <img
                  :src="require(`@/assets/img/popup/${bigSlider}`)"
                  id="img-1"
                  alt=""
                />
              </figure>
            </div>
            <div class="product__slider-small">
              <div
                class="small__slider-img"
                v-for="image in images"
                :key="image.id"
              >
                <img
                  @mouseenter="$emit('changeImagePath', image.src)"
                  :src="require(`@/assets/img/popup/${image.src}`)"
                  alt=""
                />
              </div>
            </div>
          </div>
          <div class="product__datas">
            <h4 class="pop-up-product__title">Набор карандашей Bianyo</h4>
            <p class="product__text">
              Набор карандашей Bianyoчёрнографитных в коробке 12 шт
            </p>
            <span class="old__price price-old">213 manat</span>
            <span class="new__price price-new">120.10 manat</span>
            <div class="product__btns">
              <div class="btn__count-box">
                <button>
                  <svg
                    width="15"
                    height="4"
                    viewBox="0 0 25 4"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      d="M0.833 3.668C0.551 3.35467 0.41 2.75933 0.41 1.882C0.41 0.973333 0.582333 0.362333 0.927 0.0489995H24.192C24.4427 0.393666 24.568 1.00467 24.568 1.882C24.568 2.75933 24.427 3.35467 24.145 3.668H0.833Z"
                      fill="#FD5E29"
                    />
                  </svg>
                </button>
                <p>1</p>
                <button>
                  <svg
                    width="15"
                    height="15"
                    viewBox="0 0 23 23"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      d="M0.833 13.276C0.551 12.9627 0.41 12.3673 0.41 11.49C0.41 10.6127 0.582333 10.0017 0.927 9.657H9.857V0.585998C10.233 0.335331 10.844 0.209998 11.69 0.209998C12.536 0.209998 13.147 0.350998 13.523 0.632998V9.657H22.594C22.8447 10.0017 22.97 10.6127 22.97 11.49C22.97 12.3673 22.8133 12.9627 22.5 13.276H13.523V22.3C13.1783 22.582 12.5673 22.723 11.69 22.723C10.8127 22.723 10.2017 22.5507 9.857 22.206V13.276H0.833Z"
                      fill="#FD5E29"
                    />
                  </svg>
                </button>
              </div>
              <button class="product__add-btn">
                <svg
                  width="27"
                  height="27"
                  viewBox="0 0 27 27"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M8.32902 21.2573C7.81756 21.2573 7.31758 21.409 6.89232 21.6931C6.46705 21.9773 6.1356 22.3811 5.93987 22.8537C5.74414 23.3262 5.69293 23.8462 5.79271 24.3478C5.89249 24.8494 6.13878 25.3102 6.50044 25.6719C6.8621 26.0335 7.32288 26.2798 7.82452 26.3796C8.32615 26.4794 8.84611 26.4282 9.31864 26.2324C9.79117 26.0367 10.195 25.7053 10.4792 25.28C10.7634 24.8547 10.915 24.3548 10.915 23.8433C10.9155 23.5035 10.849 23.167 10.7192 22.853C10.5895 22.5391 10.399 22.2538 10.1588 22.0135C9.91854 21.7733 9.63325 21.5828 9.31926 21.4531C9.00528 21.3233 8.66877 21.2568 8.32902 21.2573ZM0.572021 0.571289V3.15729H3.15802L7.81202 12.9713L6.06702 16.1373C5.85329 16.5158 5.74196 16.9436 5.74402 17.3783C5.74587 18.0636 6.01892 18.7203 6.50348 19.2048C6.98805 19.6894 7.64474 19.9624 8.33002 19.9643H23.843V17.3783H8.87202C8.82949 17.3787 8.78731 17.3706 8.74794 17.3545C8.70857 17.3384 8.67281 17.3147 8.64273 17.2846C8.61266 17.2545 8.58888 17.2187 8.57279 17.1794C8.5567 17.14 8.54862 17.0978 8.54902 17.0553L8.58802 16.9003L9.75202 14.7933H19.384C19.8459 14.7947 20.2997 14.6718 20.6977 14.4375C21.0957 14.2031 21.4232 13.8659 21.646 13.4613L26.272 5.07129C26.3762 4.88103 26.4295 4.66716 26.427 4.45029C26.4252 4.10793 26.2884 3.78012 26.0463 3.53803C25.8042 3.29594 25.4764 3.15913 25.134 3.15729H6.01502L4.80002 0.571289H0.572021ZM21.258 21.2573C20.7466 21.2573 20.2466 21.409 19.8213 21.6931C19.3961 21.9773 19.0646 22.3811 18.8689 22.8537C18.6731 23.3262 18.6219 23.8462 18.7217 24.3478C18.8215 24.8494 19.0678 25.3102 19.4294 25.6719C19.7911 26.0335 20.2519 26.2798 20.7535 26.3796C21.2552 26.4794 21.7751 26.4282 22.2476 26.2324C22.7202 26.0367 23.124 25.7053 23.4082 25.28C23.6924 24.8547 23.844 24.3548 23.844 23.8433C23.8445 23.5035 23.778 23.167 23.6482 22.853C23.5185 22.5391 23.328 22.2538 23.0878 22.0135C22.8475 21.7733 22.5623 21.5828 22.2483 21.4531C21.9343 21.3233 21.5978 21.2568 21.258 21.2573Z"
                    fill="white"
                  />
                </svg>
                <h4>Sebede goş</h4>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    isProduct: {
      type: Boolean,
      default: false,
    },
    bigSlider: {
      type: String,
      default: () => '',
    },
    images: {
      type: Array,
      default: () => [],
    },
  },
  methods: {
    imageZoom(e) {
      let zoomer = e.currentTarget
      let offsetX, offsetY
      e.offsetX ? (offsetX = e.offsetX) : (offsetX = e.pageX)
      e.offsetY ? (offsetY = e.offsetY) : (offsetX = e.pageX)
      let x = (offsetX / zoomer.offsetWidth) * 100
      let y = (offsetY / zoomer.offsetHeight) * 100
      zoomer.style.backgroundPosition = x + '% ' + y + '%'
    },
  },
}
</script>

<style scoped>
figure.zoom {
  box-shadow: 0 10px 40px 0 rgba(47, 47, 47, 0.1);
  margin: 0 auto;
  width: 100%;
  height: 100%;
  background-position: 50% 50%;
  overflow: hidden;
  cursor: zoom-in;
}
figure.zoom img:hover {
  opacity: 0;
}
figure.zoom img {
  transition: opacity 0.3s;
  display: block;
  width: inherit;
}
</style>
