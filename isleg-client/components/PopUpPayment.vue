<template>
  <div class="pop-up pop-up_product">
    <div class="pop-up__product-body" style="width: 900px">
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
        <div class="payment__container">
          <div class="payment__text">
            <p>
              Eltip bermek hyzmaty Aşgabat şäheriniň çägi bilen bir hatarda
              Büzmeýine we Änew şäherine hem elýeterlidir. Hyzmat mugt amala
              aşyrylýar; <br />
              <br />
              Saýtdan sargyt edeniňizden soňra operator size jaň edip sargydy
              tassyklar (eger hemişelik müşderi bolsaňyz sargytlaryňyz
              islegiňize görä awtomatik usulda hem tassyklanýar); <br />
              <br />
              Sargydy barlap alanyňyzdan soňra töleg amala aşyrylýar. Eltip
              berijiniň size gowşurýan töleg resminamasynda siziň tölemeli
              puluňyz bellenendir. Töleg nagt we nagt däl görnüşde milli manatda
              amala aşyrylýar. Kabul edip tölegini geçiren harydyňyz yzyna
              alynmaýar;
            </p>
            <br />
          </div>
          <div class="payment__settings">
            <div class="settings__chekbox-container">
              <div class="payment__chekbox">
                <div class="payment__chekbox-title">
                  <h3>Toleg sekili</h3>
                </div>
                <div class="payment__chekbox-input">
                  <input class="top__input" name="top" id="nagt" type="radio" />
                  <label for="nagt">Nagt</label>
                </div>
                <div class="payment__chekbox-input">
                  <input
                    class="top__input"
                    name="top"
                    id="toleg"
                    type="radio"
                  />
                  <label for="toleg">Toleg terminaly</label>
                </div>
                <div class="payment__chekbox-input">
                  <input class="top__input" name="top" id="qr" type="radio" />
                  <label for="qr">QR kod (rysgal pay)</label>
                </div>
              </div>
              <div class="payment__chekbox">
                <div class="payment__chekbox-title">
                  <h3>Eltip berme wagtyny saylan</h3>
                </div>
                <div class="payment__chekbox-subtitle">
                  <h4>Ertir(19.06.2022)</h4>
                </div>
                <div class="payment__chekbox-input">
                  <input
                    checked
                    class="bottom__input"
                    name="bottom"
                    id="afternoon"
                    type="radio"
                  />
                  <label for="afternoon">12.00-15.00</label>
                </div>
                <div class="payment__chekbox-input">
                  <input
                    class="bottom__input"
                    name="bottom"
                    id="after"
                    type="radio"
                  />
                  <label for="after">15.00-18.00</label>
                </div>
                <div class="payment__chekbox-input">
                  <input
                    class="bottom__input"
                    name="bottom"
                    id="evening"
                    type="radio"
                  />
                  <label for="evening">18.00-21.00</label>
                </div>
              </div>
            </div>
            <div class="payment__form">
              <div class="payment__form-box">
                <label for="">Doly Adynyz <span>*</span></label>
                <input checked type="text" placeholder="Doly Adynyz" />
              </div>
              <div class="payment__form-box">
                <label for="">Telefon <span>*</span></label>
                <input type="number" placeholder="+993" />
              </div>
              <div class="payment__form-box">
                <label for="">Salgynyz <span>*</span></label>
                <input type="text" placeholder="Salgynyz" />
              </div>
              <div class="payment__form-box">
                <label for="">Bellik <span>*</span></label>
                <input type="text" placeholder="Bellik" />
              </div>
              <div class="payment__form-btn">
                <button class="disable__btn" disabled>Sayta agza bol</button>
                <button class="order__btn">Sargyt et</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {}
</script>

<style lang="scss" scoped></style>
