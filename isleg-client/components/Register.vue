<template>
  <div
    :class="['pop-up', 'pb-50', { active: isOpenRegister }]"
    @click="closeRegister"
  >
    <div class="pop-up__body" @click.stop>
      <div class="pop-up__wrapper">
        <div class="pop-up__close">
          <img src="@/assets/img/close.svg" alt="" @click="closeRegister" />
        </div>
        <div class="pop-up_form">
          <div class="form__input">
            <h4>Adyňyz</h4>
            <div class="form__input-container">
              <img src="@/assets/img/account.svg" alt="" />
              <input
                type="text"
                placeholder="Adyňyz"
                v-model.trim="$v.register.name.$model"
              />
            </div>
            <span
              class="error"
              v-if="$v.register.name.$error && !$v.register.name.required"
            >
              {{ $t('register.nameIsRequired') }}
            </span>
            <span class="error" v-if="!$v.register.name.minLength">
              {{ $t('register.nameMustHavetletters') }}</span
            >
          </div>
          <div class="form__input">
            <h4>Email</h4>
            <div class="form__input-container">
              <img src="@/assets/img/lock.svg" alt="" />
              <input
                type="text"
                placeholder="Email"
                v-model.trim="$v.register.email.$model"
              />
            </div>
            <span
              class="error"
              v-if="$v.register.email.$error && !$v.register.email.required"
            >
              {{ $t('register.emailIsRequired') }}
            </span>
            <span class="error" v-if="inValidEmail">
              {{ $t('register.invalidEmail') }}
            </span>
          </div>
          <div class="form__input">
            <h4>Telefon</h4>
            <div class="form__input-container">
              <img src="@/assets/img/tel.svg" alt="" />
              <input
                type="tel"
                v-model="$v.register.phone_number.$model"
                @input="enforcePhoneFormat"
              />
            </div>
            <span class="error" v-if="isPhoneNumber">
              {{ $t('register.phoneNumberIsRequired') }}
            </span>
          </div>
          <div class="form__input">
            <h4>Açar sözi</h4>
            <div class="form__input-container">
              <img src="@/assets/img/lock.svg" alt="" />
              <input
                :type="showPass ? 'password' : 'text'"
                placeholder="Açar sözi"
                v-model="$v.register.password.$model"
              />
              <img
                @click="showPass = !showPass"
                :src="showPass ? '/img/Hide.svg' : '/img/Show.svg'"
                alt=""
              />
            </div>
            <span
              class="error"
              v-if="
                $v.register.password.$error && !$v.register.password.required
              "
            >
              {{ $t('register.passwordIsRequired') }}
            </span>
            <span class="error" v-if="!$v.register.password.minLength">
              {{ $t('register.passwordMustHavetletters') }}
            </span>
          </div>
          <div class="form__input">
            <h4>Açar sözini tassykla</h4>
            <div class="form__input-container">
              <img src="@/assets/img/lock.svg" alt="" />
              <input
                :type="showConfirmPass ? 'password' : 'text'"
                placeholder="Açar sözini tassykla"
                v-model="$v.register.repeatPassword.$model"
              />
              <img
                @click="showConfirmPass = !showConfirmPass"
                :src="showConfirmPass ? '/img/Hide.svg' : '/img/Show.svg'"
                alt=""
              />
            </div>
            <span
              class="error"
              v-if="
                $v.register.repeatPassword.$error &&
                !$v.register.repeatPassword.required
              "
            >
              {{ $t('register.confirmPasswordIsRequired') }}
            </span>
            <span
              class="error"
              v-if="
                confirm &&
                $v.register.password.$model !== '' &&
                $v.register.repeatPassword.$model !== ''
              "
            >
              {{ $t('register.confirmPasswordMustHavetletters') }}
            </span>
          </div>
          <div class="confirm__chekbox">
            <input
              class="confirm__chekbox-input"
              id="confirm"
              type="checkbox"
              v-model="$v.register.checked.$model"
              @change="termsOfService($event, $v.register.checked.$model)"
            />
            <label for="confirm" class="confirm__chekbox-label"
              >Ulanyş Düzgünlerini we Gizlinlik Şertnamasyny okadym we kabul
              edýärin!</label
            >
          </div>
          <span class="error" v-if="isTearmsOfServices">
            {{ $t('register.checkedIsRequired') }}
          </span>
          <div class="pop-up__btns">
            <button class="left_btn" @click="openSignUp">Ulgama girmek</button>
            <button type="button" @click="signUp" class="right__btn">
              Agza bolmak
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { required, sameAs, minLength } from 'vuelidate/lib/validators'
export default {
  props: {
    isOpenRegister: {
      type: Boolean,
      default: () => false,
    },
  },
  data() {
    return {
      show: true,
      showPass: true,
      inValidEmail: false,
      showConfirmPass: true,
      confirm: false,
      isPhoneNumber: false,
      isTearmsOfServices: false,
      register: {
        name: '',
        phone_number: '+993 6',
        email: '',
        password: '',
        repeatPassword: '',
        checked: false,
      },
    }
  },
  validations: {
    register: {
      name: {
        required,
        minLength: minLength(4),
      },
      phone_number: {
        required,
      },
      email: {
        required,
      },
      password: {
        required,
        minLength: minLength(6),
      },
      repeatPassword: {
        required,
        sameAsPassword: sameAs('password'),
      },
      checked: {
        required,
      },
    },
  },
  watch: {
    '$v.register.repeatPassword.$model': function (value) {
      if (value === this.$v.register.password.$model) {
        this.confirm = false
      } else {
        this.confirm = true
      }
    },
    '$v.register.password.$model': function (newVal) {
      if (
        this.$v.register.repeatPassword.sameAsPassword == false &&
        newVal !== this.$v.register.repeatPassword.$model
      ) {
        this.confirm = true
      } else {
        this.confirm = false
      }
    },
    '$v.register.email.$model': function (val) {
      if (val === '') {
        if (this.inValidEmail) {
          this.inValidEmail = false
        }
      }
      if (/^[a-z0-9._-]{2,}@[a-z0-9]{2,}\.[a-z]{2,}$/i.test(val)) {
        if (this.inValidEmail) {
          this.inValidEmail = false
        }
      }
    },
  },
  computed: {
    checkValidate() {
      if (
        /^[a-z0-9._-]{2,}@[a-z0-9]{2,}\.[a-z]{2,}$/i.test(
          this.$v.register.email.$model
        )
      ) {
        return true
      } else {
        return false
      }
    },
  },
  methods: {
    enforcePhoneFormat() {
      this.isPhoneNumber = false
      let x = this.register.phone_number
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
      if (!x[2]) {
        this.register.phone_number = '+993 6'
      } else {
        this.register.phone_number =
          '+993 6' +
          (x[3] ? x[3] : '') +
          (x[4] ? ' ' + x[4] : '') +
          (x[5] ? '-' + x[5] : '') +
          (x[6] ? '-' + x[6] : '')
      }
    },
    termsOfService(e, val) {
      val = !val
      this.isTearmsOfServices = false
      if (!this.register.checked) {
        this.isTearmsOfServices = true
      } else {
        this.isTearmsOfServices = false
      }
    },
    signUp: async function () {
      this.$v.$touch()
      if (!this.register.checked) {
        this.isTearmsOfServices = true
      } else {
        this.isTearmsOfServices = false
      }
      if (this.register.phone_number.length < 16) {
        this.isPhoneNumber = true
      } else {
        this.isPhoneNumber = false
      }
      if (this.$v.$invalid) {
        if (
          this.$v.register.email.$model !== '' &&
          this.checkValidate === false
        ) {
          this.inValidEmail = true
        }

        if (
          this.$v.register.password.$model !==
          this.$v.register.repeatPassword.$model
        ) {
          if (this.confirm === false) {
            this.confirm = true
          }
        }
      } else {
        if (
          this.checkValidate == true &&
          this.register.phone_number.length >= 16 &&
          this.$v.register.password.$model ==
            this.$v.register.repeatPassword.$model
        ) {
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
          if (this.checkValidate == false) {
            this.inValidEmail = true
          }
          if (this.register.phone_number.length < 16) {
            this.isPhoneNumber = true
          }
        }
      }
    },
    openSignUp() {
      this.clear()
      this.$emit('openSignUpPopUp')
    },
    closeRegister() {
      this.clear()
      this.$emit('closeRegisterPopUp')
    },
    clear() {
      this.register.name = ''
      this.register.phone_number = '+993 6'
      this.register.email = ''
      this.register.password = ''
      this.register.repeatPassword = ''
      this.register.checked = false
      this.show = true
      this.showPass = true
      this.inValidEmail = false
      this.showConfirmPass = true
      this.confirm = false
      this.isPhoneNumber = false
      this.isTearmsOfServices = false
      this.$v.$reset()
    },
  },
}
</script>
