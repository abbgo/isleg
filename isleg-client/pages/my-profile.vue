<template>
  <div class="__container">
    <div class="myprofile">
      <div class="myprofile__wrapper">
        <div class="myprofile__content">
          <div class="myprofile__title">
            <h4>Maglumatlarym</h4>
          </div>
          <div class="myprofile__form">
            <text-filed
              placeholder="Kemal Hanow "
              label="Doly Adyňyz"
              :value="myProfileDatas.fullName"
              @updateValue="(val) => (myProfileDatas.fullName = val)"
            ></text-filed>
            <text-filed
              placeholder="Email"
              label="Email"
              :value="myProfileDatas.email"
              @updateValue="(val) => (myProfileDatas.email = val)"
            ></text-filed>
            <text-filed
              label="Telefon"
              :type="'tel'"
              :value="myProfileDatas.phone_number"
              @updateValue="(val) => update(val)"
            ></text-filed>
            <text-filed
              label="Doglan senäňiz"
              :value="some"
              type="date"
              @updateValue="(val) => (some = val)"
            ></text-filed>
            <text-filed
              placeholder=""
              label="Salgyňyz"
              :value="myProfileDatas.address"
              prepenedIcon="chevron-left.svg"
              appendIcon="chevron-left.svg"
              @updateValue="(val) => (myProfileDatas.address = val)"
            ></text-filed>
            <div class="form__box gender__input">
              <span>Jynsy</span>
              <div class="gender__checkbox">
                <div>
                  <input
                    class="custom"
                    v-model="myProfileDatas.male.boy"
                    id="male"
                    name="gender"
                    type="radio"
                  />
                  <label for="male">Erkek</label>
                </div>
                <div>
                  <input
                    class="custom"
                    v-model="myProfileDatas.male.girl"
                    id="woman"
                    name="gender"
                    type="radio"
                  />
                  <label for="woman"> Aýal</label>
                </div>
              </div>
            </div>
          </div>
          <div class="myprofile__btns">
            <button class="myprofile__btn btn__key">
              <img src="@/assets/img/key.png" alt="" />
              <span>Açar söz üýtget</span>
            </button>
            <button class="myprofile__btn">Ýatda sakla</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- <section class="agreement __container" v-if="myProfileDatas.fillName">
    <h4 class="agreement__title">Profilim</h4>
    <div class="communication_form">
      <div class="form__box">
        <span>Doly adyňyz</span>
        <input
          type="text"
          v-model="myProfileDatas.fillName"
          placeholder="Doly adyňyz"
        />
      </div>
      <div class="form__box">
        <span>Email</span>
        <input type="text" v-model="myProfileDatas.email" placeholder="Email" />
      </div>
      <div class="form__box">
        <span>Telefon</span>
        <input
          type="text"
          v-model="myProfileDatas.phone_number"
          placeholder="+993"
        />
      </div>
      <div class="form__box">
        <span>Salgyňyz</span>
        <input
          type="text"
          v-model="myProfileDatas.address"
          placeholder="Salgyňyz"
        />
      </div>
      <div class="form__box born__date">
        <span>Doglan senäňiz</span>
        <input
          v-model="myProfileDatas.birthday"
          id="born"
          type="date"
          placeholder="Doglan senäňiz"
        />
      </div>
      <div class="form__box sex__input">
        <span>Jynsy</span>
        <div>
          <input
            class="custom"
            v-model="myProfileDatas.male.boy"
            id="male"
            name="sex"
            type="radio"
          />
          <label for="male">Erkek</label>
        </div>
        <div>
          <input
            class="custom"
            v-model="myProfileDatas.male.girl"
            id="woman"
            name="sex"
            type="radio"
          />
          <label for="woman"> Aýal</label>
        </div>
      </div>
    </div>
    <div class="profile__btns">
      <div class="datas__left">
        <button
          class="product__add-btn confirim__btn"
          @click="openChangePassword"
        >
         
          <h4 class="confirm__text">Açar sözi üýtget</h4>
        </button>
      </div>
      <div class="datas__right">
        <button class="product__add-btn send__btn" @click="postMyInformation">
          <h4>Ýatda sakla</h4>
        </button>
      </div>
    </div>
    <pop-up-change-password
      :isChangePassword="isChangePassword"
      @closeChangePassword="closeChangePassword"
    ></pop-up-change-password>
  </section> -->
</template>

<script>
import { mapGetters } from 'vuex'
import { getMyProfile } from '@/api/myProfile.api'
import { getRefreshToken, getMyInformation } from '@/api/user.api'
import TextFiled from '@/components/app/TextFiled.vue'

export default {
  components: { TextFiled },
  //   middleware: ['check-auth', 'user-auth'],
  data() {
    return {
      some: '',
      isChangePassword: false,
      myProfileDatas: {
        fullName: '',
        phone_number: '+9936',
        email: '',
        address: '',
        male: {
          boy: false,
          girl: false,
        },
        birthday: null,
      },
    }
  },
  computed: {
    ...mapGetters('ui', ['myProfile']),
  },
  async mounted() {
    const cart = await JSON.parse(localStorage.getItem('lorem'))
    if (!cart) {
      this.$router.replace(this.localeLocation('/'))
    } else if (!cart.auth) {
      this.$router.replace(this.localeLocation('/'))
    } else if (!cart.auth.accessToken) {
      this.$router.replace(this.localeLocation('/'))
    } else {
      await this.fetchMyProfile()
    }
  },
  methods: {
    update(val) {
      this.myProfileDatas.phone_number = val
    },
    async fetchMyProfile() {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      try {
        const { customer_informations, status } = (
          await getMyProfile({
            url: `${this.$i18n.locale}/my-information`,
            accessToken: `Bearer ${cart.auth.accessToken}`,
          })
        ).data
        if (status) {
          this.myProfileDatas.fullName = customer_informations.full_name
          this.myProfileDatas.phone_number = customer_informations.phone_number
          this.myProfileDatas.email = customer_informations.email
          this.myProfileDatas.birthday = customer_informations.birthday
            ? new Date(customer_informations.birthday.Time)
            : null
        }
      } catch (error) {
        console.log('getMyProfile1', error)
        if (error?.response?.status === 403) {
          try {
            const { access_token, refresh_token, status } = (
              await getRefreshToken({
                url: `auth/refresh`,
                refreshToken: `Bearer ${cart.auth.refreshToken}`,
              })
            ).data
            console.log('new', access_token, refresh_token, status)
            if (status) {
              const lorem = await JSON.parse(localStorage.getItem('lorem'))
              if (lorem) {
                lorem.auth = {
                  accessToken: access_token,
                  refreshToken: refresh_token,
                }
                localStorage.setItem('lorem', JSON.stringify(lorem))
              } else {
                localStorage.setItem(
                  'lorem',
                  JSON.stringify({
                    auth: {
                      accessToken: access_token,
                      refreshToken: refresh_token,
                    },
                  })
                )
              }
              try {
                const { customer_informations, status } = (
                  await getMyProfile({
                    url: `${this.$i18n.locale}/my-information`,
                    accessToken: `Bearer ${access_token}`,
                  })
                ).data
                if (status) {
                  this.myProfileDatas.fullName = customer_informations.full_name
                  this.myProfileDatas.phone_number =
                    customer_informations.phone_number
                  this.myProfileDatas.email = customer_informations.email
                  this.myProfileDatas.birthday = customer_informations.birthday
                    ? new Date(customer_informations.birthday.Time)
                    : null
                }
              } catch (error) {
                console.log('getMyProfile2', error)
              }
            }
          } catch (error) {
            console.log('ref', error.response.status)
            if (error.response.status === 403) {
              cart.auth.accessToken = null
              cart.auth.refreshToken = null
              localStorage.setItem('lorem', JSON.stringify(cart))
              this.$router.push(this.localeLocation('/'))
            }
          }
        }
      }
    },
    async postMyInformation() {
      const cart = await JSON.parse(localStorage.getItem('lorem'))
      try {
        const { customer_informations, status } = (
          await getMyInformation({
            url: `${this.$i18n.locale}/my-information`,
            accessToken: `Bearer ${cart.auth.accessToken}`,
          })
        ).data
        if (status) {
          this.myProfileDatas.fullName = customer_informations.full_name
          this.myProfileDatas.phone_number = customer_informations.phone_number
          this.myProfileDatas.email = customer_informations.email
        }
      } catch (error) {
        console.log('getMyProfile1', error)
      }
    },
    openChangePassword() {
      this.isChangePassword = true
      document.body.classList.add('_lock')
    },
    closeChangePassword() {
      this.isChangePassword = false
      document.body.classList.remove('_lock')
    },
  },
}
</script>
<style lang="scss">
.myprofile {
  width: 100%;
  @media (max-width: 950px) {
    margin-bottom: 100px;
  }
  &__wrapper {
    background: #fff;
    display: flex;
    align-items: center;
    justify-content: center;
    @media (max-width: 950px) {
      margin: 20px;
      border-radius: 5px;
    }
  }
  &__content {
    width: 75%;
    // padding: 10px 150px 30px 150px;

    @media (max-width: 950px) {
      width: 80%;
    }
    @media (max-width: 530px) {
      width: 100%;
    }
  }
  &__title {
    h4 {
      padding: 15px 0;
      font-family: TTNormsPro-Bold;
      font-style: normal;
      font-weight: 800;
      font-size: 18px;
      line-height: 120%;
      text-align: center;
      text-decoration-line: underline;
      color: #1b3254;
    }
  }
  &__form {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
  }
  &__btns {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 40px 20px;
    @media (max-width: 435px) {
      padding: 40px 0 20px 0;
    }
    @media (max-width: 435px) {
      flex-direction: column;
      align-items: flex-start;
      padding-left: 20px;
      padding-right: 20px;
    }
  }
  &__btn {
    border-radius: 5px;
    padding: 8px 5px;
    background: #fd5e29;
    font-family: TTNormsPro;
    font-style: normal;
    font-weight: 700;
    font-size: 18px;
    line-height: 120%;
    color: #ffffff;
    margin-bottom: 10px;
    @media (max-width: 435px) {
      width: 100%;
    }
    img {
      width: 16px;
      height: 16px;
      object-fit: contain;
      object-position: center;
    }
  }
}
.btn__key {
  background: #8d98a9;
}
</style>
