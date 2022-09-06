// Styles
import "@mdi/font/css/materialdesignicons.css"
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import 'vuetify/styles'

// Vuetify
import { createVuetify } from 'vuetify'

export default createVuetify({
    icons: {
      iconfont: 'md',
    },
    theme: {
      defaultTheme: 'light'
    }
    // https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides

})
  
