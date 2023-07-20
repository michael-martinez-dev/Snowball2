<template>
<div class="debt-table">
    <DebtsTable
        :bills="bills"
        :useActual="useActual"
        @addDebt="createDebt"
        @updateDebt="updateDebt"
        @deleteDebt="deleteDebt"
    />
    <DebtExport
        :bills="bills"
    />
    <TotalsDisplay
        :total="total"
        :monthly="monthly"
        :useActual="useActual"
        @useActual="useActual = !useActual"
    />
</div>
</template>

<script lang="ts">
import DebtsTable from "./DebtsTable.vue";
import TotalsDisplay from "./TotalsDisplay.vue";
import DebtExport from "./DebtExport.vue";
import {defineComponent} from "vue";
import {Bill} from "../models/Bills";
import { getBills, createBill, updateBill, deleteBill } from "../api/debts";

export default defineComponent({
  name: "DebtsMain",
  data() {
    return {
      bills: [] as Bill[],
      useActual: true,
    }
  },
  computed: {
    total() : number {
      return this.bills.reduce((total, bill) => total + bill.total, 0);
    },
    monthly() : number {
      if (this.useActual) {
        return this.bills.reduce((total, bill) => total + bill.monthlyActual, 0);
      } else {
        return this.bills.reduce((total, bill) => total + bill.monthlyMin, 0);
      }
    },
  },
  components: {
    DebtsTable,
    TotalsDisplay,
    DebtExport,
  },
  created() {
    this.getDebts();
  },
  methods: {
    async getDebts() {
      console.log("getDebts");
      this.bills = await getBills();
    },
    async createDebt(bill: Bill) {
      console.log("createDebt");
      await createBill({
        id: bill.id,
        name: bill.name,
        type: bill.type,
        total: bill.total,
        interest: bill.interest,
        monthlyMin: bill.monthlyMin,
        monthlyActual: bill.monthlyActual,
        dueDay: bill.dueDay,
      }).then(
          () => this.getDebts()
      );
    },
    async updateDebt(bill: Bill) {
      console.log("updateDebt");
      await updateBill({
        id: bill.id,
        name: bill.name,
        type: bill.type,
        total: bill.total,
        interest: bill.interest,
        monthlyMin: bill.monthlyMin,
        monthlyActual: bill.monthlyActual,
        dueDay: bill.dueDay,
      }).then(
          () => this.getDebts()
      );
    },
    async deleteDebt(bill: Bill) {
      console.log("deleteDebt");
      await deleteBill(bill.id).then(
          () => this.getDebts()
      );
    },
  },
})
</script>

<style>


</style>
