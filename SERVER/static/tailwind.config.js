const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}"],
  theme: {
    extend: {
      inset: {
        '70': '278px',
    }
  },
    fontFamily: {
    Graphik: ['Graphik', 'sans-serif'],
    Merriweather: ['Merriweather', 'sans-serif'],
    Josefin: ['Josefin Sans', 'sans-serif'],
  },
plugins: []
}  
}
