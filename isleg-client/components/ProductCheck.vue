<template>
  <div class="product__chek">
    <div class="product__chek-img">
      <img
        :data-src="`${imgURL}/${
          getBasketProduct &&
          getBasketProduct.main_image &&
          getBasketProduct.main_image.medium
        }`"
        loading="lazy"
        alt="isleg"
      />
    </div>
    <div class="product__chek-text">
      <span> {{ translationProductName(getBasketProduct.translations) }}</span>
      <div class="new__price product__chek-new-price">
        {{ getBasketProduct && getBasketProduct.price }}
        TMT
      </div>
      <div class="chek__count">
        <button @click="fromBasketRemove(getBasketProduct)">
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
        <p>{{ getBasketProduct && getBasketProduct.quantity }}</p>
        <button @click="fromBasketAdd(getBasketProduct)">
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
    </div>
    <div class="product__chek-close">
      <div class="close" @click="$emit('popUpSureOpen', getBasketProduct)">
        <img src="@/assets/img/close.svg" alt="" />
      </div>
      <span v-if="getBasketProduct && getBasketProduct.price"
        >{{
          parseFloat(
            getBasketProduct.price * getBasketProduct.quantity
          ).toFixed(2)
        }}
        TMT</span
      >
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import observer from '@/mixins/observer'
import translation from '@/mixins/translation'
import { productAdd, getRefreshToken } from '@/api/user.api'

export default {
  mixins: [observer, translation],
  props: {
    basketProduct: {
      type: Object,
      default: () => {},
    },
  },
  data() {
    return {
      basketProductQuantity: 0,
      isDisabled: false,
      count: 0,
    }
  },
  computed: {
    ...mapGetters('card', ['imgURL', 'productTotal']),
    getBasketProduct() {
      this.basketProductQuantity = this.basketProduct.quantity
      if (
        this.basketProductQuantity === this.basketProduct.limit_amount ||
        this.basketProductQuantity === this.basketProduct.amount
      ) {
        this.isDisabled = true
      }
      return this.basketProduct
    },
  },
  methods: {
    async fromBasketAdd(data) {
      const cart = JSON.parse(localStorage.getItem('lorem'))
      const array = []
      if (this.isDisabled) {
        if (this.count === 0) {
          if (this.basketProductQuantity === data.limit_amount) {
            this.$toast(`Harydyn satyn alma mukdary ${data.limit_amount} !`)
          } else if (this.basketProductQuantity === data.amount) {
            this.$toast(`Harydyn stock  mukdary ${data.amount} !`)
          }
        }
        this.count++
      } else {
        this.basketProductQuantity += 1
        this.$store.commit('products/SET_PRODUCT_TOTAL_INCREMENT', {
          data: data,
          quantity: this.basketProductQuantity,
        })
        if (cart) {
          const findProduct = cart.cart?.find(
            (product) => product.id === data.id
          )
          if (findProduct) {
            findProduct.quantity = this.basketProductQuantity
            localStorage.setItem('lorem', JSON.stringify(cart))
          } else {
            cart.cart?.push(data)
            localStorage.setItem('lorem', JSON.stringify(cart))
          }
        } else {
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
                    quantity_of_product: this.basketProductQuantity,
                  },
                ],
                accessToken: `Bearer ${cart?.auth?.accessToken}`,
              })
            ).data
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
                            quantity_of_product: this.basketProductQuantity,
                          },
                        ],
                        accessToken: `Bearer ${access_token}`,
                      })
                    ).data
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
        if (
          this.basketProductQuantity === data.limit_amount ||
          this.basketProductQuantity === data.amount
        ) {
          this.isDisabled = true
        }
      }
    },
    async fromBasketRemove(data) {
      const cart = JSON.parse(localStorage.getItem('lorem'))
      this.isDisabled = false
      this.count = 0
      if (this.basketProductQuantity === 1) {
        this.$emit('popUpSureOpen', data)
      } else {
        this.basketProductQuantity -= 1
        this.$store.commit('products/SET_PRODUCT_TOTAL_DECREMENT', {
          data,
          quantity: this.basketProductQuantity,
        })
        const findProduct = cart.cart.find((product) => product.id === data.id)
        findProduct.quantity = this.basketProductQuantity
        localStorage.setItem('lorem', JSON.stringify(cart))
        if (cart && cart?.auth?.accessToken) {
          try {
            const res = (
              await productAdd({
                url: `${this.$i18n.locale}/add-cart`,
                data: [
                  {
                    product_id: data.id,
                    quantity_of_product: this.basketProductQuantity,
                  },
                ],
                accessToken: `Bearer ${cart?.auth?.accessToken}`,
              })
            ).data
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
                            quantity_of_product: this.basketProductQuantity,
                          },
                        ],
                        accessToken: `Bearer ${access_token}`,
                      })
                    ).data
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
      }
    },
  },
}
</script>
