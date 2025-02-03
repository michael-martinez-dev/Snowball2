import { defineStore } from "pinia";
import { Bill } from "../../models/Bills";
import { getBills, createBill, updateBill, deleteBill } from "../../api/debts";

export const useDebtsStore = defineStore("debts", {
  state: () => ({
    debts: [] as Bill[],
  }),

  getters: {
    totalDebt(state): number {
      return state.debts.reduce((total, debt) => total + debt.total, 0);
    },
    totalMonthlyMin(state): number {
      return state.debts.reduce((sum, debt) => sum + debt.monthlyMin, 0);
    },
    totalMonthlyActual(state): number {
      return state.debts.reduce((sum, debt) => sum + debt.monthlyActual, 0);
    },
  },

  actions: {
    async fetchDebts() {
      console.log("fetchDebts...");
      const debts = await getBills();
      this.debts = debts;
    },

    addDebt(debt: Bill) {
      createBill(debt).then(this.fetchDebts);
    },

    updateDebt(debt: Bill) {
      updateBill(debt).then(this.fetchDebts);
    },

    removeDebt(debtId: number) {
      deleteBill(debtId).then(this.fetchDebts);
    },
  },
});
