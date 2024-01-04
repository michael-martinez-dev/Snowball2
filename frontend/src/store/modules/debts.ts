import { Module } from 'vuex';
import { StateInterface } from '../index';
import { Bill } from '../../models/Bills';
import { getBills, createBill, updateBill, deleteBill } from '../../api/debts';

interface DebtsState {
    debts: Bill[];
}

const debtsModule: Module<DebtsState, StateInterface> = {
    namespaced: true,
    state: {
        debts: []
    },
    mutations: {
        SET_DEBTS(state, debts: Bill[]) {
            state.debts = debts;
            },
        ADD_DEBT(state, debt: Bill) {
            state.debts.push(debt);
            },
        UPDATE_DEBT(state, updatedDebt: Bill) {
            const index = state.debts.findIndex(debt => debt.id === updatedDebt.id);
            if (index !== -1) {
                state.debts.splice(index, 1, updatedDebt);
            }
        },
        DELETE_DEBT(state, debtId: number) {
            state.debts = state.debts.filter(debt => debt.id !== debtId);
        }
    },
    actions: {
        async fetchDebts({ commit }) {
            const debts = await getBills();
            commit('SET_DEBTS', debts);
            },
        async addDebt({ commit }, debt: Bill) {
            const newDebt = await createBill(debt);
            commit('ADD_DEBT', newDebt);
            },
        async updateDebt({ commit }, debt: Bill) {
            const updatedDebt = await updateBill(debt);
            commit('UPDATE_DEBT', updatedDebt);
            },
        async removeDebt({ commit }, debtId: number) {
            await deleteBill(debtId);
            commit('DELETE_DEBT', debtId);
        }
    },
    getters: {
        totalDebt(state): number {
            return state.debts.reduce((total, debt) => total + debt.total, 0);
        }
    }
};

export default debtsModule;
