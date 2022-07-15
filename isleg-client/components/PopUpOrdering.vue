<template>
  <div class="pop-up pop-up_product">
    <div class="arriving__body">
      <div class="arriving__header">
        <span
          ><svg
            width="14"
            height="26"
            viewBox="0 0 24 36"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <g>
              <g>
                <path
                  d="M16.8184 9.7038L9.63736 2.4248L2.00036 9.7038"
                  stroke="#FD5E29"
                  stroke-width="3"
                />
                <path
                  d="M9.58789 3.77686L9.58789 18.4769"
                  stroke="#FD5E29"
                  stroke-width="3"
                />
              </g>
              <g>
                <path
                  d="M22.9 25.7957L15.719 33.0747L8.08203 25.7957"
                  stroke="#FD5E29"
                  stroke-width="3"
                />
                <path
                  d="M15.7188 31.9165L15.7188 17.2165"
                  stroke="#FD5E29"
                  stroke-width="3"
                />
              </g>
            </g>
          </svg>
          Tertipleme</span
        >
        <div
          style="cursor: pointer"
          class="arriving__close"
          @click="$emit('close')"
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
      <div class="arriving__radio-container">
        <div class="arriving__radio-input">
          <input
            class="arriving__input"
            id="cheapest"
            type="radio"
            name="arriving"
          />
          <label for="cheapest"> Iň arzandan</label>
        </div>
        <div class="arriving__radio-input">
          <input
            class="arriving__input"
            id="expensive"
            type="radio"
            name="arriving"
          />
          <label for="expensive"> Iň gymmatdan</label>
        </div>
        <div class="arriving__radio-input">
          <input
            class="arriving__input"
            id="theBest"
            type="radio"
            name="arriving"
          />
          <label for="theBest"> Iň gowulary</label>
        </div>
        <div class="arriving__radio-input">
          <input
            class="arriving__input"
            id="new"
            type="radio"
            name="arriving"
          />
          <label for="new"> Täzeleri</label>
        </div>
        <div class="arriving__radio-input">
          <input
            checked
            class="arriving__input"
            id="old"
            type="radio"
            name="arriving"
          />
          <label for="old"> Köneleri</label>
        </div>
      </div>
      <div class="arriving__btn">
        <button>Tertiplemek</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {}
</script>
