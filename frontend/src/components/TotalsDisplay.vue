<template>
    <v-container fluid>
        <v-card
            class="totals-card"
            elevation="5"
            width="100vw"
            style="align-self: start"
        >
            <v-card-title>
                <span class="text-h5">Debt Totals</span>
            </v-card-title>
            <v-card-text>
                <v-container fluid>
                    <v-switch
                        v-model="useActual"
                        label="Use Actual Monthly Payments"
                    />
                    <v-row>
                        <v-col>
                            <strong>Total Debt: </strong>
                            - {{ totalFormatted }}
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                            <strong>Monthly Payment: </strong>
                            - {{ monthlyFormatted }}
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                            <strong>Payments Left: </strong>
                            {{ paymentsLeft }} ({{
                                YearsAndMonthsLeft.years
                            }}
                            years, {{ YearsAndMonthsLeft.months }} months)*
                        </v-col>
                    </v-row>
                    <v-row>* doesn't consider interest</v-row>
                </v-container>
            </v-card-text>
        </v-card>
    </v-container>
</template>

<script lang="ts">
import { ref, computed, defineComponent } from "vue";
import { useDebtsStore } from "../store/modules/debts";

export default defineComponent({
    name: "TotalsDisplay",
    setup() {
        const debtsStore = useDebtsStore();
        const useActual = ref(true);
        debtsStore.fetchDebts();

        const monthly = computed(() => {
            console.log("useActual: ", useActual.value);
            debtsStore.fetchDebts();
            let amt = useActual.value
                ? debtsStore.totalMonthlyActual
                : debtsStore.totalMonthlyMin;
            console.log("amt = ", amt);
            return amt;
        });

        const totalFormatted = computed(() => {
            return debtsStore.totalDebt.toLocaleString("en-US", {
                style: "currency",
                currency: "USD",
            });
        });

        const monthlyFormatted = computed(() => {
            return monthly.value.toLocaleString("en-US", {
                style: "currency",
                currency: "USD",
            });
        });

        const paymentsLeft = computed(() => {
            const val = debtsStore.totalDebt / (monthly.value || 1);
            const payments = Math.ceil(val);
            return !Number.isNaN(payments) ? payments : 1;
        });

        const YearsAndMonthsLeft = computed(() => {
            let payments = paymentsLeft.value;
            let years = Math.floor(payments / 12);
            let months = payments % 12;
            return { years, months };
        });

        return {
            useActual,
            totalFormatted,
            monthlyFormatted,
            paymentsLeft,
            YearsAndMonthsLeft,
        };
    },
});
</script>
