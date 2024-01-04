// frontend/src/store/index.ts
import { createStore } from 'vuex';
import debts from './modules/debts';
import { Debt } from "../types/Debt";

export interface StateInterface {
  debts: Debt[]
}

export default createStore({
  modules: {
    debts
  }
});
