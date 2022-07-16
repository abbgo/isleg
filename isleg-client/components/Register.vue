<template>
  <div class="pop-up" @click="$emit('closeRegisterPopUp')">
    <div class="pop-up__body" @click.stop>
      <div class="pop-up__wrapper">
        <div class="pop-up__close">
          <img
            src="@/assets/img/close.svg"
            alt=""
            @click="$emit('closeRegister')"
          />
        </div>
        <div class="pop-up_form">
          <div class="form__input">
            <h4>Adyňyz</h4>
            <input type="text" placeholder="Adyňyz" v-model="register.name" />
          </div>
          <div class="form__input">
            <h4>Telefon</h4>
            <input
              type="tel"
              v-model="register.phone_number"
              @input="enforcePhoneFormat"
            />
          </div>
          <div class="form__input">
            <h4>Açar sözi</h4>
            <input
              type="text"
              placeholder="Açar sözi"
              v-model="register.password"
            />
          </div>
          <div class="pop-up__btns">
            <button class="left_btn" @click="$emit('openSignUpPopUp')">
              Ulgama girmek
            </button>
            <button @click="$emit('registerPost')" class="right__btn">
              Agza bolmak
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    register: {
      type: Object,
      default: () => {},
    },
  },
  methods: {
    enforcePhoneFormat() {
      let x = this.register.phone_number
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)

      this.register.phone_number = !x[2]
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
