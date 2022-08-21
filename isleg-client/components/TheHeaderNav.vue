<template>
  <nav class="menu-header__nav __container">
    <span @click="scrollLeft" class="nav__arrow nav__arrow-left"
      ><img src="/img/arrow.svg" alt=""
    /></span>
    <div class="menu-header__wrapper">
      <ul class="menu-header__list">
        <li
          class="menu-header__item"
          v-for="category in categories"
          :key="category.id"
        >
          <a href="#" class="menu-header__link"
            ><span><img :src="`${imgURL}/${category.image}`" alt="" /> </span>
            <h4 class="nav__item">{{ category.name }}</h4></a
          >
          <div class="header__sub-menu">
            <div class="header__sub-scroll-content">
              <div class="header__sub-content">
                <div
                  class="sub__category"
                  v-for="subCategory in category.child_category"
                  :key="subCategory.id"
                >
                  <div class="sub__category-title">{{ subCategory.name }}</div>
                  <div class="sub__category-items">
                    <span
                      v-for="nestedSubCategory in subCategory.child_category"
                      :key="nestedSubCategory.id"
                      >{{ nestedSubCategory.name }}</span
                    >
                  </div>
                </div>
              </div>
            </div>
          </div>
        </li>
      </ul>
    </div>
    <span @click="scrollRight" class="nav__arrow nav__arrow-right"
      ><img src="/img/arrow.svg" alt=""
    /></span>
  </nav>
</template>

<script>
export default {
  props: {
    categories: {
      type: Array,
      default: () => [],
    },
    imgURL: {
      type: String,
      default: () => '',
    },
  },
  mounted() {
    console.log(this.categories)
  },
  methods: {
    scrollLeft() {
      let scrollContainer = document.querySelector('.menu-header__wrapper')
      scrollContainer.scrollLeft -= 150
    },
    scrollRight() {
      let scrollContainer = document.querySelector('.menu-header__wrapper')
      scrollContainer.scrollLeft += 150
    },
  },
}
</script>
