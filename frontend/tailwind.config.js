/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],

  theme: {
    extend: {
      colors: {
        'regal-blue': '#243c5a',
        appgrey: '#5C5C80'
      }
    }
  },
  plugins: []
}
