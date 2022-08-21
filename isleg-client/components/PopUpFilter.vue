<template>
  <div :class="['pop-up', { active: isFilter }]" @click="closeFilter">
    <div class="filter__body" @click.stop>
      <div class="arriving__header">
        <div class="filter__title">
          <span>
            <svg
              width="22"
              height="22"
              viewBox="0 0 26 26"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M16.4591 11.3041H9.80908L1.83008 1.99512H24.4381L16.4591 11.3041Z"
                fill="#FD5E29"
              />
              <path
                d="M16.4576 21.2782L9.80859 25.2682V11.3042H16.4576V21.2782Z"
                fill="#FD5E29"
              />
              <path
                d="M24.771 1.995H1.5C1.23747 1.98919 0.987219 1.88265 0.801069 1.69743C0.614919 1.51221 0.507123 1.2625 0.5 1C0.505858 0.736623 0.613095 0.485659 0.799377 0.299377C0.985659 0.113095 1.23662 0.00585781 1.5 0H24.771C25.0344 0.00585781 25.2853 0.113095 25.4716 0.299377C25.6579 0.485659 25.7651 0.736623 25.771 1C25.7639 1.2625 25.6561 1.51221 25.4699 1.69743C25.2838 1.88265 25.0335 1.98919 24.771 1.995Z"
                fill="#FD5E29"
              />
            </svg>
          </span>
          <span>Filter</span>
        </div>
        <div
          style="cursor: pointer"
          class="arriving__close"
          @click="closeFilter"
        >
          <svg
            width="40"
            height="40"
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
      </div>
      <div class="arriving__line"></div>
      <div class="filter__range">
        <div class="range__title">Baha</div>
        <div class="range__slider"></div>
        <div class="range__inputs">
          <div class="input__container">
            <label>TMT</label>
            <input id="input0" type="number" readonly placeholder="0" />
          </div>
          <div class="input__container">
            <label>TMT</label>
            <input id="input1" type="number" readonly placeholder="3000" />
          </div>
        </div>
      </div>
      <div class="filter__switch">
        <input class="switcher__input" type="checkbox" />
        <span class="switcher"></span>
        <label>Diňe arzanladyşdaky</label>
      </div>
      <div class="range__title">Brendlar</div>
      <div class="range__company" :class="{ active: brend }">
        <span @click="brend = !brend" class="brend__active"
          ><h4>Brends</h4>
          <img src="/img/select.svg" alt="" />
        </span>
        <div class="brend__items">
          <span class="brend__search"
            ><input type="text" placeholder="Gozle" />
            <img src="/img/search.svg" alt="" />
          </span>
          <div class="brend__container">
            <span class="brend">Dolhin</span>
            <span class="brend">Dolhin</span>
            <span class="brend">Dolhin</span>
            <span class="brend">Dolhin</span>
            <span class="brend">Dolhin</span>
            <span class="brend">Dolhin</span>
            <span class="brend">Dolhin</span>
          </div>
        </div>
      </div>
      <div class="filter__btn" @click="$emit('filterPost')">
        <button>Filter etmek</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    isFilter: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      brend: false,
    }
  },
  methods: {
    closeFilter() {
      this.brend = false
      this.$emit('close')
    },
  },
}
</script>
