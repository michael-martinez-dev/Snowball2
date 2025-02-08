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
            <strong>Years and Months Left (W/ Interest): </strong>
            {{
                monthsAndYearsLeftWithInterest.years === Infinity
                    ? "∞"
                    : monthsAndYearsLeftWithInterest.years
            }}
            years,
            {{
                monthsAndYearsLeftWithInterest.months === Infinity
                    ? "∞"
                    : monthsAndYearsLeftWithInterest.months
            }}
            months
        </div>
        <div>
            <strong>Years and Months Left (No Interest): </strong>
            {{ monthsAndYearsLeftNoInterest.years }} years,
            {{ monthsAndYearsLeftNoInterest.months }} months
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
} from "../utils/amortization";

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
        const monthsAndYearsLeftWithInterest = computed(() => {
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

            // return { years: 0, months: 0 };

            // // Get the maximum months to pay off all debts
            const maxMonths = getMaximumMonthsToPayOffAllDebts(shaped);

            if (maxMonths === Infinity) {
                return { years: Infinity, months: Infinity };
            }

            return {
                years: Math.floor(maxMonths / 12),
                months: maxMonths % 12,
            };
        });

        const monthsAndYearsLeftNoInterest = computed(() => {
            const months = totalPrincipal.value / totalMonthlyPayment.value;
            const years = Math.floor(months / 12);
            const monthsLeft = Math.round(months % 12);
            return { years, months: monthsLeft };
        });

        return {
            useActual,
            totalPrincipal,
            totalMonthlyPayment,
            monthsAndYearsLeftWithInterest,
            monthsAndYearsLeftNoInterest,
        };
    },
});
</script>
