<template>
  <section class="agreement __container">
    <h4 class="agreement__title">Aragatnaşyk</h4>
    <div class="communication_form">
      <div class="form__box">
        <span>{{ translationsContact.full_name }}</span>
        <input
          type="text"
          :placeholder="translationsContact.full_name"
          v-model="$v.communicationForm.fullName.$model"
        />
        <div
          class="error"
          v-if="
            $v.communicationForm.fullName.$error &&
            !$v.communicationForm.fullName.required
          "
        >
          {{ $t('register.nameIsRequired') }}
        </div>
        <div class="error" v-if="!$v.communicationForm.fullName.minLength">
          {{ $t('register.nameMustHavetletters') }}
        </div>
      </div>
      <div class="form__box">
        <span>{{ translationsContact.email }}</span>
        <input
          type="email"
          v-model="$v.communicationForm.email.$model"
          :placeholder="translationsContact.email"
        />
        <div
          class="error"
          v-if="
            $v.communicationForm.email.$error &&
            !$v.communicationForm.email.required
          "
        >
          {{ $t('register.emailIsRequired') }}
        </div>
        <div class="error" v-if="inValidEmail">
          {{ $t('register.invalidEmail') }}
        </div>
      </div>
      <div class="form__box">
        <span>Telefon</span>
        <input
          type="tel"
          v-model="$v.communicationForm.tel.$model"
          @input="enforcePhoneFormat"
          placeholder="+993"
        />
        <div class="error" v-if="isPhoneNumber">
          {{ $t('register.phoneNumberIsRequired') }}
        </div>
      </div>
      <div class="form__box">
        <span>Hatynyz</span>
        <textarea
          type="text"
          v-model="communicationForm.text"
          placeholder="Hatynyz"
        />
        <div
          class="error"
          v-if="
            $v.communicationForm.text.$error &&
            !$v.communicationForm.text.required
          "
        >
          {{ $t('textIsRequired') }}
        </div>
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
        <button :disabled="disabled" class="product__add-btn send__btn">
          <h4>Ugrat</h4>
        </button>
      </div>
    </div>
  </section>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators'
import { sendMail, translationContact } from '@/api/user.api'

export default {
  data() {
    return {
      disabled: false,
      inValidEmail: false,
      isPhoneNumber: false,
      communicationForm: {
        fullName: '',
        email: '',
        tel: '+9936',
        text: '',
      },
      translationsContact: null,
    }
  },
  validations: {
    communicationForm: {
      fullName: {
        required,
        minLength: minLength(2),
      },
      tel: {
        required,
      },
      email: {
        required,
      },
      text: {
        required,
      },
    },
  },
  watch: {
    '$v.communicationForm.email.$model': function (val) {
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
          this.$v.communicationForm.email.$model
        )
      ) {
        return true
      } else {
        return false
      }
    },
  },
  async fetch() {
    try {
      const { data, status } = await translationContact({
        url: `${this.$i18n.locale}/translation-contact`,
      })
      console.log(data, status)
      if (status === 200) {
        this.translationsContact = data.translation_contact
      }
    } catch (error) {
      console.log(error)
    }
  },
  methods: {
    enforcePhoneFormat() {
      this.isPhoneNumber = false
      let x = this.communicationForm.tel
        .replace(/\D/g, '')
        .match(/(\d{0,3})(\d{0,1})(\d{0,1})(\d{0,2})(\d{0,2})(\d{0,2})/)
      this.communicationForm.tel = !x[2]
        ? '+9936'
        : '+9936' +
          (x[3] ? x[3] : '') +
          (x[4] ? x[4] : '') +
          (x[5] ? x[5] : '') +
          (x[6] ? x[6] : '')
    },
    async postCommunication() {
      this.$v.$touch()
      if (this.communicationForm.tel.length < 12) {
        this.isPhoneNumber = true
      } else {
        this.isPhoneNumber = false
      }
      if (this.$v.$invalid) {
        if (
          this.$v.communicationForm.email.$model !== '' &&
          this.checkValidate === false
        ) {
          this.inValidEmail = true
        }
      } else {
        if (
          this.checkValidate == true &&
          this.communicationForm.tel.length >= 12
        ) {
          this.disabled = true
          try {
            const { data, status } = await sendMail({
              url: `${this.$i18n.locale}/send-mail`,
              data: {
                full_name: this.communicationForm.fullName,
                phone_number: this.communicationForm.tel,
                letter: this.communicationForm.text,
                email: this.communicationForm.email,
              },
            })
            if (status) {
              this.$toast(this.$t('mailSendedSuccess'))
              this.clear()
            }
          } catch (err) {
            console.log(err)
          } finally {
            this.disabled = false
          }
        } else {
          if (this.checkValidate == false) {
            this.inValidEmail = true
          }
          if (this.communicationForm.tel.length < 12) {
            this.isPhoneNumber = true
          }
        }
      }
    },
    clear() {
      this.communicationForm.fullName = ''
      this.communicationForm.tel = '+9936'
      this.communicationForm.email = ''
      this.communicationForm.text = ''
      this.$v.$reset()
    },
  },
}
</script>
