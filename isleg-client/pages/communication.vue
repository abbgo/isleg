<template>
  <div class="__container">
    <div class="communication">
      <div class="communication__wrapper">
        <div class="communication__form">
          <div class="communication__form-title">
            <h4>{{ contact }}</h4>
          </div>
          <div class="communication__form-control">
            <div
              :class="[
                'control__input',
                {
                  error: !communicationForm.fullName && !error,
                },
              ]"
            >
              <div class="control__input-label">
                <span>{{
                  translationsContact && translationsContact.full_name
                }}</span>
              </div>
              <div class="control__input-input">
                <input
                  type="text"
                  :placeholder="
                    translationsContact && translationsContact.full_name
                  "
                  v-model="communicationForm.fullName"
                />
              </div>
            </div>
            <div
              :class="[
                'control__input',
                {
                  error: !inValidEmail,
                },
              ]"
            >
              <div class="control__input-label">
                <span>{{
                  translationsContact && translationsContact.email
                }}</span>
              </div>
              <div class="control__input-input">
                <input
                  type="email"
                  v-model="communicationForm.email"
                  :placeholder="
                    translationsContact && translationsContact.email
                  "
                />
              </div>
            </div>
            <div
              :class="[
                'control__input',
                {
                  error: !isPhoneNumber,
                },
              ]"
            >
              <div class="control__input-label">
                <span>{{
                  translationsContact && translationsContact.phone
                }}</span>
              </div>
              <div class="control__input-input">
                <input
                  type="tel"
                  v-model="communicationForm.tel"
                  @input="enforcePhoneFormat"
                  placeholder="+993"
                />
              </div>
            </div>
            <div
              :class="[
                'control__input',
                {
                  error: !communicationForm.text && !error,
                },
              ]"
            >
              <div class="control__input-label">
                <span>{{
                  translationsContact && translationsContact.letter
                }}</span>
              </div>
              <div class="control__input-input">
                <input
                  type="text"
                  v-model="communicationForm.text"
                  :placeholder="
                    translationsContact && translationsContact.letter
                  "
                />
              </div>
            </div>
          </div>
          <div class="communication__info">
            <div class="communication__info-text">
              <span
                >{{ translationsContact && translationsContact.company_phone }}:
                <span
                  v-for="(phone, i) in companyInformations &&
                  companyInformations.company_phones"
                  :key="i"
                  >{{ phone.split(', ')[i] }}</span
                >
              </span>
              <span
                >{{ translationsContact && translationsContact.imo }}:
                {{
                  companyInformations &&
                  companyInformations.company_setting &&
                  companyInformations.company_setting.imo
                }}</span
              >
              <span
                >{{ translationsContact && translationsContact.company_email }}:
                {{
                  companyInformations &&
                  companyInformations.company_setting &&
                  companyInformations.company_setting.email
                }}</span
              >
              <span
                >{{ translationsContact && translationsContact.instagram }}:
                {{
                  companyInformations &&
                  companyInformations.company_setting &&
                  companyInformations.company_setting.instagram
                }}</span
              >
            </div>
            <div class="communication__info-button" @click="postCommunication">
              <button :disabled="disabled">
                {{ translationsContact && translationsContact.button_text }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { sendMail, translationContact, companyPhones } from '@/api/user.api'
import { mapGetters } from 'vuex'
export default {
  data() {
    return {
      disabled: false,
      inValidEmail: true,
      isPhoneNumber: true,
      error: true,
      isPost: false,
      communicationForm: {
        fullName: null,
        email: null,
        tel: '+9936',
        text: null,
      },
      translationsContact: null,
      companyInformations: null,
    }
  },
  watch: {
    'communicationForm.email': function (val) {
      if (this.isPost) {
        if (val) {
          if (!this.inValidEmail) {
            this.inValidEmail = true
          }
        }
      }
    },
    'communicationForm.tel': function (val) {
      if (this.isPost === true) {
        if (val.length === 12) {
          this.isPhoneNumber = true
        } else {
          this.isPhoneNumber = false
        }
      }
    },
  },
  computed: {
    ...mapGetters('ui', ['contact']),
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
  async fetch() {
    try {
      const { data, status } = await translationContact({
        url: `${this.$i18n.locale}/translation-contact`,
      })
      if (status === 200) {
        this.translationsContact = data.translation_contact
      }
    } catch (error) {
      console.log(error)
    }
    try {
      const { data, status } = await companyPhones({
        url: `${this.$i18n.locale}/company-phones`,
      })
      if (status === 200) {
        this.companyInformations = data
      }
    } catch (error) {
      console.log(error)
    }
  },
  methods: {
    enforcePhoneFormat() {
      this.isPhoneNumber = true
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
      this.error = true
      this.isPhoneNumber = true
      this.inValidEmail = true
      this.isPost = true
      if (!this.communicationForm.fullName || !this.communicationForm.text) {
        this.error = false
      }
      if (this.communicationForm.tel.length !== 12) {
        this.isPhoneNumber = false
      }
      if (!this.communicationForm.email || !this.checkValidate) {
        this.inValidEmail = false
      }
      if (this.inValidEmail && this.error && this.isPhoneNumber) {
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
      }
    },
    clear() {
      this.communicationForm.fullName = null
      this.communicationForm.tel = '+9936'
      this.communicationForm.email = null
      this.communicationForm.text = null
      this.isPhoneNumber = true
      this.inValidEmail = true
    },
  },
}
</script>
<style>
.communication {
  width: 100%;
}
.communication__wrapper {
  background: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
}
.communication__form {
  width: 75%;
  padding: 10px 40px 30px 40px;
}
.communication__form-title {
  padding: 15px 0;
  font-family: TTNormsPro-Bold;
  font-style: normal;
  font-weight: 800;
  font-size: 19px;
  line-height: 120%;
  text-align: center;
  text-decoration-line: underline;
  color: #1b3254;
}

.communication__form-control {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
}
.control__input {
  padding: 10px 20px;
  width: 50%;
}

.control__input-label span {
  font-family: TTNormsPro;
  font-style: normal;
  font-weight: 800;
  font-size: 16px;
  line-height: 120%;
  text-align: center;
  color: #1b3254;
}
.control__input-input {
  border-radius: 5px;
  background: #eee;
  padding: 15px 8px;
  margin-top: 5px;
}
.control__input-input input {
  background: transparent;
  width: 100%;
  height: 100%;
  font-size: 16px;
}

.communication__info {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  width: 100%;
  padding: 10px 20px;
}
.communication__info-text {
  display: flex;
  flex-direction: column;
}
.communication__info-text span {
  margin: 5px;
  font-family: TTNormsPro;
  font-size: 16px;
  line-height: 120%;
  color: #1b3254;
}
.communication__info-button button {
  border-radius: 6px;
  background: #fd5e29;
  padding: 7px 25px;
  font-family: TTNormsPro-Bold;
  font-style: normal;
  font-weight: 600;
  font-size: 18px;
  line-height: 120%;
  color: #ffffff;
  width: 100%;
}

@media (max-width: 950px) {
  .communication {
    margin-bottom: 100px;
  }
  .communication__wrapper {
    margin: 0 20px;
    border-radius: 5px;
  }
  .communication__form {
    padding: 10px;
  }
}
@media (max-width: 950px) {
  .communication__form {
    width: 100%;
  }
}
@media (max-width: 555px) {
  .control__input {
    width: 100%;
  }
  .communication__info {
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
  }
  .communication__info-button {
    width: 100%;
    margin-top: 10px;
  }
}
@media (max-width: 320px) {
  .communication__wrapper {
    margin: 0 10px;
  }
  .communication__form {
    padding: 0px;
  }
  .control__input {
    padding: 10px;
  }
}
.control__input.error {
  animation: 0.2s invalid forwards;
}
.control__input.error .control__input-label span {
  color: #aa0000;
}
.control__input.error .control__input-input {
  border: 1px solid #aa0000;
}

@keyframes invalid {
  0% {
    transform: translateX(0px);
  }
  25% {
    transform: translateX(5px);
  }
  50% {
    transform: translateX(0px);
  }
  75% {
    transform: translateX(-5px);
  }
  100% {
    transform: translateX(0px);
  }
}
</style>
