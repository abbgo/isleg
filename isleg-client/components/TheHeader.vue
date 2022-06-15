<template>
  <header class="header __container">
    <div class="header__container">
      <div class="serach">
        <span class="search__icon">
          <svg
            width="22"
            height="21"
            viewBox="0 0 22 21"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M19.7734 20.712L14.6728 15.623C13.2047 16.6821 11.4401 17.2514 9.62982 17.25C7.34162 17.2476 5.1479 16.3381 3.52924 14.7207C1.91058 13.1034 0.99929 10.9103 0.995056 8.62207C0.999557 6.33441 1.91147 4.14189 3.53021 2.52539C5.14896 0.908894 7.34216 0.000342889 9.62982 -0.000976562C11.9177 7.75772e-05 14.1123 0.908528 15.7314 2.52502C17.3505 4.14152 18.2618 6.33415 18.2665 8.62207C18.2679 10.5813 17.5997 12.4821 16.372 14.009L21.4228 19.056C21.6214 19.2756 21.7255 19.565 21.7109 19.8607C21.6963 20.1564 21.564 20.4342 21.3447 20.6331C21.1171 20.8638 20.8084 20.9956 20.4843 21C20.3527 21.0012 20.2216 20.9763 20.0995 20.9269C19.9775 20.8775 19.8671 20.8045 19.7734 20.712ZM3.37299 8.59998C3.37616 10.2666 4.04028 11.8639 5.21967 13.0414C6.39906 14.2189 7.99742 14.8805 9.664 14.881C11.3288 14.8783 12.9245 14.2157 14.1015 13.0383C15.2785 11.8609 15.9412 10.2648 15.9433 8.59998C15.9409 6.93569 15.2786 5.34037 14.1015 4.16382C12.9244 2.98727 11.3283 2.32556 9.664 2.32397C7.99865 2.3245 6.40121 2.98556 5.2226 4.16211C4.04398 5.33866 3.38033 6.93463 3.37689 8.59998H3.37299ZM7.07318 6.97705C6.98247 6.81369 6.93451 6.62985 6.93451 6.44299C6.93451 6.25613 6.98247 6.07242 7.07318 5.90906C7.44372 5.53671 7.8846 5.24157 8.37006 5.04089C8.85551 4.84022 9.37601 4.73795 9.90131 4.73999C10.9733 4.73531 12.0028 5.15573 12.7656 5.90906C12.8622 6.0704 12.914 6.2549 12.914 6.44299C12.914 6.63108 12.8622 6.8157 12.7656 6.97705C12.6621 7.10898 12.5246 7.20937 12.3671 7.26697C12.2097 7.32457 12.0391 7.33701 11.8749 7.30298C11.6316 7.01683 11.3293 6.78712 10.9882 6.62964C10.6472 6.47216 10.276 6.39069 9.90033 6.39099C9.52513 6.39165 9.15399 6.47356 8.81342 6.63098C8.47284 6.78841 8.17028 7.01761 7.9267 7.30298C7.90386 7.32512 7.87669 7.34205 7.84662 7.35242C7.81655 7.36279 7.7845 7.3664 7.75287 7.36304C7.61839 7.35482 7.48806 7.31577 7.37103 7.24902C7.254 7.18227 7.1536 7.08959 7.07806 6.97803L7.07318 6.97705Z"
              fill="#FD5E29"
            />
          </svg>
        </span>
        <input class="input__border" type="text" placeholder="Gözleg" />
      </div>
      <a href="#" class="header__logo">
        <img src="@/assets/img/logo.svg" alt="" />
      </a>
      <div class="header__items">
        <span class="item___box">
          <img src="@/assets/img/flag.png" alt="" />
        </span>
        <span class="item__border"></span>
        <span class="item___box">
          <svg
            width="30"
            height="22"
            viewBox="0 0 30 22"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M0.99292 21.6281V19.05C0.99292 15.62 7.86737 13.8911 11.3074 13.8911C14.7474 13.8911 21.6257 15.62 21.6257 19.05V21.6281H0.99292ZM6.1521 6.15405C6.15183 4.78641 6.6942 3.47464 7.66089 2.5072C8.62758 1.53976 9.93973 0.995913 11.3074 0.995117C11.9849 0.994986 12.655 1.12823 13.281 1.38745C13.907 1.64667 14.4757 2.02678 14.9548 2.50586C15.4339 2.98494 15.8138 3.55371 16.073 4.17969C16.3322 4.80567 16.4657 5.47653 16.4656 6.15405C16.4657 6.83158 16.3322 7.50256 16.073 8.12854C15.8138 8.75452 15.4339 9.32328 14.9548 9.80237C14.4757 10.2815 13.907 10.6614 13.281 10.9207C12.655 11.1799 11.9849 11.3132 11.3074 11.3131C9.94025 11.3115 8.62889 10.7672 7.66284 9.7998C6.6968 8.83245 6.15476 7.52117 6.15503 6.15405H6.1521Z"
              fill="#8D98A9"
            />
            <path
              d="M11.313 11.3121C10.2928 11.3121 9.29557 11.0096 8.44734 10.4428C7.59911 9.87605 6.938 9.07048 6.5476 8.12798C6.1572 7.18548 6.05506 6.14837 6.25408 5.14782C6.4531 4.14727 6.94436 3.2282 7.66571 2.50684C8.38707 1.78548 9.30614 1.29423 10.3067 1.09521C11.3073 0.896183 12.3444 0.998329 13.2869 1.38873C14.2294 1.77912 15.0349 2.44024 15.6017 3.28846C16.1685 4.13669 16.471 5.13394 16.471 6.1541C16.4711 6.83149 16.3378 7.50228 16.0786 8.12813C15.8194 8.75399 15.4395 9.32266 14.9605 9.80165C14.4815 10.2806 13.9129 10.6606 13.287 10.9197C12.6612 11.1789 11.9904 11.3122 11.313 11.3121ZM22.919 8.7331V4.8641H25.496V8.7331H29.365V11.3121H25.496V15.1811H22.919V11.3121H19.05V8.7321L22.919 8.7331ZM11.313 13.8951C14.756 13.8951 21.63 15.6231 21.63 19.0531V21.6321H0.995972V19.0491C0.995972 15.6191 7.86897 13.8951 11.313 13.8951Z"
              fill="#8D98A9"
            />
          </svg>
        </span>
        <span class="item__border"></span>
        <span class="item___box shop">
          <svg
            width="22"
            height="21"
            viewBox="0 0 22 21"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M7.27123 16.5C6.8572 16.5 6.45248 16.621 6.10823 16.8477C5.76398 17.0744 5.49567 17.3966 5.33723 17.7735C5.17879 18.1505 5.13734 18.5653 5.21811 18.9655C5.29888 19.3657 5.49825 19.7332 5.79101 20.0218C6.08377 20.3103 6.45677 20.5068 6.86284 20.5864C7.2689 20.666 7.6898 20.6251 8.07231 20.469C8.45482 20.3128 8.78175 20.0484 9.01177 19.7091C9.24179 19.3699 9.36456 18.971 9.36456 18.563C9.36496 18.292 9.31109 18.0235 9.20603 17.7731C9.10098 17.5226 8.9468 17.295 8.75233 17.1034C8.55787 16.9117 8.32694 16.7598 8.07279 16.6562C7.81863 16.5527 7.54624 16.4996 7.27123 16.5V16.5ZM0.991211 0V2.063H3.08455L6.85215 9.892L5.43968 12.419C5.26654 12.7209 5.1763 13.0622 5.17789 13.409C5.17949 13.9557 5.40056 14.4795 5.79279 14.866C6.18502 15.2526 6.71653 15.4704 7.27123 15.472H19.8302V13.409H7.71059C7.6761 13.4094 7.64187 13.403 7.60993 13.3902C7.57798 13.3774 7.54896 13.3584 7.52457 13.3343C7.50018 13.3103 7.48091 13.2817 7.4679 13.2502C7.45488 13.2187 7.44839 13.185 7.4488 13.151L7.48026 13.027L8.4219 11.346H16.2189C16.5928 11.3473 16.9602 11.2493 17.2824 11.0625C17.6047 10.8756 17.87 10.6067 18.0504 10.284L21.7978 3.59C21.8829 3.43689 21.9263 3.26459 21.9236 3.09C21.9217 2.81784 21.811 2.55739 21.6156 2.36513C21.4201 2.17287 21.1556 2.06431 20.8794 2.063H5.39706L4.41381 0H0.991211ZM17.7339 16.5C17.3198 16.5 16.9151 16.621 16.5709 16.8477C16.2266 17.0744 15.9583 17.3966 15.7999 17.7735C15.6414 18.1505 15.6 18.5653 15.6807 18.9655C15.7615 19.3657 15.9609 19.7332 16.2536 20.0218C16.5464 20.3103 16.9194 20.5068 17.3255 20.5864C17.7315 20.666 18.1524 20.6251 18.5349 20.469C18.9175 20.3128 19.2444 20.0484 19.4744 19.7091C19.7044 19.3699 19.8272 18.971 19.8272 18.563C19.8276 18.292 19.7737 18.0235 19.6687 17.7731C19.5636 17.5226 19.4094 17.295 19.215 17.1034C19.0205 16.9117 18.7896 16.7598 18.5354 16.6562C18.2813 16.5527 18.0089 16.4996 17.7339 16.5V16.5Z"
              fill="#8D98A9"
            />
          </svg>
        </span>
      </div>
    </div>
    <nav class="menu-header__nav">
      <ul class="menu-header__list">
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/idea.png" alt="" /> </span>
            <h4 class="nav__item">Elektrik</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link menu-header__link__active"
            ><span><img src="@/assets/img/menu/kanselaria.png" alt="" /> </span>
            <h4 class="nav__item">Kanselýar</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/balloons.png" alt="" /> </span>
            <h4 class="nav__item">Sowgatlyk</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/carrot.png" alt="" /> </span>
            <h4 class="nav__item">Iýmit</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/armchair.png" alt="" /> </span>
            <h4 class="nav__item">Mebel</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/men-hoodie.png" alt="" /> </span>
            <h4 class="nav__item">Egin-eşik</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span
              ><img src="@/assets/img/menu/ps4-controller.png" alt="" />
            </span>
            <h4 class="nav__item">Oýun</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/nike-dunk.png" alt="" /> </span>
            <h4 class="nav__item">Aýakgap</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/support.png" alt="" /></span>
            <h4 class="nav__item">Gurluşyk</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/nike-dunk.png" alt="" /> </span>
            <h4 class="nav__item">Aýakgap</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/support.png" alt="" /></span>
            <h4 class="nav__item">Gurluşyk</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/nike-dunk.png" alt="" /> </span>
            <h4 class="nav__item">Aýakgap</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/support.png" alt="" /></span>
            <h4 class="nav__item">Gurluşyk</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/nike-dunk.png" alt="" /> </span>
            <h4 class="nav__item">Aýakgap</h4></a
          >
        </li>
        <li class="menu-header__item">
          <a href="#" class="menu-header__link"
            ><span><img src="@/assets/img/menu/support.png" alt="" /></span>
            <h4 class="nav__item">Gurluşyk</h4></a
          >
        </li>
      </ul>
    </nav>
    <div class="header__sub-menu">
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
      <span class="sub-menu__items">Lampa</span>
      <span class="sub-menu__items">Rozetka</span>
      <span class="sub-menu__items">Hojalyk tehnika</span>
      <span class="sub-menu__items">Ütükler</span>
      <span class="sub-menu__items">Kolonkalar</span>
      <span class="sub-menu__items">Sowadyjylar</span>
    </div>
  </header>
</template>

<script>
export default {
  mounted() {
    let search = document.querySelector('.serach')
    let inputFocus = document.querySelector('.input__border')

    window.addEventListener('click', (event) => {
      search.classList.add('search__focus')
      inputFocus.classList.add('serach__input-focus')
      if (!event.target.classList.contains('input__border')) {
        search.classList.remove('search__focus')
        inputFocus.classList.remove('serach__input-focus')
      }
    })
  },
}
</script>
