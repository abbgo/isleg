<template>
  <section class="agreement __container">
    <h4 class="agreement__title">Aragatnaşyk</h4>
    <div class="communication_form">
      <div class="form__box">
        <span>Doly adyňyz</span>
        <input
          type="text"
          v-model="communicationForm.fullName"
          placeholder="Doly adyňyz"
        />
      </div>
      <div class="form__box">
        <span>Email</span>
        <input
          type="email"
          v-model="communicationForm.email"
          placeholder="Email"
        />
      </div>
      <div class="form__box">
        <span>Telefon</span>
        <input
          type="tel"
          v-model="communicationForm.tel"
          @input="enforcePhoneFormat"
          placeholder="+993"
        />
      </div>
      <div class="form__box">
        <span>Hatynyz</span>
        <textarea
          type="text"
          v-model="communicationForm.text"
          placeholder="Hatynyz"
        />
      </div>
    </div>
    <div class="communication__datas">
      <div class="datas__left">
        <span>Telefon: +993 62 766780 </span>
        <span>imo: +993 62 766780</span>
        <span>E-mail: info@isleg.com</span>
        <span>Instagram: @isleg_com</span>
      </div>
      <div class="datas__right" @click="postCommunication">
        <button class="product__add-btn send__btn">
          <h4>Ugrat</h4>
        </button>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  data() {
    return {
      communicationForm: {
        fullName: '',
        email: '',
        tel: '+993 6',
        text: '',
      },
    }
  },
  computed: {
    checkValidate() {
      if (
        /^[a-z0-9._-]{2,}@[a-z0-9]{2,}\.[a-z]{2,}$/i.test(
          this.communicationForm.email
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
      let x = this.communicationForm.tel
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
      this.communicationForm.tel = !x[2]
        ? '+993 6'
        : '+993 6' +
          (x[3] ? x[3] : '') +
          (x[4] ? ' ' + x[4] : '') +
          (x[5] ? '-' + x[5] : '') +
          (x[6] ? '-' + x[6] : '')
    },
    postCommunication() {
      this.$toast('wkednewj')
    },
  },
}
</script>
