import axios from 'axios'
import jwt_decode from 'jwt-decode'

// import VuexPersistence from 'vuex-persistedstate'
/**
 * This function below clean the mess when you're using window API, it can't do it on server side
 * so this is the workaround I found on Github issues.
 * https://github.com/championswimmer/vuex-persist/issues/88#issuecomment-606881735 
 */
/*
    function getPlugins() {
    let plugins = []

    if (process.browser) {
        const vuexLocal = new VuexPersistence({
            storage: window.localStorage
        })
        plugins.push(vuexLocal.plugin)
        return plugins
    }
    }

    export const plugins = VuexPersistence()
*/

/**
 * State in Vuex is like Vue's data that can store.. well, data. 
 * But it's not just for a specific page. Instead, it's globally enabled.
 * So you can access the data (in this case--state) from any page.
 * 
 * You can access state from any page (in /pages folder) with:
 * this.$store.state.(filename in store).(variable name)
 * Example: this.$store.state.auth.authorized   // this will return the value 'false'
 * 
 * For learn more about Vue data: https://vuejs.org/v2/guide/index.html#Declarative-Rendering
 * For learn more about Vuex state: https://vuex.vuejs.org/guide/state.html 
 */
export const state = () => ({
    authorized: false,
    userData: [],
    userEmail: null,
    loginStatus: null,
    version: null
})

/**
 * Mutation in Vuex is the only way for changing values in State.
 * You can't change State values with something like:
 * this.authorized = true; 
 * But you have to do it through mutation. I think this will provide more security(?)
 * 
 * You can do mutation from any page (in /pages folder, once again) with:
 * this.$store.commit('(filename)/(mutation name)', values)
 * Example: this.$store.commit('auth/setAuthorized', true)
 * 
 * For learn more about Vuex Mutation: https://vuex.vuejs.org/guide/mutations.html
 */
export const mutations = {
    setAuth(state, payload) {
        state.authorized = payload
    },
    setLoginStatus(state, payload) {
        state.loginStatus = payload
    },
    setUserData(state, payload) {
        state.userData = payload
    },
    setUserEmail(state, payload) {
        state.userEmail = payload
    },
    setAuthorized(state, payload) {
        state.authorized = payload
    },
    setVersion(state, version) {
        state.version = version;
    }
}

/**
 * Actions is like function that you can do before commit data through Mutations. 
 * It uses async function so you can return a Promise from it.
 * 
 * You can call mutation from any page with:
 * this.$store.dispatch('(filename)/(function)', values)
 * 
 * Commit object below when called will call a Mutation on this same file. 
 * Dispatch object below when called will call another Actions function on this same file. 
 * 
 * For learn more about Vuex Actions: https://vuex.vuejs.org/guide/actions.html
 */
export const actions = {
    async nuxtServerInit({ commit }, { req }) {
        if (req.cookies.vuex) {
          const STATE = JSON.parse(req.cookies.vuex);
    
          commit('setVersion', STATE.version);
        }
    },
    async login({commit, dispatch}, {email, password, csrf}) {
        try {
            const res = await axios.post('http://localhost:3000/api/login', 
            { email, password, key: '2spX^cJdJ8wArpEei!AQ4nH%Y&%5s*bdmNoZuMPkyrqLjO*lO8t3ZY^EvSX^Nx!r#9t!BtXw0S%j5ranpR*7mL1&' },
            { headers: { 'csrf-token': csrf } } )
            console.log(res)
            if (res.data.message == 'failed') {
                commit('setLoginStatus', 'Wrong password')
                return Promise.reject(`Failed: wrong password. ${res.data['detail']}`)
            } else if (res.data.message = 'success') {
                commit('setLoginStatus', 'Login success')
                commit('setAuth', true)
                commit('setUserEmail', email)
                return dispatch('getUserData', {email: email})
                        .then(res => commit('setUserData', res))
                        .catch(err => console.log(err))
            } else {
                commit('setLoginStatus', 'Login failed')
                return Promise.reject(`Failed: login just failed. ${res.data['detail']}`)
            }
        } catch (error) {
            return Promise.reject(error + ` \nvalues:  ${email} ${password} ${csrf}`)
        }
    },
    async logout({commit, state}, {csrf}) {
        try {
            const res = await axios.post('http://localhost:3000/api/logout',
            {email: state.userEmail, key: '2spX^cJdJ8wArpEei!AQ4nH%Y&%5s*bdmNoZuMPkyrqLjO*lO8t3ZY^EvSX^Nx!r#9t!BtXw0S%j5ranpR*7mL1&'},
            {headers: { 'csrf-token': csrf } } )
            console.log(res)
            if (res.data.message == "failed") {
                return Promise.reject(`Logout failed. ${res.data['detail']}`)
            } else if (res.data.message == "success") {
                commit('setLoginStatus', null)
                commit('setAuth', false)
                commit('setUserEmail', null)
                commit('setUserData', null)
                return Promise.resolve("Logout success")
            } else {
                return Promise.reject(`Logout failed. ${res.data['detail']}`)
            }
        } catch (error) {
            return Promise.reject(error)
        }
    },
    async getUserData({commit}, {email}) {
        try {
            const res = await axios.post('http://localhost:3000/api/userdata', 
            { email, key: '2spX^cJdJ8wArpEei!AQ4nH%Y&%5s*bdmNoZuMPkyrqLjO*lO8t3ZY^EvSX^Nx!r#9t!BtXw0S%j5ranpR*7mL1&' }
            )
            const decoded = jwt_decode(res.data.data)
            return Promise.resolve(decoded)
        } catch (error) {
            return Promise.reject(error)
        }
    },
    async register({ commit }, { name, email, password, csrf }) {
        try {
            const res = await axios.post('http://localhost:3000/api/register', 
            { name, email, password, key: "2spX^cJdJ8wArpEei!AQ4nH%Y&%5s*bdmNoZuMPkyrqLjO*lO8t3ZY^EvSX^Nx!r#9t!BtXw0S%j5ranpR*7mL1&"},
            { headers: { 'csrf-token': csrf } })
            console.log(res)
            if (res.data.message == "failed") {
                return Promise.reject(`Failed: Server error. ${res.data['detail']}`)
            } else if (res.data.message == "success") {
                return Promise.resolve("Success: Please login")
            } else {
                return Promise.reject(`Failed: Server error. ${res.data['detail']}`)
            }
        } catch (error) {
            return Promise.reject(error)
        }
    }
}