<template>
  <div class="pages__container">
    <section class="menu__filter __container">
      <span @click="openOrdering"
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
      >
      <span
        ><svg
          width="18"
          height="30"
          viewBox="0 0 26 32"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M11 7H8C7.73478 7 7.48043 6.89464 7.29289 6.70711C7.10536 6.51957 7 6.26522 7 6C7 5.73478 7.10536 5.48043 7.29289 5.29289C7.48043 5.10536 7.73478 5 8 5H11C11.2652 5 11.5196 5.10536 11.7071 5.29289C11.8946 5.48043 12 5.73478 12 6C12 6.26522 11.8946 6.51957 11.7071 6.70711C11.5196 6.89464 11.2652 7 11 7Z"
            fill="#FD5E29"
          />
          <path
            d="M11 12H8C7.73478 12 7.48043 11.8946 7.29289 11.7071C7.10536 11.5196 7 11.2652 7 11C7 10.7348 7.10536 10.4804 7.29289 10.2929C7.48043 10.1054 7.73478 10 8 10H11C11.2652 10 11.5196 10.1054 11.7071 10.2929C11.8946 10.4804 12 10.7348 12 11C12 11.2652 11.8946 11.5196 11.7071 11.7071C11.5196 11.8946 11.2652 12 11 12Z"
            fill="#FD5E29"
          />
          <path
            d="M11 17H8C7.73478 17 7.48043 16.8946 7.29289 16.7071C7.10536 16.5196 7 16.2652 7 16C7 15.7348 7.10536 15.4804 7.29289 15.2929C7.48043 15.1054 7.73478 15 8 15H11C11.2652 15 11.5196 15.1054 11.7071 15.2929C11.8946 15.4804 12 15.7348 12 16C12 16.2652 11.8946 16.5196 11.7071 16.7071C11.5196 16.8946 11.2652 17 11 17Z"
            fill="#FD5E29"
          />
          <path
            d="M11 22H8C7.73478 22 7.48043 21.8946 7.29289 21.7071C7.10536 21.5196 7 21.2652 7 21C7 20.7348 7.10536 20.4804 7.29289 20.2929C7.48043 20.1054 7.73478 20 8 20H11C11.2652 20 11.5196 20.1054 11.7071 20.2929C11.8946 20.4804 12 20.7348 12 21C12 21.2652 11.8946 21.5196 11.7071 21.7071C11.5196 21.8946 11.2652 22 11 22Z"
            fill="#FD5E29"
          />
          <path
            d="M11 27H8C7.73478 27 7.48043 26.8946 7.29289 26.7071C7.10536 26.5196 7 26.2652 7 26C7 25.7348 7.10536 25.4804 7.29289 25.2929C7.48043 25.1054 7.73478 25 8 25H11C11.2652 25 11.5196 25.1054 11.7071 25.2929C11.8946 25.4804 12 25.7348 12 26C12 26.2652 11.8946 26.5196 11.7071 26.7071C11.5196 26.8946 11.2652 27 11 27Z"
            fill="#FD5E29"
          />
          <path
            d="M9 32H3C2.20435 32 1.44129 31.6839 0.87868 31.1213C0.316071 30.5587 0 29.7956 0 29V25C0 24.7348 0.105357 24.4804 0.292893 24.2929C0.48043 24.1054 0.734784 24 1 24C1.26522 24 1.51957 24.1054 1.70711 24.2929C1.89464 24.4804 2 24.7348 2 25V29C2 29.2652 2.10536 29.5196 2.29289 29.7071C2.48043 29.8946 2.73478 30 3 30H9C9.26522 30 9.51957 29.8946 9.70711 29.7071C9.89464 29.5196 10 29.2652 10 29V3C10 2.73478 9.89464 2.48043 9.70711 2.29289C9.51957 2.10536 9.26522 2 9 2H3C2.73478 2 2.48043 2.10536 2.29289 2.29289C2.10536 2.48043 2 2.73478 2 3V21C2 21.2652 1.89464 21.5196 1.70711 21.7071C1.51957 21.8946 1.26522 22 1 22C0.734784 22 0.48043 21.8946 0.292893 21.7071C0.105357 21.5196 0 21.2652 0 21V3C0 2.20435 0.316071 1.44129 0.87868 0.87868C1.44129 0.31607 2.20435 0 3 0L9 0C9.79565 0 10.5587 0.31607 11.1213 0.87868C11.6839 1.44129 12 2.20435 12 3V29C12 29.7956 11.6839 30.5587 11.1213 31.1213C10.5587 31.6839 9.79565 32 9 32V32Z"
            fill="#FD5E29"
          />
          <path
            d="M22 32C21.8274 32 21.6578 31.9553 21.5077 31.8703C21.3575 31.7853 21.2319 31.6629 21.143 31.515L18.143 26.515C18.0495 26.3595 18.0001 26.1815 18 26V4C18 2.93913 18.4214 1.92172 19.1716 1.17157C19.9217 0.421427 20.9391 0 22 0C23.0609 0 24.0783 0.421427 24.8284 1.17157C25.5786 1.92172 26 2.93913 26 4V6C26 6.26522 25.8946 6.51957 25.7071 6.70711C25.5196 6.89464 25.2652 7 25 7C24.7348 7 24.4804 6.89464 24.2929 6.70711C24.1054 6.51957 24 6.26522 24 6V4C24 3.46957 23.7893 2.96086 23.4142 2.58579C23.0391 2.21071 22.5304 2 22 2C21.4696 2 20.9609 2.21071 20.5858 2.58579C20.2107 2.96086 20 3.46957 20 4V25.723L22 29.056L24 25.723V10C24 9.73478 24.1054 9.48043 24.2929 9.29289C24.4804 9.10536 24.7348 9 25 9C25.2652 9 25.5196 9.10536 25.7071 9.29289C25.8946 9.48043 26 9.73478 26 10V26C25.9999 26.1815 25.9505 26.3595 25.857 26.515L22.857 31.515C22.7681 31.6629 22.6425 31.7853 22.4923 31.8703C22.3422 31.9553 22.1726 32 22 32Z"
            fill="#FD5E29"
          />
          <path
            d="M25 27H19C18.7348 27 18.4804 26.8946 18.2929 26.7071C18.1054 26.5196 18 26.2652 18 26C18 25.7348 18.1054 25.4804 18.2929 25.2929C18.4804 25.1054 18.7348 25 19 25H25C25.2652 25 25.5196 25.1054 25.7071 25.2929C25.8946 25.4804 26 25.7348 26 26C26 26.2652 25.8946 26.5196 25.7071 26.7071C25.5196 26.8946 25.2652 27 25 27Z"
            fill="#FD5E29"
          />
          <path
            d="M25 7H19C18.7348 7 18.4804 6.89464 18.2929 6.70711C18.1054 6.51957 18 6.26522 18 6C18 5.73478 18.1054 5.48043 18.2929 5.29289C18.4804 5.10536 18.7348 5 19 5H25C25.2652 5 25.5196 5.10536 25.7071 5.29289C25.8946 5.48043 26 5.73478 26 6C26 6.26522 25.8946 6.51957 25.7071 6.70711C25.5196 6.89464 25.2652 7 25 7Z"
            fill="#FD5E29"
          />
        </svg>
        Ýazuw</span
      >
      <span @click="openFilter"
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
      </span>
    </section>
    <section class="product__categoty __container">
      <div class="category__section">
        <products></products>
        <div class="pagination">
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
        </div>
      </div>
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
import noUiSlider from '@/plugins/nouislider.min'
export default {
  components: {
    Products,
  },
  data() {
    return {
      isOrdering: false,
      isFilter: false,
    }
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
  },
}
</script>
