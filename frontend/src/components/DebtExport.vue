<template>
    <v-btn
        flat
        color="primary"
        @click="exportDebts"
    >Export Debts</v-btn>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Bill } from "../models/Bills";

export default defineComponent({
    name: "DebtsTable",
    props: {
        bills: {
            type: Array as () => Bill[],
            required: true,
        },
    },
    methods: {
        exportDebts() {
            console.log("exportDebts");
            const csv = this.bills.map((bill) => {
                return [
                    bill.name,
                    bill.type,
                    bill.total,
                    bill.interest,
                    bill.monthlyMin,
                    bill.monthlyActual,
                ].join(",");
            }).join("\n");
            const blob = new Blob([csv], { type: "text/csv" });
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement("a");
            a.setAttribute("hidden", "");
            a.setAttribute("href", url);
            a.setAttribute("download", "debts.csv");
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
        }
    }
})
</script>