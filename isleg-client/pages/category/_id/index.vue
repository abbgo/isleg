<template>
  <div class="__container">
    <section class="menu__filter">
      <!-- <span @click="openOrdering"
        ><svg
          width="18"
          height="30"
          viewBox="0 0 24 36"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <g opacity="0.8">
            <g opacity="0.8">
              <path
                opacity="0.8"
                d="M16.8184 9.7038L9.63736 2.4248L2.00036 9.7038"
                stroke="#FD5E29"
                stroke-width="3"
              />
              <path
                opacity="0.8"
                d="M9.58789 3.77686L9.58789 18.4769"
                stroke="#FD5E29"
                stroke-width="3"
              />
            </g>
            <g opacity="0.8">
              <path
                opacity="0.8"
                d="M22.9 25.7957L15.719 33.0747L8.08203 25.7957"
                stroke="#FD5E29"
                stroke-width="3"
              />
              <path
                opacity="0.8"
                d="M15.7188 31.9165L15.7188 17.2165"
                stroke="#FD5E29"
                stroke-width="3"
              />
            </g>
          </g>
        </svg>
        Tertipleme</span
      > -->
      <span
        ><img
          v-if="categoryProducts && categoryProducts.image"
          style="margin-right: 8px; width: 30px; height: 30px"
          :src="`${imgURL}/${categoryProducts && categoryProducts.image}`"
          alt="isleg"
        />
        {{ categoryProducts && categoryProducts.name }}</span
      >
      <!-- <span @click="openFilter"
        >Filter
        <svg
          width="22"
          height="22"
          viewBox="0 0 27 26"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <g opacity="0.8">
            <path
              opacity="0.8"
              d="M16.6915 11.3041H10.0415L2.0625 1.99512H24.6705L16.6915 11.3041Z"
              fill="#FD5E29"
            />
            <g opacity="0.8">
              <path
                opacity="0.9"
                d="M16.69 21.2782L10.041 25.2682V11.3042H16.69V21.2782Z"
                fill="#FD5E29"
              />
              <path
                opacity="0.9"
                d="M25.0034 1.995H1.73242C1.46989 1.98919 1.21964 1.88265 1.03349 1.69743C0.847341 1.51221 0.739545 1.2625 0.732422 1C0.73828 0.736623 0.845517 0.485659 1.0318 0.299377C1.21808 0.113095 1.46904 0.00585781 1.73242 0H25.0034C25.2668 0.00585781 25.5178 0.113095 25.704 0.299377C25.8903 0.485659 25.9976 0.736623 26.0034 1C25.9963 1.2625 25.8885 1.51221 25.7024 1.69743C25.5162 1.88265 25.266 1.98919 25.0034 1.995Z"
                fill="#FD5E29"
              />
            </g>
          </g>
        </svg>
      </span> -->
    </section>
    <section class="product__categoty">
      <client-only
        ><div class="category__section">
          <Products :products="categoryProducts && categoryProducts.products" />
          <!-- <pagination
            :modelValue="page"
            @clickPage="(pagination) => updatePage(pagination)"
            :pageCount="paginationCount"
          ></pagination> -->
          <!-- <div class="pagination">
          <div class="pagination__box">
            <div class="left__arrows">
              <svg
                width="12"
                height="12"
                viewBox="0 0 17 17"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M6.12267 15.4759L1.51367 8.1059L6.12267 0.899902"
                  stroke="#8D98A9"
                  stroke-width="2"
                />
                <path
                  d="M15.8883 15.4759L11.2793 8.1059L15.8883 0.899902"
                  stroke="#8D98A9"
                  stroke-width="2"
                />
              </svg>
            </div>
            <div class="pagination__numbers">
              <span class="active__pagination">1</span>
              <span>2</span>
              <span>3</span>
              <span>4</span>
              <span>5</span>
              <span>6</span>
            </div>
            <div class="right__arrows">
              <svg
                width="12"
                height="12"
                viewBox="0 0 17 17"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M6.12267 15.4759L1.51367 8.1059L6.12267 0.899902"
                  stroke="#8D98A9"
                  stroke-width="2"
                />
                <path
                  d="M15.8883 15.4759L11.2793 8.1059L15.8883 0.899902"
                  stroke="#8D98A9"
                  stroke-width="2"
                />
              </svg>
            </div>
          </div>
        </div> -->
        </div></client-only
      >
    </section>
    <pop-up-ordering
      :isOrdering="isOrdering"
      @close="closeOrdering"
    ></pop-up-ordering>
    <pop-up-filter
      :isFilter="isFilter"
      @filterPost="filterPost"
      @close="closeFilter"
    ></pop-up-filter>
  </div>
</template>

<script>
import Products from '@/components/app/Products.vue'
import Pagination from '@/components/app/Pagination.vue'
import noUiSlider from '@/plugins/nouislider.min'
import { getCategoryProducts } from '@/api/categories.api'
import { mapGetters } from 'vuex'
export default {
  components: {
    Products,
    Pagination,
  },
  data() {
    return {
      isOrdering: false,
      isFilter: false,
      limit: 20,
      page: 1,
      paginationCount: 0,
      categoryProducts: null,
    }
  },
  async fetch() {
    await this.fetchCategoryProducts()
  },
  computed: {
    ...mapGetters('ui', ['imgURL']),
  },
  mounted() {
    let rangeSlider = document.querySelector('.range__slider')
    if (rangeSlider) {
      noUiSlider.create(rangeSlider, {
        start: [0, 3000],
        connect: true,
        step: 1,
        range: {
          min: 0,
          max: 3000,
        },
      })
      let input0 = document.getElementById('input0')
      let input1 = document.getElementById('input1')
      let inputs = [input0, input1]

      rangeSlider.noUiSlider.on('update', function (value, handle) {
        inputs[handle].value = Math.round(value[handle])
      })
    }
  },
  methods: {
    async fetchCategoryProducts() {
      try {
        const res = (
          await getCategoryProducts({
            url: `${this.$i18n.locale}/category/${this.$route.params?.id}/${this.limit}/${this.page}`,
          })
        ).data
        if (res.status) {
          this.paginationCount = Math.ceil(res.count_of_products / this.limit)
          this.categoryProducts = res.category
        }
      } catch (e) {
        console.log(e)
      }
    },
    openOrdering() {
      this.isOrdering = true
      document.body.classList.add('_lock')
    },
    openFilter() {
      this.isFilter = true
      document.body.classList.add('_lock')
    },
    closeOrdering() {
      this.isOrdering = false
      document.body.classList.remove('_lock')
    },
    closeFilter() {
      this.isFilter = false
      document.body.classList.remove('_lock')
    },
    filterPost() {
      let input0 = document.getElementById('input0')
      let input1 = document.getElementById('input1')
      console.log('input0', input0.value)
      console.log('input1', input1.value)
    },
    async updatePage(p) {
      this.page = p
      await this.fetchCategoryProducts()
    },
  },
}
</script>
