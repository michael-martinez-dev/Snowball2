<template>
  <v-container fluid>
    <v-card
        class="totals-card"
        elevation="5"
    >
      <v-card-title>
        <span class="text-h5">Debt Totals</span>
      </v-card-title>
      <v-card-text>
        <v-container fluid>
          <v-switch
              v-model="useActual"
              label="Use Actual Monthly Payments"
              @change="$emit('useActual')"
          />
          <v-row>
            <v-col>
              <strong>Total Debt: </strong>
              - {{ totalFormatted }}
            </v-col>
            <v-col
                cols="12"
                sm="6"
                md="4"
            >
              <strong>Monthly Payment: </strong>
              - {{ monthlyFormatted }}
            </v-col>
            <v-col
                cols="12"
                sm="6"
                md="4"
            >
              <strong>Payments Left: </strong>
              {{ paymentsLeft }} ({{ YearsAndMonthsLeft.years }} years, {{ YearsAndMonthsLeft.months }} months)
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
  name: "TotalsDisplay",
  props: {
    total: {
      type: Number,
      required: true,
    },
    monthly: {
      type: Number,
      required: true,
    },
    useActual: {
      type: Boolean,
      required: true,
    },
  },
  computed: {
    totalFormatted(): string {
      return this.total.toLocaleString("en-US", {
        style: "currency",
        currency: "USD",
      });
    },
    monthlyFormatted(): string {
      return this.monthly.toLocaleString("en-US", {
        style: "currency",
        currency: "USD",
      });
    },
    paymentsLeft(): number {
      let payments = Math.ceil(this.total / this.monthly);
      return !Number.isNaN(payments) ? payments : 0;
    },
    YearsAndMonthsLeft(): { years: number; months: number } {
      let years = Math.floor(this.paymentsLeft / 12);
      let months = this.paymentsLeft % 12;
      return { years, months };
    },
  },
});
</script>

