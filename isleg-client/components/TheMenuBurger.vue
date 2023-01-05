<template>
  <div class="menu__burger" @click.stop="close">
    <div class="menu__burger-wrapper" @click.stop>
      <div class="menu__burger-overlay">
        <div class="menu__burger-header">
          <span
            class="menu__burger-chevron"
            v-if="subCategoryId"
            @click.stop="comeBack"
          >
            <img src="@/assets/img/chevron-mobile.svg" alt=""
          /></span>
          <h2>isleg</h2>
          <span class="menu__burger-close" @click="close"></span>
        </div>
        <div :class="['menu__burger-items', { active: subCategoryId }]">
          <ul class="menu__burger-list">
            <li
              v-for="category in categories"
              :key="category.id"
              :class="[
                'menu__burger-item',
                { active: subCategoryId == category.id },
              ]"
              @click.stop="subCategory(category.id, category.child_category)"
            >
              <div class="menu__burger-item-img" v-if="category.child_category">
                <img :src="`${imgURL}/${category.image}`" alt="" />
              </div>
              <h3 class="menu__burger-item-title">{{ category.name }}</h3>
              <span class="menu__burger-item-arrow">
                <img src="@/assets/img/chevron-right.svg" alt="" />
              </span>
              <div class="menu__burger-subitems">
                <ul class="menu__burger-sublist">
                  <li
                    v-for="subCategory in category.child_category"
                    :key="subCategory.id"
                    :class="[
                      'menu__burger-subitem',
                      { active: nestedSubCategoryId == subCategory.id },
                    ]"
                    @click.stop="
                      nestedSubCategory(
                        subCategory.id,
                        subCategory.child_category
                      )
                    "
                  >
                    <h3>{{ subCategory.name }}</h3>
                    <span
                      class="menu__burger-item-arrow"
                      v-if="subCategory.child_category"
                    >
                      <img src="@/assets/img/chevron-right.svg" alt="" />
                    </span>
                    <div class="menu__burger-nesteditems">
                      <ul class="menu__burger-nestedlist">
                        <li
                          v-for="nestedSubCategory in subCategory.child_category"
                          :key="nestedSubCategory.id"
                          class="menu__burger-nesteditem"
                          @click.stop="nestedCategory(nestedSubCategory.id)"
                        >
                          <h3>{{ nestedSubCategory.name }}</h3>
                        </li>
                      </ul>
                    </div>
                  </li>
                </ul>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
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
  data() {
    return {
      subCategoryId: null,
      nestedSubCategoryId: null,
    }
  },
  mounted() {
    console.log(this.categories)
  },
  methods: {
    subCategory(id, child) {
      document.body.classList.add('_lock')
      this.$router.push(this.localeLocation(`/category/${id}`))
      if (child) {
        this.subCategoryId = id
      } else {
        this.close()
      }
    },
    nestedSubCategory(id, child) {
      this.$router.push(this.localeLocation(`/category/${id}`))
      if (child) {
        this.nestedSubCategoryId = id
      } else {
        this.close()
      }
    },
    nestedCategory(id) {
      this.$router.push(this.localeLocation(`/category/${id}`))
      this.close()
    },
    comeBack() {
      if (this.nestedSubCategoryId) {
        this.nestedSubCategoryId = null
        this.$router.go(-1)
      } else {
        this.subCategoryId = null
        this.$router.go(-1)
      }
    },
    close() {
      this.$emit('close')
      document.body.classList.remove('_lock')
    },
  },
}
</script>

<style lang="scss" scoped>
.menu__burger {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  background: rgba(0, 0, 0, 0.3);
  &-wrapper {
    width: 80%;
    height: 100%;
    background: #fff;
    overflow: hidden;
  }
  &-overlay {
    height: 100%;
    overflow: hidden;
  }
  &-header {
    display: flex;
    align-items: center;
    padding: 25px 10px;
    background: #f7f7f7;

    h2 {
      flex: 1 1 auto;
      font-family: TTNormsPro-ExtraBold;
      font-size: 28px;
      color: #fd5e29;
      letter-spacing: 3px;
      cursor: pointer;
    }
  }
  &-close {
    width: 22px;
    height: 22px;
    position: relative;
    cursor: pointer;
    &::after {
      position: absolute;
      content: '';
      width: 100%;
      height: 3px;
      background: #fd5e29;
      top: 50%;
      transform: rotate(45deg);
    }
    &::before {
      position: absolute;
      content: '';
      width: 100%;
      height: 3px;
      background: #fd5e29;
      top: 50%;
      transform: rotate(-45deg);
    }
  }
  &-chevron {
    width: 21px;
    height: 21px;
    border-radius: 50%;
    border: 1px solid #fd5e29;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-right: 15px;
    padding-right: 3px;
    cursor: pointer;
  }
  &-items {
    height: calc(100% - 72px);
    overflow-y: auto;
    &.active {
      overflow: hidden;
    }
  }
  &-list {
    position: relative;
  }
  &-item {
    display: flex;
    align-items: center;
    width: 100%;
    padding: 15px 10px;
    border-bottom: 1px solid hsla(0, 0%, 60%, 0.2);
    cursor: pointer;
    &.active {
      .menu__burger-subitems {
        display: block;
      }
    }
    &-img {
      width: 21px;
      height: 21px;
      margin-right: 15px;
      img {
        width: 100%;
        height: 100%;
      }
    }
    &-title {
      flex: 1 1 auto;
      font-family: TTNormsPro;
      font-size: 18px;
      color: #3d3d3d;
    }
  }
  &-subitems {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: calc(100vh - 72px);
    overflow: hidden;
    background: #fff;
    display: none;
  }
  &-sublist {
    overflow-y: auto;
    height: 100%;
    &.active {
      overflow: hidden;
    }
  }
  &-subitem {
    display: flex;
    align-items: center;
    padding: 15px 10px;
    border-bottom: 1px solid hsla(0, 0%, 60%, 0.2);
    // position: relative;
    &.active {
      .menu__burger-nesteditems {
        display: block;
      }
    }
    h3 {
      flex: 1 1 auto;
      font-family: TTNormsPro;
      font-size: 18px;
      color: #3d3d3d;
    }
  }
  &-nesteditems {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: calc(100vh - 72px);
    overflow: hidden;
    z-index: 3;
    background: #fff;
    display: none;
  }
  &-nestedlist {
    overflow-y: auto;
    height: 100%;
  }
  &-nesteditem {
    display: flex;
    align-items: center;
    padding: 15px 10px;
    border-bottom: 1px solid hsla(0, 0%, 60%, 0.2);
    h3 {
      flex: 1 1 auto;
      font-family: TTNormsPro;
      font-size: 18px;
      color: #3d3d3d;
    }
  }
}
</style>
