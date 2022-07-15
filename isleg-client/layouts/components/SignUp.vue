<template>
  <div class="pop-up" @click="$emit('closeSignUpPopUp')">
    <div class="pop-up__body" @click.stop>
      <div class="pop-up__wrapper">
        <div class="pop-up__close" @click="$emit('closeSignUp')">
          <img
            src="@/assets/img/close.svg"
            alt=""
            @click="$emit('closeSignUp')"
          />
        </div>
        <div class="pop-up_form">
          <div class="form__input">
            <h4>Telefon</h4>
            <input
              type="tel"
              v-model="signUp.phone_number"
              @input="enforcePhoneFormat"
            />
          </div>
          <div class="form__input">
            <h4>Açar sözi</h4>
            <input
              type="text"
              v-model="signUp.password"
              placeholder="Açar sözi"
            />
          </div>
          <div class="form-chek__password">
            <p>Açar sözi unutdym</p>
          </div>
          <div class="pop-up__btns">
            <button class="left_btn" @click="$emit('openRegister')">
              Agza bolmak
            </button>
            <button class="right__btn">Ulgama girmek</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    signUp: {
      type: Object,
      default: () => {},
    },
  },
  methods: {
    enforcePhoneFormat() {
      let x = this.signUp.phone_number
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
      this.signUp.newPhoneVal = this.signUp.phone_number = !x[2]
        ? '+993 6'
        : '+993 6' +
          (x[3] ? x[3] : '') +
          (x[4] ? ' ' + x[4] : '') +
          (x[5] ? '-' + x[5] : '') +
          (x[6] ? '-' + x[6] : '')
    },
  },
}
</script>
