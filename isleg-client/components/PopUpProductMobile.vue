<template>
  <div class="mobile__product-box">
    <div class="mobile__product-wrapper">
      <div class="mobile__product-content">
        <div class="mobile__product-img">
          <div class="mobile__product-slider slider__product">
            <div class="slider__product-content">
              <swiper
                ref="mySwiper"
                :options="swiperOptions"
                class="slider__product-wrapper"
              >
                <swiper-slide class="slider__product-slide">
                  <img
                    :src="`${imgURL}/${productData.main_image.small}`"
                    alt=""
                  />
                </swiper-slide>
                <swiper-slide
                  class="slider__product-slide"
                  v-for="(image, i) in productData.images"
                  :key="i"
                >
                  <img :src="`${imgURL}/${image.small}`" alt="" />
                </swiper-slide>
                <div
                  class="swiper-pagination slider__product-pagination"
                  slot="pagination"
                />
              </swiper>
            </div>
          </div>
          <div class="mobile__product-arrow" @click.stop="$emit('close')">
            <span>
              <img src="@/assets/img/chevron-left.svg" alt="" />
            </span>
          </div>
        </div>
        <div class="mobile__product-datas">
          <div class="mobile__product-title">
            <p>{{ translationProductDescription(productData.translations) }}</p>
          </div>
          <div class="mobile__product-data">
            <div class="mobile__product-item">
              <span>Mocberi</span>
              <span>{{ productData.amount }} sany</span>
            </div>
          </div>
          <div class="mobile__product-prices">
            <div class="mobile__product-price">
              <span
                class="price__old"
                v-if="
                  productData &&
                  productData.old_price &&
                  productData.old_price > 0
                "
                >{{ productData && productData.old_price }} TMT</span
              >
              <span class="price__new"
                >{{ productData && productData.price }} TMT</span
              >
            </div>
            <div class="mobile__product-basket">
              <button v-if="quantity > 0">
                <span @click="removeFromBasket(productData)"><p></p></span>
                <span>{{ quantity }}</span>
                <span @click.stop="addToBasket(productData)">+</span>
              </button>
              <div
                class="basket__button"
                v-else
                @click.stop="addToBasket(productData)"
              >
                <img src="@/assets/img/basket.svg" alt="" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import translation from '@/mixins/translation'

export default {
  mixins: [translation],
  props: {
    isProductMobile: {
      type: Boolean,
      default: false,
    },
    quantity: {
      type: Number,
      default: 0,
    },
    productData: {
      type: Object,
      default: () => {},
    },
  },
  data() {
    return {
      swiperOptions: {
        slidesPerView: 1,
        centeredSlides: true,
        loop: false,
        pagination: {
          el: '.swiper-pagination',
          type: 'bullets',
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
    ...mapGetters('ui', ['imgURL']),
  },
  methods: {
    addToBasket(data) {
      this.$emit('add', data)
    },
    removeFromBasket(data) {
      this.$emit('remove', data)
    },
  },
}
</script>

<style lang="scss">
.mobile {
  @media (max-width: 950px) {
    &__product-box {
      position: fixed;
      width: 100vw;
      height: 100vh;
      background-color: #f5f5f5;
      top: 0;
      left: 0;
      z-index: 15;
      padding-top: 100px;
    }
    &__product {
      &-wrapper {
        padding: 20px;
        height: 100%;
        overflow-y: scroll;
        overflow-x: hidden;
        padding-bottom: 100px;
      }
      &-content {
      }
      &-img {
        position: relative;
        background: #fff;
        width: 100%;
        border-radius: 10px;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-bottom: 20px;
        .slider__product {
          padding: 30px 30px 5px 30px;
          overflow: hidden;
          &-content {
            height: 250px;
          }
          &-wrapper {
            height: 100% !important;
          }
          &-slide {
            height: 200px;
            width: 200px;
            img {
              width: 100%;
              height: 100%;
              object-fit: contain;
              object-position: center;
            }
          }
          &-pagination {
            display: flex;
            left: 50% !important;
            transform: translateX(-50%) !important;
            bottom: 10px !important;
            background: #eee;
            border-radius: 10px;
            width: fit-content !important;
            padding: 4px 2px;
            .swiper-pagination-bullet {
              width: 10px;
              height: 10px;
            }
            .swiper-pagination-bullet-active {
              background-color: #fd5e29;
            }
            .swiper-container {
            }
          }
        }
      }
      &-datas {
        position: relative;
        background: #fff;
        width: 100%;
        border-radius: 10px;
        padding: 10px;
        line-height: 120%;
        font-family: TTNormsPro;
        font-weight: 400;
        color: #1b3254;
      }
      &-title {
        font-size: 14px;
        padding-bottom: 10px;
        margin-bottom: 10px;
        border-bottom: 2px solid #eee;
      }
      &-data {
        padding-bottom: 10px;
        border-bottom: 2px solid #eee;
      }
      &-item {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin-bottom: 5px;
      }
      &-prices {
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 10px 0;
      }
      &-price {
        display: flex;
        flex-direction: column;
        line-height: 150%;
        .price__new {
          color: #fd5e29;
          font-size: 20px;
        }
        .price__old {
          text-decoration: line-through;
          font-size: 18px;
          opacity: 0.5;
        }
      }
      &-basket {
        .basket__button {
          background-color: #fd5e29;
          border-radius: 50%;
          padding: 5px;
          img {
            width: 22px;
            height: 22px;
          }
        }
        button {
          background-color: #fd5e29;
          padding: 5px;
          border-radius: 10px;
          display: flex;
          align-items: center;
          span {
            border-radius: 10px;
            display: flex;
            align-items: center;
            justify-content: center;

            &:nth-child(1) {
              background-color: #fff;
              //   margin-right: 15px;
              height: 26px;
              width: 26px;
              font-size: 24px;
              color: #fd5e29;
              p {
                width: 10px;
                height: 2px;
                background-color: #fd5e29;
              }
            }
            &:nth-child(2) {
              color: #fff;
              font-size: 20px;
              padding: 0 15px;
              margin: 0 5px;
            }
            &:nth-child(3) {
              background-color: #fff;
              //   margin-left: 15px;
              font-size: 22px;
              color: #fd5e29;
              height: 26px;
              width: 26px;
            }
          }
        }
      }
      &-arrow {
        position: absolute;
        top: 10px;
        left: 10px;
        span {
          border-radius: 50%;
          background: #eee;
          padding: 5px;
          display: flex;
          align-items: center;
          img {
            width: 14px;
            height: 14px;
            object-fit: contain;
            object-position: center;
          }
        }
      }
    }
  }
}
</style>
