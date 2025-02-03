<template>
    <v-container>
        <v-switch v-model="useActual" label="Use Actual?" />
        <div>
            <strong>Total Principal: </strong>
            {{
                totalPrincipal.toLocaleString("en-US", {
                    style: "currency",
                    currency: "USD",
                })
            }}
        </div>
        <div>
            <strong>Total Monthly Payment: </strong>
            {{
                totalMonthlyPayment.toLocaleString("en-US", {
                    style: "currency",
                    currency: "USD",
                })
            }}
        </div>
        <div>
            <strong>Months Left: </strong>
            {{ monthsLeft === Infinity ? "∞" : monthsLeft }}
        </div>
    </v-container>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from "vue";
import { useDebtsStore } from "../store/modules/debts";

// Import your new amortization utils
import {
    monthsToPayOff,
    getMaximumMonthsToPayOffAllDebts,
} from "../utils/amortization"; // <-- Path depends on your folder structure

interface DebtItem {
    id: number;
    name: string;
    type: string;
    total: string; // principal
    interest: string; // annual interest rate (%)
    monthlyMin: string;
    monthlyActual: string;
    dueDay: number;
}

export default defineComponent({
    name: "TotalsDisplay",
    setup() {
        const debtsStore = useDebtsStore();
        const useActual = ref(true);

        // If you want the total principal
        const totalPrincipal = computed(() => {
            return debtsStore.debts.reduce((acc: number, d: DebtItem) => {
                return acc + (parseFloat(d.total) || 0);
            }, 0);
        });

        // Summing monthly payments (choose min or actual)
        const totalMonthlyPayment = computed(() => {
            return debtsStore.debts.reduce((acc: number, d: DebtItem) => {
                const pay = useActual.value
                    ? parseFloat(d.monthlyActual) || 0
                    : parseFloat(d.monthlyMin) || 0;
                return acc + pay;
            }, 0);
        });

        // For the "max months" approach
        const monthsLeft = computed(() => {
            // Convert each debt into the shape needed by getMaximumMonthsToPayOffAllDebts
            const shaped = debtsStore.debts.map((d: DebtItem) => {
                return {
                    principal: parseFloat(d.total) || 0,
                    annualInterestRate: parseFloat(d.interest) || 0,
                    monthlyPayment: useActual.value
                        ? parseFloat(d.monthlyActual) || 0
                        : parseFloat(d.monthlyMin) || 0,
                };
            });

            return getMaximumMonthsToPayOffAllDebts(shaped);
        });

        return {
            useActual,
            totalPrincipal,
            totalMonthlyPayment,
            monthsLeft,
        };
    },
});
</script>
