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
