<template>
  <div class="profile__box">
    <div
      class="profile__item"
      @click="$router.push(localeLocation('/my-profile'))"
    >
      <span>
        <svg
          width="20"
          height="20"
          viewBox="0 0 20 20"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M9.688 9.688C10.6461 9.688 11.5826 9.40391 12.3792 8.87164C13.1758 8.33938 13.7966 7.58285 14.1633 6.69772C14.5299 5.8126 14.6258 4.83863 14.4389 3.89898C14.252 2.95934 13.7907 2.09622 13.1132 1.41878C12.4358 0.741331 11.5727 0.279985 10.633 0.0930783C9.69337 -0.0938285 8.71941 0.00209886 7.83428 0.36873C6.94916 0.735361 6.19263 1.35623 5.66036 2.15282C5.1281 2.94941 4.844 3.88595 4.844 4.844C4.84387 5.48016 4.96907 6.11012 5.21246 6.69788C5.45585 7.28564 5.81265 7.81969 6.26248 8.26952C6.71231 8.71935 7.24637 9.07615 7.83413 9.31954C8.42189 9.56293 9.05184 9.68813 9.688 9.688ZM9.688 12.11C6.454 12.11 0 13.733 0 16.954V19.376H19.377V16.955C19.377 13.733 12.922 12.11 9.688 12.11Z"
            fill="#1B3254"
          />
        </svg>
      </span>
      <h4>Maglumatym</h4>
    </div>
    <div
      class="profile__item"
      @click="$router.push(localeLocation('/my-orders'))"
    >
      <span>
        <svg
          width="18"
          height="21"
          viewBox="0 0 18 21"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M16 5H14C14 3.67392 13.4732 2.40215 12.5355 1.46447C11.5979 0.526784 10.3261 0 9 0C7.67392 0 6.40215 0.526784 5.46447 1.46447C4.52678 2.40215 4 3.67392 4 5H2C1.4713 5.00264 0.965159 5.21452 0.592247 5.58931C0.219334 5.9641 0.00999338 6.4713 0.00999999 7L0 19C0.00158273 19.5299 0.212805 20.0377 0.587535 20.4125C0.962265 20.7872 1.47005 20.9984 2 21H16C16.5299 20.9984 17.0377 20.7872 17.4125 20.4125C17.7872 20.0377 17.9984 19.5299 18 19V7C17.9984 6.47005 17.7872 5.96227 17.4125 5.58753C17.0377 5.2128 16.5299 5.00158 16 5ZM9 2C9.79565 2 10.5587 2.31607 11.1213 2.87868C11.6839 3.44129 12 4.20435 12 5H6C6 4.20435 6.31607 3.44129 6.87868 2.87868C7.44129 2.31607 8.20435 2 9 2ZM9 12C7.67392 12 6.40215 11.4732 5.46447 10.5355C4.52678 9.59785 4 8.32608 4 7H6C6 7.79565 6.31607 8.55871 6.87868 9.12132C7.44129 9.68393 8.20435 10 9 10C9.79565 10 10.5587 9.68393 11.1213 9.12132C11.6839 8.55871 12 7.79565 12 7H14C14 8.32608 13.4732 9.59785 12.5355 10.5355C11.5979 11.4732 10.3261 12 9 12Z"
            fill="#1B3254"
          />
        </svg>
      </span>
      <h4>Sargytlarym</h4>
    </div>
    <div class="profile__item profile__item-active">
      <span>
        <svg
          width="23"
          height="21"
          viewBox="0 0 23 21"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M5.16973 6.99398H14.8007V2.58398L22.5527 10.337L14.8007 18.09V14.418L5.17073 14.418L5.16973 6.99398Z"
            fill="#FD5E29"
          />
          <path
            d="M10.3379 0L0.000890732 0L0.000890732 20.673H10.3379V17.228H3.44689L3.44689 3.446L10.3379 3.446V0Z"
            fill="#FD5E29"
          />
        </svg>
      </span>
      <h4>Çyk</h4>
    </div>
  </div>
</template>

<script>
export default {}
</script>
