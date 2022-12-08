<template>
  <div class="product__box" @click.stop="openPopUpPoduct">
    <div class="product__img">
      <img
        loading="lazy"
        :data-src="`${imgURL}/${getProduct && getProduct.main_image.medium}`"
      />
    </div>
    <div class="product__description">
      <p>
        {{ translationProductName(getProduct.translations) }}
      </p>
    </div>
    <div class="product__price">
      <div class="new__price">
        {{ getProduct && getProduct.price }}
        TMT
      </div>
      <div class="old__price" v-if="getProduct && getProduct.old_price > 0">
        <span>
          {{ getProduct && getProduct.old_price }}
          TMT</span
        >
        <span>-15%</span>
      </div>
    </div>
    <div class="chek__count box__count" @click.stop v-if="quantity > 0">
      <button @click.stop="basketRemove(getProduct)">
        <svg
          width="8"
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
      <p>
        {{ getProduct && getProduct.quantity }}
      </p>
      <button @click.stop="basketAdd(getProduct)">
        <svg
          width="10"
          height="10"
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
    <button class="product__add-btn" v-else @click.stop="basketAdd(getProduct)">
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
    <div class="product__like">
      <svg
        @click.stop="productLike(getProduct)"
        width="26"
        height="25"
        viewBox="0 0 26 25"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M13 23L11.2602 21.4185C5.08064 15.8191 1.0003 12.1266 1.0003 7.5928C0.991958 6.7246 1.15698 5.86346 1.48571 5.05975C1.81445 4.25603 2.30029 3.52589 2.91483 2.91203C3.52936 2.29817 4.26025 1.81292 5.06473 1.48467C5.8692 1.15642 6.73112 0.991759 7.60005 1.00032C8.62705 1.00808 9.64046 1.23574 10.572 1.66794C11.5035 2.10014 12.3314 2.72684 13 3.5058C13.6686 2.72684 14.4965 2.10014 15.428 1.66794C16.3595 1.23574 17.3729 1.00808 18.3999 1.00032C19.269 0.991755 20.1311 1.15647 20.9357 1.48483C21.7402 1.81318 22.4712 2.29858 23.0858 2.91262C23.7003 3.52666 24.1861 4.25701 24.5148 5.06091C24.8434 5.86482 25.0083 6.72615 24.9997 7.59449C24.9997 12.1266 20.9193 15.8191 14.7398 21.4303L13 23Z"
          :fill="isFavorite ? fillColor : fillEmpty"
          stroke="#FD5E29"
          stroke-width="2"
        />
      </svg>
    </div>
    <div class="product__new" v-if="getProduct && getProduct.is_new">täze</div>
    <LazyPopUpProduct
      v-if="isProduct"
      :isProduct="isProduct"
      :productData="getProduct"
      :quantity="quantity"
      @add="basketAdd"
      @remove="basketRemove"
      @close="closePopUpPoduct"
    />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import observer from '@/mixins/observer'
import translation from '@/mixins/translation'
import { productAdd, productLike, getRefreshToken } from '@/api/user.api'
export default {
  mixins: [observer, translation],
  props: {
    product: {
      type: Object,
      default: () => {},
    },
  },
  data() {
    return {
      fillEmpty: null,
      fillColor: '#FD5E29',
      quantity: 0,
      isFavorite: false,
      isProduct: false,
      isDisabled: false,
      count: 0,
    }
  },
  computed: {
    ...mapGetters('products', ['imgURL', 'removedFromBasket']),
    ...mapGetters('ui', ['isUserLoggined']),
    getProduct() {
      if (this.isUserLoggined) {
        const cart = JSON.parse(localStorage.getItem('lorem'))
        console.log('computed propertyyy')
        if (cart) {
          if (this.$route.name === `wishlist___${this.$i18n.locale}`) {
            this.isFavorite = this.product.is_favorite
            this.quantity = this.product.quantity
            if (
              this.quantity === this.product.limit_amount ||
              this.quantity === this.product.amount
            ) {
              this.isDisabled = true
            }
            return this.product
          } else {
            if (cart?.cart || cart?.cart?.length > 0) {
              const findProduct = cart.cart.find(
                (product) => product.id === this.product.id
              )
              if (!findProduct) {
                return this.product
              } else {
                let totalCount = cart.cart.reduce((total, num) => {
                  return total + num.quantity
                }, 0)
                this.quantity = findProduct.quantity
                this.isFavorite = findProduct.is_favorite
                this.$store.commit(
                  'products/SET_PRODUCT_COUNT',
                  totalCount == 0 ? null : totalCount
                )
                if (
                  this.quantity === findProduct.limit_amount ||
                  this.quantity === findProduct.amount
                ) {
                  this.isDisabled = true
                }
                return findProduct
              }
            } else if (cart.wishlist) {
              const findProduct = cart.wishlist.find(
                (product) => product.id === this.product.id
              )
              if (!findProduct) {
                return this.product
              } else {
                this.isFavorite = findProduct.is_favorite
                return findProduct
              }
            } else {
              return this.product
            }
          }
        } else {
          return this.product
        }
      } else {
        const isServer = typeof window === 'undefined'
        if (!isServer) {
          const cart = JSON.parse(localStorage.getItem('lorem'))
          if (cart) {
            if (this.$route.name === `wishlist___${this.$i18n.locale}`) {
              this.isFavorite = this.product.is_favorite
              this.quantity = this.product.quantity
              if (
                this.quantity === this.product.limit_amount ||
                this.quantity === this.product.amount
              ) {
                this.isDisabled = true
              }
              return this.product
            } else {
              if (cart?.cart || cart?.cart?.length > 0) {
                const findProduct = cart.cart.find(
                  (product) => product.id === this.product.id
                )
                if (!findProduct) {
                  return this.product
                } else {
                  let totalCount = cart.cart.reduce((total, num) => {
                    return total + num.quantity
                  }, 0)
                  this.quantity = findProduct.quantity
                  this.isFavorite = findProduct.is_favorite
                  this.$store.commit(
                    'products/SET_PRODUCT_COUNT',
                    totalCount == 0 ? null : totalCount
                  )
                  if (
                    this.quantity === findProduct.limit_amount ||
                    this.quantity === findProduct.amount
                  ) {
                    this.isDisabled = true
                  }
                  return findProduct
                }
              } else if (cart.wishlist) {
                const findProduct = cart.wishlist.find(
                  (product) => product.id === this.product.id
                )
                if (!findProduct) {
                  return this.product
                } else {
                  this.isFavorite = findProduct.is_favorite
                  return findProduct
                }
              } else {
                return this.product
              }
            }
          } else {
            return this.product
          }
        }
      }
    },
  },
  methods: {
    async productLike(data) {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      const array = []
      this.isFavorite = !this.isFavorite
      if (this.$route.name === `wishlist___${this.$i18n.locale}`) {
        if (this.isFavorite === false) {
          this.$emit('removeFromWishlist', data)
        }
      }
      console.log(this.isFavorite)
      if (this.isFavorite) {
        console.log('deken')

        this.$store.commit('products/SET_LIKES_COUNT_INCREMENT')
      } else {
        this.$store.commit('products/SET_LIKES_COUNT_DECREMENT')
      }
      if (cart && cart.cart) {
        const findProduct = cart.cart.find((product) => product.id === data.id)
        if (findProduct) {
          this.$store.commit('products/SET_PRODUCT_FAVORITE', {
            data: data,
            isFavorite: this.isFavorite,
          })
          findProduct.is_favorite = this.isFavorite
          if (findProduct.is_favorite === false) {
            if (findProduct.quantity === 0) {
              cart.cart = cart.cart.filter(
                (product) => product.id !== findProduct.id
              )
              localStorage.setItem('lorem', JSON.stringify(cart))
            } else {
              localStorage.setItem('lorem', JSON.stringify(cart))
            }
          } else {
            localStorage.setItem('lorem', JSON.stringify(cart))
          }
        } else {
          if (this.isFavorite) {
            this.$store.commit('products/SET_PRODUCT_FAVORITE', {
              data: data,
              isFavorite: this.isFavorite,
            })
            cart.cart.push(data)
            localStorage.setItem('lorem', JSON.stringify(cart))
          }
        }
      } else {
        this.$store.commit('products/SET_PRODUCT_FAVORITE', {
          data: data,
          isFavorite: this.isFavorite,
        })
        array.push(data)
        localStorage.setItem(
          'lorem',
          JSON.stringify({
            cart: [...array],
          })
        )
      }
      console.log(cart, this.isFavorite, [data.id])
      if (cart && cart?.auth?.accessToken) {
        try {
          const res = await this.$axios.$post(
            `/${this.$i18n.locale}/like?status=${this.isFavorite}`,
            { product_ids: [data.id] },
            {
              headers: {
                Authorization: `Bearer ${cart.auth.accessToken}`,
              },
            }
          )
          console.log(res)
          console.log('productLike', res)
        } catch (error) {
          console.log(error.response)
          if (error?.response?.status === 403) {
            try {
              const { access_token, refresh_token, status } = (
                await getRefreshToken({
                  url: `auth/refresh`,
                  refreshToken: `Bearer ${cart.auth.refreshToken}`,
                })
              ).data
              console.log('new', access_token, refresh_token, status)
              if (status) {
                const lorem = await JSON.parse(localStorage.getItem('lorem'))
                if (lorem) {
                  lorem.auth = {
                    accessToken: access_token,
                    refreshToken: refresh_token,
                  }
                  localStorage.setItem('lorem', JSON.stringify(lorem))
                } else {
                  localStorage.setItem(
                    'lorem',
                    JSON.stringify({
                      auth: {
                        accessToken: access_token,
                        refreshToken: refresh_token,
                      },
                    })
                  )
                }
                try {
                  const response = await this.$axios.$post(
                    `/${this.$i18n.locale}/like?status=${this.isFavorite}`,
                    { product_ids: [data.id] },
                    {
                      headers: {
                        Authorization: `Bearer ${access_token}`,
                      },
                    }
                  )
                  console.log(res)
                  console.log('productLike', response)
                } catch (error) {
                  console.log('productLike1', error)
                }
              }
            } catch (error) {
              console.log('ref', error.response.status)
              if (error.response.status === 403) {
                cart.auth.accessToken = null
                cart.auth.refreshToken = null
                localStorage.setItem('lorem', JSON.stringify(cart))
                this.$router.push(this.localeLocation('/'))
              }
            }
          }
        }
      }
    },
    async basketAdd(data) {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      const array = []
      if (this.isDisabled === true) {
        if (this.count === 0) {
          if (this.quantity === data.limit_amount) {
            this.$toast(`Harydyn satyn alma mukdary ${data.limit_amount} !`)
          } else if (this.quantity === data.amount) {
            this.$toast(`Harydyn stock  mukdary ${data.amount} !`)
          }
        }
        this.count++
      } else {
        if (cart && cart.cart) {
          const findProduct = cart.cart.find(
            (product) => product.id === data.id
          )
          if (findProduct) {
            this.quantity += 1
            this.$store.commit('products/SET_PRODUCT_TOTAL_INCREMENT', {
              data: data,
              quantity: this.quantity,
            })
            findProduct.quantity = this.quantity
            localStorage.setItem('lorem', JSON.stringify(cart))
          } else {
            if (this.removedFromBasket && this.quantity > 0) {
              this.quantity = 0
              setTimeout(() => {
                this.$store.commit('products/SET_REMOVED_FROM_BASKET', false)
              }, 0)
            } else {
              this.quantity += 1
              this.$store.commit('products/SET_PRODUCT_TOTAL_INCREMENT', {
                data: data,
                quantity: this.quantity,
              })
              cart.cart.push(data)
              localStorage.setItem('lorem', JSON.stringify(cart))
            }
          }
        } else {
          this.quantity += 1
          this.$store.commit('products/SET_PRODUCT_TOTAL_INCREMENT', {
            data: data,
            quantity: this.quantity,
          })
          array.push(data)
          localStorage.setItem(
            'lorem',
            JSON.stringify({
              cart: [...array],
            })
          )
        }
        if (cart && cart?.auth?.accessToken) {
          try {
            const res = (
              await productAdd({
                url: `${this.$i18n.locale}/add-cart`,
                data: [
                  {
                    product_id: data.id,
                    quantity_of_product: this.quantity,
                  },
                ],
                accessToken: `Bearer ${cart?.auth?.accessToken}`,
              })
            ).data
            console.log('productAdd', res)
          } catch (error) {
            console.log(error.response)
            if (error?.response?.status === 403) {
              try {
                const { access_token, refresh_token, status } = (
                  await getRefreshToken({
                    url: `auth/refresh`,
                    refreshToken: `Bearer ${cart.auth.refreshToken}`,
                  })
                ).data
                console.log('new', access_token, refresh_token, status)
                if (status) {
                  const lorem = await JSON.parse(localStorage.getItem('lorem'))
                  if (lorem) {
                    lorem.auth = {
                      accessToken: access_token,
                      refreshToken: refresh_token,
                    }
                    localStorage.setItem('lorem', JSON.stringify(lorem))
                  } else {
                    localStorage.setItem(
                      'lorem',
                      JSON.stringify({
                        auth: {
                          accessToken: access_token,
                          refreshToken: refresh_token,
                        },
                      })
                    )
                  }
                  try {
                    const response = (
                      await productAdd({
                        url: `${this.$i18n.locale}/add-cart`,
                        data: [
                          {
                            product_id: data.id,
                            quantity_of_product: this.quantity,
                          },
                        ],
                        accessToken: `Bearer ${access_token}`,
                      })
                    ).data
                    console.log('productAdd1', response)
                  } catch (error) {
                    console.log('productAdd1', error)
                  }
                }
              } catch (error) {
                console.log('ref', error.response.status)
                if (error.response.status === 403) {
                  cart.auth.accessToken = null
                  cart.auth.refreshToken = null
                  localStorage.setItem('lorem', JSON.stringify(cart))
                  this.$router.push(this.localeLocation('/'))
                }
              }
            }
          }
        }
        console.log(this.quantity)
        if (
          this.quantity === data.limit_amount ||
          this.quantity === data.amount
        ) {
          this.isDisabled = true
        }
      }
    },
    async basketRemove(data) {
      this.isDisabled = false
      this.count = 0
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      const findProduct = cart?.cart.find((product) => product.id === data.id)
      if (findProduct) {
        this.quantity -= 1
        this.$store.commit('products/SET_PRODUCT_TOTAL_DECREMENT', {
          data,
          quantity: this.quantity,
        })
        findProduct.quantity = this.quantity
        if (findProduct.quantity === 0) {
          cart.cart = cart?.cart.filter((product) => product.id !== data.id)
          this.$store.commit('products/SET_REMOVED_FROM_BASKET', true)
        }
        localStorage.setItem('lorem', JSON.stringify(cart))
      } else {
        this.quantity = 0
      }
      console.log(this.quantity)
      if (cart && cart?.auth?.accessToken) {
        try {
          const res = (
            await productAdd({
              url: `${this.$i18n.locale}/add-cart`,
              data: [
                {
                  product_id: data.id,
                  quantity_of_product: this.quantity,
                },
              ],
              accessToken: `Bearer ${cart?.auth?.accessToken}`,
            })
          ).data
          console.log('productAdd', res)
        } catch (error) {
          console.log(error.response)
          if (error?.response?.status === 403) {
            try {
              const { access_token, refresh_token, status } = (
                await getRefreshToken({
                  url: `auth/refresh`,
                  refreshToken: `Bearer ${cart.auth.refreshToken}`,
                })
              ).data
              console.log('new', access_token, refresh_token, status)
              if (status) {
                const lorem = await JSON.parse(localStorage.getItem('lorem'))
                if (lorem) {
                  lorem.auth = {
                    accessToken: access_token,
                    refreshToken: refresh_token,
                  }
                  localStorage.setItem('lorem', JSON.stringify(lorem))
                } else {
                  localStorage.setItem(
                    'lorem',
                    JSON.stringify({
                      auth: {
                        accessToken: access_token,
                        refreshToken: refresh_token,
                      },
                    })
                  )
                }
                try {
                  const response = (
                    await productAdd({
                      url: `${this.$i18n.locale}/add-cart`,
                      data: [
                        {
                          product_id: data.id,
                          quantity_of_product: this.quantity,
                        },
                      ],
                      accessToken: `Bearer ${access_token}`,
                    })
                  ).data
                  console.log('productAdd1', response)
                } catch (error) {
                  console.log('productAdd1', error)
                }
              }
            } catch (error) {
              console.log('ref', error.response.status)
              if (error.response.status === 403) {
                cart.auth.accessToken = null
                cart.auth.refreshToken = null
                localStorage.setItem('lorem', JSON.stringify(cart))
                this.$router.push(this.localeLocation('/'))
              }
            }
          }
        }
      }
    },
    openPopUpPoduct() {
      this.isProduct = true
      document.body.classList.add('_lock')
    },
    closePopUpPoduct() {
      this.isProduct = false
      document.body.classList.remove('_lock')
    },
  },
}
</script>
