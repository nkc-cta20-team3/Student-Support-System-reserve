import { createStore } from 'vuex'

const store = createStore({
  state() {
    return {
      userId: null,
      ClassId: null,
      DocId: null
    }
  },
  mutations: {
    setUserId(state, userId) {
      state.userId = userId
    },
    setClassId(state, ClassId) {
      state.ClassId = ClassId
    },
    setDocId(state, DocId) {
      state.DocId = DocId
    }
  },
  actions: {
    updateUserId({ commit }, userId) {
      // ユーザIDをVuexのstateに保存する
      commit('setUserId', userId)
    },
    updateClassId({ commit }, ClassId) {
      commit('setClassId', ClassId)
    },
    updateDocId({ commit }, DocId) {
      commit('setDocId', DocId)
    }
  },
  getters: {
    getUserId(state) {
      return state.userId
    },
    getClassId(state) {
      return state.ClassId
    },
    getDocId(state) {
      return state.DocId
    }
  }
})

export default store