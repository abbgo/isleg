<template>
  <div :class="['pop-up', { active: isOpenSignUp }]" @click="closeSignUp">
    <div class="pop-up__body" @click.stop>
      <div class="pop-up__wrapper">
        <div class="pop-up__close" @click="closeSignUp">
          <img src="@/assets/img/close.svg" alt="" />
        </div>
        <div class="pop-up_form">
          <div class="form__input">
            <h4>Telefon</h4>
            <div class="form__input-container">
              <img src="@/assets/img/tel.svg" alt="" />
              <input
                type="tel"
                v-model="$v.signUp.phone_number.$model"
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
                v-model="$v.signUp.password.$model"
              />
              <img
                @click="showPass = !showPass"
                :src="showPass ? '/img/Hide.svg' : '/img/Show.svg'"
                alt=""
              />
            </div>
            <span
              class="error"
              v-if="$v.signUp.password.$error && !$v.signUp.password.required"
            >
              {{ $t('register.passwordIsRequired') }}
            </span>
            <span class="error" v-if="!$v.signUp.password.minLength">
              {{ $t('register.passwordMustHavetletters') }}
            </span>
          </div>
          <div class="form-chek__password">
            <p>Açar sözini unutdym</p>
          </div>
          <div class="pop-up__btns">
            <button class="left_btn" @click="openRegister">Agza bolmak</button>
            <button type="button" class="right__btn" @click="logIn">
              Ulgama girmek
            </button>
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
    isOpenSignUp: {
      type: Boolean,
      default: () => false,
    },
  },
  data() {
    return {
      showPass: true,
      isPhoneNumber: false,
      signUp: {
        phone_number: '+993 6',
        password: '',
      },
    }
  },
  validations: {
    signUp: {
      phone_number: {
        required,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    enforcePhoneFormat() {
      this.isPhoneNumber = false
      let x = this.signUp.phone_number
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
      this.signUp.phone_number = !x[2]
        ? '+993 6'
        : '+993 6' +
          (x[3] ? x[3] : '') +
          (x[4] ? ' ' + x[4] : '') +
          (x[5] ? '-' + x[5] : '') +
          (x[6] ? '-' + x[6] : '')
    },
    logIn: async function () {
      this.$v.$touch()
      if (this.signUp.phone_number.length >= 16) {
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
        this.isPhoneNumber = true
      }
    },
    openRegister() {
      this.clear()
      this.$emit('openRegisterPopUp')
    },
    closeSignUp() {
      this.clear()
      this.$emit('closeSignUpPopUp')
    },
    clear() {
      this.signUp.phone_number = '+993 6'
      this.signUp.password = ''
      this.isPhoneNumber = false
      this.$v.$reset()
    },
  },
}
</script>
