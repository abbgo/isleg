<template>
  <div :class="['pop-up', 'pt-20', { active: isPayment }]">
    <div class="pop-up__body pop-up__product-body" style="width: 900px">
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
                <div
                  class="payment__chekbox-input"
                  v-for="item in payment.paymentForm"
                  :key="item.id"
                >
                  <input
                    :checked="item.checked"
                    class="top__input"
                    name="top"
                    :id="item.depends"
                    type="radio"
                    @change="paymentChecked(item)"
                  />
                  <label :for="item.depends">{{ item.name }}</label>
                </div>
                <span class="error" v-if="isPaymentForm">
                  {{ $t('payment.paymentForm') }}
                </span>
              </div>
              <div class="payment__chekbox">
                <div class="payment__chekbox-title">
                  <h3>Eltip bermek wagtyny saylan</h3>
                </div>
                <div class="payment__chekbox-subtitle">
                  <h4>Ertir(19.06.2022)</h4>
                </div>
                <div
                  class="payment__chekbox-input"
                  v-for="item in payment.theDeliveryTime"
                  :key="item.id"
                >
                  <input
                    :checked="item.checked"
                    class="bottom__input"
                    name="bottom"
                    :id="item.depends"
                    type="radio"
                    @change="theDeliveryTimeChecked(item)"
                  />
                  <label :for="item.depends">{{ item.time }}</label>
                </div>
                <span class="error" v-if="isTheDeliveryTime">
                  {{ $t('payment.theDeliveryTime') }}
                </span>
              </div>
            </div>
            <div class="payment__form">
              <div class="payment__form-box">
                <label for="">Doly Adynyz <span>*</span></label>
                <input
                  type="text"
                  v-model.trim="$v.payment.fullName.$model"
                  placeholder="Doly Adynyz"
                />
                <span
                  class="error"
                  v-if="
                    $v.payment.fullName.$error && !$v.payment.fullName.required
                  "
                >
                  {{ $t('register.nameIsRequired') }}
                </span>
                <span class="error" v-if="!$v.payment.fullName.minLength">
                  {{ $t('register.nameMustHavetletters') }}</span
                >
              </div>
              <div class="payment__form-box">
                <label for="">Telefon <span>*</span></label>
                <input
                  type="tel"
                  v-model="$v.payment.phone_number.$model"
                  @input="enforcePhoneFormat"
                />
                <span class="error" v-if="isPhoneNumber">
                  {{ $t('register.phoneNumberIsRequired') }}
                </span>
              </div>
              <div class="payment__form-box">
                <label for="">Salgynyz <span>*</span></label>
                <input
                  type="text"
                  v-model.trim="$v.payment.address.$model"
                  placeholder="Salgynyz"
                />
                <span
                  class="error"
                  v-if="
                    $v.payment.address.$error && !$v.payment.address.required
                  "
                >
                  {{ $t('payment.address') }}
                </span>
              </div>
              <div class="payment__form-box">
                <label for="">Bellik</label>
                <input type="text" placeholder="Bellik" />
              </div>
              <div class="payment__form-btn">
                <button class="disable__btn" disabled>Sayta agza bol</button>
                <button class="order__btn" @click="order">Sargyt et</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators'
export default {
  props: {
    isPayment: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      isPhoneNumber: false,
      isPaymentForm: false,
      isTheDeliveryTime: false,
      payment: {
        fullName: '',
        phone_number: '+993 6',
        address: '',
        note: '',
        paymentForm: [
          { id: 1, checked: false, depends: 'nagt', name: 'Nagt' },
          { id: 2, checked: false, depends: 'toleg', name: 'Toleg terminaly' },
          { id: 3, checked: false, depends: 'qr', name: 'QR kod (rysgal pay)' },
        ],
        theDeliveryTime: [
          { id: 4, checked: false, depends: 'afternoon', time: '12.00-15.00' },
          { id: 5, checked: false, depends: 'after', time: '15.00-18.00' },
          { id: 6, checked: false, depends: 'evening', time: '18.00-21.00' },
        ],
      },
    }
  },
  validations: {
    payment: {
      fullName: {
        required,
        minLength: minLength(4),
      },
      phone_number: {
        required,
      },
      address: {
        required,
      },
      note: {
        required,
      },
    },
  },
  mounted() {
    console.log(this.$v)
  },
  methods: {
    enforcePhoneFormat() {
      this.isPhoneNumber = false
      let x = this.payment.phone_number
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
      if (!x[2]) {
        this.payment.phone_number = '+993 6'
      } else {
        this.payment.phone_number =
          '+993 6' +
          (x[3] ? x[3] : '') +
          (x[4] ? ' ' + x[4] : '') +
          (x[5] ? '-' + x[5] : '') +
          (x[6] ? '-' + x[6] : '')
      }
    },
    paymentChecked(payload) {
      const findItem = this.payment.paymentForm.find(
        (item) => item.checked == true
      )
      if (findItem) {
        findItem.checked = false
      }
      payload.checked = true
      this.isPaymentForm = false
      console.log(this.payment.paymentForm)
    },
    theDeliveryTimeChecked(payload) {
      const findItem = this.payment.paymentForm.find(
        (item) => item.checked == true
      )
      if (findItem) {
        findItem.checked = false
      }
      payload.checked = true
      this.isTheDeliveryTime = false
      console.log(this.payment.theDeliveryTime)
    },
    order: async function () {
      this.$v.$touch()
      console.log(this.payment.paymentForm)
      const paymentForm = this.payment.paymentForm.filter(
        (pay) => pay.checked == true
      )
      const theDeliveryTime = this.payment.theDeliveryTime.filter(
        (pay) => pay.checked == true
      )
      console.log('paymentForm', paymentForm)
      console.log('theDeliveryTime', theDeliveryTime)
      if (paymentForm.length == 0) {
        this.isPaymentForm = true
      }
      if (theDeliveryTime.length == 0) {
        this.isTheDeliveryTime = true
      }
      if (this.$v.$invalid) {
        if (this.payment.phone_number.length < 16) {
          this.isPhoneNumber = true
        } else {
          this.isPhoneNumber = false
        }
      } else {
        if (this.payment.phone_number.length >= 16) {
          // try {
          //   const res = await this.$axios.$post('/api/login/user/register', {
          //     username: this.$v.userRegister.fullName.$model,
          //     email: this.$v.userRegister.email.$model,
          //     password: this.$v.userRegister.password.$model,
          //   })
          //   if (res.status === 201) {
          //     await this.$auth.loginWith('userLogin', {
          //       data: {
          //         email: this.$v.userRegister.email.$model,
          //         password: this.$v.userRegister.password.$model,
          //       },
          //     })
          //     this.$router.push(this.localeLocation('/user-profile/trades'))
          //   }
          // } catch (e) {
          //   this.$toast(this.$t('register.error'))
          // }
        } else {
          if (this.payment.phone_number.length < 16) {
            this.isPhoneNumber = true
          }
        }
      }
    },
  },
}
</script>
