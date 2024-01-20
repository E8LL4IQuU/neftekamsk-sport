/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,vue}",
  ],
  theme: {
    extend: {},
    screens: {
      'pc': {'max': '1300px'},
      'laptop':{'max':'935px'},
      'tablet':{'max':'760px'},
      'mobile':{'max':'520px'},
      'sm': {'min': '640px'},
      'md': {'min': '768px'},
      'lg': {'min': '1024px'},
      'xl': {'min': '1280px'},
      '2xl': {'min': '1536px'}
    }
  },
  plugins: [
  require('@tailwindcss/forms'),
  ],
}

